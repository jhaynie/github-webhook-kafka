// Code generated by protoc-gen-go.
// source: watch_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WatchEvent struct {
	Action     string      `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
}

func (m *WatchEvent) Reset()                    { *m = WatchEvent{} }
func (m *WatchEvent) String() string            { return proto.CompactTextString(m) }
func (*WatchEvent) ProtoMessage()               {}
func (*WatchEvent) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

func (m *WatchEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *WatchEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*WatchEvent)(nil), "github.WatchEvent")
}

func init() { proto.RegisterFile("watch_event.proto", fileDescriptor35) }

var fileDescriptor35 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4f, 0x2c, 0x49,
	0xce, 0x88, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d, 0x82, 0x88, 0x49, 0x09, 0x14,
	0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55, 0x42, 0x44, 0x94, 0xea, 0xb8, 0xb8, 0xc2,
	0x41, 0x5a, 0x5d, 0x41, 0x3a, 0x85, 0xc4, 0xb8, 0xd8, 0x12, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x23, 0x2e, 0x2e, 0x84, 0x4e, 0x09, 0x26,
	0xa0, 0x1c, 0xb7, 0x91, 0x90, 0x1e, 0xc4, 0x02, 0xbd, 0x20, 0xb8, 0x4c, 0x10, 0x92, 0x2a, 0x21,
	0x15, 0x2e, 0xb6, 0xe2, 0xd4, 0xbc, 0x94, 0xd4, 0x22, 0x09, 0x66, 0xb0, 0x7a, 0x1e, 0x98, 0xfa,
	0x50, 0xa0, 0x7b, 0x82, 0xa0, 0x72, 0x49, 0x6c, 0x60, 0x67, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x14, 0x46, 0xc8, 0xc1, 0x00, 0x00, 0x00,
}