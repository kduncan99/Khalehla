// Khalehla Project
// Copyright © 2023 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

// Describes the console service for Khalehla.
// The service is defined by a client (the operating system) and a server (the console).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: exec/msg/console.proto

package msg

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{0}
}

type ClearReadReplyMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId uint32 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *ClearReadReplyMessageRequest) Reset() {
	*x = ClearReadReplyMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearReadReplyMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearReadReplyMessageRequest) ProtoMessage() {}

func (x *ClearReadReplyMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearReadReplyMessageRequest.ProtoReflect.Descriptor instead.
func (*ClearReadReplyMessageRequest) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{1}
}

func (x *ClearReadReplyMessageRequest) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type SolicitedInputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId uint32 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SolicitedInputMessage) Reset() {
	*x = SolicitedInputMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitedInputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitedInputMessage) ProtoMessage() {}

func (x *SolicitedInputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitedInputMessage.ProtoReflect.Descriptor instead.
func (*SolicitedInputMessage) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{2}
}

func (x *SolicitedInputMessage) GetMessageId() uint32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *SolicitedInputMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UnsolicitedInputMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UnsolicitedInputMessage) Reset() {
	*x = UnsolicitedInputMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsolicitedInputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsolicitedInputMessage) ProtoMessage() {}

func (x *UnsolicitedInputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsolicitedInputMessage.ProtoReflect.Descriptor instead.
func (*UnsolicitedInputMessage) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{3}
}

func (x *UnsolicitedInputMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PollInputMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasInput bool `protobuf:"varint,1,opt,name=hasInput,proto3" json:"hasInput,omitempty"`
	// Types that are assignable to InputSpecifier:
	//
	//	*PollInputMessageResponse_SolicitedMessage
	//	*PollInputMessageResponse_UnsolicitedMessage
	InputSpecifier isPollInputMessageResponse_InputSpecifier `protobuf_oneof:"inputSpecifier"`
}

func (x *PollInputMessageResponse) Reset() {
	*x = PollInputMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollInputMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollInputMessageResponse) ProtoMessage() {}

func (x *PollInputMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollInputMessageResponse.ProtoReflect.Descriptor instead.
func (*PollInputMessageResponse) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{4}
}

func (x *PollInputMessageResponse) GetHasInput() bool {
	if x != nil {
		return x.HasInput
	}
	return false
}

func (m *PollInputMessageResponse) GetInputSpecifier() isPollInputMessageResponse_InputSpecifier {
	if m != nil {
		return m.InputSpecifier
	}
	return nil
}

func (x *PollInputMessageResponse) GetSolicitedMessage() *SolicitedInputMessage {
	if x, ok := x.GetInputSpecifier().(*PollInputMessageResponse_SolicitedMessage); ok {
		return x.SolicitedMessage
	}
	return nil
}

func (x *PollInputMessageResponse) GetUnsolicitedMessage() *UnsolicitedInputMessage {
	if x, ok := x.GetInputSpecifier().(*PollInputMessageResponse_UnsolicitedMessage); ok {
		return x.UnsolicitedMessage
	}
	return nil
}

type isPollInputMessageResponse_InputSpecifier interface {
	isPollInputMessageResponse_InputSpecifier()
}

type PollInputMessageResponse_SolicitedMessage struct {
	SolicitedMessage *SolicitedInputMessage `protobuf:"bytes,2,opt,name=solicitedMessage,proto3,oneof"`
}

type PollInputMessageResponse_UnsolicitedMessage struct {
	UnsolicitedMessage *UnsolicitedInputMessage `protobuf:"bytes,3,opt,name=unsolicitedMessage,proto3,oneof"`
}

func (*PollInputMessageResponse_SolicitedMessage) isPollInputMessageResponse_InputSpecifier() {}

func (*PollInputMessageResponse_UnsolicitedMessage) isPollInputMessageResponse_InputSpecifier() {}

type ReadOnlyMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Text   []string `protobuf:"bytes,4,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *ReadOnlyMessageRequest) Reset() {
	*x = ReadOnlyMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOnlyMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOnlyMessageRequest) ProtoMessage() {}

func (x *ReadOnlyMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOnlyMessageRequest.ProtoReflect.Descriptor instead.
func (*ReadOnlyMessageRequest) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{5}
}

func (x *ReadOnlyMessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ReadOnlyMessageRequest) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

type ReadReplyMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId          int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Sender             string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Text               []string `protobuf:"bytes,4,rep,name=text,proto3" json:"text,omitempty"`
	MaxReplyCharacters int32    `protobuf:"varint,5,opt,name=maxReplyCharacters,proto3" json:"maxReplyCharacters,omitempty"`
}

func (x *ReadReplyMessageRequest) Reset() {
	*x = ReadReplyMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReplyMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReplyMessageRequest) ProtoMessage() {}

func (x *ReadReplyMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReplyMessageRequest.ProtoReflect.Descriptor instead.
func (*ReadReplyMessageRequest) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{6}
}

func (x *ReadReplyMessageRequest) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReadReplyMessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ReadReplyMessageRequest) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *ReadReplyMessageRequest) GetMaxReplyCharacters() int32 {
	if x != nil {
		return x.MaxReplyCharacters
	}
	return 0
}

type StatusMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []string `protobuf:"bytes,4,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *StatusMessageRequest) Reset() {
	*x = StatusMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exec_msg_console_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusMessageRequest) ProtoMessage() {}

