// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: example.proto

package conflictingenums

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

type RootEnum int32

const (
	RootEnum_UNKNOWN RootEnum = 0
	RootEnum_VAL1    RootEnum = 1
)

// Enum value maps for RootEnum.
var (
	RootEnum_name = map[int32]string{
		0: "UNKNOWN",
		1: "VAL1",
	}
	RootEnum_value = map[string]int32{
		"UNKNOWN": 0,
		"VAL1":    1,
	}
)

func (x RootEnum) Enum() *RootEnum {
	p := new(RootEnum)
	*p = x
	return p
}

func (x RootEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RootEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_enumTypes[0].Descriptor()
}

func (RootEnum) Type() protoreflect.EnumType {
	return &file_example_proto_enumTypes[0]
}

func (x RootEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RootEnum.Descriptor instead.
func (RootEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

type SampleMessage1_InnerEnum int32

const (
	SampleMessage1_UNKNOWN SampleMessage1_InnerEnum = 0
	SampleMessage1_VAL2    SampleMessage1_InnerEnum = 1
)

// Enum value maps for SampleMessage1_InnerEnum.
var (
	SampleMessage1_InnerEnum_name = map[int32]string{
		0: "UNKNOWN",
		1: "VAL2",
	}
	SampleMessage1_InnerEnum_value = map[string]int32{
		"UNKNOWN": 0,
		"VAL2":    1,
	}
)

func (x SampleMessage1_InnerEnum) Enum() *SampleMessage1_InnerEnum {
	p := new(SampleMessage1_InnerEnum)
	*p = x
	return p
}

func (x SampleMessage1_InnerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SampleMessage1_InnerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_enumTypes[1].Descriptor()
}

func (SampleMessage1_InnerEnum) Type() protoreflect.EnumType {
	return &file_example_proto_enumTypes[1]
}

func (x SampleMessage1_InnerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SampleMessage1_InnerEnum.Descriptor instead.
func (SampleMessage1_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0, 0}
}

type SampleMessage2_InnerEnum int32

const (
	SampleMessage2_UNKNOWN SampleMessage2_InnerEnum = 0
	SampleMessage2_VAL3    SampleMessage2_InnerEnum = 1
)

// Enum value maps for SampleMessage2_InnerEnum.
var (
	SampleMessage2_InnerEnum_name = map[int32]string{
		0: "UNKNOWN",
		1: "VAL3",
	}
	SampleMessage2_InnerEnum_value = map[string]int32{
		"UNKNOWN": 0,
		"VAL3":    1,
	}
)

func (x SampleMessage2_InnerEnum) Enum() *SampleMessage2_InnerEnum {
	p := new(SampleMessage2_InnerEnum)
	*p = x
	return p
}

func (x SampleMessage2_InnerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SampleMessage2_InnerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_proto_enumTypes[2].Descriptor()
}

func (SampleMessage2_InnerEnum) Type() protoreflect.EnumType {
	return &file_example_proto_enumTypes[2]
}

func (x SampleMessage2_InnerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SampleMessage2_InnerEnum.Descriptor instead.
func (SampleMessage2_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1, 0}
}

type SampleMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SampleMessage1) Reset() {
	*x = SampleMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleMessage1) ProtoMessage() {}

func (x *SampleMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleMessage1.ProtoReflect.Descriptor instead.
func (*SampleMessage1) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

type SampleMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SampleMessage2) Reset() {
	*x = SampleMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleMessage2) ProtoMessage() {}

func (x *SampleMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleMessage2.ProtoReflect.Descriptor instead.
func (*SampleMessage2) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x22, 0x22, 0x0a, 0x09, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x41, 0x4c, 0x32, 0x10, 0x01, 0x22, 0x34,
	0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x22, 0x22, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x41,
	0x4c, 0x33, 0x10, 0x01, 0x2a, 0x21, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x56, 0x41, 0x4c, 0x31, 0x10, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_goTypes = []interface{}{
	(RootEnum)(0),                 // 0: example.RootEnum
	(SampleMessage1_InnerEnum)(0), // 1: example.SampleMessage1.InnerEnum
	(SampleMessage2_InnerEnum)(0), // 2: example.SampleMessage2.InnerEnum
	(*SampleMessage1)(nil),        // 3: example.SampleMessage1
	(*SampleMessage2)(nil),        // 4: example.SampleMessage2
}
var file_example_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleMessage1); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleMessage2); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		EnumInfos:         file_example_proto_enumTypes,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
