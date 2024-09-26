// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: strawhouse-proto/_driver/transfer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DriverTransfer_FileUpload_FullMethodName       = "/proto.DriverTransfer/FileUpload"
	DriverTransfer_FileDownloadHash_FullMethodName = "/proto.DriverTransfer/FileDownloadHash"
)

// DriverTransferClient is the client API for DriverTransfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverTransferClient interface {
	FileUpload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FileDownloadHash(ctx context.Context, in *DownloadHashRequest, opts ...grpc.CallOption) (*DownloadHashResponse, error)
}

type driverTransferClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverTransferClient(cc grpc.ClientConnInterface) DriverTransferClient {
	return &driverTransferClient{cc}
}

func (c *driverTransferClient) FileUpload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverTransfer_FileUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverTransferClient) FileDownloadHash(ctx context.Context, in *DownloadHashRequest, opts ...grpc.CallOption) (*DownloadHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadHashResponse)
	err := c.cc.Invoke(ctx, DriverTransfer_FileDownloadHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverTransferServer is the server API for DriverTransfer service.
// All implementations must embed UnimplementedDriverTransferServer
// for forward compatibility.
type DriverTransferServer interface {
	FileUpload(context.Context, *UploadRequest) (*emptypb.Empty, error)
	FileDownloadHash(context.Context, *DownloadHashRequest) (*DownloadHashResponse, error)
	mustEmbedUnimplementedDriverTransferServer()
}

// UnimplementedDriverTransferServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverTransferServer struct{}

func (UnimplementedDriverTransferServer) FileUpload(context.Context, *UploadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUpload not implemented")
}
func (UnimplementedDriverTransferServer) FileDownloadHash(context.Context, *DownloadHashRequest) (*DownloadHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDownloadHash not implemented")
}
func (UnimplementedDriverTransferServer) mustEmbedUnimplementedDriverTransferServer() {}
func (UnimplementedDriverTransferServer) testEmbeddedByValue()                        {}

// UnsafeDriverTransferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverTransferServer will
// result in compilation errors.
type UnsafeDriverTransferServer interface {
	mustEmbedUnimplementedDriverTransferServer()
}

func RegisterDriverTransferServer(s grpc.ServiceRegistrar, srv DriverTransferServer) {
	// If the following call pancis, it indicates UnimplementedDriverTransferServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverTransfer_ServiceDesc, srv)
}

func _DriverTransfer_FileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTransferServer).FileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverTransfer_FileUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTransferServer).FileUpload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverTransfer_FileDownloadHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTransferServer).FileDownloadHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverTransfer_FileDownloadHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTransferServer).FileDownloadHash(ctx, req.(*DownloadHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverTransfer_ServiceDesc is the grpc.ServiceDesc for DriverTransfer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverTransfer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DriverTransfer",
	HandlerType: (*DriverTransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileUpload",
			Handler:    _DriverTransfer_FileUpload_Handler,
		},
		{
			MethodName: "FileDownloadHash",
			Handler:    _DriverTransfer_FileDownloadHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strawhouse-proto/_driver/transfer.proto",
}
