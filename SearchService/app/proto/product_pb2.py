# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: product.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'product.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rproduct.proto\x12\x07product\"\x1e\n\x0eProductRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x8c\x01\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05price\x18\x04 \x01(\x01\x12\x10\n\x08quantity\x18\x05 \x01(\x01\x12\x0c\n\x04unit\x18\x06 \x01(\t\x12\r\n\x05store\x18\x07 \x01(\t\x12\x14\n\x0cpricePerUnit\x18\x08 \x01(\x01\"4\n\x0fProductResponse\x12!\n\x07product\x18\x01 \x01(\x0b\x32\x10.product.Product\"1\n\x0bProductList\x12\"\n\x08products\x18\x01 \x03(\x0b\x32\x10.product.Product\"+\n\x0cStoreRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05store\x18\x02 \x01(\t\"\x1e\n\x10ProductIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x8a\x01\n\x11\x41\x64\x64ProductRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x01\x12\x10\n\x08quantity\x18\x04 \x01(\x01\x12\x0c\n\x04unit\x18\x05 \x01(\t\x12\r\n\x05store\x18\x06 \x01(\t\x12\x14\n\x0cpricePerUnit\x18\x07 \x01(\x01\x32\xea\x02\n\x0eProductService\x12G\n\x12GetCheapestProduct\x12\x17.product.ProductRequest\x1a\x18.product.ProductResponse\x12=\n\x0cGetAllPrices\x12\x17.product.ProductRequest\x1a\x14.product.ProductList\x12\x45\n\x12GetCheapestByStore\x12\x15.product.StoreRequest\x1a\x18.product.ProductResponse\x12\x45\n\x0eGetProductById\x12\x19.product.ProductIdRequest\x1a\x18.product.ProductResponse\x12\x42\n\nAddProduct\x12\x1a.product.AddProductRequest\x1a\x18.product.ProductResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'product_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_PRODUCTREQUEST']._serialized_start=26
  _globals['_PRODUCTREQUEST']._serialized_end=56
  _globals['_PRODUCT']._serialized_start=59
  _globals['_PRODUCT']._serialized_end=199
  _globals['_PRODUCTRESPONSE']._serialized_start=201
  _globals['_PRODUCTRESPONSE']._serialized_end=253
  _globals['_PRODUCTLIST']._serialized_start=255
  _globals['_PRODUCTLIST']._serialized_end=304
  _globals['_STOREREQUEST']._serialized_start=306
  _globals['_STOREREQUEST']._serialized_end=349
  _globals['_PRODUCTIDREQUEST']._serialized_start=351
  _globals['_PRODUCTIDREQUEST']._serialized_end=381
  _globals['_ADDPRODUCTREQUEST']._serialized_start=384
  _globals['_ADDPRODUCTREQUEST']._serialized_end=522
  _globals['_PRODUCTSERVICE']._serialized_start=525
  _globals['_PRODUCTSERVICE']._serialized_end=887
# @@protoc_insertion_point(module_scope)
