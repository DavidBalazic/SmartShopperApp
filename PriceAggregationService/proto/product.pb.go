// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/product.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	mi := &file_proto_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      float64                `protobuf:"fixed64,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Unit          string                 `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Store         string                 `protobuf:"bytes,7,opt,name=store,proto3" json:"store,omitempty"`
	PricePerUnit  float64                `protobuf:"fixed64,8,opt,name=pricePerUnit,proto3" json:"pricePerUnit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_proto_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Product) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *Product) GetPricePerUnit() float64 {
	if x != nil {
		return x.PricePerUnit
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_proto_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProductList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	mi := &file_proto_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductList) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type StoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Store         string                 `protobuf:"bytes,2,opt,name=store,proto3" json:"store,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	mi := &file_proto_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{4}
}

func (x *StoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

var File_proto_product_proto protoreflect.FileDescriptor

const file_proto_product_proto_rawDesc = "" +
	"\n" +
	"\x13proto/product.proto\x12\aproduct\"$\n" +
	"\x0eProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"\xcf\x01\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x1a\n" +
	"\bquantity\x18\x05 \x01(\x01R\bquantity\x12\x12\n" +
	"\x04unit\x18\x06 \x01(\tR\x04unit\x12\x14\n" +
	"\x05store\x18\a \x01(\tR\x05store\x12\"\n" +
	"\fpricePerUnit\x18\b \x01(\x01R\fpricePerUnit\"=\n" +
	"\x0fProductResponse\x12*\n" +
	"\aproduct\x18\x01 \x01(\v2\x10.product.ProductR\aproduct\";\n" +
	"\vProductList\x12,\n" +
	"\bproducts\x18\x01 \x03(\v2\x10.product.ProductR\bproducts\"8\n" +
	"\fStoreRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05store\x18\x02 \x01(\tR\x05store2\xdf\x01\n" +
	"\x0eProductService\x12G\n" +
	"\x12GetCheapestProduct\x12\x17.product.ProductRequest\x1a\x18.product.ProductResponse\x12=\n" +
	"\fGetAllPrices\x12\x17.product.ProductRequest\x1a\x14.product.ProductList\x12E\n" +
	"\x12GetCheapestByStore\x12\x15.product.StoreRequest\x1a\x18.product.ProductResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData []byte
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_product_proto_rawDesc), len(file_proto_product_proto_rawDesc)))
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_product_proto_goTypes = []any{
	(*ProductRequest)(nil),  // 0: product.ProductRequest
	(*Product)(nil),         // 1: product.Product
	(*ProductResponse)(nil), // 2: product.ProductResponse
	(*ProductList)(nil),     // 3: product.ProductList
	(*StoreRequest)(nil),    // 4: product.StoreRequest
}
var file_proto_product_proto_depIdxs = []int32{
	1, // 0: product.ProductResponse.product:type_name -> product.Product
	1, // 1: product.ProductList.products:type_name -> product.Product
	0, // 2: product.ProductService.GetCheapestProduct:input_type -> product.ProductRequest
	0, // 3: product.ProductService.GetAllPrices:input_type -> product.ProductRequest
	4, // 4: product.ProductService.GetCheapestByStore:input_type -> product.StoreRequest
	2, // 5: product.ProductService.GetCheapestProduct:output_type -> product.ProductResponse
	3, // 6: product.ProductService.GetAllPrices:output_type -> product.ProductList
	2, // 7: product.ProductService.GetCheapestByStore:output_type -> product.ProductResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_product_proto_rawDesc), len(file_proto_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
}
