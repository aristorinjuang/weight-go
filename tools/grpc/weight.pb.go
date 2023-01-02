// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: weight.proto

package grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeightParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Max  float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Min  float64 `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
}

func (x *WeightParams) Reset() {
	*x = WeightParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightParams) ProtoMessage() {}

func (x *WeightParams) ProtoReflect() protoreflect.Message {
	mi := &file_weight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightParams.ProtoReflect.Descriptor instead.
func (*WeightParams) Descriptor() ([]byte, []int) {
	return file_weight_proto_rawDescGZIP(), []int{0}
}

func (x *WeightParams) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *WeightParams) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *WeightParams) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

type Weight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Max        float64                `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Min        float64                `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	Difference float64                `protobuf:"fixed64,4,opt,name=difference,proto3" json:"difference,omitempty"`
}

func (x *Weight) Reset() {
	*x = Weight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weight) ProtoMessage() {}

func (x *Weight) ProtoReflect() protoreflect.Message {
	mi := &file_weight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weight.ProtoReflect.Descriptor instead.
func (*Weight) Descriptor() ([]byte, []int) {
	return file_weight_proto_rawDescGZIP(), []int{1}
}

func (x *Weight) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Weight) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Weight) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Weight) GetDifference() float64 {
	if x != nil {
		return x.Difference
	}
	return 0
}

type Weights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weights           []*Weight `protobuf:"bytes,1,rep,name=weights,proto3" json:"weights,omitempty"`
	AverageMax        float64   `protobuf:"fixed64,2,opt,name=average_max,json=averageMax,proto3" json:"average_max,omitempty"`
	AverageMin        float64   `protobuf:"fixed64,3,opt,name=average_min,json=averageMin,proto3" json:"average_min,omitempty"`
	AverageDifference float64   `protobuf:"fixed64,4,opt,name=average_difference,json=averageDifference,proto3" json:"average_difference,omitempty"`
}

func (x *Weights) Reset() {
	*x = Weights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weights) ProtoMessage() {}

func (x *Weights) ProtoReflect() protoreflect.Message {
	mi := &file_weight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weights.ProtoReflect.Descriptor instead.
func (*Weights) Descriptor() ([]byte, []int) {
	return file_weight_proto_rawDescGZIP(), []int{2}
}

func (x *Weights) GetWeights() []*Weight {
	if x != nil {
		return x.Weights
	}
	return nil
}

func (x *Weights) GetAverageMax() float64 {
	if x != nil {
		return x.AverageMax
	}
	return 0
}

func (x *Weights) GetAverageMin() float64 {
	if x != nil {
		return x.AverageMin
	}
	return 0
}

func (x *Weights) GetAverageDifference() float64 {
	if x != nil {
		return x.AverageDifference
	}
	return 0
}

var File_weight_proto protoreflect.FileDescriptor

var file_weight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x22, 0x7c, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x9d, 0x01, 0x0a, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x07,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x78,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x69,
	0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x32, 0x8f, 0x03, 0x0a, 0x0d, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0d, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0d, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x07, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x12, 0x54, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0d, 0x2e, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x12,
	0x51, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x0d, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74,
	0x65, 0x7d, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x6a, 0x75, 0x61, 0x6e, 0x67, 0x2f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weight_proto_rawDescOnce sync.Once
	file_weight_proto_rawDescData = file_weight_proto_rawDesc
)

func file_weight_proto_rawDescGZIP() []byte {
	file_weight_proto_rawDescOnce.Do(func() {
		file_weight_proto_rawDescData = protoimpl.X.CompressGZIP(file_weight_proto_rawDescData)
	})
	return file_weight_proto_rawDescData
}

var file_weight_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_weight_proto_goTypes = []interface{}{
	(*WeightParams)(nil),          // 0: WeightParams
	(*Weight)(nil),                // 1: Weight
	(*Weights)(nil),               // 2: Weights
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_weight_proto_depIdxs = []int32{
	3, // 0: Weight.date:type_name -> google.protobuf.Timestamp
	1, // 1: Weights.weights:type_name -> Weight
	4, // 2: WeightService.ListWeights:input_type -> google.protobuf.Empty
	0, // 3: WeightService.CreateWeight:input_type -> WeightParams
	0, // 4: WeightService.ReadWeight:input_type -> WeightParams
	0, // 5: WeightService.UpdateWeight:input_type -> WeightParams
	0, // 6: WeightService.DeleteWeight:input_type -> WeightParams
	2, // 7: WeightService.ListWeights:output_type -> Weights
	4, // 8: WeightService.CreateWeight:output_type -> google.protobuf.Empty
	1, // 9: WeightService.ReadWeight:output_type -> Weight
	4, // 10: WeightService.UpdateWeight:output_type -> google.protobuf.Empty
	4, // 11: WeightService.DeleteWeight:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weight_proto_init() }
func file_weight_proto_init() {
	if File_weight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightParams); i {
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
		file_weight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weight); i {
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
		file_weight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weights); i {
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
			RawDescriptor: file_weight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weight_proto_goTypes,
		DependencyIndexes: file_weight_proto_depIdxs,
		MessageInfos:      file_weight_proto_msgTypes,
	}.Build()
	File_weight_proto = out.File
	file_weight_proto_rawDesc = nil
	file_weight_proto_goTypes = nil
	file_weight_proto_depIdxs = nil
}