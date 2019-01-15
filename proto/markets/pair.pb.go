// Code generated by protoc-gen-go. DO NOT EDIT.
// source: markets/pair.proto

package ProtobufMarkets

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PairUpdateMessage struct {
	Pair uint64 `protobuf:"varint,1,opt,name=pair,proto3" json:"pair,omitempty"`
	// Types that are valid to be assigned to Update:
	//	*PairUpdateMessage_VwapUpdate
	//	*PairUpdateMessage_PerformanceUpdate
	//	*PairUpdateMessage_TrendlineUpdate
	Update               isPairUpdateMessage_Update `protobuf_oneof:"Update"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PairUpdateMessage) Reset()         { *m = PairUpdateMessage{} }
func (m *PairUpdateMessage) String() string { return proto.CompactTextString(m) }
func (*PairUpdateMessage) ProtoMessage()    {}
func (*PairUpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f15de32d6316fa, []int{0}
}

func (m *PairUpdateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairUpdateMessage.Unmarshal(m, b)
}
func (m *PairUpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairUpdateMessage.Marshal(b, m, deterministic)
}
func (m *PairUpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairUpdateMessage.Merge(m, src)
}
func (m *PairUpdateMessage) XXX_Size() int {
	return xxx_messageInfo_PairUpdateMessage.Size(m)
}
func (m *PairUpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PairUpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PairUpdateMessage proto.InternalMessageInfo

func (m *PairUpdateMessage) GetPair() uint64 {
	if m != nil {
		return m.Pair
	}
	return 0
}

type isPairUpdateMessage_Update interface {
	isPairUpdateMessage_Update()
}

type PairUpdateMessage_VwapUpdate struct {
	VwapUpdate *PairVwapUpdate `protobuf:"bytes,2,opt,name=vwapUpdate,proto3,oneof"`
}

type PairUpdateMessage_PerformanceUpdate struct {
	PerformanceUpdate *PairPerformanceUpdate `protobuf:"bytes,3,opt,name=performanceUpdate,proto3,oneof"`
}

type PairUpdateMessage_TrendlineUpdate struct {
	TrendlineUpdate *PairTrendlineUpdate `protobuf:"bytes,4,opt,name=trendlineUpdate,proto3,oneof"`
}

func (*PairUpdateMessage_VwapUpdate) isPairUpdateMessage_Update() {}

func (*PairUpdateMessage_PerformanceUpdate) isPairUpdateMessage_Update() {}

func (*PairUpdateMessage_TrendlineUpdate) isPairUpdateMessage_Update() {}

func (m *PairUpdateMessage) GetUpdate() isPairUpdateMessage_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *PairUpdateMessage) GetVwapUpdate() *PairVwapUpdate {
	if x, ok := m.GetUpdate().(*PairUpdateMessage_VwapUpdate); ok {
		return x.VwapUpdate
	}
	return nil
}

func (m *PairUpdateMessage) GetPerformanceUpdate() *PairPerformanceUpdate {
	if x, ok := m.GetUpdate().(*PairUpdateMessage_PerformanceUpdate); ok {
		return x.PerformanceUpdate
	}
	return nil
}

func (m *PairUpdateMessage) GetTrendlineUpdate() *PairTrendlineUpdate {
	if x, ok := m.GetUpdate().(*PairUpdateMessage_TrendlineUpdate); ok {
		return x.TrendlineUpdate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PairUpdateMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PairUpdateMessage_OneofMarshaler, _PairUpdateMessage_OneofUnmarshaler, _PairUpdateMessage_OneofSizer, []interface{}{
		(*PairUpdateMessage_VwapUpdate)(nil),
		(*PairUpdateMessage_PerformanceUpdate)(nil),
		(*PairUpdateMessage_TrendlineUpdate)(nil),
	}
}

func _PairUpdateMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PairUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *PairUpdateMessage_VwapUpdate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VwapUpdate); err != nil {
			return err
		}
	case *PairUpdateMessage_PerformanceUpdate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PerformanceUpdate); err != nil {
			return err
		}
	case *PairUpdateMessage_TrendlineUpdate:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TrendlineUpdate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PairUpdateMessage.Update has unexpected type %T", x)
	}
	return nil
}

func _PairUpdateMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PairUpdateMessage)
	switch tag {
	case 2: // Update.vwapUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PairVwapUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &PairUpdateMessage_VwapUpdate{msg}
		return true, err
	case 3: // Update.performanceUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PairPerformanceUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &PairUpdateMessage_PerformanceUpdate{msg}
		return true, err
	case 4: // Update.trendlineUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PairTrendlineUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &PairUpdateMessage_TrendlineUpdate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PairUpdateMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PairUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *PairUpdateMessage_VwapUpdate:
		s := proto.Size(x.VwapUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PairUpdateMessage_PerformanceUpdate:
		s := proto.Size(x.PerformanceUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PairUpdateMessage_TrendlineUpdate:
		s := proto.Size(x.TrendlineUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PairVwapUpdate struct {
	Vwap                 float64  `protobuf:"fixed64,1,opt,name=vwap,proto3" json:"vwap,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairVwapUpdate) Reset()         { *m = PairVwapUpdate{} }
func (m *PairVwapUpdate) String() string { return proto.CompactTextString(m) }
func (*PairVwapUpdate) ProtoMessage()    {}
func (*PairVwapUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f15de32d6316fa, []int{1}
}

func (m *PairVwapUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairVwapUpdate.Unmarshal(m, b)
}
func (m *PairVwapUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairVwapUpdate.Marshal(b, m, deterministic)
}
func (m *PairVwapUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairVwapUpdate.Merge(m, src)
}
func (m *PairVwapUpdate) XXX_Size() int {
	return xxx_messageInfo_PairVwapUpdate.Size(m)
}
func (m *PairVwapUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PairVwapUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PairVwapUpdate proto.InternalMessageInfo

func (m *PairVwapUpdate) GetVwap() float64 {
	if m != nil {
		return m.Vwap
	}
	return 0
}

func (m *PairVwapUpdate) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PairPerformanceUpdate struct {
	Window               string   `protobuf:"bytes,1,opt,name=window,proto3" json:"window,omitempty"`
	Performance          float64  `protobuf:"fixed64,2,opt,name=performance,proto3" json:"performance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairPerformanceUpdate) Reset()         { *m = PairPerformanceUpdate{} }
func (m *PairPerformanceUpdate) String() string { return proto.CompactTextString(m) }
func (*PairPerformanceUpdate) ProtoMessage()    {}
func (*PairPerformanceUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f15de32d6316fa, []int{2}
}

func (m *PairPerformanceUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairPerformanceUpdate.Unmarshal(m, b)
}
func (m *PairPerformanceUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairPerformanceUpdate.Marshal(b, m, deterministic)
}
func (m *PairPerformanceUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairPerformanceUpdate.Merge(m, src)
}
func (m *PairPerformanceUpdate) XXX_Size() int {
	return xxx_messageInfo_PairPerformanceUpdate.Size(m)
}
func (m *PairPerformanceUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PairPerformanceUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PairPerformanceUpdate proto.InternalMessageInfo

func (m *PairPerformanceUpdate) GetWindow() string {
	if m != nil {
		return m.Window
	}
	return ""
}

func (m *PairPerformanceUpdate) GetPerformance() float64 {
	if m != nil {
		return m.Performance
	}
	return 0
}

type PairTrendlineUpdate struct {
	Window               string   `protobuf:"bytes,1,opt,name=window,proto3" json:"window,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Volume               string   `protobuf:"bytes,4,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairTrendlineUpdate) Reset()         { *m = PairTrendlineUpdate{} }
func (m *PairTrendlineUpdate) String() string { return proto.CompactTextString(m) }
func (*PairTrendlineUpdate) ProtoMessage()    {}
func (*PairTrendlineUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f15de32d6316fa, []int{3}
}

func (m *PairTrendlineUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairTrendlineUpdate.Unmarshal(m, b)
}
func (m *PairTrendlineUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairTrendlineUpdate.Marshal(b, m, deterministic)
}
func (m *PairTrendlineUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairTrendlineUpdate.Merge(m, src)
}
func (m *PairTrendlineUpdate) XXX_Size() int {
	return xxx_messageInfo_PairTrendlineUpdate.Size(m)
}
func (m *PairTrendlineUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PairTrendlineUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PairTrendlineUpdate proto.InternalMessageInfo

func (m *PairTrendlineUpdate) GetWindow() string {
	if m != nil {
		return m.Window
	}
	return ""
}

func (m *PairTrendlineUpdate) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PairTrendlineUpdate) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *PairTrendlineUpdate) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func init() {
	proto.RegisterType((*PairUpdateMessage)(nil), "ProtobufMarkets.PairUpdateMessage")
	proto.RegisterType((*PairVwapUpdate)(nil), "ProtobufMarkets.PairVwapUpdate")
	proto.RegisterType((*PairPerformanceUpdate)(nil), "ProtobufMarkets.PairPerformanceUpdate")
	proto.RegisterType((*PairTrendlineUpdate)(nil), "ProtobufMarkets.PairTrendlineUpdate")
}

func init() { proto.RegisterFile("markets/pair.proto", fileDescriptor_f0f15de32d6316fa) }

var fileDescriptor_f0f15de32d6316fa = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4f, 0xb3, 0x30,
	0x14, 0xc6, 0x61, 0xe3, 0x25, 0x2f, 0x67, 0x89, 0xcb, 0x8e, 0x7f, 0xb2, 0x0b, 0x13, 0x09, 0x31,
	0x66, 0x57, 0x98, 0xe8, 0x27, 0x70, 0x57, 0xdc, 0x2c, 0xc1, 0x46, 0x77, 0xdf, 0x8d, 0xce, 0x34,
	0x0e, 0xda, 0x94, 0x6e, 0x7c, 0x1b, 0x3f, 0xab, 0x69, 0x8b, 0xca, 0x10, 0xaf, 0x78, 0x4e, 0x79,
	0xce, 0x2f, 0x7d, 0x1e, 0x00, 0x2c, 0xa9, 0x7a, 0x67, 0xba, 0xbe, 0x97, 0x94, 0xab, 0x54, 0x2a,
	0xa1, 0x05, 0x4e, 0x73, 0xf3, 0xd8, 0x1c, 0x76, 0x2b, 0xf7, 0x2e, 0xf9, 0x18, 0xc1, 0x2c, 0xa7,
	0x5c, 0xbd, 0xca, 0x82, 0x6a, 0xb6, 0x62, 0x75, 0x4d, 0xdf, 0x18, 0x22, 0x04, 0x66, 0x69, 0xee,
	0xc7, 0xfe, 0x22, 0x20, 0x56, 0xe3, 0x13, 0xc0, 0xb1, 0xa1, 0xd2, 0x19, 0xe7, 0xa3, 0xd8, 0x5f,
	0x4c, 0x1e, 0x6e, 0xd2, 0x1e, 0x2f, 0x35, 0xac, 0xf5, 0xb7, 0x2d, 0xf3, 0x48, 0x67, 0x09, 0xd7,
	0x30, 0x93, 0x4c, 0xed, 0x84, 0x2a, 0x69, 0xb5, 0x65, 0x2d, 0x69, 0x6c, 0x49, 0x77, 0x83, 0xa4,
	0xbc, 0xef, 0xce, 0x3c, 0xf2, 0x1b, 0x81, 0x39, 0x4c, 0xb5, 0x62, 0x55, 0xb1, 0xe7, 0xd5, 0x17,
	0x35, 0xb0, 0xd4, 0xdb, 0x41, 0xea, 0xcb, 0xa9, 0x37, 0xf3, 0x48, 0x7f, 0x7d, 0xf9, 0x1f, 0x42,
	0xa7, 0x92, 0x25, 0x9c, 0x9d, 0x66, 0x32, 0xe5, 0x98, 0x4c, 0xb6, 0x1c, 0x9f, 0x58, 0x8d, 0xd7,
	0x10, 0x69, 0x5e, 0xb2, 0x5a, 0xd3, 0x52, 0xda, 0x6e, 0xc6, 0xe4, 0xe7, 0x20, 0x79, 0x86, 0xcb,
	0xc1, 0x34, 0x78, 0x05, 0x61, 0xc3, 0xab, 0x42, 0x34, 0x16, 0x16, 0x91, 0x76, 0xc2, 0x18, 0x26,
	0x9d, 0x94, 0x16, 0xe8, 0x93, 0xee, 0x51, 0x22, 0xe0, 0x7c, 0x20, 0xca, 0x9f, 0x40, 0x84, 0xc0,
	0x5c, 0xa7, 0xbd, 0x9a, 0xd5, 0x78, 0x01, 0xff, 0xa4, 0xe2, 0x5b, 0xf7, 0x05, 0x22, 0xe2, 0x06,
	0x43, 0x38, 0x8a, 0xfd, 0xa1, 0x74, 0x15, 0x46, 0xa4, 0x9d, 0x36, 0xa1, 0xfd, 0x81, 0x1e, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x7d, 0xde, 0x7c, 0x56, 0x02, 0x00, 0x00,
}
