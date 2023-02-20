// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/video_service.proto

package video_service

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

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	// Creates Video
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
	// Deletes Video
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error)
	// Creates Annotation
	CreateAnnotation(ctx context.Context, in *CreateAnnotationRequest, opts ...grpc.CallOption) (*CreateAnnotationResponse, error)
	// Returns all Video's Annotation
	GetAnnotations(ctx context.Context, in *GetAnnotationsRequest, opts ...grpc.CallOption) (*GetAnnotationsResponse, error)
	// Updates Annotation
	UpdateAnnotation(ctx context.Context, in *UpdateAnnotationRequest, opts ...grpc.CallOption) (*UpdateAnnotationResponse, error)
	// Deletes Annotation
	DeleteAnnotation(ctx context.Context, in *DeleteAnnotationRequest, opts ...grpc.CallOption) (*DeleteAnnotationResponse, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	out := new(CreateVideoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/CreateVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error) {
	out := new(DeleteVideoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/DeleteVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) CreateAnnotation(ctx context.Context, in *CreateAnnotationRequest, opts ...grpc.CallOption) (*CreateAnnotationResponse, error) {
	out := new(CreateAnnotationResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/CreateAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetAnnotations(ctx context.Context, in *GetAnnotationsRequest, opts ...grpc.CallOption) (*GetAnnotationsResponse, error) {
	out := new(GetAnnotationsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/GetAnnotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) UpdateAnnotation(ctx context.Context, in *UpdateAnnotationRequest, opts ...grpc.CallOption) (*UpdateAnnotationResponse, error) {
	out := new(UpdateAnnotationResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/UpdateAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) DeleteAnnotation(ctx context.Context, in *DeleteAnnotationRequest, opts ...grpc.CallOption) (*DeleteAnnotationResponse, error) {
	out := new(DeleteAnnotationResponse)
	err := c.cc.Invoke(ctx, "/api.v1.VideoService/DeleteAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations should embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	// Creates Video
	CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error)
	// Deletes Video
	DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error)
	// Creates Annotation
	CreateAnnotation(context.Context, *CreateAnnotationRequest) (*CreateAnnotationResponse, error)
	// Returns all Video's Annotation
	GetAnnotations(context.Context, *GetAnnotationsRequest) (*GetAnnotationsResponse, error)
	// Updates Annotation
	UpdateAnnotation(context.Context, *UpdateAnnotationRequest) (*UpdateAnnotationResponse, error)
	// Deletes Annotation
	DeleteAnnotation(context.Context, *DeleteAnnotationRequest) (*DeleteAnnotationResponse, error)
}

// UnimplementedVideoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedVideoServiceServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoServiceServer) CreateAnnotation(context.Context, *CreateAnnotationRequest) (*CreateAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnotation not implemented")
}
func (UnimplementedVideoServiceServer) GetAnnotations(context.Context, *GetAnnotationsRequest) (*GetAnnotationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnotations not implemented")
}
func (UnimplementedVideoServiceServer) UpdateAnnotation(context.Context, *UpdateAnnotationRequest) (*UpdateAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnnotation not implemented")
}
func (UnimplementedVideoServiceServer) DeleteAnnotation(context.Context, *DeleteAnnotationRequest) (*DeleteAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnotation not implemented")
}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/CreateVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/DeleteVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_CreateAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).CreateAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/CreateAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).CreateAnnotation(ctx, req.(*CreateAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetAnnotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetAnnotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/GetAnnotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetAnnotations(ctx, req.(*GetAnnotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_UpdateAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).UpdateAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/UpdateAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).UpdateAnnotation(ctx, req.(*UpdateAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_DeleteAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).DeleteAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.VideoService/DeleteAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).DeleteAnnotation(ctx, req.(*DeleteAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVideo",
			Handler:    _VideoService_CreateVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _VideoService_DeleteVideo_Handler,
		},
		{
			MethodName: "CreateAnnotation",
			Handler:    _VideoService_CreateAnnotation_Handler,
		},
		{
			MethodName: "GetAnnotations",
			Handler:    _VideoService_GetAnnotations_Handler,
		},
		{
			MethodName: "UpdateAnnotation",
			Handler:    _VideoService_UpdateAnnotation_Handler,
		},
		{
			MethodName: "DeleteAnnotation",
			Handler:    _VideoService_DeleteAnnotation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/video_service.proto",
}
