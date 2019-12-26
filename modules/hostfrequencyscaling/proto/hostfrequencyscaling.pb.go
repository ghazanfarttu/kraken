// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hostfrequencyscaling.proto

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

type HostFreqScalingConfig struct {
	FreqScalPolicies     map[string]*HostFreqScalingPolicy `protobuf:"bytes,1,rep,name=freq_scal_policies,json=freqScalPolicies,proto3" json:"freq_scal_policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FreqSensorUrl        string                            `protobuf:"bytes,2,opt,name=freq_sensor_url,json=freqSensorUrl,proto3" json:"freq_sensor_url,omitempty"`
	ScalingFreqPolicy    string                            `protobuf:"bytes,3,opt,name=scaling_freq_policy,json=scalingFreqPolicy,proto3" json:"scaling_freq_policy,omitempty"`
	LowToHighScaler      string                            `protobuf:"bytes,4,opt,name=low_to_high_scaler,json=lowToHighScaler,proto3" json:"low_to_high_scaler,omitempty"`
	HighToLowScaler      string                            `protobuf:"bytes,5,opt,name=high_to_low_scaler,json=highToLowScaler,proto3" json:"high_to_low_scaler,omitempty"`
	LowFreqScalerDur     int32                             `protobuf:"varint,6,opt,name=low_freq_scaler_dur,json=lowFreqScalerDur,proto3" json:"low_freq_scaler_dur,omitempty"`
	EnforceLowFreqScaler bool                              `protobuf:"varint,7,opt,name=enforce_low_freq_scaler,json=enforceLowFreqScaler,proto3" json:"enforce_low_freq_scaler,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HostFreqScalingConfig) Reset()         { *m = HostFreqScalingConfig{} }
func (m *HostFreqScalingConfig) String() string { return proto.CompactTextString(m) }
func (*HostFreqScalingConfig) ProtoMessage()    {}
func (*HostFreqScalingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_hostfrequencyscaling_5197ba616350a4a5, []int{0}
}
func (m *HostFreqScalingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostFreqScalingConfig.Unmarshal(m, b)
}
func (m *HostFreqScalingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostFreqScalingConfig.Marshal(b, m, deterministic)
}
func (dst *HostFreqScalingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostFreqScalingConfig.Merge(dst, src)
}
func (m *HostFreqScalingConfig) XXX_Size() int {
	return xxx_messageInfo_HostFreqScalingConfig.Size(m)
}
func (m *HostFreqScalingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HostFreqScalingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HostFreqScalingConfig proto.InternalMessageInfo

func (m *HostFreqScalingConfig) GetFreqScalPolicies() map[string]*HostFreqScalingPolicy {
	if m != nil {
		return m.FreqScalPolicies
	}
	return nil
}

func (m *HostFreqScalingConfig) GetFreqSensorUrl() string {
	if m != nil {
		return m.FreqSensorUrl
	}
	return ""
}

func (m *HostFreqScalingConfig) GetScalingFreqPolicy() string {
	if m != nil {
		return m.ScalingFreqPolicy
	}
	return ""
}

func (m *HostFreqScalingConfig) GetLowToHighScaler() string {
	if m != nil {
		return m.LowToHighScaler
	}
	return ""
}

func (m *HostFreqScalingConfig) GetHighToLowScaler() string {
	if m != nil {
		return m.HighToLowScaler
	}
	return ""
}

func (m *HostFreqScalingConfig) GetLowFreqScalerDur() int32 {
	if m != nil {
		return m.LowFreqScalerDur
	}
	return 0
}

func (m *HostFreqScalingConfig) GetEnforceLowFreqScaler() bool {
	if m != nil {
		return m.EnforceLowFreqScaler
	}
	return false
}

type HostFreqScalingPolicy struct {
	ScalingGovernor      string   `protobuf:"bytes,1,opt,name=scaling_governor,json=scalingGovernor,proto3" json:"scaling_governor,omitempty"`
	ScalingMinFreq       string   `protobuf:"bytes,2,opt,name=scaling_min_freq,json=scalingMinFreq,proto3" json:"scaling_min_freq,omitempty"`
	ScalingMaxFreq       string   `protobuf:"bytes,3,opt,name=scaling_max_freq,json=scalingMaxFreq,proto3" json:"scaling_max_freq,omitempty"`
	NodeArch             string   `protobuf:"bytes,4,opt,name=node_arch,json=nodeArch,proto3" json:"node_arch,omitempty"`
	NodePlatform         string   `protobuf:"bytes,5,opt,name=node_platform,json=nodePlatform,proto3" json:"node_platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostFreqScalingPolicy) Reset()         { *m = HostFreqScalingPolicy{} }
func (m *HostFreqScalingPolicy) String() string { return proto.CompactTextString(m) }
func (*HostFreqScalingPolicy) ProtoMessage()    {}
func (*HostFreqScalingPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_hostfrequencyscaling_5197ba616350a4a5, []int{1}
}
func (m *HostFreqScalingPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostFreqScalingPolicy.Unmarshal(m, b)
}
func (m *HostFreqScalingPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostFreqScalingPolicy.Marshal(b, m, deterministic)
}
func (dst *HostFreqScalingPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostFreqScalingPolicy.Merge(dst, src)
}
func (m *HostFreqScalingPolicy) XXX_Size() int {
	return xxx_messageInfo_HostFreqScalingPolicy.Size(m)
}
func (m *HostFreqScalingPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_HostFreqScalingPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_HostFreqScalingPolicy proto.InternalMessageInfo

func (m *HostFreqScalingPolicy) GetScalingGovernor() string {
	if m != nil {
		return m.ScalingGovernor
	}
	return ""
}

func (m *HostFreqScalingPolicy) GetScalingMinFreq() string {
	if m != nil {
		return m.ScalingMinFreq
	}
	return ""
}

func (m *HostFreqScalingPolicy) GetScalingMaxFreq() string {
	if m != nil {
		return m.ScalingMaxFreq
	}
	return ""
}

func (m *HostFreqScalingPolicy) GetNodeArch() string {
	if m != nil {
		return m.NodeArch
	}
	return ""
}

func (m *HostFreqScalingPolicy) GetNodePlatform() string {
	if m != nil {
		return m.NodePlatform
	}
	return ""
}

func init() {
	proto.RegisterType((*HostFreqScalingConfig)(nil), "proto.HostFreqScalingConfig")
	proto.RegisterMapType((map[string]*HostFreqScalingPolicy)(nil), "proto.HostFreqScalingConfig.FreqScalPoliciesEntry")
	proto.RegisterType((*HostFreqScalingPolicy)(nil), "proto.HostFreqScalingPolicy")
}

func init() {
	proto.RegisterFile("hostfrequencyscaling.proto", fileDescriptor_hostfrequencyscaling_5197ba616350a4a5)
}

var fileDescriptor_hostfrequencyscaling_5197ba616350a4a5 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x86, 0x94, 0xd6, 0xa5, 0x6c, 0x70, 0xa9, 0x88, 0x0a, 0x87, 0xa8, 0x48, 0x28,
	0x08, 0x91, 0x43, 0x10, 0x12, 0xe2, 0x86, 0xf8, 0xd7, 0x43, 0x91, 0xaa, 0xb4, 0x9c, 0x4d, 0x48,
	0x9d, 0xc4, 0xc2, 0xf5, 0x6c, 0x27, 0xc9, 0x66, 0xf3, 0xa6, 0x3c, 0x08, 0x0f, 0x80, 0x3c, 0xf1,
	0x02, 0xbb, 0xec, 0x29, 0xd1, 0x37, 0xbf, 0xcf, 0x33, 0xf3, 0x0d, 0x3b, 0x69, 0xa0, 0xed, 0x2a,
	0x94, 0xb7, 0xbd, 0x34, 0xe5, 0xd8, 0x96, 0x85, 0x56, 0xa6, 0x4e, 0xe7, 0x08, 0x1d, 0xf0, 0x80,
	0x3e, 0xa7, 0xbf, 0x7c, 0x76, 0x7c, 0x06, 0x6d, 0xf7, 0x09, 0xe5, 0xed, 0xe5, 0x04, 0xbc, 0x07,
	0x53, 0xa9, 0x9a, 0x7f, 0x63, 0xdc, 0x5a, 0x85, 0xb5, 0x89, 0x39, 0x68, 0x55, 0x2a, 0xd9, 0x46,
	0x5e, 0xec, 0x27, 0x07, 0x59, 0x36, 0x3d, 0x92, 0x6e, 0x75, 0xa6, 0x2b, 0xe5, 0xc2, 0x99, 0x3e,
	0x9a, 0x0e, 0xc7, 0x3c, 0xac, 0x36, 0x64, 0xfe, 0x8c, 0xcd, 0xa6, 0x0e, 0xd2, 0xb4, 0x80, 0xa2,
	0x47, 0x1d, 0xed, 0xc4, 0x5e, 0xb2, 0x9f, 0x1f, 0x12, 0x4a, 0xea, 0x57, 0xd4, 0x3c, 0x65, 0x47,
	0x6e, 0x76, 0x41, 0x3c, 0x0d, 0x33, 0x46, 0x3e, 0xb1, 0x0f, 0x5c, 0xc9, 0x36, 0xa5, 0x97, 0x47,
	0xfe, 0x82, 0x71, 0x0d, 0x83, 0xe8, 0x40, 0x34, 0xaa, 0x6e, 0x68, 0x01, 0x89, 0xd1, 0x1d, 0xc2,
	0x67, 0x1a, 0x86, 0x2b, 0x38, 0x53, 0x75, 0x73, 0x49, 0xb2, 0x85, 0x89, 0xea, 0x40, 0x58, 0x93,
	0x83, 0x83, 0x09, 0xb6, 0x95, 0x2b, 0x38, 0x87, 0xc1, 0xc1, 0x2f, 0xd9, 0x91, 0x85, 0xfe, 0xe4,
	0x22, 0x51, 0x5c, 0xf7, 0x18, 0xed, 0xc6, 0x5e, 0x12, 0xe4, 0xa1, 0x86, 0x61, 0xb5, 0xba, 0xc4,
	0x0f, 0x3d, 0xf2, 0xd7, 0xec, 0x91, 0x34, 0x15, 0x60, 0x29, 0xc5, 0x86, 0x2d, 0xba, 0x1b, 0x7b,
	0xc9, 0x5e, 0xfe, 0xd0, 0x95, 0xcf, 0xff, 0x75, 0x9e, 0x14, 0xec, 0x78, 0x6b, 0x84, 0x3c, 0x64,
	0xfe, 0x0f, 0x39, 0x46, 0x1e, 0x0d, 0x67, 0x7f, 0x79, 0xc6, 0x82, 0x45, 0xa1, 0x7b, 0x49, 0xc1,
	0x1d, 0x64, 0x4f, 0xb6, 0xdf, 0x65, 0xca, 0x25, 0x9f, 0xd0, 0xb7, 0x3b, 0x6f, 0xbc, 0xd3, 0x9f,
	0xde, 0x7f, 0x67, 0x77, 0xe1, 0x3d, 0x67, 0xe1, 0x2a, 0xec, 0x1a, 0x16, 0x12, 0x0d, 0xa0, 0x6b,
	0x38, 0x73, 0xfa, 0x67, 0x27, 0xf3, 0xe4, 0x2f, 0x7a, 0xa3, 0x0c, 0xad, 0xe7, 0x0e, 0x78, 0xdf,
	0xe9, 0x5f, 0x94, 0xb1, 0x1d, 0xd6, 0xc8, 0x62, 0x39, 0x91, 0xfe, 0x3a, 0x59, 0x2c, 0x89, 0x7c,
	0xcc, 0xf6, 0x0d, 0x5c, 0x4b, 0x51, 0x60, 0xd9, 0xb8, 0x93, 0xed, 0x59, 0xe1, 0x1d, 0x96, 0x0d,
	0x7f, 0xca, 0x0e, 0xa9, 0x38, 0xd7, 0x45, 0x57, 0x01, 0xde, 0xb8, 0x33, 0xdd, 0xb3, 0xe2, 0x85,
	0xd3, 0xbe, 0xef, 0x52, 0x04, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x70, 0x05, 0xd3, 0xed,
	0xfd, 0x02, 0x00, 0x00,
}
