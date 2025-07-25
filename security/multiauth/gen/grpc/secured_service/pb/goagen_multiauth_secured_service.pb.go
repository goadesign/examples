// Code generated with goa v3.21.5, DO NOT EDIT.
//
// secured_service protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: goagen_multiauth_secured_service.proto

package secured_servicepb

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

type SigninRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{0}
}

type SigninResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// JWT token
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// API Key
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// OAuth2 token
	OauthToken    string `protobuf:"bytes,3,opt,name=oauth_token,json=oauthToken,proto3" json:"oauth_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{1}
}

func (x *SigninResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *SigninResponse) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *SigninResponse) GetOauthToken() string {
	if x != nil {
		return x.OauthToken
	}
	return ""
}

type SecureRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Whether to force auth failure even with a valid JWT
	Fail          *bool `protobuf:"varint,1,opt,name=fail,proto3,oneof" json:"fail,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecureRequest) Reset() {
	*x = SecureRequest{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureRequest) ProtoMessage() {}

func (x *SecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureRequest.ProtoReflect.Descriptor instead.
func (*SecureRequest) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{2}
}

func (x *SecureRequest) GetFail() bool {
	if x != nil && x.Fail != nil {
		return *x.Fail
	}
	return false
}

type SecureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecureResponse) Reset() {
	*x = SecureResponse{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureResponse) ProtoMessage() {}

func (x *SecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureResponse.ProtoReflect.Descriptor instead.
func (*SecureResponse) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{3}
}

func (x *SecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type DoublySecureRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API key
	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoublySecureRequest) Reset() {
	*x = DoublySecureRequest{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoublySecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoublySecureRequest) ProtoMessage() {}

func (x *DoublySecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoublySecureRequest.ProtoReflect.Descriptor instead.
func (*DoublySecureRequest) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{4}
}

func (x *DoublySecureRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DoublySecureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoublySecureResponse) Reset() {
	*x = DoublySecureResponse{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoublySecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoublySecureResponse) ProtoMessage() {}

func (x *DoublySecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoublySecureResponse.ProtoReflect.Descriptor instead.
func (*DoublySecureResponse) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{5}
}

func (x *DoublySecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type AlsoDoublySecureRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Username used to perform signin
	Username *string `protobuf:"bytes,1,opt,name=username,proto3,oneof" json:"username,omitempty"`
	// Password used to perform signin
	Password *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
	// API key
	Key           *string `protobuf:"bytes,3,opt,name=key,proto3,oneof" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlsoDoublySecureRequest) Reset() {
	*x = AlsoDoublySecureRequest{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlsoDoublySecureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlsoDoublySecureRequest) ProtoMessage() {}

func (x *AlsoDoublySecureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlsoDoublySecureRequest.ProtoReflect.Descriptor instead.
func (*AlsoDoublySecureRequest) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{6}
}

func (x *AlsoDoublySecureRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *AlsoDoublySecureRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *AlsoDoublySecureRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

type AlsoDoublySecureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlsoDoublySecureResponse) Reset() {
	*x = AlsoDoublySecureResponse{}
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlsoDoublySecureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlsoDoublySecureResponse) ProtoMessage() {}

func (x *AlsoDoublySecureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goagen_multiauth_secured_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlsoDoublySecureResponse.ProtoReflect.Descriptor instead.
func (*AlsoDoublySecureResponse) Descriptor() ([]byte, []int) {
	return file_goagen_multiauth_secured_service_proto_rawDescGZIP(), []int{7}
}

func (x *AlsoDoublySecureResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_goagen_multiauth_secured_service_proto protoreflect.FileDescriptor

const file_goagen_multiauth_secured_service_proto_rawDesc = "" +
	"\n" +
	"&goagen_multiauth_secured_service.proto\x12\x0fsecured_service\"\x0f\n" +
	"\rSigninRequest\"\\\n" +
	"\x0eSigninResponse\x12\x10\n" +
	"\x03jwt\x18\x01 \x01(\tR\x03jwt\x12\x17\n" +
	"\aapi_key\x18\x02 \x01(\tR\x06apiKey\x12\x1f\n" +
	"\voauth_token\x18\x03 \x01(\tR\n" +
	"oauthToken\"1\n" +
	"\rSecureRequest\x12\x17\n" +
	"\x04fail\x18\x01 \x01(\bH\x00R\x04fail\x88\x01\x01B\a\n" +
	"\x05_fail\"&\n" +
	"\x0eSecureResponse\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"'\n" +
	"\x13DoublySecureRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\",\n" +
	"\x14DoublySecureResponse\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\"\x94\x01\n" +
	"\x17AlsoDoublySecureRequest\x12\x1f\n" +
	"\busername\x18\x01 \x01(\tH\x00R\busername\x88\x01\x01\x12\x1f\n" +
	"\bpassword\x18\x02 \x01(\tH\x01R\bpassword\x88\x01\x01\x12\x15\n" +
	"\x03key\x18\x03 \x01(\tH\x02R\x03key\x88\x01\x01B\v\n" +
	"\t_usernameB\v\n" +
	"\t_passwordB\x06\n" +
	"\x04_key\"0\n" +
	"\x18AlsoDoublySecureResponse\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field2\xec\x02\n" +
	"\x0eSecuredService\x12I\n" +
	"\x06Signin\x12\x1e.secured_service.SigninRequest\x1a\x1f.secured_service.SigninResponse\x12I\n" +
	"\x06Secure\x12\x1e.secured_service.SecureRequest\x1a\x1f.secured_service.SecureResponse\x12[\n" +
	"\fDoublySecure\x12$.secured_service.DoublySecureRequest\x1a%.secured_service.DoublySecureResponse\x12g\n" +
	"\x10AlsoDoublySecure\x12(.secured_service.AlsoDoublySecureRequest\x1a).secured_service.AlsoDoublySecureResponseB\x14Z\x12/secured_servicepbb\x06proto3"

var (
	file_goagen_multiauth_secured_service_proto_rawDescOnce sync.Once
	file_goagen_multiauth_secured_service_proto_rawDescData []byte
)

func file_goagen_multiauth_secured_service_proto_rawDescGZIP() []byte {
	file_goagen_multiauth_secured_service_proto_rawDescOnce.Do(func() {
		file_goagen_multiauth_secured_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_goagen_multiauth_secured_service_proto_rawDesc), len(file_goagen_multiauth_secured_service_proto_rawDesc)))
	})
	return file_goagen_multiauth_secured_service_proto_rawDescData
}

var file_goagen_multiauth_secured_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goagen_multiauth_secured_service_proto_goTypes = []any{
	(*SigninRequest)(nil),            // 0: secured_service.SigninRequest
	(*SigninResponse)(nil),           // 1: secured_service.SigninResponse
	(*SecureRequest)(nil),            // 2: secured_service.SecureRequest
	(*SecureResponse)(nil),           // 3: secured_service.SecureResponse
	(*DoublySecureRequest)(nil),      // 4: secured_service.DoublySecureRequest
	(*DoublySecureResponse)(nil),     // 5: secured_service.DoublySecureResponse
	(*AlsoDoublySecureRequest)(nil),  // 6: secured_service.AlsoDoublySecureRequest
	(*AlsoDoublySecureResponse)(nil), // 7: secured_service.AlsoDoublySecureResponse
}
var file_goagen_multiauth_secured_service_proto_depIdxs = []int32{
	0, // 0: secured_service.SecuredService.Signin:input_type -> secured_service.SigninRequest
	2, // 1: secured_service.SecuredService.Secure:input_type -> secured_service.SecureRequest
	4, // 2: secured_service.SecuredService.DoublySecure:input_type -> secured_service.DoublySecureRequest
	6, // 3: secured_service.SecuredService.AlsoDoublySecure:input_type -> secured_service.AlsoDoublySecureRequest
	1, // 4: secured_service.SecuredService.Signin:output_type -> secured_service.SigninResponse
	3, // 5: secured_service.SecuredService.Secure:output_type -> secured_service.SecureResponse
	5, // 6: secured_service.SecuredService.DoublySecure:output_type -> secured_service.DoublySecureResponse
	7, // 7: secured_service.SecuredService.AlsoDoublySecure:output_type -> secured_service.AlsoDoublySecureResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_multiauth_secured_service_proto_init() }
func file_goagen_multiauth_secured_service_proto_init() {
	if File_goagen_multiauth_secured_service_proto != nil {
		return
	}
	file_goagen_multiauth_secured_service_proto_msgTypes[2].OneofWrappers = []any{}
	file_goagen_multiauth_secured_service_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_goagen_multiauth_secured_service_proto_rawDesc), len(file_goagen_multiauth_secured_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_multiauth_secured_service_proto_goTypes,
		DependencyIndexes: file_goagen_multiauth_secured_service_proto_depIdxs,
		MessageInfos:      file_goagen_multiauth_secured_service_proto_msgTypes,
	}.Build()
	File_goagen_multiauth_secured_service_proto = out.File
	file_goagen_multiauth_secured_service_proto_goTypes = nil
	file_goagen_multiauth_secured_service_proto_depIdxs = nil
}
