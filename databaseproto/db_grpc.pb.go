// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package databaseproto

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

// DatabaseAccessServiceClient is the client API for DatabaseAccessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseAccessServiceClient interface {
	//Stream response
	CourseRequest(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BigResponse, error)
}

type databaseAccessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseAccessServiceClient(cc grpc.ClientConnInterface) DatabaseAccessServiceClient {
	return &databaseAccessServiceClient{cc}
}

func (c *databaseAccessServiceClient) CourseRequest(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BigResponse, error) {
	out := new(BigResponse)
	err := c.cc.Invoke(ctx, "/mailer.DatabaseAccessService/CourseRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseAccessServiceServer is the server API for DatabaseAccessService service.
// All implementations must embed UnimplementedDatabaseAccessServiceServer
// for forward compatibility
type DatabaseAccessServiceServer interface {
	//Stream response
	CourseRequest(context.Context, *BookRequest) (*BigResponse, error)
	mustEmbedUnimplementedDatabaseAccessServiceServer()
}

// UnimplementedDatabaseAccessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseAccessServiceServer struct {
}

func (UnimplementedDatabaseAccessServiceServer) CourseRequest(context.Context, *BookRequest) (*BigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CourseRequest not implemented")
}
func (UnimplementedDatabaseAccessServiceServer) mustEmbedUnimplementedDatabaseAccessServiceServer() {}

// UnsafeDatabaseAccessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseAccessServiceServer will
// result in compilation errors.
type UnsafeDatabaseAccessServiceServer interface {
	mustEmbedUnimplementedDatabaseAccessServiceServer()
}

func RegisterDatabaseAccessServiceServer(s grpc.ServiceRegistrar, srv DatabaseAccessServiceServer) {
	s.RegisterService(&DatabaseAccessService_ServiceDesc, srv)
}

func _DatabaseAccessService_CourseRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseAccessServiceServer).CourseRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.DatabaseAccessService/CourseRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseAccessServiceServer).CourseRequest(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseAccessService_ServiceDesc is the grpc.ServiceDesc for DatabaseAccessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseAccessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mailer.DatabaseAccessService",
	HandlerType: (*DatabaseAccessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CourseRequest",
			Handler:    _DatabaseAccessService_CourseRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "databaseproto/db.proto",
}