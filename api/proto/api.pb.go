// Code generated by protoc-gen-go.
// source: github.com/divisionone/micro/api/proto/api.proto
// DO NOT EDIT!

/*
Package go_micro_api is a generated protocol buffer package.

It is generated from these files:
	github.com/divisionone/micro/api/proto/api.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package go_micro_api

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

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.Response")
}

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x52, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0xcd, 0x28, 0x4c, 0xf8, 0xf0, 0x60, 0x7a, 0x30, 0x0d, 0x07, 0xc4, 0x9d, 0x88, 0x87, 0x42,
	0xe0, 0xa2, 0x5c, 0x8d, 0xd1, 0x68, 0x4c, 0x48, 0xff, 0x41, 0x61, 0x0d, 0x5b, 0x14, 0x3a, 0xd7,
	0x62, 0xc2, 0x4f, 0xf4, 0xec, 0x1f, 0xb2, 0xfd, 0x56, 0x08, 0x9a, 0xdd, 0x88, 0x97, 0xe5, 0xf5,
	0xdb, 0x7b, 0x6f, 0xaf, 0xef, 0x1b, 0xdc, 0xac, 0x72, 0x9b, 0x6d, 0x17, 0x7c, 0xa9, 0xd7, 0xa3,
	0x75, 0xbe, 0x2c, 0x75, 0x78, 0xca, 0x22, 0x1f, 0x15, 0xa5, 0xb6, 0x88, 0x38, 0x22, 0x7a, 0xbe,
	0xd2, 0x1c, 0xdf, 0x72, 0x37, 0x4b, 0xc6, 0xd0, 0x9c, 0xcb, 0xbc, 0xa4, 0x17, 0x40, 0xde, 0xd4,
	0x8e, 0x45, 0x83, 0x68, 0xd8, 0x11, 0x1e, 0xd2, 0x4b, 0x88, 0x3f, 0xe5, 0xfb, 0x56, 0x19, 0xd6,
	0x18, 0x10, 0x37, 0x0c, 0xa7, 0xe4, 0x9b, 0xc0, 0x99, 0x50, 0x1f, 0x0e, 0x5a, 0xcf, 0x59, 0x2b,
	0x9b, 0xe9, 0x34, 0x08, 0xc3, 0x89, 0x52, 0x68, 0x16, 0xd2, 0x66, 0x4e, 0xe9, 0xa7, 0x88, 0xe9,
	0x1d, 0xc4, 0x99, 0x92, 0xa9, 0x2a, 0x19, 0x71, 0x7e, 0xdd, 0xc9, 0x35, 0x3f, 0x0e, 0xc2, 0x83,
	0x25, 0x7f, 0x42, 0xce, 0xc3, 0xc6, 0x96, 0x3b, 0x11, 0x04, 0x74, 0x0c, 0x64, 0xa5, 0x2c, 0x6b,
	0xa2, 0xae, 0x5f, 0xaf, 0x7b, 0x54, 0xb6, 0x12, 0x79, 0x2a, 0x9d, 0xba, 0x00, 0xda, 0x58, 0xd6,
	0x42, 0xc9, 0x55, 0xbd, 0x64, 0xee, 0x18, 0x95, 0x06, 0xc9, 0x3e, 0xf5, 0x42, 0xa7, 0x3b, 0x16,
	0x57, 0xa9, 0x3d, 0xee, 0xbd, 0x42, 0xf7, 0x28, 0x51, 0x4d, 0x4d, 0x43, 0x68, 0x61, 0x31, 0x78,
	0xd7, 0xee, 0x84, 0xfe, 0xfe, 0x94, 0xef, 0x56, 0x54, 0x84, 0x59, 0xe3, 0x36, 0xea, 0x3d, 0x43,
	0x7b, 0x1f, 0xf4, 0x64, 0xaf, 0x17, 0xe8, 0x1c, 0x6e, 0x70, 0xaa, 0x59, 0xf2, 0x15, 0x41, 0x5b,
	0x28, 0x53, 0xe8, 0x8d, 0x51, 0xb4, 0x0f, 0x60, 0xac, 0xb4, 0x5b, 0x73, 0xaf, 0x53, 0x85, 0x9e,
	0x2d, 0x71, 0x34, 0xa1, 0xb3, 0xc3, 0x2a, 0x1b, 0xd8, 0x6f, 0xf2, 0xb7, 0xdf, 0xca, 0xa7, 0x76,
	0x97, 0xfb, 0x92, 0xc9, 0xbf, 0x95, 0xbc, 0x88, 0xf1, 0x47, 0x9f, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x11, 0xec, 0xf6, 0x16, 0x03, 0x00, 0x00,
}
