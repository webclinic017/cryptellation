// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candlesticks.proto

package candlesticks

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Candlestick struct {
	Open                 float32  `protobuf:"fixed32,1,opt,name=open,proto3" json:"open,omitempty"`
	High                 float32  `protobuf:"fixed32,2,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float32  `protobuf:"fixed32,3,opt,name=low,proto3" json:"low,omitempty"`
	Close                float32  `protobuf:"fixed32,4,opt,name=close,proto3" json:"close,omitempty"`
	Volume               float32  `protobuf:"fixed32,5,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candlestick) Reset()         { *m = Candlestick{} }
func (m *Candlestick) String() string { return proto.CompactTextString(m) }
func (*Candlestick) ProtoMessage()    {}
func (*Candlestick) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a8f911941630527, []int{0}
}

func (m *Candlestick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candlestick.Unmarshal(m, b)
}
func (m *Candlestick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candlestick.Marshal(b, m, deterministic)
}
func (m *Candlestick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candlestick.Merge(m, src)
}
func (m *Candlestick) XXX_Size() int {
	return xxx_messageInfo_Candlestick.Size(m)
}
func (m *Candlestick) XXX_DiscardUnknown() {
	xxx_messageInfo_Candlestick.DiscardUnknown(m)
}

var xxx_messageInfo_Candlestick proto.InternalMessageInfo

func (m *Candlestick) GetOpen() float32 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candlestick) GetHigh() float32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candlestick) GetLow() float32 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candlestick) GetClose() float32 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candlestick) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type ReadCandlesticksRequest struct {
	ExchangeName         string   `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	PairSymbol           string   `protobuf:"bytes,2,opt,name=pair_symbol,json=pairSymbol,proto3" json:"pair_symbol,omitempty"`
	PeriodSymbol         string   `protobuf:"bytes,3,opt,name=period_symbol,json=periodSymbol,proto3" json:"period_symbol,omitempty"`
	Start                string   `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Limit                int64    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCandlesticksRequest) Reset()         { *m = ReadCandlesticksRequest{} }
func (m *ReadCandlesticksRequest) String() string { return proto.CompactTextString(m) }
func (*ReadCandlesticksRequest) ProtoMessage()    {}
func (*ReadCandlesticksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a8f911941630527, []int{1}
}

func (m *ReadCandlesticksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCandlesticksRequest.Unmarshal(m, b)
}
func (m *ReadCandlesticksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCandlesticksRequest.Marshal(b, m, deterministic)
}
func (m *ReadCandlesticksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCandlesticksRequest.Merge(m, src)
}
func (m *ReadCandlesticksRequest) XXX_Size() int {
	return xxx_messageInfo_ReadCandlesticksRequest.Size(m)
}
func (m *ReadCandlesticksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCandlesticksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCandlesticksRequest proto.InternalMessageInfo

func (m *ReadCandlesticksRequest) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *ReadCandlesticksRequest) GetPairSymbol() string {
	if m != nil {
		return m.PairSymbol
	}
	return ""
}

func (m *ReadCandlesticksRequest) GetPeriodSymbol() string {
	if m != nil {
		return m.PeriodSymbol
	}
	return ""
}

