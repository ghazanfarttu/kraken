// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hostthermaldiscovery.proto

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

type HostDiscoveryConfig struct {
	PollingInterval      string                            `protobuf:"bytes,1,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	TempSensorPath       string                            `protobuf:"bytes,2,opt,name=temp_sensor_path,json=tempSensorPath,proto3" json:"temp_sensor_path,omitempty"`
	ThermalThresholds    map[string]*HostThermalThresholds `protobuf:"bytes,3,rep,name=thermal_thresholds,json=thermalThresholds,proto3" json:"thermal_thresholds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HostDiscoveryConfig) Reset()         { *m = HostDiscoveryConfig{} }
func (m *HostDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*HostDiscoveryConfig) ProtoMessage()    {}
func (*HostDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_hostthermaldiscovery_063889daa550b856, []int{0}
}
func (m *HostDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostDiscoveryConfig.Unmarshal(m, b)
}
func (m *HostDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostDiscoveryConfig.Marshal(b, m, deterministic)
}
func (dst *HostDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostDiscoveryConfig.Merge(dst, src)
}
func (m *HostDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_HostDiscoveryConfig.Size(m)
}
func (m *HostDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HostDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HostDiscoveryConfig proto.InternalMessageInfo

func (m *HostDiscoveryConfig) GetPollingInterval() string {
	if m != nil {
		return m.PollingInterval
	}
	return ""
}

func (m *HostDiscoveryConfig) GetTempSensorPath() string {
	if m != nil {
		return m.TempSensorPath
	}
	return ""
}

func (m *HostDiscoveryConfig) GetThermalThresholds() map[string]*HostThermalThresholds {
	if m != nil {
		return m.ThermalThresholds
	}
	return nil
}

type HostThermalThresholds struct {
	LowerNormal          int32    `protobuf:"varint,1,opt,name=lower_normal,json=lowerNormal,proto3" json:"lower_normal,omitempty"`
	UpperNormal          int32    `protobuf:"varint,2,opt,name=upper_normal,json=upperNormal,proto3" json:"upper_normal,omitempty"`
	LowerHigh            int32    `protobuf:"varint,3,opt,name=lower_high,json=lowerHigh,proto3" json:"lower_high,omitempty"`
	UpperHigh            int32    `protobuf:"varint,4,opt,name=upper_high,json=upperHigh,proto3" json:"upper_high,omitempty"`
	LowerCritical        int32    `protobuf:"varint,5,opt,name=lower_critical,json=lowerCritical,proto3" json:"lower_critical,omitempty"`
	UpperCritical        int32    `protobuf:"varint,6,opt,name=upper_critical,json=upperCritical,proto3" json:"upper_critical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostThermalThresholds) Reset()         { *m = HostThermalThresholds{} }
func (m *HostThermalThresholds) String() string { return proto.CompactTextString(m) }
func (*HostThermalThresholds) ProtoMessage()    {}
func (*HostThermalThresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_hostthermaldiscovery_063889daa550b856, []int{1}
}
func (m *HostThermalThresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostThermalThresholds.Unmarshal(m, b)
}
func (m *HostThermalThresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostThermalThresholds.Marshal(b, m, deterministic)
}
func (dst *HostThermalThresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostThermalThresholds.Merge(dst, src)
}
func (m *HostThermalThresholds) XXX_Size() int {
	return xxx_messageInfo_HostThermalThresholds.Size(m)
}
func (m *HostThermalThresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_HostThermalThresholds.DiscardUnknown(m)
}

var xxx_messageInfo_HostThermalThresholds proto.InternalMessageInfo

func (m *HostThermalThresholds) GetLowerNormal() int32 {
	if m != nil {
		return m.LowerNormal
	}
	return 0
}

func (m *HostThermalThresholds) GetUpperNormal() int32 {
	if m != nil {
		return m.UpperNormal
	}
	return 0
}

func (m *HostThermalThresholds) GetLowerHigh() int32 {
	if m != nil {
		return m.LowerHigh
	}
	return 0
}

func (m *HostThermalThresholds) GetUpperHigh() int32 {
	if m != nil {
		return m.UpperHigh
	}
	return 0
}

func (m *HostThermalThresholds) GetLowerCritical() int32 {
	if m != nil {
		return m.LowerCritical
	}
	return 0
}

func (m *HostThermalThresholds) GetUpperCritical() int32 {
	if m != nil {
		return m.UpperCritical
	}
	return 0
}

func init() {
	proto.RegisterType((*HostDiscoveryConfig)(nil), "proto.HostDiscoveryConfig")
	proto.RegisterMapType((map[string]*HostThermalThresholds)(nil), "proto.HostDiscoveryConfig.ThermalThresholdsEntry")
	proto.RegisterType((*HostThermalThresholds)(nil), "proto.HostThermalThresholds")
}

func init() {
	proto.RegisterFile("hostthermaldiscovery.proto", fileDescriptor_hostthermaldiscovery_063889daa550b856)
}

var fileDescriptor_hostthermaldiscovery_063889daa550b856 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0x6b, 0x07, 0x3b, 0xd3, 0x39, 0x23, 0x4a, 0x19, 0x0a, 0x73, 0x20, 0xcc, 0x9b,
	0x81, 0xf3, 0x46, 0xbc, 0x9d, 0xc2, 0xbc, 0x11, 0xa9, 0xbb, 0xaf, 0xdd, 0x16, 0x9b, 0x60, 0xd6,
	0x94, 0xe4, 0x6c, 0xb2, 0xb7, 0xf2, 0xb5, 0x7c, 0x0b, 0xe9, 0x49, 0x37, 0x45, 0x7b, 0xd5, 0xf2,
	0xfd, 0xdf, 0x39, 0xf9, 0x13, 0xe8, 0x0a, 0x6d, 0x11, 0x05, 0x37, 0xcb, 0x54, 0x2d, 0xa4, 0x9d,
	0xeb, 0x35, 0x37, 0x9b, 0x61, 0x61, 0x34, 0x6a, 0x16, 0xd2, 0xa7, 0xff, 0xe9, 0xc3, 0xf1, 0x44,
	0x5b, 0xbc, 0xdf, 0xc6, 0x63, 0x9d, 0xbf, 0xc9, 0x8c, 0x5d, 0x41, 0xa7, 0xd0, 0x4a, 0xc9, 0x3c,
	0x4b, 0x64, 0x8e, 0xdc, 0xac, 0x53, 0x15, 0x79, 0x3d, 0x6f, 0xd0, 0x8c, 0x0f, 0x2b, 0xfe, 0x58,
	0x61, 0x36, 0x80, 0x0e, 0xf2, 0x65, 0x91, 0x58, 0x9e, 0x5b, 0x6d, 0x92, 0x22, 0x45, 0x11, 0xf9,
	0xa4, 0xb6, 0x4b, 0xfe, 0x42, 0xf8, 0x39, 0x45, 0xc1, 0x5e, 0x81, 0x55, 0x6d, 0x12, 0x14, 0x86,
	0x5b, 0xa1, 0xd5, 0xc2, 0x46, 0x41, 0x2f, 0x18, 0xb4, 0x46, 0xd7, 0xae, 0xd7, 0xb0, 0xa6, 0xcc,
	0x70, 0xea, 0x86, 0xa6, 0xbb, 0x99, 0x87, 0x1c, 0xcd, 0x26, 0x3e, 0xc2, 0xbf, 0xbc, 0x3b, 0x83,
	0xd3, 0x7a, 0x99, 0x75, 0x20, 0x78, 0xe7, 0x9b, 0xea, 0x0e, 0xe5, 0x2f, 0x1b, 0x41, 0xb8, 0x4e,
	0xd5, 0x8a, 0x53, 0xd9, 0xd6, 0xe8, 0xec, 0x57, 0x81, 0x7f, 0x3b, 0x62, 0xa7, 0xde, 0xf9, 0xb7,
	0x5e, 0xff, 0xcb, 0x83, 0x93, 0x5a, 0x89, 0x5d, 0xc0, 0xbe, 0xd2, 0x1f, 0xdc, 0x24, 0xb9, 0x2e,
	0x13, 0x3a, 0x2c, 0x8c, 0x5b, 0xc4, 0x9e, 0x08, 0x95, 0xca, 0xaa, 0x28, 0x7e, 0x14, 0xdf, 0x29,
	0xc4, 0x2a, 0xe5, 0x1c, 0xc0, 0x6d, 0x11, 0x32, 0x13, 0x51, 0x40, 0x42, 0x93, 0xc8, 0x44, 0x66,
	0xa2, 0x8c, 0xdd, 0x06, 0x8a, 0xf7, 0x5c, 0x4c, 0x84, 0xe2, 0x4b, 0x68, 0xbb, 0xe9, 0xb9, 0x91,
	0x28, 0xe7, 0xa9, 0x8a, 0x42, 0x52, 0x0e, 0x88, 0x8e, 0x2b, 0x58, 0x6a, 0x6e, 0xcb, 0x4e, 0x6b,
	0x38, 0x8d, 0xe8, 0x56, 0x9b, 0x35, 0xe8, 0x4d, 0x6e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x5a, 0x78, 0xab, 0x4a, 0x02, 0x00, 0x00,
}
