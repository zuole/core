// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/DOSNetwork/core/share/dkg/pedersen/dkg.proto

package dkg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pedersen "github.com/DOSNetwork/core/share/vss/pedersen"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Deal struct {
	Index                uint32                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Deal                 *pedersen.EncryptedDeal `protobuf:"bytes,2,opt,name=deal,proto3" json:"deal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7eac839653a0a37, []int{0}
}
func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (dst *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(dst, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Deal) GetDeal() *pedersen.EncryptedDeal {
	if m != nil {
		return m.Deal
	}
	return nil
}

type Response struct {
	Index                uint32             `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Response             *pedersen.Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7eac839653a0a37, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response) GetResponse() *pedersen.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Deal)(nil), "dkg.Deal")
	proto.RegisterType((*Response)(nil), "dkg.Response")
}

func init() {
	proto.RegisterFile("github.com/DOSNetwork/core/share/dkg/pedersen/dkg.proto", fileDescriptor_b7eac839653a0a37)
}

var fileDescriptor_b7eac839653a0a37 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0xf1, 0x0f, 0xf6, 0x4b, 0x2d, 0x29, 0xcf, 0x2f,
	0xca, 0xd6, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0xd5, 0x4f, 0xc9, 0x4e,
	0xd7, 0x2f, 0x48, 0x4d, 0x49, 0x2d, 0x2a, 0x4e, 0xcd, 0x03, 0x71, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x98, 0x53, 0xb2, 0xd3, 0xa5, 0x08, 0xeb, 0x2e, 0x2b, 0x2e, 0x46, 0xe8, 0x2e, 0x2b,
	0x2e, 0x86, 0xe8, 0x56, 0x72, 0xe1, 0x62, 0x71, 0x49, 0x4d, 0xcc, 0x11, 0x12, 0xe1, 0x62, 0xcd,
	0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x70, 0x84, 0xd4, 0xb8,
	0x58, 0x52, 0x52, 0x13, 0x73, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf4, 0x40, 0xfa,
	0x5c, 0xf3, 0x92, 0x8b, 0x2a, 0x0b, 0x4a, 0x52, 0x53, 0x40, 0xfa, 0x82, 0xc0, 0xf2, 0x4a, 0xde,
	0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x38, 0x4c, 0xd2, 0xe4, 0xe2, 0x28,
	0x82, 0xaa, 0x80, 0x9a, 0xc6, 0x0b, 0x36, 0x0d, 0xa6, 0x2d, 0x08, 0x2e, 0x9d, 0xc4, 0x06, 0x76,
	0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x79, 0x17, 0x8b, 0x12, 0x01, 0x00, 0x00,
}
