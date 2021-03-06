// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: email.proto

package ms_email_grpc_api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// SendEmailRequest represent SendEmail Request
type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateCode string          `protobuf:"bytes,1,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
	From         *EmailAddress   `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	To           *EmailAddress   `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	Cc           []*EmailAddress `protobuf:"bytes,4,rep,name=Cc,proto3" json:"Cc,omitempty"`
	Bcc          []*EmailAddress `protobuf:"bytes,5,rep,name=Bcc,proto3" json:"Bcc,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

func (x *SendEmailRequest) GetFrom() *EmailAddress {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *SendEmailRequest) GetTo() *EmailAddress {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendEmailRequest) GetCc() []*EmailAddress {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *SendEmailRequest) GetBcc() []*EmailAddress {
	if x != nil {
		return x.Bcc
	}
	return nil
}

// EmailAddress represent EmailAddress
type EmailAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *EmailAddress) Reset() {
	*x = EmailAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAddress) ProtoMessage() {}

func (x *EmailAddress) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAddress.ProtoReflect.Descriptor instead.
func (*EmailAddress) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{1}
}

func (x *EmailAddress) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SendEmailResponse represent SendEmail Response
type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateCode string `protobuf:"bytes,1,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailResponse) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe8, 0x01, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x29, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x29, 0x0a,
	0x02, 0x43, 0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x02, 0x43, 0x63, 0x12, 0x2b, 0x0a, 0x03, 0x42, 0x63, 0x63, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x03, 0x42, 0x63, 0x63, 0x22, 0x38, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x37, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x3b, 0x6d, 0x73,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_email_proto_goTypes = []interface{}{
	(*SendEmailRequest)(nil),  // 0: email.email.SendEmailRequest
	(*EmailAddress)(nil),      // 1: email.email.EmailAddress
	(*SendEmailResponse)(nil), // 2: email.email.SendEmailResponse
}
var file_email_proto_depIdxs = []int32{
	1, // 0: email.email.SendEmailRequest.From:type_name -> email.email.EmailAddress
	1, // 1: email.email.SendEmailRequest.To:type_name -> email.email.EmailAddress
	1, // 2: email.email.SendEmailRequest.Cc:type_name -> email.email.EmailAddress
	1, // 3: email.email.SendEmailRequest.Bcc:type_name -> email.email.EmailAddress
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAddress); i {
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
		file_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}
