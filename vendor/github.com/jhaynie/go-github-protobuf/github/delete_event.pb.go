// Code generated by protoc-gen-go.
// source: delete_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeleteEvent struct {
	Ref        string      `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	RefType    string      `protobuf:"bytes,2,opt,name=ref_type,json=refType" json:"ref_type,omitempty"`
	PusherType string      `protobuf:"bytes,3,opt,name=pusher_type,json=pusherType" json:"pusher_type,omitempty"`
	Repository *Repository `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,7,opt,name=sender" json:"sender,omitempty"`
}

func (m *DeleteEvent) Reset()                    { *m = DeleteEvent{} }
func (m *DeleteEvent) String() string            { return proto.CompactTextString(m) }
func (*DeleteEvent) ProtoMessage()               {}
func (*DeleteEvent) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DeleteEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *DeleteEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteEvent)(nil), "github.DeleteEvent")
}

func init() { proto.RegisterFile("delete_event.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d, 0x82, 0x88, 0x49, 0x09,
	0x14, 0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55, 0x42, 0x44, 0x94, 0xb6, 0x32, 0x72,
	0x71, 0xbb, 0x80, 0x35, 0xbb, 0x82, 0xf4, 0x0a, 0x09, 0x70, 0x31, 0x17, 0xa5, 0xa6, 0x49, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x92, 0x5c, 0x1c, 0x40, 0x2a, 0xbe, 0xa4, 0xb2,
	0x20, 0x55, 0x82, 0x09, 0x2c, 0xcc, 0x0e, 0xe4, 0x87, 0x00, 0xb9, 0x42, 0xf2, 0x5c, 0xdc, 0x05,
	0xa5, 0xc5, 0x19, 0xa9, 0x45, 0x10, 0x59, 0x66, 0xb0, 0x2c, 0x17, 0x44, 0x08, 0xac, 0xc0, 0x88,
	0x8b, 0x0b, 0x61, 0xa3, 0x04, 0x0b, 0x50, 0x9e, 0xdb, 0x48, 0x48, 0x0f, 0xe2, 0x30, 0xbd, 0x20,
	0xb8, 0x4c, 0x10, 0x92, 0x2a, 0x21, 0x15, 0x2e, 0xb6, 0xe2, 0xd4, 0xbc, 0x94, 0xd4, 0x22, 0x09,
	0x76, 0xb0, 0x7a, 0x1e, 0x98, 0xfa, 0x50, 0xa0, 0x3f, 0x82, 0xa0, 0x72, 0x49, 0x6c, 0x60, 0xe7,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0xf7, 0x39, 0x35, 0xfa, 0x00, 0x00, 0x00,
}
