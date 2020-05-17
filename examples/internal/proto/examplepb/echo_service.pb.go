// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: examples/internal/proto/examplepb/echo_service.proto

// Echo Service
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Embedded represents a message embedded in SimpleMessage.
type Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mark:
	//	*Embedded_Progress
	//	*Embedded_Note
	Mark isEmbedded_Mark `protobuf_oneof:"mark"`
}

func (x *Embedded) Reset() {
	*x = Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embedded) ProtoMessage() {}

func (x *Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embedded.ProtoReflect.Descriptor instead.
func (*Embedded) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{0}
}

func (m *Embedded) GetMark() isEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (x *Embedded) GetProgress() int64 {
	if x, ok := x.GetMark().(*Embedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (x *Embedded) GetNote() string {
	if x, ok := x.GetMark().(*Embedded_Note); ok {
		return x.Note
	}
	return ""
}

type isEmbedded_Mark interface {
	isEmbedded_Mark()
}

type Embedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type Embedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

func (*Embedded_Progress) isEmbedded_Mark() {}

func (*Embedded_Note) isEmbedded_Mark() {}

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// Types that are assignable to Code:
	//	*SimpleMessage_LineNum
	//	*SimpleMessage_Lang
	Code   isSimpleMessage_Code `protobuf_oneof:"code"`
	Status *Embedded            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to Ext:
	//	*SimpleMessage_En
	//	*SimpleMessage_No
	Ext isSimpleMessage_Ext `protobuf_oneof:"ext"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (m *SimpleMessage) GetCode() isSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (x *SimpleMessage) GetLineNum() int64 {
	if x, ok := x.GetCode().(*SimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (x *SimpleMessage) GetLang() string {
	if x, ok := x.GetCode().(*SimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (x *SimpleMessage) GetStatus() *Embedded {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *SimpleMessage) GetExt() isSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (x *SimpleMessage) GetEn() int64 {
	if x, ok := x.GetExt().(*SimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (x *SimpleMessage) GetNo() *Embedded {
	if x, ok := x.GetExt().(*SimpleMessage_No); ok {
		return x.No
	}
	return nil
}

type isSimpleMessage_Code interface {
	isSimpleMessage_Code()
}

type SimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type SimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,proto3,oneof"`
}

func (*SimpleMessage_LineNum) isSimpleMessage_Code() {}

func (*SimpleMessage_Lang) isSimpleMessage_Code() {}

type isSimpleMessage_Ext interface {
	isSimpleMessage_Ext()
}

type SimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,proto3,oneof"`
}

type SimpleMessage_No struct {
	No *Embedded `protobuf:"bytes,7,opt,name=no,proto3,oneof"`
}

func (*SimpleMessage_En) isSimpleMessage_Ext() {}

func (*SimpleMessage_No) isSimpleMessage_Ext() {}

var File_examples_internal_proto_examplepb_echo_service_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_echo_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xa3, 0x02, 0x0a,
	0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x02, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x01, 0x52, 0x02, 0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x48, 0x01, 0x52,
	0x02, 0x6e, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x65,
	0x78, 0x74, 0x32, 0xa7, 0x05, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xbc, 0x02, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x3d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0xae, 0x01, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x5a, 0x1d, 0x12, 0x1b, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6e, 0x75, 0x6d, 0x7d, 0x5a, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x7b, 0x6e, 0x75, 0x6d, 0x7d, 0x2f, 0x7b, 0x6c, 0x61, 0x6e, 0x67, 0x7d, 0x5a,
	0x31, 0x12, 0x2f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65,
	0x63, 0x68, 0x6f, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x7d, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x7d, 0x5a, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x32, 0x2f, 0x7b, 0x6e, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x7d, 0x12, 0xaa, 0x01, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x3d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0xab,
	0x01, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescData = file_examples_internal_proto_examplepb_echo_service_proto_rawDesc
)

func file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_echo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_echo_service_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescData
}

var file_examples_internal_proto_examplepb_echo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_internal_proto_examplepb_echo_service_proto_goTypes = []interface{}{
	(*Embedded)(nil),      // 0: grpc.gateway.examples.internal.proto.examplepb.Embedded
	(*SimpleMessage)(nil), // 1: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
}
var file_examples_internal_proto_examplepb_echo_service_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage.status:type_name -> grpc.gateway.examples.internal.proto.examplepb.Embedded
	0, // 1: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage.no:type_name -> grpc.gateway.examples.internal.proto.examplepb.Embedded
	1, // 2: grpc.gateway.examples.internal.proto.examplepb.EchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1, // 3: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoBody:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1, // 4: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoDelete:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1, // 5: grpc.gateway.examples.internal.proto.examplepb.EchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1, // 6: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoBody:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1, // 7: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoDelete:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_echo_service_proto_init() }
func file_examples_internal_proto_examplepb_echo_service_proto_init() {
	if File_examples_internal_proto_examplepb_echo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embedded); i {
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
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
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
	file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Embedded_Progress)(nil),
		(*Embedded_Note)(nil),
	}
	file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SimpleMessage_LineNum)(nil),
		(*SimpleMessage_Lang)(nil),
		(*SimpleMessage_En)(nil),
		(*SimpleMessage_No)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_echo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_echo_service_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_echo_service_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_echo_service_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_echo_service_proto = out.File
	file_examples_internal_proto_examplepb_echo_service_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_echo_service_proto_goTypes = nil
	file_examples_internal_proto_examplepb_echo_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.EchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.proto.examplepb.EchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) Echo(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedEchoServiceServer) EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (*UnimplementedEchoServiceServer) EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.proto.examplepb.EchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoDelete(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _EchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/echo_service.proto",
}
