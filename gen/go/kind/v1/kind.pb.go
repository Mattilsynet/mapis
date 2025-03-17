// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: kind/v1/kind.proto

package kindv1

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

type Kind int32

const (
	Kind_UnknownKind           Kind = 0
	Kind_IdportenClient        Kind = 1
	Kind_MaskinportenClient    Kind = 2
	Kind_Event                 Kind = 3
	Kind_AppRegistrationClient Kind = 4
	Kind_NatsAccount           Kind = 5
	Kind_DigdirScope           Kind = 6
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "UnknownKind",
		1: "IdportenClient",
		2: "MaskinportenClient",
		3: "Event",
		4: "AppRegistrationClient",
		5: "NatsAccount",
		6: "DigdirScope",
	}
	Kind_value = map[string]int32{
		"UnknownKind":           0,
		"IdportenClient":        1,
		"MaskinportenClient":    2,
		"Event":                 3,
		"AppRegistrationClient": 4,
		"NatsAccount":           5,
		"DigdirScope":           6,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_kind_v1_kind_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_kind_v1_kind_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_kind_v1_kind_proto_rawDescGZIP(), []int{0}
}

var File_kind_v1_kind_proto protoreflect.FileDescriptor

var file_kind_v1_kind_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0x8b, 0x01,
	0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x64, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x19,
	0x0a, 0x15, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x61, 0x74,
	0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x69,
	0x67, 0x64, 0x69, 0x72, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x10, 0x06, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x74, 0x74, 0x69, 0x6c,
	0x73, 0x79, 0x6e, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x69, 0x6e, 0x64, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kind_v1_kind_proto_rawDescOnce sync.Once
	file_kind_v1_kind_proto_rawDescData = file_kind_v1_kind_proto_rawDesc
)

func file_kind_v1_kind_proto_rawDescGZIP() []byte {
	file_kind_v1_kind_proto_rawDescOnce.Do(func() {
		file_kind_v1_kind_proto_rawDescData = protoimpl.X.CompressGZIP(file_kind_v1_kind_proto_rawDescData)
	})
	return file_kind_v1_kind_proto_rawDescData
}

var file_kind_v1_kind_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kind_v1_kind_proto_goTypes = []any{
	(Kind)(0), // 0: kind.v1.Kind
}
var file_kind_v1_kind_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kind_v1_kind_proto_init() }
func file_kind_v1_kind_proto_init() {
	if File_kind_v1_kind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kind_v1_kind_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kind_v1_kind_proto_goTypes,
		DependencyIndexes: file_kind_v1_kind_proto_depIdxs,
		EnumInfos:         file_kind_v1_kind_proto_enumTypes,
	}.Build()
	File_kind_v1_kind_proto = out.File
	file_kind_v1_kind_proto_rawDesc = nil
	file_kind_v1_kind_proto_goTypes = nil
	file_kind_v1_kind_proto_depIdxs = nil
}
