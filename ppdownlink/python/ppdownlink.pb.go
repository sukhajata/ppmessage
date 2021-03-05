// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: ppdownlink.proto

package ppdownlink

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

type ConfigDownlinkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceeui  string `protobuf:"bytes,1,opt,name=deviceeui,proto3" json:"deviceeui,omitempty"`
	Slot       uint32 `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Index      uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Firmware   string `protobuf:"bytes,4,opt,name=firmware,proto3" json:"firmware,omitempty"`
	Value      []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Numretries uint32 `protobuf:"varint,6,opt,name=numretries,proto3" json:"numretries,omitempty"`
}

func (x *ConfigDownlinkMessage) Reset() {
	*x = ConfigDownlinkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppdownlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDownlinkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDownlinkMessage) ProtoMessage() {}

func (x *ConfigDownlinkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppdownlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDownlinkMessage.ProtoReflect.Descriptor instead.
func (*ConfigDownlinkMessage) Descriptor() ([]byte, []int) {
	return file_ppdownlink_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigDownlinkMessage) GetDeviceeui() string {
	if x != nil {
		return x.Deviceeui
	}
	return ""
}

func (x *ConfigDownlinkMessage) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ConfigDownlinkMessage) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ConfigDownlinkMessage) GetFirmware() string {
	if x != nil {
		return x.Firmware
	}
	return ""
}

func (x *ConfigDownlinkMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ConfigDownlinkMessage) GetNumretries() uint32 {
	if x != nil {
		return x.Numretries
	}
	return 0
}

type FunctionDownlinkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceeui string `protobuf:"bytes,1,opt,name=deviceeui,proto3" json:"deviceeui,omitempty"`
	Slot      uint32 `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Index     uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Firmware  string `protobuf:"bytes,4,opt,name=firmware,proto3" json:"firmware,omitempty"`
	Param1    []byte `protobuf:"bytes,5,opt,name=param1,proto3" json:"param1,omitempty"`
	Param2    []byte `protobuf:"bytes,6,opt,name=param2,proto3" json:"param2,omitempty"`
	Param3    []byte `protobuf:"bytes,7,opt,name=param3,proto3" json:"param3,omitempty"`
	Param4    []byte `protobuf:"bytes,8,opt,name=param4,proto3" json:"param4,omitempty"`
}

func (x *FunctionDownlinkMessage) Reset() {
	*x = FunctionDownlinkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppdownlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDownlinkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDownlinkMessage) ProtoMessage() {}

func (x *FunctionDownlinkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppdownlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDownlinkMessage.ProtoReflect.Descriptor instead.
func (*FunctionDownlinkMessage) Descriptor() ([]byte, []int) {
	return file_ppdownlink_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionDownlinkMessage) GetDeviceeui() string {
	if x != nil {
		return x.Deviceeui
	}
	return ""
}

func (x *FunctionDownlinkMessage) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *FunctionDownlinkMessage) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *FunctionDownlinkMessage) GetFirmware() string {
	if x != nil {
		return x.Firmware
	}
	return ""
}

func (x *FunctionDownlinkMessage) GetParam1() []byte {
	if x != nil {
		return x.Param1
	}
	return nil
}

func (x *FunctionDownlinkMessage) GetParam2() []byte {
	if x != nil {
		return x.Param2
	}
	return nil
}

func (x *FunctionDownlinkMessage) GetParam3() []byte {
	if x != nil {
		return x.Param3
	}
	return nil
}

func (x *FunctionDownlinkMessage) GetParam4() []byte {
	if x != nil {
		return x.Param4
	}
	return nil
}

type ResendRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceeui   string `protobuf:"bytes,1,opt,name=deviceeui,proto3" json:"deviceeui,omitempty"`
	Timesent    uint64 `protobuf:"varint,2,opt,name=timesent,proto3" json:"timesent,omitempty"`
	Messageid   uint32 `protobuf:"varint,3,opt,name=messageid,proto3" json:"messageid,omitempty"`
	Messagetype uint32 `protobuf:"varint,4,opt,name=messagetype,proto3" json:"messagetype,omitempty"`
	Spare       int32  `protobuf:"varint,5,opt,name=spare,proto3" json:"spare,omitempty"`
	State       int32  `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ResendRequestMessage) Reset() {
	*x = ResendRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppdownlink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendRequestMessage) ProtoMessage() {}

func (x *ResendRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppdownlink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendRequestMessage.ProtoReflect.Descriptor instead.
func (*ResendRequestMessage) Descriptor() ([]byte, []int) {
	return file_ppdownlink_proto_rawDescGZIP(), []int{2}
}

func (x *ResendRequestMessage) GetDeviceeui() string {
	if x != nil {
		return x.Deviceeui
	}
	return ""
}

func (x *ResendRequestMessage) GetTimesent() uint64 {
	if x != nil {
		return x.Timesent
	}
	return 0
}

func (x *ResendRequestMessage) GetMessageid() uint32 {
	if x != nil {
		return x.Messageid
	}
	return 0
}

func (x *ResendRequestMessage) GetMessagetype() uint32 {
	if x != nil {
		return x.Messagetype
	}
	return 0
}

func (x *ResendRequestMessage) GetSpare() int32 {
	if x != nil {
		return x.Spare
	}
	return 0
}

func (x *ResendRequestMessage) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type RelayCommandMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceeui string `protobuf:"bytes,1,opt,name=deviceeui,proto3" json:"deviceeui,omitempty"`
	Timesent  uint64 `protobuf:"varint,2,opt,name=timesent,proto3" json:"timesent,omitempty"`
	Index     uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Command   uint32 `protobuf:"varint,4,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *RelayCommandMessage) Reset() {
	*x = RelayCommandMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppdownlink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayCommandMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayCommandMessage) ProtoMessage() {}

