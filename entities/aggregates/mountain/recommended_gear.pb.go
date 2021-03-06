// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/recommended_gear.proto

package mountain

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

type RecommendedGear struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Mountain             *Mountain            `protobuf:"bytes,4,opt,name=mountain,proto3" json:"mountain,omitempty"`
	Upvote               int32                `protobuf:"varint,5,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote             int32                `protobuf:"varint,6,opt,name=downvote,proto3" json:"downvote,omitempty"`
	Seen                 int32                `protobuf:"varint,7,opt,name=seen,proto3" json:"seen,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecommendedGear) Reset()         { *m = RecommendedGear{} }
func (m *RecommendedGear) String() string { return proto.CompactTextString(m) }
func (*RecommendedGear) ProtoMessage()    {}
func (*RecommendedGear) Descriptor() ([]byte, []int) {
	return fileDescriptor_50de5b97abeab0b3, []int{0}
}

func (m *RecommendedGear) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendedGear.Unmarshal(m, b)
}
func (m *RecommendedGear) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendedGear.Marshal(b, m, deterministic)
}
func (m *RecommendedGear) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendedGear.Merge(m, src)
}
func (m *RecommendedGear) XXX_Size() int {
	return xxx_messageInfo_RecommendedGear.Size(m)
}
func (m *RecommendedGear) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendedGear.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendedGear proto.InternalMessageInfo

func (m *RecommendedGear) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecommendedGear) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RecommendedGear) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *RecommendedGear) GetMountain() *Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *RecommendedGear) GetUpvote() int32 {
	if m != nil {
		return m.Upvote
	}
	return 0
}

func (m *RecommendedGear) GetDownvote() int32 {
	if m != nil {
		return m.Downvote
	}
	return 0
}

func (m *RecommendedGear) GetSeen() int32 {
	if m != nil {
		return m.Seen
	}
	return 0
}

func (m *RecommendedGear) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendedGear)(nil), "aggregates.mountain.RecommendedGear")
}

func init() {
	proto.RegisterFile("aggregates/mountain/recommended_gear.proto", fileDescriptor_50de5b97abeab0b3)
}

var fileDescriptor_50de5b97abeab0b3 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0x3b, 0x31,
	0x14, 0x64, 0xb7, 0x7f, 0x7e, 0xdb, 0xfc, 0x40, 0x21, 0x82, 0x2c, 0x2b, 0xe2, 0xd2, 0xd3, 0x22,
	0x34, 0x01, 0x3d, 0xf5, 0xa8, 0x17, 0x11, 0xf1, 0xb2, 0xe8, 0xc5, 0x4b, 0x49, 0x37, 0xcf, 0x34,
	0xb4, 0xd9, 0x94, 0xec, 0x4b, 0xfd, 0x1a, 0x7e, 0x64, 0x69, 0xda, 0xb4, 0x97, 0xbd, 0xcd, 0xbc,
	0x99, 0xe4, 0x4d, 0x26, 0xe4, 0x5e, 0x28, 0xe5, 0x40, 0x09, 0x84, 0x8e, 0x1b, 0xeb, 0x5b, 0x14,
	0xba, 0xe5, 0x0e, 0x1a, 0x6b, 0x0c, 0xb4, 0x12, 0xe4, 0x42, 0x81, 0x70, 0x6c, 0xeb, 0x2c, 0x5a,
	0x7a, 0x75, 0xf6, 0xb2, 0xe8, 0x2d, 0xa6, 0x7d, 0x17, 0x44, 0x70, 0x38, 0x58, 0xdc, 0x29, 0x6b,
	0xd5, 0x06, 0x78, 0x60, 0x4b, 0xff, 0xcd, 0x51, 0x1b, 0xe8, 0x50, 0x98, 0xed, 0xc1, 0x30, 0xfd,
	0x4d, 0xc9, 0x65, 0x7d, 0x5e, 0xfa, 0x02, 0xc2, 0xd1, 0x0b, 0x92, 0x6a, 0x99, 0x27, 0x65, 0x52,
	0x0d, 0xea, 0x54, 0x4b, 0x4a, 0xc9, 0xd0, 0x7b, 0x2d, 0xf3, 0xb4, 0x4c, 0xaa, 0x49, 0x1d, 0x30,
	0xbd, 0x21, 0x13, 0xdf, 0x81, 0x5b, 0x04, 0x61, 0x10, 0x84, 0x6c, 0x3f, 0xf8, 0xdc, 0x8b, 0x73,
	0x92, 0xc5, 0x1c, 0xf9, 0xb0, 0x4c, 0xaa, 0xff, 0x0f, 0xb7, 0xac, 0xe7, 0x05, 0xec, 0xfd, 0x08,
	0xea, 0x93, 0x9d, 0x5e, 0x93, 0xb1, 0xdf, 0xee, 0x2c, 0x42, 0x3e, 0x2a, 0x93, 0x6a, 0x54, 0x1f,
	0x19, 0x2d, 0x48, 0x26, 0xed, 0x4f, 0x1b, 0x94, 0x71, 0x50, 0x4e, 0x7c, 0x9f, 0xaf, 0x03, 0x68,
	0xf3, 0x7f, 0x61, 0x1e, 0x30, 0x9d, 0x13, 0xd2, 0x38, 0x10, 0x08, 0x72, 0x21, 0x30, 0xcf, 0x42,
	0x88, 0x82, 0x1d, 0xda, 0x60, 0xb1, 0x0d, 0xf6, 0x11, 0xdb, 0xa8, 0x27, 0x47, 0xf7, 0x13, 0x3e,
	0xbf, 0x7d, 0xbd, 0x2a, 0x8d, 0x2b, 0xbf, 0x64, 0x8d, 0x35, 0x5c, 0x6f, 0x56, 0xc2, 0x98, 0x95,
	0x94, 0x7c, 0xed, 0xa5, 0x58, 0xeb, 0x59, 0x8c, 0x3a, 0xeb, 0xc0, 0xed, 0x74, 0x03, 0x1c, 0x5a,
	0xd4, 0xa8, 0xa1, 0xe3, 0x3d, 0xff, 0xb1, 0x1c, 0x87, 0x5d, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xcb, 0xf9, 0x89, 0xee, 0x01, 0x00, 0x00,
}
