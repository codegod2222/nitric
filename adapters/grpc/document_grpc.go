// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"context"

	pb "github.com/nitric-dev/membrane/interfaces/nitric/v1"
	"github.com/nitric-dev/membrane/sdk"
	"google.golang.org/protobuf/types/known/structpb"
)

// GRPC Interface for registered Nitric Document Plugin
type DocumentServer struct {
	pb.UnimplementedDocumentServiceServer
	// TODO: Support multiple plugin registrations
	// Just need to settle on a way of addressing them on calls
	documentPlugin sdk.DocumentService
}

func (s *DocumentServer) Set(ctx context.Context, req *pb.DocumentSetRequest) (*pb.DocumentSetResponse, error) {
	key := toSdkKey(req.Key)

	err := s.documentPlugin.Set(key, req.GetContent().AsMap())
	if err != nil {
		return nil, err
	}

	return &pb.DocumentSetResponse{}, nil
}

func (s *DocumentServer) Get(ctx context.Context, req *pb.DocumentGetRequest) (*pb.DocumentGetResponse, error) {
	key := toSdkKey(req.Key)

	doc, err := s.documentPlugin.Get(key)
	if err != nil {
		return nil, err
	}

	pbDoc, err := toPbDoc(doc)
	if err != nil {
		return nil, err
	}

	return &pb.DocumentGetResponse{
		Document: pbDoc,
	}, nil
}

func (s *DocumentServer) Delete(ctx context.Context, req *pb.DocumentDeleteRequest) (*pb.DocumentDeleteResponse, error) {
	key := toSdkKey(req.Key)

	err := s.documentPlugin.Delete(key)
	if err != nil {
		return nil, err
	}

	return &pb.DocumentDeleteResponse{}, nil
}

func (s *DocumentServer) Query(ctx context.Context, req *pb.DocumentQueryRequest) (*pb.DocumentQueryResponse, error) {
	collection := toSdkCollection(req.Collection)
	expressions := make([]sdk.QueryExpression, len(req.GetExpressions()))
	for i, exp := range req.GetExpressions() {
		expressions[i] = sdk.QueryExpression{
			Operand:  exp.GetOperand(),
			Operator: exp.GetOperator(),
			Value:    toExpValue(exp.GetValue()),
		}
	}
	limit := int(req.GetLimit())
	pagingMap := req.GetPagingToken()

	qr, err := s.documentPlugin.Query(collection, expressions, limit, pagingMap)
	if err != nil {
		return nil, err
	}

	pbDocuments := make([]*pb.Document, len(qr.Documents))
	for _, doc := range qr.Documents {
		pbDoc, err := toPbDoc(&doc)
		if err != nil {
			return nil, err
		}

		pbDocuments = append(pbDocuments, pbDoc)
	}

	return &pb.DocumentQueryResponse{
		Documents:   pbDocuments,
		PagingToken: qr.PagingToken,
	}, nil
}

func NewDocumentServer(docPlugin sdk.DocumentService) pb.DocumentServiceServer {
	return &DocumentServer{
		documentPlugin: docPlugin,
	}
}

func toPbDoc(doc *sdk.Document) (*pb.Document, error) {
	valStruct, err := structpb.NewStruct(doc.Content)
	if err != nil {
		return nil, err
	}

	return &pb.Document{
		Content: valStruct,
	}, nil
}

func toSdkKey(key *pb.Key) *sdk.Key {
	if key == nil {
		return nil
	}

	sdkKey := &sdk.Key{
		Collection: sdk.Collection{Name: key.Collection.Name},
		Id:         key.Id,
	}

	parentKey := key.Collection.Parent
	if parentKey != nil {
		sdkKey.Collection.Parent = &sdk.Key{
			Collection: sdk.Collection{Name: parentKey.Collection.Name},
			Id:         parentKey.Id,
		}
	}

	return sdkKey
}

func toSdkCollection(coll *pb.Collection) *sdk.Collection {
	if coll == nil {
		return nil
	}

	sdkCol := &sdk.Collection{Name: coll.Name}

	if coll.Parent != nil {
		sdkCol.Parent = &sdk.Key{
			Collection: sdk.Collection{Name: coll.Parent.Collection.Name},
			Id:         coll.Parent.Id,
		}
	}
	return sdkCol
}

func toExpValue(x *pb.ExpressionValue) interface{} {
	if x, ok := x.GetKind().(*pb.ExpressionValue_IntValue); ok {
		return x.IntValue
	}
	if x, ok := x.GetKind().(*pb.ExpressionValue_DoubleValue); ok {
		return x.DoubleValue
	}
	if x, ok := x.GetKind().(*pb.ExpressionValue_StringValue); ok {
		return x.StringValue
	}
	if x, ok := x.GetKind().(*pb.ExpressionValue_BoolValue); ok {
		return x.BoolValue
	}
	return nil
}
