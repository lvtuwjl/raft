// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: request_vote_responses.proto

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

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        *uint64 `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	VoteGranted *bool   `protobuf:"varint,2,req,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_vote_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_vote_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_request_vote_responses_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteResponse) GetTerm() uint64 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil && x.VoteGranted != nil {
		return *x.VoteGranted
	}
	return false
}

var File_request_vote_responses_proto protoreflect.FileDescriptor

var file_request_vote_responses_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x6f, 0x74,
	0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0b,
	0x56, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x42, 0x39, 0xd8, 0xe1, 0x1e,
	0x00, 0xe0, 0xe1, 0x1e, 0x01, 0xf0, 0xe1, 0x1e, 0x01, 0xf8, 0xe1, 0x1e, 0x01, 0x80, 0xe2, 0x1e,
	0x01, 0xa8, 0xe2, 0x1e, 0x01, 0xb8, 0xe2, 0x1e, 0x01, 0xc0, 0xe2, 0x1e, 0x01, 0xc8, 0xe2, 0x1e,
	0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_request_vote_responses_proto_rawDescOnce sync.Once
	file_request_vote_responses_proto_rawDescData = file_request_vote_responses_proto_rawDesc
)

func file_request_vote_responses_proto_rawDescGZIP() []byte {
	file_request_vote_responses_proto_rawDescOnce.Do(func() {
		file_request_vote_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_vote_responses_proto_rawDescData)
	})
	return file_request_vote_responses_proto_rawDescData
}

var file_request_vote_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_vote_responses_proto_goTypes = []interface{}{
	(*RequestVoteResponse)(nil), // 0: protobuf.RequestVoteResponse
}
var file_request_vote_responses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_vote_responses_proto_init() }
func file_request_vote_responses_proto_init() {
	if File_request_vote_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_vote_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
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
			RawDescriptor: file_request_vote_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_vote_responses_proto_goTypes,
		DependencyIndexes: file_request_vote_responses_proto_depIdxs,
		MessageInfos:      file_request_vote_responses_proto_msgTypes,
	}.Build()
	File_request_vote_responses_proto = out.File
	file_request_vote_responses_proto_rawDesc = nil
	file_request_vote_responses_proto_goTypes = nil
	file_request_vote_responses_proto_depIdxs = nil
}
