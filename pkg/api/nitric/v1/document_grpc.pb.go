// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: document/v1/document.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	// Get an existing document
	Get(ctx context.Context, in *DocumentGetRequest, opts ...grpc.CallOption) (*DocumentGetResponse, error)
	// Create a new or overwrite an existing document
	Set(ctx context.Context, in *DocumentSetRequest, opts ...grpc.CallOption) (*DocumentSetResponse, error)
	// Delete an existing document
	Delete(ctx context.Context, in *DocumentDeleteRequest, opts ...grpc.CallOption) (*DocumentDeleteResponse, error)
	// Query the document collection (supports pagination)
	Query(ctx context.Context, in *DocumentQueryRequest, opts ...grpc.CallOption) (*DocumentQueryResponse, error)
	// Query the document collection (supports streaming)
	QueryStream(ctx context.Context, in *DocumentQueryStreamRequest, opts ...grpc.CallOption) (DocumentService_QueryStreamClient, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) Get(ctx context.Context, in *DocumentGetRequest, opts ...grpc.CallOption) (*DocumentGetResponse, error) {
	out := new(DocumentGetResponse)
	err := c.cc.Invoke(ctx, "/nitric.document.v1.DocumentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Set(ctx context.Context, in *DocumentSetRequest, opts ...grpc.CallOption) (*DocumentSetResponse, error) {
	out := new(DocumentSetResponse)
	err := c.cc.Invoke(ctx, "/nitric.document.v1.DocumentService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Delete(ctx context.Context, in *DocumentDeleteRequest, opts ...grpc.CallOption) (*DocumentDeleteResponse, error) {
	out := new(DocumentDeleteResponse)
	err := c.cc.Invoke(ctx, "/nitric.document.v1.DocumentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Query(ctx context.Context, in *DocumentQueryRequest, opts ...grpc.CallOption) (*DocumentQueryResponse, error) {
	out := new(DocumentQueryResponse)
	err := c.cc.Invoke(ctx, "/nitric.document.v1.DocumentService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) QueryStream(ctx context.Context, in *DocumentQueryStreamRequest, opts ...grpc.CallOption) (DocumentService_QueryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[0], "/nitric.document.v1.DocumentService/QueryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &documentServiceQueryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DocumentService_QueryStreamClient interface {
	Recv() (*DocumentQueryStreamResponse, error)
	grpc.ClientStream
}

type documentServiceQueryStreamClient struct {
	grpc.ClientStream
}

func (x *documentServiceQueryStreamClient) Recv() (*DocumentQueryStreamResponse, error) {
	m := new(DocumentQueryStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	// Get an existing document
	Get(context.Context, *DocumentGetRequest) (*DocumentGetResponse, error)
	// Create a new or overwrite an existing document
	Set(context.Context, *DocumentSetRequest) (*DocumentSetResponse, error)
	// Delete an existing document
	Delete(context.Context, *DocumentDeleteRequest) (*DocumentDeleteResponse, error)
	// Query the document collection (supports pagination)
	Query(context.Context, *DocumentQueryRequest) (*DocumentQueryResponse, error)
	// Query the document collection (supports streaming)
	QueryStream(*DocumentQueryStreamRequest, DocumentService_QueryStreamServer) error
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) Get(context.Context, *DocumentGetRequest) (*DocumentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDocumentServiceServer) Set(context.Context, *DocumentSetRequest) (*DocumentSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedDocumentServiceServer) Delete(context.Context, *DocumentDeleteRequest) (*DocumentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDocumentServiceServer) Query(context.Context, *DocumentQueryRequest) (*DocumentQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDocumentServiceServer) QueryStream(*DocumentQueryStreamRequest, DocumentService_QueryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryStream not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.document.v1.DocumentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Get(ctx, req.(*DocumentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.document.v1.DocumentService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Set(ctx, req.(*DocumentSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.document.v1.DocumentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Delete(ctx, req.(*DocumentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.document.v1.DocumentService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Query(ctx, req.(*DocumentQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_QueryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DocumentQueryStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServiceServer).QueryStream(m, &documentServiceQueryStreamServer{stream})
}

type DocumentService_QueryStreamServer interface {
	Send(*DocumentQueryStreamResponse) error
	grpc.ServerStream
}

type documentServiceQueryStreamServer struct {
	grpc.ServerStream
}

func (x *documentServiceQueryStreamServer) Send(m *DocumentQueryStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nitric.document.v1.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DocumentService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _DocumentService_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DocumentService_Delete_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _DocumentService_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryStream",
			Handler:       _DocumentService_QueryStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "document/v1/document.proto",
}
