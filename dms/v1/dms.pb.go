// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dms/v1/dms.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	stubs "github.com/oswee/stubs"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateDeliveryOrderRequest struct {
	DeliveryOrder        *DeliveryOrder `protobuf:"bytes,1,opt,name=delivery_order,json=deliveryOrder,proto3" json:"delivery_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateDeliveryOrderRequest) Reset()         { *m = CreateDeliveryOrderRequest{} }
func (m *CreateDeliveryOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeliveryOrderRequest) ProtoMessage()    {}
func (*CreateDeliveryOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{0}
}

func (m *CreateDeliveryOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeliveryOrderRequest.Unmarshal(m, b)
}
func (m *CreateDeliveryOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeliveryOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateDeliveryOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeliveryOrderRequest.Merge(m, src)
}
func (m *CreateDeliveryOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeliveryOrderRequest.Size(m)
}
func (m *CreateDeliveryOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeliveryOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeliveryOrderRequest proto.InternalMessageInfo

func (m *CreateDeliveryOrderRequest) GetDeliveryOrder() *DeliveryOrder {
	if m != nil {
		return m.DeliveryOrder
	}
	return nil
}

type ListDeliveryOrdersRequest struct {
	StakeholderId        int64    `protobuf:"varint,1,opt,name=stakeholder_id,json=stakeholderId,proto3" json:"stakeholder_id,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeliveryOrdersRequest) Reset()         { *m = ListDeliveryOrdersRequest{} }
func (m *ListDeliveryOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeliveryOrdersRequest) ProtoMessage()    {}
func (*ListDeliveryOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{1}
}

