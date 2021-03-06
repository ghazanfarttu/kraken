// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dummy.proto

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

type DummyConfig struct {
	Say                  []string `protobuf:"bytes,1,rep,name=say,proto3" json:"say,omitempty"`
	Wait                 string   `protobuf:"bytes,2,opt,name=wait,proto3" json:"wait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyConfig) Reset()         { *m = DummyConfig{} }
func (m *DummyConfig) String() string { return proto.CompactTextString(m) }
func (*DummyConfig) ProtoMessage()    {}
func (*DummyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_7d8d04d8665a6408, []int{0}
}
func (m *DummyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyConfig.Unmarshal(m, b)
}
func (m *DummyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyConfig.Marshal(b, m, deterministic)
}
func (dst *DummyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyConfig.Merge(dst, src)
}
func (m *DummyConfig) XXX_Size() int {
	return xxx_messageInfo_DummyConfig.Size(m)
}
func (m *DummyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DummyConfig proto.InternalMessageInfo

func (m *DummyConfig) GetSay() []string {
	if m != nil {
		return m.Say
	}
	return nil
}

func (m *DummyConfig) GetWait() string {
	if m != nil {
		return m.Wait
	}
	return ""
}

func init() {
	proto.RegisterType((*DummyConfig)(nil), "proto.DummyConfig")
}

func init() { proto.RegisterFile("dummy.proto", fileDescriptor_dummy_7d8d04d8665a6408) }

var fileDescriptor_dummy_7d8d04d8665a6408 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x29, 0xcd, 0xcd,
	0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xc6, 0x5c, 0xdc, 0x2e,
	0x20, 0x51, 0xe7, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0xc4, 0x4a, 0x09,
	0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x88, 0x8b, 0xa5, 0x3c, 0x31, 0xb3, 0x44,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x4e, 0x62, 0x03, 0xeb, 0x35, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x6c, 0xfc, 0x3f, 0x51, 0x00, 0x00, 0x00,
}
