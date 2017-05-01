// Code generated by protoc-gen-go.
// source: hlc.proto
// DO NOT EDIT!

/*
Package hlc is a generated protocol buffer package.

It is generated from these files:
	hlc.proto

It has these top-level messages:
	HybridLogicalClock
*/
package hlc

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

type HybridLogicalClock struct {
	LocalTime  int64 `protobuf:"varint,1,opt,name=local_time,json=localTime" json:"local_time,omitempty"`
	CausalTime int64 `protobuf:"varint,2,opt,name=causal_time,json=causalTime" json:"causal_time,omitempty"`
}

func (m *HybridLogicalClock) Reset()                    { *m = HybridLogicalClock{} }
func (m *HybridLogicalClock) String() string            { return proto.CompactTextString(m) }
func (*HybridLogicalClock) ProtoMessage()               {}
func (*HybridLogicalClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HybridLogicalClock) GetLocalTime() int64 {
	if m != nil {
		return m.LocalTime
	}
	return 0
}

func (m *HybridLogicalClock) GetCausalTime() int64 {
	if m != nil {
		return m.CausalTime
	}
	return 0
}

func init() {
	proto.RegisterType((*HybridLogicalClock)(nil), "hlc.HybridLogicalClock")
}

func init() { proto.RegisterFile("hlc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xc8, 0x49, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xc8, 0x49, 0x56, 0x0a, 0xe1, 0x12, 0xf2, 0xa8,
	0x4c, 0x2a, 0xca, 0x4c, 0xf1, 0xc9, 0x4f, 0xcf, 0x4c, 0x4e, 0xcc, 0x71, 0xce, 0xc9, 0x4f, 0xce,
	0x16, 0x92, 0xe5, 0xe2, 0xca, 0xc9, 0x4f, 0x4e, 0xcc, 0x89, 0x2f, 0xc9, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0e, 0xe2, 0x04, 0x8b, 0x84, 0x64, 0xe6, 0xa6, 0x0a, 0xc9, 0x73, 0x71,
	0x27, 0x27, 0x96, 0x16, 0xc3, 0xe4, 0x99, 0xc0, 0xf2, 0x5c, 0x10, 0x21, 0x90, 0x82, 0x24, 0x36,
	0xb0, 0x0d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x83, 0xe9, 0x2a, 0x6e, 0x00, 0x00,
	0x00,
}
