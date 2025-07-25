// Code generated with goa v3.21.5, DO NOT EDIT.
//
// chatter protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/streaming/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: goagen_streaming_chatter.proto

package chatterpb

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

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{0}
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type EchoerStreamingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoerStreamingRequest) Reset() {
	*x = EchoerStreamingRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoerStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoerStreamingRequest) ProtoMessage() {}

func (x *EchoerStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoerStreamingRequest.ProtoReflect.Descriptor instead.
func (*EchoerStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{2}
}

func (x *EchoerStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type EchoerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoerResponse) Reset() {
	*x = EchoerResponse{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoerResponse) ProtoMessage() {}

func (x *EchoerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoerResponse.ProtoReflect.Descriptor instead.
func (*EchoerResponse) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{3}
}

func (x *EchoerResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ListenerStreamingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListenerStreamingRequest) Reset() {
	*x = ListenerStreamingRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListenerStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerStreamingRequest) ProtoMessage() {}

func (x *ListenerStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerStreamingRequest.ProtoReflect.Descriptor instead.
func (*ListenerStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{4}
}

func (x *ListenerStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ListenerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListenerResponse) Reset() {
	*x = ListenerResponse{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListenerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerResponse) ProtoMessage() {}

func (x *ListenerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerResponse.ProtoReflect.Descriptor instead.
func (*ListenerResponse) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{5}
}

type SummaryStreamingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SummaryStreamingRequest) Reset() {
	*x = SummaryStreamingRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryStreamingRequest) ProtoMessage() {}

func (x *SummaryStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryStreamingRequest.ProtoReflect.Descriptor instead.
func (*SummaryStreamingRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{6}
}

func (x *SummaryStreamingRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type ChatSummaryCollection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         []*ChatSummary         `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatSummaryCollection) Reset() {
	*x = ChatSummaryCollection{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatSummaryCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSummaryCollection) ProtoMessage() {}

func (x *ChatSummaryCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSummaryCollection.ProtoReflect.Descriptor instead.
func (*ChatSummaryCollection) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{7}
}

func (x *ChatSummaryCollection) GetField() []*ChatSummary {
	if x != nil {
		return x.Field
	}
	return nil
}

type ChatSummary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length *int32 `protobuf:"zigzag32,2,opt,name=length,proto3,oneof" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt        string `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatSummary) Reset() {
	*x = ChatSummary{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSummary) ProtoMessage() {}

func (x *ChatSummary) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSummary.ProtoReflect.Descriptor instead.
func (*ChatSummary) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{8}
}

func (x *ChatSummary) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *ChatSummary) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *ChatSummary) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{9}
}

type SubscribeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// Time at which the message was added
	AddedAt       string `protobuf:"bytes,3,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{10}
}

func (x *SubscribeResponse) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *SubscribeResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SubscribeResponse) GetAddedAt() string {
	if x != nil {
		return x.AddedAt
	}
	return ""
}

type HistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryRequest) Reset() {
	*x = HistoryRequest{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRequest) ProtoMessage() {}

func (x *HistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRequest.ProtoReflect.Descriptor instead.
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{11}
}

type HistoryResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Message sent to the server
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// Length of the message sent
	Length *int32 `protobuf:"zigzag32,2,opt,name=length,proto3,oneof" json:"length,omitempty"`
	// Time at which the message was sent
	SentAt        string `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	mi := &file_goagen_streaming_chatter_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_streaming_chatter_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_goagen_streaming_chatter_proto_rawDescGZIP(), []int{12}
}

func (x *HistoryResponse) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *HistoryResponse) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *HistoryResponse) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

var File_goagen_streaming_chatter_proto protoreflect.FileDescriptor

const file_goagen_streaming_chatter_proto_rawDesc = "" +
	"\n" +
	"\x1egoagen_streaming_chatter.proto\x12\achatter\"\x0e\n" +
	"\fLoginRequest\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\".\n" +
	"\x16EchoerStreamingRequest\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"&\n" +
	"\x0eEchoerResponse\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"0\n" +
	"\x18ListenerStreamingRequest\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"\x12\n" +
	"\x10ListenerResponse\"/\n" +
	"\x17SummaryStreamingRequest\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"C\n" +
	"\x15ChatSummaryCollection\x12*\n" +
	"\x05field\x18\x01 \x03(\v2\x14.chatter.ChatSummaryR\x05field\"i\n" +
	"\vChatSummary\x12\x19\n" +
	"\bmessage_\x18\x01 \x01(\tR\amessage\x12\x1b\n" +
	"\x06length\x18\x02 \x01(\x11H\x00R\x06length\x88\x01\x01\x12\x17\n" +
	"\asent_at\x18\x03 \x01(\tR\x06sentAtB\t\n" +
	"\a_length\"\x12\n" +
	"\x10SubscribeRequest\"a\n" +
	"\x11SubscribeResponse\x12\x19\n" +
	"\bmessage_\x18\x01 \x01(\tR\amessage\x12\x16\n" +
	"\x06action\x18\x02 \x01(\tR\x06action\x12\x19\n" +
	"\badded_at\x18\x03 \x01(\tR\aaddedAt\"\x10\n" +
	"\x0eHistoryRequest\"m\n" +
	"\x0fHistoryResponse\x12\x19\n" +
	"\bmessage_\x18\x01 \x01(\tR\amessage\x12\x1b\n" +
	"\x06length\x18\x02 \x01(\x11H\x00R\x06length\x88\x01\x01\x12\x17\n" +
	"\asent_at\x18\x03 \x01(\tR\x06sentAtB\t\n" +
	"\a_length2\xaa\x03\n" +
	"\aChatter\x126\n" +
	"\x05Login\x12\x15.chatter.LoginRequest\x1a\x16.chatter.LoginResponse\x12F\n" +
	"\x06Echoer\x12\x1f.chatter.EchoerStreamingRequest\x1a\x17.chatter.EchoerResponse(\x010\x01\x12J\n" +
	"\bListener\x12!.chatter.ListenerStreamingRequest\x1a\x19.chatter.ListenerResponse(\x01\x12M\n" +
	"\aSummary\x12 .chatter.SummaryStreamingRequest\x1a\x1e.chatter.ChatSummaryCollection(\x01\x12D\n" +
	"\tSubscribe\x12\x19.chatter.SubscribeRequest\x1a\x1a.chatter.SubscribeResponse0\x01\x12>\n" +
	"\aHistory\x12\x17.chatter.HistoryRequest\x1a\x18.chatter.HistoryResponse0\x01B\fZ\n" +
	"/chatterpbb\x06proto3"

var (
	file_goagen_streaming_chatter_proto_rawDescOnce sync.Once
	file_goagen_streaming_chatter_proto_rawDescData []byte
)

func file_goagen_streaming_chatter_proto_rawDescGZIP() []byte {
	file_goagen_streaming_chatter_proto_rawDescOnce.Do(func() {
		file_goagen_streaming_chatter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_goagen_streaming_chatter_proto_rawDesc), len(file_goagen_streaming_chatter_proto_rawDesc)))
	})
	return file_goagen_streaming_chatter_proto_rawDescData
}

var file_goagen_streaming_chatter_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_goagen_streaming_chatter_proto_goTypes = []any{
	(*LoginRequest)(nil),             // 0: chatter.LoginRequest
	(*LoginResponse)(nil),            // 1: chatter.LoginResponse
	(*EchoerStreamingRequest)(nil),   // 2: chatter.EchoerStreamingRequest
	(*EchoerResponse)(nil),           // 3: chatter.EchoerResponse
	(*ListenerStreamingRequest)(nil), // 4: chatter.ListenerStreamingRequest
	(*ListenerResponse)(nil),         // 5: chatter.ListenerResponse
	(*SummaryStreamingRequest)(nil),  // 6: chatter.SummaryStreamingRequest
	(*ChatSummaryCollection)(nil),    // 7: chatter.ChatSummaryCollection
	(*ChatSummary)(nil),              // 8: chatter.ChatSummary
	(*SubscribeRequest)(nil),         // 9: chatter.SubscribeRequest
	(*SubscribeResponse)(nil),        // 10: chatter.SubscribeResponse
	(*HistoryRequest)(nil),           // 11: chatter.HistoryRequest
	(*HistoryResponse)(nil),          // 12: chatter.HistoryResponse
}
var file_goagen_streaming_chatter_proto_depIdxs = []int32{
	8,  // 0: chatter.ChatSummaryCollection.field:type_name -> chatter.ChatSummary
	0,  // 1: chatter.Chatter.Login:input_type -> chatter.LoginRequest
	2,  // 2: chatter.Chatter.Echoer:input_type -> chatter.EchoerStreamingRequest
	4,  // 3: chatter.Chatter.Listener:input_type -> chatter.ListenerStreamingRequest
	6,  // 4: chatter.Chatter.Summary:input_type -> chatter.SummaryStreamingRequest
	9,  // 5: chatter.Chatter.Subscribe:input_type -> chatter.SubscribeRequest
	11, // 6: chatter.Chatter.History:input_type -> chatter.HistoryRequest
	1,  // 7: chatter.Chatter.Login:output_type -> chatter.LoginResponse
	3,  // 8: chatter.Chatter.Echoer:output_type -> chatter.EchoerResponse
	5,  // 9: chatter.Chatter.Listener:output_type -> chatter.ListenerResponse
	7,  // 10: chatter.Chatter.Summary:output_type -> chatter.ChatSummaryCollection
	10, // 11: chatter.Chatter.Subscribe:output_type -> chatter.SubscribeResponse
	12, // 12: chatter.Chatter.History:output_type -> chatter.HistoryResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_goagen_streaming_chatter_proto_init() }
func file_goagen_streaming_chatter_proto_init() {
	if File_goagen_streaming_chatter_proto != nil {
		return
	}
	file_goagen_streaming_chatter_proto_msgTypes[8].OneofWrappers = []any{}
	file_goagen_streaming_chatter_proto_msgTypes[12].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_goagen_streaming_chatter_proto_rawDesc), len(file_goagen_streaming_chatter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_streaming_chatter_proto_goTypes,
		DependencyIndexes: file_goagen_streaming_chatter_proto_depIdxs,
		MessageInfos:      file_goagen_streaming_chatter_proto_msgTypes,
	}.Build()
	File_goagen_streaming_chatter_proto = out.File
	file_goagen_streaming_chatter_proto_goTypes = nil
	file_goagen_streaming_chatter_proto_depIdxs = nil
}
