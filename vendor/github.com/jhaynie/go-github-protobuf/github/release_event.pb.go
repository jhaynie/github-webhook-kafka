// Code generated by protoc-gen-go.
// source: release_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReleaseEvent struct {
	Action     string      `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Release    *Release    `protobuf:"bytes,2,opt,name=release" json:"release,omitempty"`
	Repository *Repository `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
}

func (m *ReleaseEvent) Reset()                    { *m = ReleaseEvent{} }
func (m *ReleaseEvent) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEvent) ProtoMessage()               {}
func (*ReleaseEvent) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *ReleaseEvent) GetRelease() *Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *ReleaseEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ReleaseEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*ReleaseEvent)(nil), "github.ReleaseEvent")
}

func init() { proto.RegisterFile("release_event.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcd, 0x49,
	0x4d, 0x2c, 0x4e, 0x8d, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d, 0x82, 0x88, 0x49,
	0xf1, 0x42, 0x15, 0x42, 0xb9, 0x02, 0x45, 0xa9, 0x05, 0xf9, 0xc5, 0x99, 0x25, 0xf9, 0x45, 0x95,
	0x10, 0x11, 0xa5, 0xd5, 0x8c, 0x5c, 0x3c, 0x41, 0x10, 0x35, 0xae, 0x20, 0xb3, 0x84, 0xc4, 0xb8,
	0xd8, 0x12, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c,
	0x21, 0x4d, 0x2e, 0x76, 0xa8, 0x59, 0x12, 0x4c, 0x40, 0x09, 0x6e, 0x23, 0x7e, 0x3d, 0x88, 0x7d,
	0x7a, 0x50, 0xed, 0x41, 0x30, 0x79, 0x21, 0x23, 0x2e, 0x2e, 0x84, 0x3d, 0x12, 0xcc, 0x60, 0xd5,
	0x42, 0x08, 0xd5, 0x30, 0x99, 0x20, 0x24, 0x55, 0x42, 0x2a, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29,
	0xa9, 0x45, 0x12, 0x2c, 0x60, 0xf5, 0x3c, 0x30, 0xf5, 0xa1, 0x40, 0xcf, 0x04, 0x41, 0xe5, 0x92,
	0xd8, 0xc0, 0x8e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xb6, 0xd8, 0x1d, 0x00, 0x01,
	0x00, 0x00,
}