func (m *ReadCandlesticksRequest) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ReadCandlesticksRequest) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *ReadCandlesticksRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadCandlesticksResponse struct {
	Candlesticks         []*Candlestick `protobuf:"bytes,1,rep,name=candlesticks,proto3" json:"candlesticks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadCandlesticksResponse) Reset()         { *m = ReadCandlesticksResponse{} }
func (m *ReadCandlesticksResponse) String() string { return proto.CompactTextString(m) }
func (*ReadCandlesticksResponse) ProtoMessage()    {}
func (*ReadCandlesticksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a8f911941630527, []int{2}
}

func (m *ReadCandlesticksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCandlesticksResponse.Unmarshal(m, b)
}
func (m *ReadCandlesticksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCandlesticksResponse.Marshal(b, m, deterministic)
}
func (m *ReadCandlesticksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCandlesticksResponse.Merge(m, src)
}
func (m *ReadCandlesticksResponse) XXX_Size() int {
	return xxx_messageInfo_ReadCandlesticksResponse.Size(m)
}
func (m *ReadCandlesticksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCandlesticksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCandlesticksResponse proto.InternalMessageInfo

func (m *ReadCandlesticksResponse) GetCandlesticks() []*Candlestick {
	if m != nil {
		return m.Candlesticks
	}
	return nil
}

func init() {
	proto.RegisterType((*Candlestick)(nil), "candlesticks.Candlestick")
	proto.RegisterType((*ReadCandlesticksRequest)(nil), "candlesticks.ReadCandlesticksRequest")
	proto.RegisterType((*ReadCandlesticksResponse)(nil), "candlesticks.ReadCandlesticksResponse")
}

func init() { proto.RegisterFile("candlesticks.proto", fileDescriptor_8a8f911941630527) }

var fileDescriptor_8a8f911941630527 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x1d, 0x03, 0x92, 0xbd, 0x60, 0x42, 0x2a, 0xd1, 0x4a, 0x4c, 0x24, 0x33, 0x1a, 0x4e,
	0x90, 0xe0, 0xd9, 0x8b, 0xde, 0x3d, 0x94, 0x93, 0x5e, 0x48, 0xd9, 0xde, 0x8c, 0x86, 0xad, 0x9d,
	0x6b, 0xc1, 0x3f, 0x1f, 0xcf, 0x4f, 0x66, 0xda, 0x42, 0xdc, 0x34, 0xc6, 0xdb, 0xfb, 0xfc, 0xf2,
	0xf4, 0xe9, 0xd3, 0x77, 0x03, 0x92, 0x70, 0x99, 0xe6, 0xa8, 0x8d, 0x48, 0x36, 0x7a, 0x5a, 0x56,
	0xca, 0x28, 0xd2, 0xaf, 0xb3, 0x78, 0x0b, 0xbd, 0x87, 0x6f, 0x4d, 0x08, 0xb4, 0x55, 0x89, 0x92,
	0x06, 0xe3, 0x60, 0xd2, 0x62, 0x6e, 0xb6, 0x6c, 0x2d, 0xb2, 0x35, 0x6d, 0x79, 0x66, 0x67, 0x32,
	0x80, 0x30, 0x57, 0xaf, 0x34, 0x74, 0xc8, 0x8e, 0x64, 0x08, 0x9d, 0x24, 0x57, 0x1a, 0x69, 0xdb,
	0x31, 0x2f, 0xc8, 0x29, 0x74, 0x77, 0x2a, 0xdf, 0x16, 0x48, 0x3b, 0x0e, 0xef, 0x55, 0xfc, 0x19,
	0xc0, 0x19, 0x43, 0x9e, 0xd6, 0xee, 0xd6, 0x0c, 0x5f, 0xb6, 0xa8, 0x0d, 0xb9, 0x82, 0x63, 0x7c,
	0x4b, 0xd6, 0x5c, 0x66, 0xb8, 0x94, 0xbc, 0x40, 0x57, 0x26, 0x62, 0xfd, 0x03, 0x7c, 0xe4, 0x05,
	0x92, 0x4b, 0xe8, 0x95, 0x5c, 0x54, 0x4b, 0xfd, 0x5e, 0xac, 0x54, 0xee, 0xba, 0x45, 0x0c, 0x2c,
	0x5a, 0x38, 0x62, 0x53, 0x4a, 0xac, 0x84, 0x4a, 0x0f, 0x96, 0xd0, 0xa7, 0x78, 0xb8, 0x37, 0x0d,
	0xa1, 0xa3, 0x0d, 0xaf, 0x8c, 0x2b, 0x1d, 0x31, 0x2f, 0xec, 0xe3, 0x50, 0xa6, 0xae, 0x71, 0xc4,
	0xec, 0x68, 0x7d, 0xb9, 0x28, 0x84, 0xa1, 0xdd, 0x71, 0x30, 0x09, 0x99, 0x17, 0xf1, 0x13, 0xd0,
	0xdf, 0x6f, 0xd0, 0xa5, 0x92, 0x1a, 0xc9, 0x1d, 0x34, 0xf6, 0x4c, 0x83, 0x71, 0x38, 0xe9, 0xcd,
	0xcf, 0xa7, 0x8d, 0x0f, 0x52, 0x3b, 0xc9, 0x1a, 0xf6, 0xf9, 0x07, 0x9c, 0xd4, 0x63, 0x17, 0x58,
	0xed, 0x44, 0x82, 0x24, 0x81, 0xc1, 0xcf, 0x1b, 0xc9, 0x75, 0x33, 0xf3, 0x8f, 0xad, 0x8e, 0x6e,
	0xfe, 0xb3, 0xf9, 0xe2, 0xf1, 0xd1, 0xfd, 0xc5, 0xf3, 0x68, 0x56, 0x6e, 0xb2, 0x59, 0x86, 0xd2,
	0xfd, 0x31, 0xb3, 0xfa, 0xc1, 0x55, 0xd7, 0xb1, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x87, 0x8b, 0x09, 0x5b, 0x02, 0x00, 0x00,
}
