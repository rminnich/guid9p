// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: proto/ninep.proto

package ninep

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NinePService_Walk_FullMethodName   = "/ninep.NinePService/Walk"
	NinePService_Open_FullMethodName   = "/ninep.NinePService/Open"
	NinePService_Read_FullMethodName   = "/ninep.NinePService/Read"
	NinePService_Write_FullMethodName  = "/ninep.NinePService/Write"
	NinePService_Create_FullMethodName = "/ninep.NinePService/Create"
	NinePService_Stat_FullMethodName   = "/ninep.NinePService/Stat"
	NinePService_Clunk_FullMethodName  = "/ninep.NinePService/Clunk"
	NinePService_Flush_FullMethodName  = "/ninep.NinePService/Flush"
)

// NinePServiceClient is the client API for NinePService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the gRPC service.
type NinePServiceClient interface {
	Walk(ctx context.Context, in *WalkRequest, opts ...grpc.CallOption) (*WalkResponse, error)
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	Clunk(ctx context.Context, in *ClunkRequest, opts ...grpc.CallOption) (*ClunkResponse, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type ninePServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNinePServiceClient(cc grpc.ClientConnInterface) NinePServiceClient {
	return &ninePServiceClient{cc}
}

func (c *ninePServiceClient) Walk(ctx context.Context, in *WalkRequest, opts ...grpc.CallOption) (*WalkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalkResponse)
	err := c.cc.Invoke(ctx, NinePService_Walk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, NinePService_Open_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, NinePService_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, NinePService_Write_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, NinePService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NinePService_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Clunk(ctx context.Context, in *ClunkRequest, opts ...grpc.CallOption) (*ClunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClunkResponse)
	err := c.cc.Invoke(ctx, NinePService_Clunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ninePServiceClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, NinePService_Flush_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NinePServiceServer is the server API for NinePService service.
// All implementations must embed UnimplementedNinePServiceServer
// for forward compatibility.
//
// Define the gRPC service.
type NinePServiceServer interface {
	Walk(context.Context, *WalkRequest) (*WalkResponse, error)
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Stat(context.Context, *StatRequest) (*StatResponse, error)
	Clunk(context.Context, *ClunkRequest) (*ClunkResponse, error)
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	mustEmbedUnimplementedNinePServiceServer()
}

// UnimplementedNinePServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNinePServiceServer struct{}

func (UnimplementedNinePServiceServer) Walk(context.Context, *WalkRequest) (*WalkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Walk not implemented")
}
func (UnimplementedNinePServiceServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedNinePServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedNinePServiceServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedNinePServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNinePServiceServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedNinePServiceServer) Clunk(context.Context, *ClunkRequest) (*ClunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clunk not implemented")
}
func (UnimplementedNinePServiceServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedNinePServiceServer) mustEmbedUnimplementedNinePServiceServer() {}
func (UnimplementedNinePServiceServer) testEmbeddedByValue()                      {}

// UnsafeNinePServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NinePServiceServer will
// result in compilation errors.
type UnsafeNinePServiceServer interface {
	mustEmbedUnimplementedNinePServiceServer()
}

func RegisterNinePServiceServer(s grpc.ServiceRegistrar, srv NinePServiceServer) {
	// If the following call pancis, it indicates UnimplementedNinePServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NinePService_ServiceDesc, srv)
}

func _NinePService_Walk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Walk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Walk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Walk(ctx, req.(*WalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Open_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Clunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Clunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Clunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Clunk(ctx, req.(*ClunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NinePService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NinePServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NinePService_Flush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NinePServiceServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NinePService_ServiceDesc is the grpc.ServiceDesc for NinePService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NinePService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ninep.NinePService",
	HandlerType: (*NinePServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Walk",
			Handler:    _NinePService_Walk_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _NinePService_Open_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _NinePService_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _NinePService_Write_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NinePService_Create_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _NinePService_Stat_Handler,
		},
		{
			MethodName: "Clunk",
			Handler:    _NinePService_Clunk_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _NinePService_Flush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ninep.proto",
}