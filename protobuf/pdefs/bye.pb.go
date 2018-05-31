// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/pdefs/bye.proto

package pdefs

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

type GoodBye struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodBye) Reset()         { *m = GoodBye{} }
func (m *GoodBye) String() string { return proto.CompactTextString(m) }
func (*GoodBye) ProtoMessage()    {}
func (*GoodBye) Descriptor() ([]byte, []int) {
	return fileDescriptor_bye_6654826d0b404a6d, []int{0}
}
func (m *GoodBye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodBye.Unmarshal(m, b)
}
func (m *GoodBye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodBye.Marshal(b, m, deterministic)
}
func (dst *GoodBye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodBye.Merge(dst, src)
}
func (m *GoodBye) XXX_Size() int {
	return xxx_messageInfo_GoodBye.Size(m)
}
func (m *GoodBye) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodBye.DiscardUnknown(m)
}

var xxx_messageInfo_GoodBye proto.InternalMessageInfo

func (m *GoodBye) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GoodBye) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GoodBye)(nil), "pdefs.GoodBye")
}

func init() { proto.RegisterFile("protobuf/pdefs/bye.proto", fileDescriptor_bye_6654826d0b404a6d) }

var fileDescriptor_bye_6654826d0b404a6d = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x48, 0x49, 0x4d, 0x2b, 0xd6, 0x4f, 0xaa, 0x4c, 0xd5, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x05, 0x94, 0xac, 0xb9, 0xd8, 0xdd, 0xf3, 0xf3, 0x53, 0x9c, 0x2a, 0x53,
	0x85, 0xc4, 0xb8, 0xd8, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xa0, 0x3c, 0x21, 0x09, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26,
	0xb0, 0x04, 0x8c, 0xeb, 0xc4, 0x1e, 0x05, 0x31, 0x25, 0x89, 0x0d, 0x6c, 0xa6, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x70, 0xbf, 0xed, 0x3c, 0x6f, 0x00, 0x00, 0x00,
}
