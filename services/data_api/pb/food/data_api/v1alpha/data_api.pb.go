// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: data_api/food/data_api/v1alpha/data_api.proto

package data_api_v1alpha

import (
	v1alpha1 "HackDavisFood/services/inventory/pb/food/inventory/common/v1alpha"
	v1alpha "HackDavisFood/services/orders/pb/food/orders/common/v1alpha"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpsertOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *v1alpha.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpsertOrderRequest) Reset() {
	*x = UpsertOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertOrderRequest) ProtoMessage() {}

func (x *UpsertOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertOrderRequest.ProtoReflect.Descriptor instead.
func (*UpsertOrderRequest) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertOrderRequest) GetOrder() *v1alpha.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpsertOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsertOrderResponse) Reset() {
	*x = UpsertOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertOrderResponse) ProtoMessage() {}

func (x *UpsertOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertOrderResponse.ProtoReflect.Descriptor instead.
func (*UpsertOrderResponse) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{1}
}

type InsertProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *v1alpha1.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *InsertProductRequest) Reset() {
	*x = InsertProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertProductRequest) ProtoMessage() {}

func (x *InsertProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertProductRequest.ProtoReflect.Descriptor instead.
func (*InsertProductRequest) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{2}
}

func (x *InsertProductRequest) GetProduct() *v1alpha1.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type InsertProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertProductResponse) Reset() {
	*x = InsertProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertProductResponse) ProtoMessage() {}

func (x *InsertProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertProductResponse.ProtoReflect.Descriptor instead.
func (*InsertProductResponse) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{3}
}

type UpdateInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*v1alpha1.Item `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateInventoryRequest) Reset() {
	*x = UpdateInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryRequest) ProtoMessage() {}

func (x *UpdateInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateInventoryRequest) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInventoryRequest) GetItem() []*v1alpha1.Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInventoryResponse) Reset() {
	*x = UpdateInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryResponse) ProtoMessage() {}

func (x *UpdateInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateInventoryResponse) Descriptor() ([]byte, []int) {
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP(), []int{5}
}

var File_data_api_food_data_api_v1alpha_data_api_proto protoreflect.FileDescriptor

var file_data_api_food_data_api_v1alpha_data_api_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x30, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x66,
	0x6f, 0x6f, 0x64, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4d, 0x0a, 0x12, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x15, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x22, 0x17, 0x0a, 0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x19, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x48, 0x61, 0x63, 0x6b, 0x44,
	0x61, 0x76, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6f,
	0x6f, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_api_food_data_api_v1alpha_data_api_proto_rawDescOnce sync.Once
	file_data_api_food_data_api_v1alpha_data_api_proto_rawDescData = file_data_api_food_data_api_v1alpha_data_api_proto_rawDesc
)

func file_data_api_food_data_api_v1alpha_data_api_proto_rawDescGZIP() []byte {
	file_data_api_food_data_api_v1alpha_data_api_proto_rawDescOnce.Do(func() {
		file_data_api_food_data_api_v1alpha_data_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_api_food_data_api_v1alpha_data_api_proto_rawDescData)
	})
	return file_data_api_food_data_api_v1alpha_data_api_proto_rawDescData
}

var file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_data_api_food_data_api_v1alpha_data_api_proto_goTypes = []interface{}{
	(*UpsertOrderRequest)(nil),      // 0: food.data_api.v1alpha.UpsertOrderRequest
	(*UpsertOrderResponse)(nil),     // 1: food.data_api.v1alpha.UpsertOrderResponse
	(*InsertProductRequest)(nil),    // 2: food.data_api.v1alpha.InsertProductRequest
	(*InsertProductResponse)(nil),   // 3: food.data_api.v1alpha.InsertProductResponse
	(*UpdateInventoryRequest)(nil),  // 4: food.data_api.v1alpha.UpdateInventoryRequest
	(*UpdateInventoryResponse)(nil), // 5: food.data_api.v1alpha.UpdateInventoryResponse
	(*v1alpha.Order)(nil),           // 6: food.orders.common.v1alpha.Order
	(*v1alpha1.Product)(nil),        // 7: food.inventory.common.v1alpha.Product
	(*v1alpha1.Item)(nil),           // 8: food.inventory.common.v1alpha.Item
}
var file_data_api_food_data_api_v1alpha_data_api_proto_depIdxs = []int32{
	6, // 0: food.data_api.v1alpha.UpsertOrderRequest.order:type_name -> food.orders.common.v1alpha.Order
	7, // 1: food.data_api.v1alpha.InsertProductRequest.product:type_name -> food.inventory.common.v1alpha.Product
	8, // 2: food.data_api.v1alpha.UpdateInventoryRequest.item:type_name -> food.inventory.common.v1alpha.Item
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_data_api_food_data_api_v1alpha_data_api_proto_init() }
func file_data_api_food_data_api_v1alpha_data_api_proto_init() {
	if File_data_api_food_data_api_v1alpha_data_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertProductResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInventoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInventoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_api_food_data_api_v1alpha_data_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_api_food_data_api_v1alpha_data_api_proto_goTypes,
		DependencyIndexes: file_data_api_food_data_api_v1alpha_data_api_proto_depIdxs,
		MessageInfos:      file_data_api_food_data_api_v1alpha_data_api_proto_msgTypes,
	}.Build()
	File_data_api_food_data_api_v1alpha_data_api_proto = out.File
	file_data_api_food_data_api_v1alpha_data_api_proto_rawDesc = nil
	file_data_api_food_data_api_v1alpha_data_api_proto_goTypes = nil
	file_data_api_food_data_api_v1alpha_data_api_proto_depIdxs = nil
}
