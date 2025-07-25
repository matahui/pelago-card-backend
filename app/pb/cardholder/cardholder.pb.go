// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/cardholder.proto

package cardholder

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountReq) Reset() {
	*x = CreateAccountReq{}
	mi := &file_proto_cardholder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReq) ProtoMessage() {}

func (x *CreateAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cardholder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountReq.ProtoReflect.Descriptor instead.
func (*CreateAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_cardholder_proto_rawDescGZIP(), []int{0}
}

type CreateAccountResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResp) Reset() {
	*x = CreateAccountResp{}
	mi := &file_proto_cardholder_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResp) ProtoMessage() {}

func (x *CreateAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cardholder_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResp.ProtoReflect.Descriptor instead.
func (*CreateAccountResp) Descriptor() ([]byte, []int) {
	return file_proto_cardholder_proto_rawDescGZIP(), []int{1}
}

var File_proto_cardholder_proto protoreflect.FileDescriptor

const file_proto_cardholder_proto_rawDesc = "" +
	"\n" +
	"\x16proto/cardholder.proto\x12\x04card\"\x12\n" +
	"\x10CreateAccountReq\"\x13\n" +
	"\x11CreateAccountResp2U\n" +
	"\x11CardHolderService\x12@\n" +
	"\rCreateAccount\x12\x16.card.CreateAccountReq\x1a\x17.card.CreateAccountRespB\x0eZ\f./cardholderb\x06proto3"

var (
	file_proto_cardholder_proto_rawDescOnce sync.Once
	file_proto_cardholder_proto_rawDescData []byte
)

func file_proto_cardholder_proto_rawDescGZIP() []byte {
	file_proto_cardholder_proto_rawDescOnce.Do(func() {
		file_proto_cardholder_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cardholder_proto_rawDesc), len(file_proto_cardholder_proto_rawDesc)))
	})
	return file_proto_cardholder_proto_rawDescData
}

var file_proto_cardholder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_cardholder_proto_goTypes = []any{
	(*CreateAccountReq)(nil),  // 0: card.CreateAccountReq
	(*CreateAccountResp)(nil), // 1: card.CreateAccountResp
}
var file_proto_cardholder_proto_depIdxs = []int32{
	0, // 0: card.CardHolderService.CreateAccount:input_type -> card.CreateAccountReq
	1, // 1: card.CardHolderService.CreateAccount:output_type -> card.CreateAccountResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cardholder_proto_init() }
func file_proto_cardholder_proto_init() {
	if File_proto_cardholder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cardholder_proto_rawDesc), len(file_proto_cardholder_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cardholder_proto_goTypes,
		DependencyIndexes: file_proto_cardholder_proto_depIdxs,
		MessageInfos:      file_proto_cardholder_proto_msgTypes,
	}.Build()
	File_proto_cardholder_proto = out.File
	file_proto_cardholder_proto_goTypes = nil
	file_proto_cardholder_proto_depIdxs = nil
}
