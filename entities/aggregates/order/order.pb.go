// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/order.proto

package order

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

type OrderStatus int32

const (
	OrderStatus_PENDING     OrderStatus = 0
	OrderStatus_APPROVED    OrderStatus = 1
	OrderStatus_DISAPPROVED OrderStatus = 2
	OrderStatus_RENTED      OrderStatus = 4
	OrderStatus_DONE        OrderStatus = 5
)

var OrderStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "DISAPPROVED",
	4: "RENTED",
	5: "DONE",
}

var OrderStatus_value = map[string]int32{
	"PENDING":     0,
	"APPROVED":    1,
	"DISAPPROVED": 2,
	"RENTED":      4,
	"DONE":        5,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a6b118f4f9db3fd, []int{0}
}

type Order struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CartUuid             string               `protobuf:"bytes,3,opt,name=cart_uuid,json=cartUuid,proto3" json:"cart_uuid,omitempty"`
	OrderNum             string               `protobuf:"bytes,4,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	AddressUuid          string               `protobuf:"bytes,5,opt,name=address_uuid,json=addressUuid,proto3" json:"address_uuid,omitempty"`
	Status               OrderStatus          `protobuf:"varint,6,opt,name=status,proto3,enum=aggregates.order.OrderStatus" json:"status,omitempty"`
	ShipmentFee          int32                `protobuf:"varint,7,opt,name=shipment_fee,json=shipmentFee,proto3" json:"shipment_fee,omitempty"`
	Delivered            bool                 `protobuf:"varint,8,opt,name=delivered,proto3" json:"delivered,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	TenantUuid           string               `protobuf:"bytes,11,opt,name=tenant_uuid,json=tenantUuid,proto3" json:"tenant_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6b118f4f9db3fd, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Order) GetCartUuid() string {
	if m != nil {
		return m.CartUuid
	}
	return ""
}

func (m *Order) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *Order) GetAddressUuid() string {
	if m != nil {
		return m.AddressUuid
	}
	return ""
}

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_PENDING
}

func (m *Order) GetShipmentFee() int32 {
	if m != nil {
		return m.ShipmentFee
	}
	return 0
}

