//
// Steward Strategy Provider API
//
// This proto defines the service/methods used by Strategy Providers to interact with Cellars through the Sommelier chain.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: steward.proto

package steward_proto

import (
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

//
// Represents a single function call on a particular Cellar
type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID (currently simply an Ethereum address) of the target Cellar
	CellarId string `protobuf:"bytes,1,opt,name=cellar_id,json=cellarId,proto3" json:"cellar_id,omitempty"`
	// The data from which the desired contract function will be encoded
	//
	// Types that are assignable to CallData:
	//	*SubmitRequest_AaveV2Stablecoin
	//	*SubmitRequest_CellarV1
	CallData isSubmitRequest_CallData `protobuf_oneof:"call_data"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitRequest) GetCellarId() string {
	if x != nil {
		return x.CellarId
	}
	return ""
}

func (m *SubmitRequest) GetCallData() isSubmitRequest_CallData {
	if m != nil {
		return m.CallData
	}
	return nil
}

func (x *SubmitRequest) GetAaveV2Stablecoin() *AaveV2Stablecoin {
	if x, ok := x.GetCallData().(*SubmitRequest_AaveV2Stablecoin); ok {
		return x.AaveV2Stablecoin
	}
	return nil
}

func (x *SubmitRequest) GetCellarV1() *CellarV1 {
	if x, ok := x.GetCallData().(*SubmitRequest_CellarV1); ok {
		return x.CellarV1
	}
	return nil
}

type isSubmitRequest_CallData interface {
	isSubmitRequest_CallData()
}

type SubmitRequest_AaveV2Stablecoin struct {
	AaveV2Stablecoin *AaveV2Stablecoin `protobuf:"bytes,2,opt,name=aave_v2_stablecoin,json=aaveV2Stablecoin,proto3,oneof"`
}

type SubmitRequest_CellarV1 struct {
	CellarV1 *CellarV1 `protobuf:"bytes,3,opt,name=cellar_v1,json=cellarV1,proto3,oneof"`
}

func (*SubmitRequest_AaveV2Stablecoin) isSubmitRequest_CallData() {}

func (*SubmitRequest_CellarV1) isSubmitRequest_CallData() {}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{1}
}

var File_steward_proto protoreflect.FileDescriptor

var file_steward_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x1a, 0x18, 0x61, 0x61, 0x76,
	0x65, 0x5f, 0x76, 0x32, 0x5f, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x6c,
	0x6c, 0x61, 0x72, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x12, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32,
	0x5f, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x48,
	0x00, 0x52, 0x10, 0x61, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63,
	0x6f, 0x69, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x5f, 0x76, 0x31,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x56, 0x31, 0x48, 0x00, 0x52, 0x08,
	0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x56, 0x31, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x51, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_steward_proto_rawDescOnce sync.Once
	file_steward_proto_rawDescData = file_steward_proto_rawDesc
)

func file_steward_proto_rawDescGZIP() []byte {
	file_steward_proto_rawDescOnce.Do(func() {
		file_steward_proto_rawDescData = protoimpl.X.CompressGZIP(file_steward_proto_rawDescData)
	})
	return file_steward_proto_rawDescData
}

var file_steward_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_steward_proto_goTypes = []interface{}{
	(*SubmitRequest)(nil),    // 0: steward.v2.SubmitRequest
	(*SubmitResponse)(nil),   // 1: steward.v2.SubmitResponse
	(*AaveV2Stablecoin)(nil), // 2: steward.v2.AaveV2Stablecoin
	(*CellarV1)(nil),         // 3: steward.v2.CellarV1
}
var file_steward_proto_depIdxs = []int32{
	2, // 0: steward.v2.SubmitRequest.aave_v2_stablecoin:type_name -> steward.v2.AaveV2Stablecoin
	3, // 1: steward.v2.SubmitRequest.cellar_v1:type_name -> steward.v2.CellarV1
	0, // 2: steward.v2.ContractCall.Submit:input_type -> steward.v2.SubmitRequest
	1, // 3: steward.v2.ContractCall.Submit:output_type -> steward.v2.SubmitResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steward_proto_init() }
func file_steward_proto_init() {
	if File_steward_proto != nil {
		return
	}
	file_aave_v2_stablecoin_proto_init()
	file_cellar_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_steward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
	file_steward_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SubmitRequest_AaveV2Stablecoin)(nil),
		(*SubmitRequest_CellarV1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steward_proto_goTypes,
		DependencyIndexes: file_steward_proto_depIdxs,
		MessageInfos:      file_steward_proto_msgTypes,
	}.Build()
	File_steward_proto = out.File
	file_steward_proto_rawDesc = nil
	file_steward_proto_goTypes = nil
	file_steward_proto_depIdxs = nil
}
