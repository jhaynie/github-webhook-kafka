// Code generated by protoc-gen-go.
// source: fork_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ForkEvent struct {
	Forkee     *Repository `protobuf:"bytes,1,opt,name=forkee" json:"forkee,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
}

func (m *ForkEvent) Reset()                    { *m = ForkEvent{} }
func (m *ForkEvent) String() string            { return proto.CompactTextString(m) }
func (*ForkEvent) ProtoMessage()               {}
func (*ForkEvent) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *ForkEvent) GetForkee() *Repository {
	if m != nil {
		return m.Forkee
	}
	return nil
}

func (m *ForkEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ForkEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*ForkEvent)(nil), "github.ForkEvent")
}

func init() { proto.RegisterFile("fork_event.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcb, 0x2f, 0xca,
	0x8e, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d, 0x82, 0x88, 0x49, 0x09, 0x14, 0xa5,
	0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55, 0x42, 0x44, 0x94, 0x26, 0x32, 0x72, 0x71, 0xba,
	0x01, 0xb5, 0xba, 0x82, 0x74, 0x0a, 0x69, 0x71, 0xb1, 0x81, 0xcc, 0x49, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd2, 0x83, 0x18, 0xa2, 0x17, 0x04, 0xd7, 0x17, 0x04, 0x55, 0x21,
	0x64, 0xc4, 0xc5, 0x85, 0x30, 0x4d, 0x82, 0x09, 0xa7, 0x7a, 0x24, 0x55, 0x42, 0x2a, 0x5c, 0x6c,
	0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0xcc, 0x60, 0xf5, 0x3c, 0x30, 0xf5, 0xa1, 0x40, 0x37,
	0x06, 0x41, 0xe5, 0x92, 0xd8, 0xc0, 0x4e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0x29,
	0x41, 0x13, 0xd4, 0x00, 0x00, 0x00,
}