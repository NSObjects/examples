// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/booking/srv/geo/proto/geo.proto

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/booking/srv/geo/proto/geo.proto

It has these top-level messages:
	Request
	Result
*/
package geo

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

// The latitude and longitude of the current location.
type Request struct {
	Lat float32 `protobuf:"fixed32,1,opt,name=lat" json:"lat,omitempty"`
	Lon float32 `protobuf:"fixed32,2,opt,name=lon" json:"lon,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Request) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Result struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "geo.Request")
	proto.RegisterType((*Result)(nil), "geo.Result")
}

func init() {
	proto.RegisterFile("github.com/micro/examples/booking/srv/geo/proto/geo.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0xae, 0x82, 0x30,
	0x14, 0x86, 0x03, 0x4d, 0xb8, 0xf7, 0x9e, 0xeb, 0x60, 0x3a, 0x11, 0x26, 0x82, 0x0e, 0xc4, 0x44,
	0x9a, 0xe8, 0xe4, 0x13, 0x18, 0x17, 0x87, 0xbe, 0x01, 0xc5, 0x93, 0x42, 0x2c, 0x1c, 0x6c, 0x8b,
	0xd1, 0xb7, 0x37, 0x25, 0xc4, 0xed, 0xfb, 0xbf, 0xe1, 0x7c, 0x07, 0x4e, 0xba, 0xf3, 0xed, 0xa4,
	0xaa, 0x86, 0x7a, 0xd1, 0x77, 0x8d, 0x25, 0x81, 0xaf, 0xba, 0x1f, 0x0d, 0x3a, 0xa1, 0x88, 0xee,
	0xdd, 0xa0, 0x85, 0xb3, 0x4f, 0xa1, 0x91, 0xc4, 0x68, 0xc9, 0x53, 0xa0, 0x6a, 0x26, 0xce, 0x34,
	0x52, 0xb1, 0x87, 0x1f, 0x89, 0x8f, 0x09, 0x9d, 0xe7, 0x6b, 0x60, 0xa6, 0xf6, 0x69, 0x94, 0x47,
	0x65, 0x2c, 0x03, 0xce, 0x86, 0x86, 0x34, 0x5e, 0x0c, 0x0d, 0xc5, 0x16, 0x12, 0x89, 0x6e, 0x32,
	0x9e, 0x67, 0xf0, 0xdb, 0x92, 0x47, 0x73, 0xb9, 0xb9, 0x34, 0xca, 0x59, 0xf9, 0x27, 0xbf, 0xfb,
	0xb0, 0x03, 0x76, 0x46, 0xe2, 0x1b, 0x48, 0xae, 0x58, 0x5b, 0xf5, 0xe6, 0xab, 0x2a, 0x64, 0x97,
	0x50, 0xf6, 0xbf, 0xac, 0x70, 0x47, 0x25, 0xf3, 0x33, 0xc7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x39, 0xe5, 0x14, 0x38, 0xc9, 0x00, 0x00, 0x00,
}
