// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: data_api/food/data_api/v1alpha/data_api_service.proto

package data_api_v1alpha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_data_api_food_data_api_v1alpha_data_api_service_proto protoreflect.FileDescriptor

var file_data_api_food_data_api_v1alpha_data_api_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xda, 0x02,
	0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x66, 0x0a, 0x0b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x29, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2b, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4b, 0x5a, 0x49, 0x48, 0x61,
	0x63, 0x6b, 0x44, 0x61, 0x76, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62,
	0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_data_api_food_data_api_v1alpha_data_api_service_proto_goTypes = []interface{}{
	(*UpsertOrderRequest)(nil),      // 0: food.data_api.v1alpha.UpsertOrderRequest
	(*InsertProductRequest)(nil),    // 1: food.data_api.v1alpha.InsertProductRequest
	(*UpdateInventoryRequest)(nil),  // 2: food.data_api.v1alpha.UpdateInventoryRequest
	(*UpsertOrderResponse)(nil),     // 3: food.data_api.v1alpha.UpsertOrderResponse
	(*InsertProductResponse)(nil),   // 4: food.data_api.v1alpha.InsertProductResponse
	(*UpdateInventoryResponse)(nil), // 5: food.data_api.v1alpha.UpdateInventoryResponse
}
var file_data_api_food_data_api_v1alpha_data_api_service_proto_depIdxs = []int32{
	0, // 0: food.data_api.v1alpha.DataApiService.UpsertOrder:input_type -> food.data_api.v1alpha.UpsertOrderRequest
	1, // 1: food.data_api.v1alpha.DataApiService.InsertProduct:input_type -> food.data_api.v1alpha.InsertProductRequest
	2, // 2: food.data_api.v1alpha.DataApiService.UpdateInventory:input_type -> food.data_api.v1alpha.UpdateInventoryRequest
	3, // 3: food.data_api.v1alpha.DataApiService.UpsertOrder:output_type -> food.data_api.v1alpha.UpsertOrderResponse
	4, // 4: food.data_api.v1alpha.DataApiService.InsertProduct:output_type -> food.data_api.v1alpha.InsertProductResponse
	5, // 5: food.data_api.v1alpha.DataApiService.UpdateInventory:output_type -> food.data_api.v1alpha.UpdateInventoryResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_api_food_data_api_v1alpha_data_api_service_proto_init() }
func file_data_api_food_data_api_v1alpha_data_api_service_proto_init() {
	if File_data_api_food_data_api_v1alpha_data_api_service_proto != nil {
		return
	}
	file_data_api_food_data_api_v1alpha_data_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_api_food_data_api_v1alpha_data_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_api_food_data_api_v1alpha_data_api_service_proto_goTypes,
		DependencyIndexes: file_data_api_food_data_api_v1alpha_data_api_service_proto_depIdxs,
	}.Build()
	File_data_api_food_data_api_v1alpha_data_api_service_proto = out.File
	file_data_api_food_data_api_v1alpha_data_api_service_proto_rawDesc = nil
	file_data_api_food_data_api_v1alpha_data_api_service_proto_goTypes = nil
	file_data_api_food_data_api_v1alpha_data_api_service_proto_depIdxs = nil
}