func (x *RelayCommandMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppdownlink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayCommandMessage.ProtoReflect.Descriptor instead.
func (*RelayCommandMessage) Descriptor() ([]byte, []int) {
	return file_ppdownlink_proto_rawDescGZIP(), []int{3}
}

func (x *RelayCommandMessage) GetDeviceeui() string {
	if x != nil {
		return x.Deviceeui
	}
	return ""
}

func (x *RelayCommandMessage) GetTimesent() uint64 {
	if x != nil {
		return x.Timesent
	}
	return 0
}

func (x *RelayCommandMessage) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *RelayCommandMessage) GetCommand() uint32 {
	if x != nil {
		return x.Command
	}
	return 0
}

type MulticastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cg1     uint32 `protobuf:"varint,1,opt,name=cg1,proto3" json:"cg1,omitempty"`
	Cg2     uint32 `protobuf:"varint,2,opt,name=cg2,proto3" json:"cg2,omitempty"`
	Cg3     uint32 `protobuf:"varint,3,opt,name=cg3,proto3" json:"cg3,omitempty"`
	Cg4     uint32 `protobuf:"varint,4,opt,name=cg4,proto3" json:"cg4,omitempty"`
	Cg5     uint32 `protobuf:"varint,5,opt,name=cg5,proto3" json:"cg5,omitempty"`
	Cg6     uint32 `protobuf:"varint,6,opt,name=cg6,proto3" json:"cg6,omitempty"`
	Cg7     uint32 `protobuf:"varint,7,opt,name=cg7,proto3" json:"cg7,omitempty"`
	Cg8     uint32 `protobuf:"varint,8,opt,name=cg8,proto3" json:"cg8,omitempty"`
	Cg9     uint32 `protobuf:"varint,9,opt,name=cg9,proto3" json:"cg9,omitempty"`
	Cg10    uint32 `protobuf:"varint,10,opt,name=cg10,proto3" json:"cg10,omitempty"`
	Cg11    uint32 `protobuf:"varint,11,opt,name=cg11,proto3" json:"cg11,omitempty"`
	Cg12    uint32 `protobuf:"varint,12,opt,name=cg12,proto3" json:"cg12,omitempty"`
	Cg13    uint32 `protobuf:"varint,13,opt,name=cg13,proto3" json:"cg13,omitempty"`
	Cg14    uint32 `protobuf:"varint,14,opt,name=cg14,proto3" json:"cg14,omitempty"`
	Cg15    uint32 `protobuf:"varint,15,opt,name=cg15,proto3" json:"cg15,omitempty"`
	Cg16    uint32 `protobuf:"varint,16,opt,name=cg16,proto3" json:"cg16,omitempty"`
	Operand uint32 `protobuf:"varint,17,opt,name=operand,proto3" json:"operand,omitempty"`
	Command uint32 `protobuf:"varint,18,opt,name=command,proto3" json:"command,omitempty"`
	Value   uint32 `protobuf:"varint,19,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MulticastMessage) Reset() {
	*x = MulticastMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppdownlink_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MulticastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulticastMessage) ProtoMessage() {}

func (x *MulticastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppdownlink_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulticastMessage.ProtoReflect.Descriptor instead.
func (*MulticastMessage) Descriptor() ([]byte, []int) {
	return file_ppdownlink_proto_rawDescGZIP(), []int{4}
}

func (x *MulticastMessage) GetCg1() uint32 {
	if x != nil {
		return x.Cg1
	}
	return 0
}

func (x *MulticastMessage) GetCg2() uint32 {
	if x != nil {
		return x.Cg2
	}
	return 0
}

func (x *MulticastMessage) GetCg3() uint32 {
	if x != nil {
		return x.Cg3
	}
	return 0
}

func (x *MulticastMessage) GetCg4() uint32 {
	if x != nil {
		return x.Cg4
	}
	return 0
}

func (x *MulticastMessage) GetCg5() uint32 {
	if x != nil {
		return x.Cg5
	}
	return 0
}

func (x *MulticastMessage) GetCg6() uint32 {
	if x != nil {
		return x.Cg6
	}
	return 0
}

func (x *MulticastMessage) GetCg7() uint32 {
	if x != nil {
		return x.Cg7
	}
	return 0
}

func (x *MulticastMessage) GetCg8() uint32 {
	if x != nil {
		return x.Cg8
	}
	return 0
}

func (x *MulticastMessage) GetCg9() uint32 {
	if x != nil {
		return x.Cg9
	}
	return 0
}

func (x *MulticastMessage) GetCg10() uint32 {
	if x != nil {
		return x.Cg10
	}
	return 0
}

func (x *MulticastMessage) GetCg11() uint32 {
	if x != nil {
		return x.Cg11
	}
	return 0
}

func (x *MulticastMessage) GetCg12() uint32 {
	if x != nil {
		return x.Cg12
	}
	return 0
}

func (x *MulticastMessage) GetCg13() uint32 {
	if x != nil {
		return x.Cg13
	}
	return 0
}

func (x *MulticastMessage) GetCg14() uint32 {
	if x != nil {
		return x.Cg14
	}
	return 0
}

func (x *MulticastMessage) GetCg15() uint32 {
	if x != nil {
		return x.Cg15
	}
	return 0
}

func (x *MulticastMessage) GetCg16() uint32 {
	if x != nil {
		return x.Cg16
	}
	return 0
}

func (x *MulticastMessage) GetOperand() uint32 {
	if x != nil {
		return x.Operand
	}
	return 0
}

func (x *MulticastMessage) GetCommand() uint32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *MulticastMessage) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_ppdownlink_proto protoreflect.FileDescriptor

var file_ppdownlink_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xb1,
	0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x33, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x34, 0x22, 0xbc, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x70, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x7f, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x65, 0x75, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x8a, 0x03, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x33, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x67, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x34, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x67, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67,
	0x35, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x63, 0x67, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x63, 0x67, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x38, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x67, 0x39, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x67, 0x39, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x67, 0x31,
	0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x67, 0x31, 0x30, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x67, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x67, 0x31,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x67, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x67, 0x31, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x67, 0x31, 0x33, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x67, 0x31, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x67, 0x31,
	0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x67, 0x31, 0x34, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x67, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x67, 0x31,
	0x35, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x67, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x67, 0x31, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75,
	0x6b, 0x68, 0x61, 0x6a, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x70, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ppdownlink_proto_rawDescOnce sync.Once
	file_ppdownlink_proto_rawDescData = file_ppdownlink_proto_rawDesc
)

func file_ppdownlink_proto_rawDescGZIP() []byte {
	file_ppdownlink_proto_rawDescOnce.Do(func() {
		file_ppdownlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_ppdownlink_proto_rawDescData)
	})
	return file_ppdownlink_proto_rawDescData
}

var file_ppdownlink_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ppdownlink_proto_goTypes = []interface{}{
	(*ConfigDownlinkMessage)(nil),   // 0: ppdownlink.ConfigDownlinkMessage
	(*FunctionDownlinkMessage)(nil), // 1: ppdownlink.FunctionDownlinkMessage
	(*ResendRequestMessage)(nil),    // 2: ppdownlink.ResendRequestMessage
	(*RelayCommandMessage)(nil),     // 3: ppdownlink.RelayCommandMessage
	(*MulticastMessage)(nil),        // 4: ppdownlink.MulticastMessage
}
var file_ppdownlink_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ppdownlink_proto_init() }
func file_ppdownlink_proto_init() {
	if File_ppdownlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ppdownlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDownlinkMessage); i {
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
		file_ppdownlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDownlinkMessage); i {
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
		file_ppdownlink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendRequestMessage); i {
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
		file_ppdownlink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayCommandMessage); i {
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
		file_ppdownlink_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MulticastMessage); i {
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
			RawDescriptor: file_ppdownlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ppdownlink_proto_goTypes,
		DependencyIndexes: file_ppdownlink_proto_depIdxs,
		MessageInfos:      file_ppdownlink_proto_msgTypes,
	}.Build()
	File_ppdownlink_proto = out.File
	file_ppdownlink_proto_rawDesc = nil
	file_ppdownlink_proto_goTypes = nil
	file_ppdownlink_proto_depIdxs = nil
}