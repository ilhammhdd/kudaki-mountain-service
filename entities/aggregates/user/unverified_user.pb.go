// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/unverified_user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UnverifiedUser struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UnverifiedUser) Reset()         { *m = UnverifiedUser{} }
func (m *UnverifiedUser) String() string { return proto.CompactTextString(m) }
func (*UnverifiedUser) ProtoMessage()    {}
func (*UnverifiedUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66d3b9a6d578134, []int{0}
}

func (m *UnverifiedUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnverifiedUser.Unmarshal(m, b)
}
func (m *UnverifiedUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnverifiedUser.Marshal(b, m, deterministic)
}
func (m *UnverifiedUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnverifiedUser.Merge(m, src)
}
func (m *UnverifiedUser) XXX_Size() int {
	return xxx_messageInfo_UnverifiedUser.Size(m)
}
func (m *UnverifiedUser) XXX_DiscardUnknown() {
	xxx_messageInfo_UnverifiedUser.DiscardUnknown(m)
}

var xxx_messageInfo_UnverifiedUser proto.InternalMessageInfo

func (m *UnverifiedUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UnverifiedUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UnverifiedUser) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UnverifiedUser)(nil), "aggregates.user.UnverifiedUser")
}

func init() {
	proto.RegisterFile("aggregates/user/unverified_user.proto", fileDescriptor_d66d3b9a6d578134)
}

var fileDescriptor_d66d3b9a6d578134 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x2a, 0x82, 0x2b, 0x54, 0x08, 0x08, 0x21, 0x17, 0x8b, 0x20, 0xd4, 0x43, 0x77,
	0x40, 0x4f, 0x1e, 0x15, 0xc4, 0x7b, 0xb0, 0x17, 0x2f, 0x65, 0x92, 0x9d, 0x6e, 0x86, 0x76, 0xb3,
	0x65, 0x77, 0xb6, 0x3f, 0xc1, 0xdf, 0x2d, 0x49, 0x6d, 0x85, 0x7a, 0x7c, 0xcb, 0xf7, 0xf6, 0x7d,
	0xa3, 0x1e, 0xd0, 0xda, 0x40, 0x16, 0x85, 0x22, 0xa4, 0x48, 0x01, 0x52, 0xbf, 0xa7, 0xc0, 0x6b,
	0x26, 0xb3, 0x1a, 0xb2, 0xde, 0x05, 0x2f, 0xbe, 0xb8, 0xf9, 0xc3, 0xf4, 0xf0, 0x5c, 0x55, 0xff,
	0x7a, 0x27, 0xb8, 0xba, 0xb3, 0xde, 0xdb, 0x2d, 0xc1, 0x98, 0x9a, 0xb4, 0x06, 0x61, 0x47, 0x51,
	0xd0, 0xed, 0x0e, 0xc0, 0xfd, 0x77, 0xa6, 0xa6, 0xcb, 0xd3, 0xce, 0x32, 0x52, 0x28, 0xa6, 0x2a,
	0x67, 0x53, 0x66, 0xb3, 0x6c, 0x3e, 0xa9, 0x73, 0x36, 0xc5, 0xa3, 0xba, 0x18, 0x7e, 0x2c, 0xf3,
	0x59, 0x36, 0xbf, 0x7e, 0xba, 0xd5, 0x67, 0xfb, 0x7a, 0x28, 0xd5, 0x23, 0x52, 0xbc, 0x28, 0xd5,
	0x06, 0x42, 0x21, 0xb3, 0x42, 0x29, 0x27, 0x63, 0xa1, 0xd2, 0x07, 0x07, 0x7d, 0x74, 0xd0, 0x9f,
	0x47, 0x87, 0xfa, 0xea, 0x97, 0x7e, 0x95, 0xb7, 0x8f, 0xaf, 0x77, 0xcb, 0xd2, 0xa5, 0x46, 0xb7,
	0xde, 0x01, 0x6f, 0x3b, 0x74, 0xae, 0x33, 0x06, 0x36, 0xc9, 0xe0, 0x86, 0x17, 0xce, 0xa7, 0x5e,
	0x90, 0xfb, 0x45, 0xa4, 0xb0, 0xe7, 0x96, 0x80, 0x7a, 0x61, 0x61, 0x8a, 0x70, 0x76, 0x7d, 0x73,
	0x39, 0xee, 0x3c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x17, 0xd5, 0x8e, 0xbe, 0x4f, 0x01, 0x00,
	0x00,
}
