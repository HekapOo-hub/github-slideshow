// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package competition

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

// CompetitionServiceClient is the client API for CompetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompetitionServiceClient interface {
	CreateCompetition(ctx context.Context, in *Competition, opts ...grpc.CallOption) (*Response, error)
	GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	SetResult(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type competitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionServiceClient(cc grpc.ClientConnInterface) CompetitionServiceClient {
	return &competitionServiceClient{cc}
}

func (c *competitionServiceClient) CreateCompetition(ctx context.Context, in *Competition, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.competition.CompetitionService/CreateCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.competition.CompetitionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.competition.CompetitionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) SetResult(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.competition.CompetitionService/SetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompetitionServiceServer is the server API for CompetitionService service.
// All implementations must embed UnimplementedCompetitionServiceServer
// for forward compatibility
type CompetitionServiceServer interface {
	CreateCompetition(context.Context, *Competition) (*Response, error)
	GetById(context.Context, *Id) (*Response, error)
	GetAll(context.Context, *Empty) (*Response, error)
	SetResult(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCompetitionServiceServer()
}

// UnimplementedCompetitionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompetitionServiceServer struct {
}

func (UnimplementedCompetitionServiceServer) CreateCompetition(context.Context, *Competition) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) GetById(context.Context, *Id) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCompetitionServiceServer) GetAll(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCompetitionServiceServer) SetResult(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetResult not implemented")
}
func (UnimplementedCompetitionServiceServer) mustEmbedUnimplementedCompetitionServiceServer() {}

// UnsafeCompetitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompetitionServiceServer will
// result in compilation errors.
type UnsafeCompetitionServiceServer interface {
	mustEmbedUnimplementedCompetitionServiceServer()
}

func RegisterCompetitionServiceServer(s grpc.ServiceRegistrar, srv CompetitionServiceServer) {
	s.RegisterService(&CompetitionService_ServiceDesc, srv)
}

func _CompetitionService_CreateCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Competition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).CreateCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.competition.CompetitionService/CreateCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).CreateCompetition(ctx, req.(*Competition))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.competition.CompetitionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).GetById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.competition.CompetitionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_SetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).SetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.competition.CompetitionService/SetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).SetResult(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CompetitionService_ServiceDesc is the grpc.ServiceDesc for CompetitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompetitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.competition.CompetitionService",
	HandlerType: (*CompetitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompetition",
			Handler:    _CompetitionService_CreateCompetition_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _CompetitionService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CompetitionService_GetAll_Handler,
		},
		{
			MethodName: "SetResult",
			Handler:    _CompetitionService_SetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "competition.proto",
}
