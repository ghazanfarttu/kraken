// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HostFrequencyScaler.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HostFrequencyScaler_ScalerState int32

const (
	HostFrequencyScaler_NONE        HostFrequencyScaler_ScalerState = 0
	HostFrequencyScaler_POWER_SAVE  HostFrequencyScaler_ScalerState = 1
	HostFrequencyScaler_PERFORMANCE HostFrequencyScaler_ScalerState = 2
)

var HostFrequencyScaler_ScalerState_name = map[int32]string{
	0: "NONE",
	1: "POWER_SAVE",
	2: "PERFORMANCE",
}
var HostFrequencyScaler_ScalerState_value = map[string]int32{
	"NONE":        0,
	"POWER_SAVE":  1,
	"PERFORMANCE": 2,
}

func (x HostFrequencyScaler_ScalerState) String() string {
	return proto.EnumName(HostFrequencyScaler_ScalerState_name, int32(x))
}
func (HostFrequencyScaler_ScalerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_HostFrequencyScaler_2508640b1434a762, []int{0, 0}
}

type HostFrequencyScaler struct {
	State                HostFrequencyScaler_ScalerState `protobuf:"varint,1,opt,name=state,proto3,enum=proto.HostFrequencyScaler_ScalerState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *HostFrequencyScaler) Reset()         { *m = HostFrequencyScaler{} }
func (m *HostFrequencyScaler) String() string { return proto.CompactTextString(m) }
func (*HostFrequencyScaler) ProtoMessage()    {}
func (*HostFrequencyScaler) Descriptor() ([]byte, []int) {
	return fileDescriptor_HostFrequencyScaler_2508640b1434a762, []int{0}
}
func (m *HostFrequencyScaler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostFrequencyScaler.Unmarshal(m, b)
}
func (m *HostFrequencyScaler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostFrequencyScaler.Marshal(b, m, deterministic)
}
func (dst *HostFrequencyScaler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostFrequencyScaler.Merge(dst, src)
}
func (m *HostFrequencyScaler) XXX_Size() int {
	return xxx_messageInfo_HostFrequencyScaler.Size(m)
}
func (m *HostFrequencyScaler) XXX_DiscardUnknown() {
	xxx_messageInfo_HostFrequencyScaler.DiscardUnknown(m)
}

var xxx_messageInfo_HostFrequencyScaler proto.InternalMessageInfo

func (m *HostFrequencyScaler) GetState() HostFrequencyScaler_ScalerState {
	if m != nil {
		return m.State
	}
	return HostFrequencyScaler_NONE
}

func init() {
	proto.RegisterType((*HostFrequencyScaler)(nil), "proto.HostFrequencyScaler")
	proto.RegisterEnum("proto.HostFrequencyScaler_ScalerState", HostFrequencyScaler_ScalerState_name, HostFrequencyScaler_ScalerState_value)
}

func init() {
	proto.RegisterFile("HostFrequencyScaler.proto", fileDescriptor_HostFrequencyScaler_2508640b1434a762)
}

var fileDescriptor_HostFrequencyScaler_2508640b1434a762 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf4, 0xc8, 0x2f, 0x2e,
	0x71, 0x2b, 0x4a, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0xae, 0x0c, 0x4e, 0x4e, 0xcc, 0x49, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xbd, 0x8c, 0x5c, 0xc2, 0x58, 0x14,
	0x09, 0xd9, 0x70, 0xb1, 0x16, 0x97, 0x24, 0x96, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19,
	0xa9, 0x41, 0x74, 0xe9, 0x61, 0x33, 0x0f, 0x42, 0x05, 0x83, 0x54, 0x07, 0x41, 0x34, 0x29, 0x59,
	0x70, 0x71, 0x23, 0x89, 0x0a, 0x71, 0x70, 0xb1, 0xf8, 0xf9, 0xfb, 0xb9, 0x0a, 0x30, 0x08, 0xf1,
	0x71, 0x71, 0x05, 0xf8, 0x87, 0xbb, 0x06, 0xc5, 0x07, 0x3b, 0x86, 0xb9, 0x0a, 0x30, 0x0a, 0xf1,
	0x73, 0x71, 0x07, 0xb8, 0x06, 0xb9, 0xf9, 0x07, 0xf9, 0x3a, 0xfa, 0x39, 0xbb, 0x0a, 0x30, 0x25,
	0xb1, 0x81, 0xed, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb5, 0xba, 0x7f, 0xba, 0x00,
	0x00, 0x00,
}
