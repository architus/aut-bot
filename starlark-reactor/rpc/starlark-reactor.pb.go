// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: lib/ipc/proto/starlark-reactor.proto

package starlarkreactor

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

type StarlarkScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script         string   `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	TriggerMessage *Message `protobuf:"bytes,2,opt,name=trigger_message,json=triggerMessage,proto3" json:"trigger_message,omitempty"`
	Author         *Author  `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Count          uint64   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Captures       []string `protobuf:"bytes,5,rep,name=captures,proto3" json:"captures,omitempty"`
	Arguments      []string `protobuf:"bytes,6,rep,name=arguments,proto3" json:"arguments,omitempty"`
	Channel        *Channel `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *StarlarkScript) Reset() {
	*x = StarlarkScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarlarkScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarlarkScript) ProtoMessage() {}

func (x *StarlarkScript) ProtoReflect() protoreflect.Message {
	mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarlarkScript.ProtoReflect.Descriptor instead.
func (*StarlarkScript) Descriptor() ([]byte, []int) {
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP(), []int{0}
}

func (x *StarlarkScript) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *StarlarkScript) GetTriggerMessage() *Message {
	if x != nil {
		return x.TriggerMessage
	}
	return nil
}

func (x *StarlarkScript) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *StarlarkScript) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *StarlarkScript) GetCaptures() []string {
	if x != nil {
		return x.Captures
	}
	return nil
}

func (x *StarlarkScript) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *StarlarkScript) GetChannel() *Channel {
	if x != nil {
		return x.Channel
	}
	return nil
}

type ScriptOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Errno  uint64 `protobuf:"varint,3,opt,name=errno,proto3" json:"errno,omitempty"`
}

func (x *ScriptOutput) Reset() {
	*x = ScriptOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptOutput) ProtoMessage() {}

func (x *ScriptOutput) ProtoReflect() protoreflect.Message {
	mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptOutput.ProtoReflect.Descriptor instead.
func (*ScriptOutput) Descriptor() ([]byte, []int) {
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP(), []int{1}
}

func (x *ScriptOutput) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ScriptOutput) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ScriptOutput) GetErrno() uint64 {
	if x != nil {
		return x.Errno
	}
	return 0
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP(), []int{2}
}

func (x *Channel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AvatarUrl     string   `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Color         string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Discriminator uint32   `protobuf:"varint,4,opt,name=discriminator,proto3" json:"discriminator,omitempty"`
	Roles         []uint64 `protobuf:"varint,5,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Name          string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Nick          string   `protobuf:"bytes,7,opt,name=nick,proto3" json:"nick,omitempty"`
	DispName      string   `protobuf:"bytes,8,opt,name=disp_name,json=dispName,proto3" json:"disp_name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP(), []int{3}
}

func (x *Author) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Author) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Author) GetDiscriminator() uint32 {
	if x != nil {
		return x.Discriminator
	}
	return 0
}

func (x *Author) GetRoles() []uint64 {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *Author) GetDispName() string {
	if x != nil {
		return x.DispName
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Clean   string `protobuf:"bytes,3,opt,name=clean,proto3" json:"clean,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_lib_ipc_proto_starlark_reactor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetClean() string {
	if x != nil {
		return x.Clean
	}
	return ""
}

var File_lib_ipc_proto_starlark_reactor_proto protoreflect.FileDescriptor

var file_lib_ipc_proto_starlark_reactor_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6c, 0x69, 0x62, 0x2f, 0x69, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x2d, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xa0, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72,
	0x6c, 0x61, 0x72, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6c, 0x61,
	0x72, 0x6b, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x52, 0x0a, 0x0c, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x22, 0x2d,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xce, 0x01,
	0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63,
	0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x32, 0x68, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x72, 0x6c, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x11,
	0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x75, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x6c, 0x61, 0x72, 0x6b, 0x72, 0x65, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_ipc_proto_starlark_reactor_proto_rawDescOnce sync.Once
	file_lib_ipc_proto_starlark_reactor_proto_rawDescData = file_lib_ipc_proto_starlark_reactor_proto_rawDesc
)

func file_lib_ipc_proto_starlark_reactor_proto_rawDescGZIP() []byte {
	file_lib_ipc_proto_starlark_reactor_proto_rawDescOnce.Do(func() {
		file_lib_ipc_proto_starlark_reactor_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_ipc_proto_starlark_reactor_proto_rawDescData)
	})
	return file_lib_ipc_proto_starlark_reactor_proto_rawDescData
}

var file_lib_ipc_proto_starlark_reactor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_lib_ipc_proto_starlark_reactor_proto_goTypes = []interface{}{
	(*StarlarkScript)(nil), // 0: starlarkreactor.StarlarkScript
	(*ScriptOutput)(nil),   // 1: starlarkreactor.ScriptOutput
	(*Channel)(nil),        // 2: starlarkreactor.Channel
	(*Author)(nil),         // 3: starlarkreactor.Author
	(*Message)(nil),        // 4: starlarkreactor.Message
}
var file_lib_ipc_proto_starlark_reactor_proto_depIdxs = []int32{
	4, // 0: starlarkreactor.StarlarkScript.trigger_message:type_name -> starlarkreactor.Message
	3, // 1: starlarkreactor.StarlarkScript.author:type_name -> starlarkreactor.Author
	2, // 2: starlarkreactor.StarlarkScript.channel:type_name -> starlarkreactor.Channel
	0, // 3: starlarkreactor.StarlarkReactor.RunStarlarkScript:input_type -> starlarkreactor.StarlarkScript
	1, // 4: starlarkreactor.StarlarkReactor.RunStarlarkScript:output_type -> starlarkreactor.ScriptOutput
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lib_ipc_proto_starlark_reactor_proto_init() }
func file_lib_ipc_proto_starlark_reactor_proto_init() {
	if File_lib_ipc_proto_starlark_reactor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_ipc_proto_starlark_reactor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarlarkScript); i {
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
		file_lib_ipc_proto_starlark_reactor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptOutput); i {
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
		file_lib_ipc_proto_starlark_reactor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_lib_ipc_proto_starlark_reactor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_lib_ipc_proto_starlark_reactor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_lib_ipc_proto_starlark_reactor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lib_ipc_proto_starlark_reactor_proto_goTypes,
		DependencyIndexes: file_lib_ipc_proto_starlark_reactor_proto_depIdxs,
		MessageInfos:      file_lib_ipc_proto_starlark_reactor_proto_msgTypes,
	}.Build()
	File_lib_ipc_proto_starlark_reactor_proto = out.File
	file_lib_ipc_proto_starlark_reactor_proto_rawDesc = nil
	file_lib_ipc_proto_starlark_reactor_proto_goTypes = nil
	file_lib_ipc_proto_starlark_reactor_proto_depIdxs = nil
}
