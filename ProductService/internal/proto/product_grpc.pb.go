// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: internal/proto/product.proto

package proto

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
	ProductService_GetProductById_FullMethodName   = "/product.ProductService/GetProductById"
	ProductService_AddProduct_FullMethodName       = "/product.ProductService/AddProduct"
	ProductService_GetProductsByIds_FullMethodName = "/product.ProductService/GetProductsByIds"
	ProductService_AddProducts_FullMethodName      = "/product.ProductService/AddProducts"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProductsByIds(ctx context.Context, in *ProductsIdsRequest, opts ...grpc.CallOption) (*ProductList, error)
	AddProducts(ctx context.Context, in *AddProductsRequest, opts ...grpc.CallOption) (*ProductList, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductById(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProductById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProductService_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductsByIds(ctx context.Context, in *ProductsIdsRequest, opts ...grpc.CallOption) (*ProductList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductList)
	err := c.cc.Invoke(ctx, ProductService_GetProductsByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProducts(ctx context.Context, in *AddProductsRequest, opts ...grpc.CallOption) (*ProductList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductList)
	err := c.cc.Invoke(ctx, ProductService_AddProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility.
type ProductServiceServer interface {
	GetProductById(context.Context, *ProductIdRequest) (*ProductResponse, error)
	AddProduct(context.Context, *AddProductRequest) (*ProductResponse, error)
	GetProductsByIds(context.Context, *ProductsIdsRequest) (*ProductList, error)
	AddProducts(context.Context, *AddProductsRequest) (*ProductList, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) GetProductById(context.Context, *ProductIdRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProductServiceServer) AddProduct(context.Context, *AddProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProductsByIds(context.Context, *ProductsIdsRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByIds not implemented")
}
func (UnimplementedProductServiceServer) AddProducts(context.Context, *AddProductsRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProducts not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}
func (UnimplementedProductServiceServer) testEmbeddedByValue()                        {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductById(ctx, req.(*ProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductsByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductsByIds(ctx, req.(*ProductsIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProducts(ctx, req.(*AddProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductById",
			Handler:    _ProductService_GetProductById_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "GetProductsByIds",
			Handler:    _ProductService_GetProductsByIds_Handler,
		},
		{
			MethodName: "AddProducts",
			Handler:    _ProductService_AddProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/product.proto",
}