func (m *ListDeliveryOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeliveryOrdersRequest.Unmarshal(m, b)
}
func (m *ListDeliveryOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeliveryOrdersRequest.Marshal(b, m, deterministic)
}
func (m *ListDeliveryOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeliveryOrdersRequest.Merge(m, src)
}
func (m *ListDeliveryOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeliveryOrdersRequest.Size(m)
}
func (m *ListDeliveryOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeliveryOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeliveryOrdersRequest proto.InternalMessageInfo

func (m *ListDeliveryOrdersRequest) GetStakeholderId() int64 {
	if m != nil {
		return m.StakeholderId
	}
	return 0
}

func (m *ListDeliveryOrdersRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *ListDeliveryOrdersRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type ListDeliveryOrdersResponse struct {
	DeliveryOrders       []*DeliveryOrder `protobuf:"bytes,1,rep,name=delivery_orders,json=deliveryOrders,proto3" json:"delivery_orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListDeliveryOrdersResponse) Reset()         { *m = ListDeliveryOrdersResponse{} }
func (m *ListDeliveryOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeliveryOrdersResponse) ProtoMessage()    {}
func (*ListDeliveryOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{2}
}

func (m *ListDeliveryOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeliveryOrdersResponse.Unmarshal(m, b)
}
func (m *ListDeliveryOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeliveryOrdersResponse.Marshal(b, m, deterministic)
}
func (m *ListDeliveryOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeliveryOrdersResponse.Merge(m, src)
}
func (m *ListDeliveryOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeliveryOrdersResponse.Size(m)
}
func (m *ListDeliveryOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeliveryOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeliveryOrdersResponse proto.InternalMessageInfo

func (m *ListDeliveryOrdersResponse) GetDeliveryOrders() []*DeliveryOrder {
	if m != nil {
		return m.DeliveryOrders
	}
	return nil
}

type UpdateDeliveryOrderRequest struct {
	DeliveryOrder        *DeliveryOrder `protobuf:"bytes,1,opt,name=delivery_order,json=deliveryOrder,proto3" json:"delivery_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDeliveryOrderRequest) Reset()         { *m = UpdateDeliveryOrderRequest{} }
func (m *UpdateDeliveryOrderRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeliveryOrderRequest) ProtoMessage()    {}
func (*UpdateDeliveryOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{3}
}

func (m *UpdateDeliveryOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeliveryOrderRequest.Unmarshal(m, b)
}
func (m *UpdateDeliveryOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeliveryOrderRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDeliveryOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeliveryOrderRequest.Merge(m, src)
}
func (m *UpdateDeliveryOrderRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeliveryOrderRequest.Size(m)
}
func (m *UpdateDeliveryOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeliveryOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeliveryOrderRequest proto.InternalMessageInfo

func (m *UpdateDeliveryOrderRequest) GetDeliveryOrder() *DeliveryOrder {
	if m != nil {
		return m.DeliveryOrder
	}
	return nil
}

type DeleteDeliveryOrderRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeliveryOrderRequest) Reset()         { *m = DeleteDeliveryOrderRequest{} }
func (m *DeleteDeliveryOrderRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeliveryOrderRequest) ProtoMessage()    {}
func (*DeleteDeliveryOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{4}
}

func (m *DeleteDeliveryOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeliveryOrderRequest.Unmarshal(m, b)
}
func (m *DeleteDeliveryOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeliveryOrderRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeliveryOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeliveryOrderRequest.Merge(m, src)
}
func (m *DeleteDeliveryOrderRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeliveryOrderRequest.Size(m)
}
func (m *DeleteDeliveryOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeliveryOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeliveryOrderRequest proto.InternalMessageInfo

func (m *DeleteDeliveryOrderRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GeoCodeDeliveryOrderRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StakeholderId        int64    `protobuf:"varint,2,opt,name=stakeholder_id,json=stakeholderId,proto3" json:"stakeholder_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoCodeDeliveryOrderRequest) Reset()         { *m = GeoCodeDeliveryOrderRequest{} }
func (m *GeoCodeDeliveryOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GeoCodeDeliveryOrderRequest) ProtoMessage()    {}
func (*GeoCodeDeliveryOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{5}
}

func (m *GeoCodeDeliveryOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoCodeDeliveryOrderRequest.Unmarshal(m, b)
}
func (m *GeoCodeDeliveryOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoCodeDeliveryOrderRequest.Marshal(b, m, deterministic)
}
func (m *GeoCodeDeliveryOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoCodeDeliveryOrderRequest.Merge(m, src)
}
func (m *GeoCodeDeliveryOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GeoCodeDeliveryOrderRequest.Size(m)
}
func (m *GeoCodeDeliveryOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoCodeDeliveryOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeoCodeDeliveryOrderRequest proto.InternalMessageInfo

func (m *GeoCodeDeliveryOrderRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GeoCodeDeliveryOrderRequest) GetStakeholderId() int64 {
	if m != nil {
		return m.StakeholderId
	}
	return 0
}

func (m *GeoCodeDeliveryOrderRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type DeliveryOrder struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StakeholderId        int64    `protobuf:"varint,2,opt,name=stakeholder_id,json=stakeholderId,proto3" json:"stakeholder_id,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	DestinationZip       string   `protobuf:"bytes,5,opt,name=destination_zip,json=destinationZip,proto3" json:"destination_zip,omitempty"`
	DestinationLat       float64  `protobuf:"fixed64,6,opt,name=destination_lat,json=destinationLat,proto3" json:"destination_lat,omitempty"`
	DestinationLng       float64  `protobuf:"fixed64,7,opt,name=destination_lng,json=destinationLng,proto3" json:"destination_lng,omitempty"`
	TotalWeight          float64  `protobuf:"fixed64,8,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	RoutingSequence      int64    `protobuf:"varint,9,opt,name=routing_sequence,json=routingSequence,proto3" json:"routing_sequence,omitempty"`
	Uuid                 string   `protobuf:"bytes,10,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryOrder) Reset()         { *m = DeliveryOrder{} }
func (m *DeliveryOrder) String() string { return proto.CompactTextString(m) }
func (*DeliveryOrder) ProtoMessage()    {}
func (*DeliveryOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_91763f6e7a58dbd9, []int{6}
}

func (m *DeliveryOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryOrder.Unmarshal(m, b)
}
func (m *DeliveryOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryOrder.Marshal(b, m, deterministic)
}
func (m *DeliveryOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryOrder.Merge(m, src)
}
func (m *DeliveryOrder) XXX_Size() int {
	return xxx_messageInfo_DeliveryOrder.Size(m)
}
func (m *DeliveryOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryOrder.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryOrder proto.InternalMessageInfo

func (m *DeliveryOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeliveryOrder) GetStakeholderId() int64 {
	if m != nil {
		return m.StakeholderId
	}
	return 0
}

func (m *DeliveryOrder) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DeliveryOrder) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *DeliveryOrder) GetDestinationZip() string {
	if m != nil {
		return m.DestinationZip
	}
	return ""
}

func (m *DeliveryOrder) GetDestinationLat() float64 {
	if m != nil {
		return m.DestinationLat
	}
	return 0
}

func (m *DeliveryOrder) GetDestinationLng() float64 {
	if m != nil {
		return m.DestinationLng
	}
	return 0
}

func (m *DeliveryOrder) GetTotalWeight() float64 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

func (m *DeliveryOrder) GetRoutingSequence() int64 {
	if m != nil {
		return m.RoutingSequence
	}
	return 0
}

func (m *DeliveryOrder) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateDeliveryOrderRequest)(nil), "oswee.dms.v1.CreateDeliveryOrderRequest")
	proto.RegisterType((*ListDeliveryOrdersRequest)(nil), "oswee.dms.v1.ListDeliveryOrdersRequest")
	proto.RegisterType((*ListDeliveryOrdersResponse)(nil), "oswee.dms.v1.ListDeliveryOrdersResponse")
	proto.RegisterType((*UpdateDeliveryOrderRequest)(nil), "oswee.dms.v1.UpdateDeliveryOrderRequest")
	proto.RegisterType((*DeleteDeliveryOrderRequest)(nil), "oswee.dms.v1.DeleteDeliveryOrderRequest")
	proto.RegisterType((*GeoCodeDeliveryOrderRequest)(nil), "oswee.dms.v1.GeoCodeDeliveryOrderRequest")
	proto.RegisterType((*DeliveryOrder)(nil), "oswee.dms.v1.DeliveryOrder")
}

func init() { proto.RegisterFile("dms/v1/dms.proto", fileDescriptor_91763f6e7a58dbd9) }

var fileDescriptor_91763f6e7a58dbd9 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xce, 0xb6, 0xfc, 0xd8, 0x53, 0x5a, 0x70, 0x40, 0xb3, 0x2e, 0x18, 0x6a, 0x13, 0xa5, 0x24,
	0xa6, 0x1b, 0xf0, 0x09, 0x04, 0x8c, 0x31, 0x21, 0x8a, 0x25, 0x6a, 0x42, 0x8c, 0xeb, 0x94, 0x39,
	0x6c, 0x27, 0xee, 0xce, 0xac, 0x33, 0xb3, 0x25, 0x70, 0xe9, 0xad, 0x5e, 0xe9, 0xa3, 0xf9, 0x0a,
	0x3e, 0x81, 0x4f, 0x60, 0x76, 0xda, 0x42, 0x17, 0xb6, 0xc5, 0x98, 0x78, 0xd7, 0xfd, 0xce, 0x37,
	0xe7, 0x3b, 0xff, 0x85, 0x25, 0x16, 0x6b, 0xbf, 0xbf, 0xe5, 0xb3, 0x58, 0xb7, 0x13, 0x25, 0x8d,
	0x24, 0x0b, 0x52, 0x9f, 0x22, 0xb6, 0x33, 0xa0, 0xbf, 0xe5, 0xd5, 0x2d, 0xd8, 0x4d, 0x4f, 0x06,
	0x56, 0x6f, 0x2d, 0x94, 0x32, 0x8c, 0xd0, 0xa7, 0x09, 0xf7, 0xa9, 0x10, 0xd2, 0x50, 0xc3, 0xa5,
	0x18, 0xbe, 0x6d, 0x7e, 0x04, 0x6f, 0x57, 0x21, 0x35, 0xb8, 0x87, 0x11, 0xef, 0xa3, 0x3a, 0x7b,
	0xa5, 0x18, 0xaa, 0x0e, 0x7e, 0x4e, 0x51, 0x1b, 0xb2, 0x03, 0x75, 0x36, 0xc4, 0x03, 0x99, 0x19,
	0x5c, 0xa7, 0xe1, 0xb4, 0xaa, 0xdb, 0xab, 0xed, 0x71, 0xc9, 0x76, 0xfe, 0x6d, 0x8d, 0x8d, 0x7f,
	0x36, 0xbf, 0x3a, 0x70, 0x6f, 0x9f, 0x6b, 0x93, 0x23, 0xe9, 0x91, 0xc2, 0x43, 0xa8, 0x6b, 0x43,
	0x3f, 0x61, 0x4f, 0x46, 0x0c, 0x55, 0xc0, 0x99, 0x55, 0x28, 0x77, 0x6a, 0x63, 0xe8, 0x0b, 0x46,
	0xd6, 0xa1, 0x9a, 0xd0, 0x10, 0x03, 0x91, 0xc6, 0x5d, 0x54, 0x6e, 0xa9, 0xe1, 0xb4, 0x66, 0x3b,
	0x90, 0x41, 0x2f, 0x2d, 0x42, 0x1e, 0xc1, 0xa2, 0x42, 0x9d, 0x46, 0x26, 0x48, 0x50, 0x05, 0x99,
	0xc1, 0x2d, 0x5b, 0x52, 0x6d, 0x00, 0x1f, 0xa0, 0x3a, 0xa0, 0x21, 0x36, 0xbb, 0xe0, 0x15, 0x05,
	0xa3, 0x13, 0x29, 0x34, 0x92, 0x3d, 0x58, 0xcc, 0xe7, 0xab, 0x5d, 0xa7, 0x51, 0xbe, 0x29, 0xe1,
	0x7a, 0x2e, 0x61, 0x9d, 0xd5, 0xf4, 0x4d, 0xc2, 0xfe, 0x67, 0x4d, 0x1f, 0x83, 0xb7, 0x87, 0x11,
	0x4e, 0x50, 0xa8, 0x43, 0xe9, 0xa2, 0x8e, 0x25, 0xce, 0x9a, 0x02, 0x56, 0x9f, 0xa3, 0xdc, 0x95,
	0xec, 0xaf, 0xe8, 0x05, 0x2d, 0x29, 0x15, 0xb5, 0xc4, 0x85, 0x79, 0xca, 0x98, 0x42, 0xad, 0x6d,
	0xa5, 0x2b, 0x9d, 0xd1, 0x67, 0xf3, 0x77, 0x09, 0x6a, 0x39, 0xa5, 0x7f, 0x95, 0x58, 0x83, 0x8a,
	0xc2, 0x13, 0x54, 0x28, 0x8e, 0x71, 0x28, 0x72, 0x09, 0x10, 0x1f, 0x96, 0x19, 0x6a, 0xc3, 0x85,
	0x1d, 0xe8, 0x60, 0x14, 0xcc, 0x8c, 0xe5, 0x91, 0x31, 0xd3, 0xd3, 0x81, 0x85, 0x6c, 0x64, 0xdd,
	0xbd, 0x7c, 0x70, 0xce, 0x13, 0x77, 0xd6, 0x92, 0xeb, 0x63, 0xf0, 0x11, 0x4f, 0xae, 0x12, 0x23,
	0x6a, 0xdc, 0xb9, 0x86, 0xd3, 0x72, 0x72, 0xc4, 0x7d, 0x6a, 0xae, 0x11, 0x45, 0xe8, 0xce, 0x5f,
	0x27, 0x8a, 0x90, 0x3c, 0x80, 0x05, 0x23, 0x0d, 0x8d, 0x82, 0x53, 0xe4, 0x61, 0xcf, 0xb8, 0xb7,
	0x2c, 0xab, 0x6a, 0xb1, 0x77, 0x16, 0x22, 0x9b, 0xb0, 0xa4, 0x64, 0x6a, 0xb8, 0x08, 0x03, 0x9d,
	0x75, 0x26, 0xcb, 0xb9, 0x62, 0xab, 0xb2, 0x38, 0xc4, 0x0f, 0x87, 0x30, 0x21, 0x30, 0x93, 0xa6,
	0x9c, 0xb9, 0x60, 0xa3, 0xb7, 0xbf, 0xb7, 0xbf, 0x95, 0xe1, 0xee, 0x61, 0x8f, 0x27, 0x09, 0x17,
	0xe1, 0xae, 0x8c, 0x63, 0x2a, 0xd8, 0x21, 0xaa, 0x3e, 0x3f, 0x46, 0xf2, 0x1e, 0x96, 0x0b, 0x76,
	0x9c, 0xb4, 0xf2, 0x03, 0x37, 0xf9, 0x0c, 0x78, 0xd3, 0x46, 0x33, 0xf3, 0x5e, 0x30, 0xed, 0x57,
	0xbd, 0x4f, 0x5e, 0x88, 0xe9, 0xde, 0xdf, 0xc2, 0x72, 0xc1, 0xa4, 0x5f, 0xf5, 0x3e, 0x79, 0x19,
	0xbc, 0x3b, 0x43, 0xe6, 0xc5, 0x55, 0x7c, 0x16, 0x27, 0xe6, 0x8c, 0x7c, 0x80, 0x95, 0xa2, 0x9d,
	0x20, 0x9b, 0x79, 0xc7, 0x53, 0xf6, 0x66, 0x6a, 0xdc, 0xdb, 0xdf, 0x1d, 0x58, 0x19, 0xb5, 0xe3,
	0x75, 0x8a, 0xea, 0x6c, 0xd4, 0x8c, 0x73, 0x20, 0xd7, 0x0f, 0x10, 0xd9, 0xc8, 0xfb, 0x9a, 0x78,
	0x2f, 0xbd, 0xd6, 0xcd, 0xc4, 0xc1, 0x2d, 0x6b, 0xde, 0xfe, 0xf2, 0xf3, 0xd7, 0x8f, 0x52, 0x95,
	0x54, 0xfc, 0xc1, 0x25, 0xf3, 0xb7, 0x76, 0xd6, 0x8f, 0xee, 0x87, 0xdc, 0xf4, 0xd2, 0x6e, 0xfb,
	0x58, 0xc6, 0xbe, 0x75, 0xe4, 0x6b, 0x93, 0x76, 0xb5, 0x3f, 0xf8, 0x4f, 0xe9, 0xce, 0xd9, 0x2a,
	0x3d, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x46, 0x7e, 0xe7, 0x31, 0x64, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShippingCommandServiceClient is the client API for ShippingCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingCommandServiceClient interface {
	CreateDeliveryOrder(ctx context.Context, in *CreateDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error)
	UpdateDeliveryOrder(ctx context.Context, in *UpdateDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error)
	DeleteDeliveryOrder(ctx context.Context, in *DeleteDeliveryOrderRequest, opts ...grpc.CallOption) (*stubs.Empty, error)
	GeoCodeDeliveryOrder(ctx context.Context, in *GeoCodeDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error)
}

type shippingCommandServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingCommandServiceClient(cc *grpc.ClientConn) ShippingCommandServiceClient {
	return &shippingCommandServiceClient{cc}
}

func (c *shippingCommandServiceClient) CreateDeliveryOrder(ctx context.Context, in *CreateDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error) {
	out := new(DeliveryOrder)
	err := c.cc.Invoke(ctx, "/oswee.dms.v1.ShippingCommandService/CreateDeliveryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingCommandServiceClient) UpdateDeliveryOrder(ctx context.Context, in *UpdateDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error) {
	out := new(DeliveryOrder)
	err := c.cc.Invoke(ctx, "/oswee.dms.v1.ShippingCommandService/UpdateDeliveryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingCommandServiceClient) DeleteDeliveryOrder(ctx context.Context, in *DeleteDeliveryOrderRequest, opts ...grpc.CallOption) (*stubs.Empty, error) {
	out := new(stubs.Empty)
	err := c.cc.Invoke(ctx, "/oswee.dms.v1.ShippingCommandService/DeleteDeliveryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingCommandServiceClient) GeoCodeDeliveryOrder(ctx context.Context, in *GeoCodeDeliveryOrderRequest, opts ...grpc.CallOption) (*DeliveryOrder, error) {
	out := new(DeliveryOrder)
	err := c.cc.Invoke(ctx, "/oswee.dms.v1.ShippingCommandService/GeoCodeDeliveryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingCommandServiceServer is the server API for ShippingCommandService service.
type ShippingCommandServiceServer interface {
	CreateDeliveryOrder(context.Context, *CreateDeliveryOrderRequest) (*DeliveryOrder, error)
	UpdateDeliveryOrder(context.Context, *UpdateDeliveryOrderRequest) (*DeliveryOrder, error)
	DeleteDeliveryOrder(context.Context, *DeleteDeliveryOrderRequest) (*stubs.Empty, error)
	GeoCodeDeliveryOrder(context.Context, *GeoCodeDeliveryOrderRequest) (*DeliveryOrder, error)
}

func RegisterShippingCommandServiceServer(s *grpc.Server, srv ShippingCommandServiceServer) {
	s.RegisterService(&_ShippingCommandService_serviceDesc, srv)
}

func _ShippingCommandService_CreateDeliveryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeliveryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingCommandServiceServer).CreateDeliveryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.dms.v1.ShippingCommandService/CreateDeliveryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingCommandServiceServer).CreateDeliveryOrder(ctx, req.(*CreateDeliveryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingCommandService_UpdateDeliveryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeliveryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingCommandServiceServer).UpdateDeliveryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.dms.v1.ShippingCommandService/UpdateDeliveryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingCommandServiceServer).UpdateDeliveryOrder(ctx, req.(*UpdateDeliveryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingCommandService_DeleteDeliveryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeliveryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingCommandServiceServer).DeleteDeliveryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.dms.v1.ShippingCommandService/DeleteDeliveryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingCommandServiceServer).DeleteDeliveryOrder(ctx, req.(*DeleteDeliveryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingCommandService_GeoCodeDeliveryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCodeDeliveryOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingCommandServiceServer).GeoCodeDeliveryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.dms.v1.ShippingCommandService/GeoCodeDeliveryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingCommandServiceServer).GeoCodeDeliveryOrder(ctx, req.(*GeoCodeDeliveryOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.dms.v1.ShippingCommandService",
	HandlerType: (*ShippingCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeliveryOrder",
			Handler:    _ShippingCommandService_CreateDeliveryOrder_Handler,
		},
		{
			MethodName: "UpdateDeliveryOrder",
			Handler:    _ShippingCommandService_UpdateDeliveryOrder_Handler,
		},
		{
			MethodName: "DeleteDeliveryOrder",
			Handler:    _ShippingCommandService_DeleteDeliveryOrder_Handler,
		},
		{
			MethodName: "GeoCodeDeliveryOrder",
			Handler:    _ShippingCommandService_GeoCodeDeliveryOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dms/v1/dms.proto",
}

// ShippingQueryServiceClient is the client API for ShippingQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingQueryServiceClient interface {
	ListDeliveryOrders(ctx context.Context, in *ListDeliveryOrdersRequest, opts ...grpc.CallOption) (*ListDeliveryOrdersResponse, error)
}

type shippingQueryServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingQueryServiceClient(cc *grpc.ClientConn) ShippingQueryServiceClient {
	return &shippingQueryServiceClient{cc}
}

func (c *shippingQueryServiceClient) ListDeliveryOrders(ctx context.Context, in *ListDeliveryOrdersRequest, opts ...grpc.CallOption) (*ListDeliveryOrdersResponse, error) {
	out := new(ListDeliveryOrdersResponse)
	err := c.cc.Invoke(ctx, "/oswee.dms.v1.ShippingQueryService/ListDeliveryOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingQueryServiceServer is the server API for ShippingQueryService service.
type ShippingQueryServiceServer interface {
	ListDeliveryOrders(context.Context, *ListDeliveryOrdersRequest) (*ListDeliveryOrdersResponse, error)
}

func RegisterShippingQueryServiceServer(s *grpc.Server, srv ShippingQueryServiceServer) {
	s.RegisterService(&_ShippingQueryService_serviceDesc, srv)
}

func _ShippingQueryService_ListDeliveryOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveryOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingQueryServiceServer).ListDeliveryOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.dms.v1.ShippingQueryService/ListDeliveryOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingQueryServiceServer).ListDeliveryOrders(ctx, req.(*ListDeliveryOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingQueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.dms.v1.ShippingQueryService",
	HandlerType: (*ShippingQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDeliveryOrders",
			Handler:    _ShippingQueryService_ListDeliveryOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dms/v1/dms.proto",
}
