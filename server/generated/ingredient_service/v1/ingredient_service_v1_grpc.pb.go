// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ingredient_service/v1/ingredient_service_v1.proto

package ingredientpb

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
	IngredientService_CreateIngredient_FullMethodName = "/ingredient_service.v1.IngredientService/CreateIngredient"
	IngredientService_GetIngredient_FullMethodName    = "/ingredient_service.v1.IngredientService/GetIngredient"
	IngredientService_ListIngredients_FullMethodName  = "/ingredient_service.v1.IngredientService/ListIngredients"
	IngredientService_UpdateIngredient_FullMethodName = "/ingredient_service.v1.IngredientService/UpdateIngredient"
	IngredientService_DeleteIngredient_FullMethodName = "/ingredient_service.v1.IngredientService/DeleteIngredient"
)

// IngredientServiceClient is the client API for IngredientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngredientServiceClient interface {
	// Creates an ingredient
	CreateIngredient(ctx context.Context, in *CreateIngredientRequest, opts ...grpc.CallOption) (*CreateIngredientResponse, error)
	// Gets an ingredient
	GetIngredient(ctx context.Context, in *GetIngredientRequest, opts ...grpc.CallOption) (*GetIngredientResponse, error)
	// Lists ingredients
	ListIngredients(ctx context.Context, in *ListIngredientsRequest, opts ...grpc.CallOption) (*ListIngredientsResponse, error)
	// Updates an ingredient
	UpdateIngredient(ctx context.Context, in *UpdateIngredientRequest, opts ...grpc.CallOption) (*UpdateIngredientResponse, error)
	// Deletes an Ingredient
	DeleteIngredient(ctx context.Context, in *DeleteIngredientRequest, opts ...grpc.CallOption) (*DeleteIngredientResponse, error)
}

type ingredientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngredientServiceClient(cc grpc.ClientConnInterface) IngredientServiceClient {
	return &ingredientServiceClient{cc}
}

func (c *ingredientServiceClient) CreateIngredient(ctx context.Context, in *CreateIngredientRequest, opts ...grpc.CallOption) (*CreateIngredientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateIngredientResponse)
	err := c.cc.Invoke(ctx, IngredientService_CreateIngredient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingredientServiceClient) GetIngredient(ctx context.Context, in *GetIngredientRequest, opts ...grpc.CallOption) (*GetIngredientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIngredientResponse)
	err := c.cc.Invoke(ctx, IngredientService_GetIngredient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingredientServiceClient) ListIngredients(ctx context.Context, in *ListIngredientsRequest, opts ...grpc.CallOption) (*ListIngredientsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIngredientsResponse)
	err := c.cc.Invoke(ctx, IngredientService_ListIngredients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingredientServiceClient) UpdateIngredient(ctx context.Context, in *UpdateIngredientRequest, opts ...grpc.CallOption) (*UpdateIngredientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateIngredientResponse)
	err := c.cc.Invoke(ctx, IngredientService_UpdateIngredient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingredientServiceClient) DeleteIngredient(ctx context.Context, in *DeleteIngredientRequest, opts ...grpc.CallOption) (*DeleteIngredientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteIngredientResponse)
	err := c.cc.Invoke(ctx, IngredientService_DeleteIngredient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngredientServiceServer is the server API for IngredientService service.
// All implementations must embed UnimplementedIngredientServiceServer
// for forward compatibility.
type IngredientServiceServer interface {
	// Creates an ingredient
	CreateIngredient(context.Context, *CreateIngredientRequest) (*CreateIngredientResponse, error)
	// Gets an ingredient
	GetIngredient(context.Context, *GetIngredientRequest) (*GetIngredientResponse, error)
	// Lists ingredients
	ListIngredients(context.Context, *ListIngredientsRequest) (*ListIngredientsResponse, error)
	// Updates an ingredient
	UpdateIngredient(context.Context, *UpdateIngredientRequest) (*UpdateIngredientResponse, error)
	// Deletes an Ingredient
	DeleteIngredient(context.Context, *DeleteIngredientRequest) (*DeleteIngredientResponse, error)
	mustEmbedUnimplementedIngredientServiceServer()
}

// UnimplementedIngredientServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIngredientServiceServer struct{}

func (UnimplementedIngredientServiceServer) CreateIngredient(context.Context, *CreateIngredientRequest) (*CreateIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIngredient not implemented")
}
func (UnimplementedIngredientServiceServer) GetIngredient(context.Context, *GetIngredientRequest) (*GetIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIngredient not implemented")
}
func (UnimplementedIngredientServiceServer) ListIngredients(context.Context, *ListIngredientsRequest) (*ListIngredientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIngredients not implemented")
}
func (UnimplementedIngredientServiceServer) UpdateIngredient(context.Context, *UpdateIngredientRequest) (*UpdateIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngredient not implemented")
}
func (UnimplementedIngredientServiceServer) DeleteIngredient(context.Context, *DeleteIngredientRequest) (*DeleteIngredientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIngredient not implemented")
}
func (UnimplementedIngredientServiceServer) mustEmbedUnimplementedIngredientServiceServer() {}
func (UnimplementedIngredientServiceServer) testEmbeddedByValue()                           {}

// UnsafeIngredientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngredientServiceServer will
// result in compilation errors.
type UnsafeIngredientServiceServer interface {
	mustEmbedUnimplementedIngredientServiceServer()
}

func RegisterIngredientServiceServer(s grpc.ServiceRegistrar, srv IngredientServiceServer) {
	// If the following call pancis, it indicates UnimplementedIngredientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IngredientService_ServiceDesc, srv)
}

func _IngredientService_CreateIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientServiceServer).CreateIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngredientService_CreateIngredient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientServiceServer).CreateIngredient(ctx, req.(*CreateIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngredientService_GetIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientServiceServer).GetIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngredientService_GetIngredient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientServiceServer).GetIngredient(ctx, req.(*GetIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngredientService_ListIngredients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIngredientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientServiceServer).ListIngredients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngredientService_ListIngredients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientServiceServer).ListIngredients(ctx, req.(*ListIngredientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngredientService_UpdateIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientServiceServer).UpdateIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngredientService_UpdateIngredient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientServiceServer).UpdateIngredient(ctx, req.(*UpdateIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngredientService_DeleteIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientServiceServer).DeleteIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngredientService_DeleteIngredient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientServiceServer).DeleteIngredient(ctx, req.(*DeleteIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IngredientService_ServiceDesc is the grpc.ServiceDesc for IngredientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngredientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ingredient_service.v1.IngredientService",
	HandlerType: (*IngredientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIngredient",
			Handler:    _IngredientService_CreateIngredient_Handler,
		},
		{
			MethodName: "GetIngredient",
			Handler:    _IngredientService_GetIngredient_Handler,
		},
		{
			MethodName: "ListIngredients",
			Handler:    _IngredientService_ListIngredients_Handler,
		},
		{
			MethodName: "UpdateIngredient",
			Handler:    _IngredientService_UpdateIngredient_Handler,
		},
		{
			MethodName: "DeleteIngredient",
			Handler:    _IngredientService_DeleteIngredient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ingredient_service/v1/ingredient_service_v1.proto",
}
