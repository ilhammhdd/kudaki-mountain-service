// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/cart.proto

package rental

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

type Cart struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalPrice           int32                `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TotalItems           int32                `protobuf:"varint,5,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Open                 bool                 `protobuf:"varint,6,opt,name=open,proto3" json:"open,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef36b31966c981a, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cart) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Cart) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Cart) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Cart) GetTotalItems() int32 {
	if m != nil {
		return m.TotalItems
	}
	return 0
}

func (m *Cart) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *Cart) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Cart)(nil), "aggregates.rental.Cart")
}

func init() { proto.RegisterFile("aggregates/rental/cart.proto", fileDescriptor_7ef36b31966c981a) }

var fileDescriptor_7ef36b31966c981a = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x49, 0xbf, 0x68, 0x53, 0x78, 0xe1, 0x9d, 0xd5, 0x50, 0x85, 0x0e, 0xae, 0x66, 0xd3,
	0x09, 0xe8, 0xca, 0xa5, 0xba, 0x90, 0xee, 0x64, 0xd0, 0x8d, 0x9b, 0x72, 0x3b, 0xb9, 0xa6, 0x97,
	0x4e, 0x26, 0x43, 0x72, 0xe3, 0x6f, 0xf5, 0xe7, 0xc8, 0xa4, 0x96, 0x0a, 0xee, 0x2e, 0xe7, 0x79,
	0xc2, 0xe1, 0x44, 0x5e, 0x83, 0x31, 0x1e, 0x0d, 0x30, 0x06, 0xe5, 0xb1, 0x63, 0x68, 0x55, 0x03,
	0x9e, 0xab, 0xde, 0x3b, 0x76, 0xd9, 0xff, 0x0b, 0xad, 0x4e, 0x74, 0xb5, 0x36, 0xce, 0x99, 0x16,
	0x55, 0x12, 0xf6, 0xf1, 0x43, 0x31, 0x59, 0x0c, 0x0c, 0xb6, 0x3f, 0xbd, 0xb9, 0xf9, 0x12, 0x72,
	0xf2, 0x04, 0x9e, 0xb3, 0x7f, 0x72, 0x44, 0x3a, 0x17, 0x85, 0x28, 0xc7, 0xf5, 0x88, 0x74, 0x96,
	0xc9, 0x49, 0x8c, 0xa4, 0xf3, 0x51, 0x21, 0xca, 0x45, 0x9d, 0xee, 0xec, 0x4a, 0x2e, 0x62, 0x40,
	0xbf, 0x4b, 0x60, 0x9c, 0xc0, 0x7c, 0x08, 0xde, 0x06, 0xb8, 0x96, 0x4b, 0x76, 0x0c, 0xed, 0xae,
	0xf7, 0xd4, 0x60, 0x3e, 0x29, 0x44, 0x39, 0xad, 0x65, 0x8a, 0x5e, 0x86, 0xe4, 0x22, 0x10, 0xa3,
	0x0d, 0xf9, 0xf4, 0x97, 0xb0, 0x1d, 0x92, 0xa1, 0xd2, 0xf5, 0xd8, 0xe5, 0xb3, 0x42, 0x94, 0xf3,
	0x3a, 0xdd, 0xd9, 0xbd, 0x94, 0x8d, 0x47, 0x60, 0xd4, 0x3b, 0xe0, 0x7c, 0x5e, 0x88, 0x72, 0x79,
	0xbb, 0xaa, 0x4e, 0xab, 0xaa, 0xf3, 0xaa, 0xea, 0xf5, 0xbc, 0xaa, 0x5e, 0xfc, 0xd8, 0x0f, 0xfc,
	0xb8, 0x7d, 0x7f, 0x36, 0xc4, 0x87, 0xb8, 0xaf, 0x1a, 0x67, 0x15, 0xb5, 0x07, 0xb0, 0xf6, 0xa0,
	0xb5, 0x3a, 0x46, 0x0d, 0x47, 0xda, 0x58, 0x17, 0x3b, 0x06, 0xea, 0x36, 0x01, 0xfd, 0x27, 0x35,
	0xa8, 0xb0, 0x63, 0x62, 0xc2, 0xa0, 0xfe, 0x7c, 0xf2, 0x7e, 0x96, 0x9a, 0xee, 0xbe, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x56, 0xbf, 0x91, 0x5b, 0x80, 0x01, 0x00, 0x00,
}