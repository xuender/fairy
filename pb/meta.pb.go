// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/meta.proto

package pb

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

// Meta 模型.
type Meta int32

const (
	Meta_Unknown     Meta = 0
	Meta_Image       Meta = 1
	Meta_Video       Meta = 2
	Meta_Audio       Meta = 3
	Meta_Archive     Meta = 4
	Meta_Documents   Meta = 5
	Meta_Font        Meta = 6
	Meta_Application Meta = 7
	Meta_Java        Meta = 8
	Meta_Golang      Meta = 9
	Meta_JavaScript  Meta = 10
)

// Enum value maps for Meta.
var (
	Meta_name = map[int32]string{
		0:  "Unknown",
		1:  "Image",
		2:  "Video",
		3:  "Audio",
		4:  "Archive",
		5:  "Documents",
		6:  "Font",
		7:  "Application",
		8:  "Java",
		9:  "Golang",
		10: "JavaScript",
	}
	Meta_value = map[string]int32{
		"Unknown":     0,
		"Image":       1,
		"Video":       2,
		"Audio":       3,
		"Archive":     4,
		"Documents":   5,
		"Font":        6,
		"Application": 7,
		"Java":        8,
		"Golang":      9,
		"JavaScript":  10,
	}
)

func (x Meta) Enum() *Meta {
	p := new(Meta)
	*p = x
	return p
}

func (x Meta) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Meta) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_meta_proto_enumTypes[0].Descriptor()
}

func (Meta) Type() protoreflect.EnumType {
	return &file_pb_meta_proto_enumTypes[0]
}

func (x Meta) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Meta.Descriptor instead.
func (Meta) EnumDescriptor() ([]byte, []int) {
	return file_pb_meta_proto_rawDescGZIP(), []int{0}
}

var File_pb_meta_proto protoreflect.FileDescriptor

var file_pb_meta_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x2a, 0x91, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x6e, 0x74, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x07, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x61, 0x76, 0x61, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x47,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x61, 0x76, 0x61, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x10, 0x0a, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_meta_proto_rawDescOnce sync.Once
	file_pb_meta_proto_rawDescData = file_pb_meta_proto_rawDesc
)

func file_pb_meta_proto_rawDescGZIP() []byte {
	file_pb_meta_proto_rawDescOnce.Do(func() {
		file_pb_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_meta_proto_rawDescData)
	})
	return file_pb_meta_proto_rawDescData
}

var file_pb_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_meta_proto_goTypes = []interface{}{
	(Meta)(0), // 0: pb.Meta
}
var file_pb_meta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_meta_proto_init() }
func file_pb_meta_proto_init() {
	if File_pb_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_meta_proto_goTypes,
		DependencyIndexes: file_pb_meta_proto_depIdxs,
		EnumInfos:         file_pb_meta_proto_enumTypes,
	}.Build()
	File_pb_meta_proto = out.File
	file_pb_meta_proto_rawDesc = nil
	file_pb_meta_proto_goTypes = nil
	file_pb_meta_proto_depIdxs = nil
}