func (x *StatusMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exec_msg_console_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusMessageRequest.ProtoReflect.Descriptor instead.
func (*StatusMessageRequest) Descriptor() ([]byte, []int) {
	return file_exec_msg_console_proto_rawDescGZIP(), []int{7}
}

func (x *StatusMessageRequest) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_exec_msg_console_proto protoreflect.FileDescriptor

var file_exec_msg_console_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x1c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x2d, 0x0a, 0x17, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65, 0x64, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xe2,
	0x01, 0x0a, 0x18, 0x50, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68,
	0x61, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x4e, 0x0a, 0x12, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x12, 0x75,
	0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x2a, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xfb, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x15, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x21, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x10, 0x50, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x21, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x0a, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61,
	0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x4a, 0x0a, 0x2c, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x65, 0x61, 0x72, 0x73, 0x6e, 0x61, 0x6b, 0x65, 0x2e, 0x6b, 0x68, 0x61, 0x6c, 0x65,
	0x68, 0x6c, 0x61, 0x2e, 0x6b, 0x64, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x08, 0x65, 0x78, 0x65,
	0x63, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exec_msg_console_proto_rawDescOnce sync.Once
	file_exec_msg_console_proto_rawDescData = file_exec_msg_console_proto_rawDesc
)

func file_exec_msg_console_proto_rawDescGZIP() []byte {
	file_exec_msg_console_proto_rawDescOnce.Do(func() {
		file_exec_msg_console_proto_rawDescData = protoimpl.X.CompressGZIP(file_exec_msg_console_proto_rawDescData)
	})
	return file_exec_msg_console_proto_rawDescData
}

var file_exec_msg_console_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_exec_msg_console_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: msg.Empty
	(*ClearReadReplyMessageRequest)(nil), // 1: msg.ClearReadReplyMessageRequest
	(*SolicitedInputMessage)(nil),        // 2: msg.SolicitedInputMessage
	(*UnsolicitedInputMessage)(nil),      // 3: msg.UnsolicitedInputMessage
	(*PollInputMessageResponse)(nil),     // 4: msg.PollInputMessageResponse
	(*ReadOnlyMessageRequest)(nil),       // 5: msg.ReadOnlyMessageRequest
	(*ReadReplyMessageRequest)(nil),      // 6: msg.ReadReplyMessageRequest
	(*StatusMessageRequest)(nil),         // 7: msg.StatusMessageRequest
}
var file_exec_msg_console_proto_depIdxs = []int32{
	2, // 0: msg.PollInputMessageResponse.solicitedMessage:type_name -> msg.SolicitedInputMessage
	3, // 1: msg.PollInputMessageResponse.unsolicitedMessage:type_name -> msg.UnsolicitedInputMessage
	1, // 2: msg.Console.ClearReadReplyMessage:input_type -> msg.ClearReadReplyMessageRequest
	0, // 3: msg.Console.PollInputMessage:input_type -> msg.Empty
	0, // 4: msg.Console.Reset:input_type -> msg.Empty
	5, // 5: msg.Console.SendReadOnlyMessage:input_type -> msg.ReadOnlyMessageRequest
	6, // 6: msg.Console.SendReadReplyMessage:input_type -> msg.ReadReplyMessageRequest
	7, // 7: msg.Console.SendStatusMessage:input_type -> msg.StatusMessageRequest
	0, // 8: msg.Console.ClearReadReplyMessage:output_type -> msg.Empty
	4, // 9: msg.Console.PollInputMessage:output_type -> msg.PollInputMessageResponse
	0, // 10: msg.Console.Reset:output_type -> msg.Empty
	0, // 11: msg.Console.SendReadOnlyMessage:output_type -> msg.Empty
	0, // 12: msg.Console.SendReadReplyMessage:output_type -> msg.Empty
	0, // 13: msg.Console.SendStatusMessage:output_type -> msg.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exec_msg_console_proto_init() }
func file_exec_msg_console_proto_init() {
	if File_exec_msg_console_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exec_msg_console_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_exec_msg_console_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearReadReplyMessageRequest); i {
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
		file_exec_msg_console_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitedInputMessage); i {
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
		file_exec_msg_console_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsolicitedInputMessage); i {
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
		file_exec_msg_console_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollInputMessageResponse); i {
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
		file_exec_msg_console_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadOnlyMessageRequest); i {
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
		file_exec_msg_console_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReplyMessageRequest); i {
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
		file_exec_msg_console_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusMessageRequest); i {
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
	file_exec_msg_console_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PollInputMessageResponse_SolicitedMessage)(nil),
		(*PollInputMessageResponse_UnsolicitedMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exec_msg_console_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exec_msg_console_proto_goTypes,
		DependencyIndexes: file_exec_msg_console_proto_depIdxs,
		MessageInfos:      file_exec_msg_console_proto_msgTypes,
	}.Build()
	File_exec_msg_console_proto = out.File
	file_exec_msg_console_proto_rawDesc = nil
	file_exec_msg_console_proto_goTypes = nil
	file_exec_msg_console_proto_depIdxs = nil
}
