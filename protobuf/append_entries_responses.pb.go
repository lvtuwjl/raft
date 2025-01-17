// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: append_entries_responses.proto

package protobuf

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

type AppendEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        *uint64 `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	Index       *uint64 `protobuf:"varint,2,req,name=Index" json:"Index,omitempty"`
	CommitIndex *uint64 `protobuf:"varint,3,req,name=CommitIndex" json:"CommitIndex,omitempty"`
	Success     *bool   `protobuf:"varint,4,req,name=Success" json:"Success,omitempty"`
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_append_entries_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_append_entries_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_append_entries_responses_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEntriesResponse) GetTerm() uint64 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetIndex() uint64 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *AppendEntriesResponse) GetCommitIndex() uint64 {
	if x != nil && x.CommitIndex != nil {
		return *x.CommitIndex
	}
	return 0
}

func (x *AppendEntriesResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

var File_append_entries_responses_proto protoreflect.FileDescriptor

var file_append_entries_responses_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7d, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x39, 0xd8, 0xe1, 0x1e, 0x00, 0xe0, 0xe1, 0x1e, 0x01, 0xf0, 0xe1, 0x1e, 0x01, 0xf8, 0xe1, 0x1e,
	0x01, 0x80, 0xe2, 0x1e, 0x01, 0xa8, 0xe2, 0x1e, 0x01, 0xb8, 0xe2, 0x1e, 0x01, 0xc0, 0xe2, 0x1e,
	0x01, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0x5a, 0x0b, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_append_entries_responses_proto_rawDescOnce sync.Once
	file_append_entries_responses_proto_rawDescData = file_append_entries_responses_proto_rawDesc
)

func file_append_entries_responses_proto_rawDescGZIP() []byte {
	file_append_entries_responses_proto_rawDescOnce.Do(func() {
		file_append_entries_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_append_entries_responses_proto_rawDescData)
	})
	return file_append_entries_responses_proto_rawDescData
}

var file_append_entries_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_append_entries_responses_proto_goTypes = []interface{}{
	(*AppendEntriesResponse)(nil), // 0: protobuf.AppendEntriesResponse
}
var file_append_entries_responses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_append_entries_responses_proto_init() }
func file_append_entries_responses_proto_init() {
	if File_append_entries_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_append_entries_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesResponse); i {
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
			RawDescriptor: file_append_entries_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_append_entries_responses_proto_goTypes,
		DependencyIndexes: file_append_entries_responses_proto_depIdxs,
		MessageInfos:      file_append_entries_responses_proto_msgTypes,
	}.Build()
	File_append_entries_responses_proto = out.File
	file_append_entries_responses_proto_rawDesc = nil
	file_append_entries_responses_proto_goTypes = nil
	file_append_entries_responses_proto_depIdxs = nil
}