func (m *Order) GetDelivered() bool {
	if m != nil {
		return m.Delivered
	}
	return false
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetTenantUuid() string {
	if m != nil {
		return m.TenantUuid
	}
	return ""
}

func init() {
	proto.RegisterEnum("aggregates.order.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*Order)(nil), "aggregates.order.Order")
}

func init() { proto.RegisterFile("aggregates/order/order.proto", fileDescriptor_2a6b118f4f9db3fd) }

var fileDescriptor_2a6b118f4f9db3fd = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6f, 0x9b, 0x30,
	0x14, 0xc6, 0x07, 0x4d, 0x52, 0x78, 0x54, 0x1d, 0xf2, 0x09, 0x75, 0x9d, 0xca, 0x76, 0x42, 0x93,
	0x0a, 0x52, 0xa7, 0x1d, 0x76, 0xec, 0x04, 0xdd, 0x7a, 0x21, 0x11, 0xed, 0x76, 0xd8, 0x25, 0x72,
	0xe2, 0x57, 0x62, 0x35, 0x86, 0xc8, 0x7e, 0xee, 0x5f, 0xbb, 0x3f, 0x66, 0x8a, 0x21, 0xeb, 0x94,
	0x0b, 0x42, 0xbf, 0xef, 0x7b, 0xf6, 0xfb, 0x3e, 0x19, 0x2e, 0x79, 0xdb, 0x6a, 0x6c, 0x39, 0xa1,
	0x29, 0x7a, 0x2d, 0x50, 0x0f, 0xdf, 0x7c, 0xa7, 0x7b, 0xea, 0x59, 0xfc, 0xaa, 0xe6, 0x8e, 0x5f,
	0x5c, 0xb5, 0x7d, 0xdf, 0x6e, 0xb1, 0x70, 0xfa, 0xca, 0x3e, 0x15, 0x24, 0x15, 0x1a, 0xe2, 0x6a,
	0x37, 0x8c, 0x7c, 0xfc, 0xe3, 0xc3, 0x74, 0xbe, 0xb7, 0xb2, 0x73, 0xf0, 0xa5, 0x48, 0xbc, 0xd4,
	0xcb, 0x4e, 0x1a, 0x5f, 0x0a, 0xc6, 0x60, 0x62, 0xad, 0x14, 0x89, 0x9f, 0x7a, 0x59, 0xd8, 0xb8,
	0x7f, 0xf6, 0x0e, 0xc2, 0x35, 0xd7, 0xb4, 0x74, 0xc2, 0x89, 0x13, 0x82, 0x3d, 0xf8, 0x39, 0x8a,
	0xee, 0xd2, 0x65, 0x67, 0x55, 0x32, 0x19, 0x44, 0x07, 0x6a, 0xab, 0xd8, 0x07, 0x38, 0xe3, 0x42,
	0x68, 0x34, 0x66, 0x18, 0x9e, 0x3a, 0x3d, 0x1a, 0x99, 0x9b, 0xff, 0x02, 0x33, 0x43, 0x9c, 0xac,
	0x49, 0x66, 0xa9, 0x97, 0x9d, 0xdf, 0xbc, 0xcf, 0x8f, 0xe3, 0xe4, 0x6e, 0xd3, 0x07, 0x67, 0x6a,
	0x46, 0xf3, 0xfe, 0x64, 0xb3, 0x91, 0x3b, 0x85, 0x1d, 0x2d, 0x9f, 0x10, 0x93, 0xd3, 0xd4, 0xcb,
	0xa6, 0x4d, 0x74, 0x60, 0x77, 0x88, 0xec, 0x12, 0x42, 0x81, 0x5b, 0xf9, 0x82, 0x1a, 0x45, 0x12,
	0xa4, 0x5e, 0x16, 0x34, 0xaf, 0x80, 0x7d, 0x05, 0x58, 0x6b, 0xe4, 0x84, 0x62, 0xc9, 0x29, 0x09,
	0x53, 0x2f, 0x8b, 0x6e, 0x2e, 0xf2, 0xa1, 0xb8, 0xfc, 0x50, 0x5c, 0xfe, 0x78, 0x28, 0xae, 0x09,
	0x47, 0xf7, 0x2d, 0xb1, 0x2b, 0x88, 0x08, 0x3b, 0xde, 0x8d, 0x8d, 0x44, 0x2e, 0x14, 0x0c, 0x68,
	0x9f, 0xe9, 0xd3, 0x1c, 0xa2, 0xff, 0x76, 0x66, 0x11, 0x9c, 0x2e, 0xaa, 0xba, 0xbc, 0xaf, 0xbf,
	0xc7, 0x6f, 0xd8, 0x19, 0x04, 0xb7, 0x8b, 0x45, 0x33, 0xff, 0x55, 0x95, 0xb1, 0xc7, 0xde, 0x42,
	0x54, 0xde, 0x3f, 0xfc, 0x03, 0x3e, 0x03, 0x98, 0x35, 0x55, 0xfd, 0x58, 0x95, 0xf1, 0x84, 0x05,
	0x30, 0x29, 0xe7, 0x75, 0x15, 0x4f, 0xbf, 0xfd, 0xf8, 0x7d, 0xd7, 0x4a, 0xda, 0xd8, 0x55, 0xbe,
	0xee, 0x55, 0x21, 0xb7, 0x1b, 0xae, 0xd4, 0x46, 0x88, 0xe2, 0xd9, 0x0a, 0xfe, 0x2c, 0xaf, 0x55,
	0x6f, 0x3b, 0xe2, 0xb2, 0xbb, 0x36, 0xa8, 0x5f, 0xe4, 0x1a, 0x0b, 0xec, 0x48, 0x92, 0x44, 0x53,
	0x1c, 0x3f, 0x9c, 0xd5, 0xcc, 0x45, 0xfb, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x16, 0x7c, 0xd8,
	0xb3, 0x53, 0x02, 0x00, 0x00,
}
