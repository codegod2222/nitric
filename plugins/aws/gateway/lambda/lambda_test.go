package lambda_service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	events "github.com/aws/aws-lambda-go/events"
	"github.com/nitric-dev/membrane/sdk"
	"github.com/nitric-dev/membrane/sources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	plugin "github.com/nitric-dev/membrane/plugins/aws/gateway/lambda"
)

type MockSourceHandler struct {
	httpRequests []*sources.HttpRequest
	events       []*sources.Event
}

func (m *MockSourceHandler) HandleEvent(source *sources.Event) error {
	if m.events == nil {
		m.events = make([]*sources.Event, 0)
	}
	m.events = append(m.events, source)

	return nil
}

func (m *MockSourceHandler) HandleHttpRequest(source *sources.HttpRequest) *http.Response {
	if m.httpRequests == nil {
		m.httpRequests = make([]*sources.HttpRequest, 0)
	}
	m.httpRequests = append(m.httpRequests, source)

	return &http.Response{
		Status:     "OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("Mock Handled!"))),
	}
}

func (m *MockSourceHandler) reset() {
	m.httpRequests = make([]*sources.HttpRequest, 0)
	m.events = make([]*sources.Event, 0)
}

type MockLambdaRuntime struct {
	plugin.LambdaRuntimeHandler
	// FIXME: Make this a union array of stuff to send....
	eventQueue []interface{}
}

func (m *MockLambdaRuntime) Start(handler interface{}) {
	// cast the function type to what we know it will be
	typedFunc := handler.(func(ctx context.Context, event plugin.Event) (interface{}, error))
	for _, event := range m.eventQueue {

		bytes, _ := json.Marshal(event)
		evt := plugin.Event{}

		json.Unmarshal(bytes, &evt)
		// Unmarshal the thing into the event type we expect...
		// TODO: Do something with out results here...
		_, err := typedFunc(context.TODO(), evt)

		if err != nil {
			// Print the error?
		}
	}
}

var _ = Describe("Lambda", func() {
	mockHandler := &MockSourceHandler{}
	AfterEach(func() {
		mockHandler.reset()
	})

	Context("Http Events", func() {
		When("Sending a compliant HTTP Event", func() {

			runtime := MockLambdaRuntime{
				// Setup mock events for our runtime to process...
				eventQueue: []interface{}{&events.APIGatewayV2HTTPRequest{
					Headers: map[string]string{
						"User-Agent":            "Test",
						"x-nitric-payload-type": "TestPayload",
						"x-nitric-request-id":   "test-request-id",
						"Content-Type":          "text/plain",
					},
					RawPath: "/test/test",
					Body:    "Test Payload",
					RequestContext: events.APIGatewayV2HTTPRequestContext{
						HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
							Method: "GET",
						},
					},
				}},
			}

			client, _ := plugin.NewWithRuntime(runtime.Start)

			// This function will block which means we don't need to wait on processing,
			// the function will unblock once processing has finished, this is due to our mock
			// handler only looping once over each request
			It("The gateway should translate into a standard NitricRequest", func() {
				client.Start(mockHandler)

				Expect(len(mockHandler.httpRequests)).To(Equal(1))

				request := mockHandler.httpRequests[0]

				By("Retaining the body")
				//
				By("Retaining the Headers")
				//
				By("Retaining the method")
				Expect(request.Method).To(Equal("GET"))
				By("Retaining the path")
				Expect(request.Path).To(Equal("/test/test"))
			})
		})
	})

	Context("SNS Events", func() {
		When("The Lambda Gateway recieves SNS events", func() {
			topicName := "MyTopic"
			eventPayload := map[string]interface{}{
				"test": "test",
			}
			// eventBytes, _ := json.Marshal(&eventPayload)

			event := sdk.NitricEvent{
				RequestId:   "test-request-id",
				PayloadType: "test-payload",
				Payload:     eventPayload,
			}

			messageBytes, _ := json.Marshal(&event)

			runtime := MockLambdaRuntime{
				// Setup mock events for our runtime to process...
				eventQueue: []interface{}{&events.SNSEvent{
					Records: []events.SNSEventRecord{
						events.SNSEventRecord{
							EventVersion:         "",
							EventSource:          "aws:sns",
							EventSubscriptionArn: "some:arbitrary:subscription:arn:MySubscription",
							SNS: events.SNSEntity{
								TopicArn: fmt.Sprintf("some:arbitrary:topic:arn:%s", topicName),
								Message:  string(messageBytes),
							},
						},
					},
				}},
			}

			client, _ := plugin.NewWithRuntime(runtime.Start)

			It("The gateway should translate into a standard NitricRequest", func() {
				// This function will block which means we don't need to wait on processing,
				// the function will unblock once processing has finished, this is due to our mock
				// handler only looping once over each request
				client.Start(mockHandler)

				Expect(len(mockHandler.events)).To(Equal(1))

				request := mockHandler.events[0]

				By("Containing the Source Topic")
				Expect(request.Topic).To(Equal("MyTopic"))

				//Expect(request.ContentType).To(Equal("application/json"))
				//Expect(eventBytes).To(BeEquivalentTo(request.Payload))
				//Expect(context.PayloadType).To(Equal("test-payload"))
				//Expect(context.RequestId).To(Equal("test-request-id"))
				//Expect(context.SourceType).To(Equal(sdk.Subscription))
				//Expect(context.Source).To(Equal(topicName))
			})
		})
	})
})
