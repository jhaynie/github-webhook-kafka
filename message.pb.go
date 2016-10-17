// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package main

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

type Message struct {
	Version   int32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Event     string `protobuf:"bytes,4,opt,name=event" json:"event,omitempty"`
	Payload   []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Context   string `protobuf:"bytes,6,opt,name=context" json:"context,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Message)(nil), "main.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53,
	0x9a, 0xcd, 0xc8, 0xc5, 0xee, 0x0b, 0x11, 0x17, 0x92, 0xe0, 0x62, 0x2f, 0x4b, 0x2d, 0x2a, 0xce,
	0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x85, 0x64, 0xb8, 0x38, 0x4b,
	0x32, 0x81, 0xda, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x98, 0x80, 0x72, 0x9c, 0x41, 0x08, 0x01, 0x21,
	0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x66, 0xb0, 0x30, 0x90, 0x25, 0x24, 0xc2, 0xc5, 0x9a, 0x5a,
	0x96, 0x9a, 0x57, 0x22, 0xc1, 0x02, 0x16, 0x82, 0x70, 0x40, 0xa6, 0x17, 0x24, 0x56, 0xe6, 0xe4,
	0x27, 0xa6, 0x48, 0xb0, 0x02, 0xc5, 0x79, 0x82, 0x60, 0x5c, 0x90, 0x4c, 0x72, 0x7e, 0x5e, 0x49,
	0x6a, 0x45, 0x89, 0x04, 0x1b, 0x58, 0x07, 0x8c, 0x9b, 0xc4, 0x06, 0x76, 0xaa, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x71, 0x92, 0xe0, 0xb3, 0xbb, 0x00, 0x00, 0x00,
}
