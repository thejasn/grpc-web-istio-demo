// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EmojiServiceClient is the client API for EmojiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmojiServiceClient interface {
	InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error)
}

type emojiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmojiServiceClient(cc grpc.ClientConnInterface) EmojiServiceClient {
	return &emojiServiceClient{cc}
}

func (c *emojiServiceClient) InsertEmojis(ctx context.Context, in *EmojiRequest, opts ...grpc.CallOption) (*EmojiResponse, error) {
	out := new(EmojiResponse)
	err := c.cc.Invoke(ctx, "/proto.EmojiService/InsertEmojis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmojiServiceServer is the server API for EmojiService service.
// All implementations must embed UnimplementedEmojiServiceServer
// for forward compatibility
type EmojiServiceServer interface {
	InsertEmojis(context.Context, *EmojiRequest) (*EmojiResponse, error)
	mustEmbedUnimplementedEmojiServiceServer()
}

// UnimplementedEmojiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmojiServiceServer struct {
}

func (UnimplementedEmojiServiceServer) InsertEmojis(context.Context, *EmojiRequest) (*EmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmojis not implemented")
}
func (UnimplementedEmojiServiceServer) mustEmbedUnimplementedEmojiServiceServer() {}

// UnsafeEmojiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmojiServiceServer will
// result in compilation errors.
type UnsafeEmojiServiceServer interface {
	mustEmbedUnimplementedEmojiServiceServer()
}

func RegisterEmojiServiceServer(s *grpc.Server, srv EmojiServiceServer) {
	s.RegisterService(&_EmojiService_serviceDesc, srv)
}

func _EmojiService_InsertEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmojiService/InsertEmojis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).InsertEmojis(ctx, req.(*EmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmojiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmojiService",
	HandlerType: (*EmojiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertEmojis",
			Handler:    _EmojiService_InsertEmojis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emoji.proto",
}
