// Code generated by protoc-gen-go.
// source: create_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateEvent struct {
	Ref          string      `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	RefType      string      `protobuf:"bytes,2,opt,name=ref_type,json=refType" json:"ref_type,omitempty"`
	MasterBranch string      `protobuf:"bytes,3,opt,name=master_branch,json=masterBranch" json:"master_branch,omitempty"`
	Description  string      `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	PusherType   string      `protobuf:"bytes,5,opt,name=pusher_type,json=pusherType" json:"pusher_type,omitempty"`
	Repository   *Repository `protobuf:"bytes,6,opt,name=repository" json:"repository,omitempty"`
	Sender       *User       `protobuf:"bytes,7,opt,name=sender" json:"sender,omitempty"`
}

func (m *CreateEvent) Reset()                    { *m = CreateEvent{} }
func (m *CreateEvent) String() string            { return proto.CompactTextString(m) }
func (*CreateEvent) ProtoMessage()               {}
func (*CreateEvent) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *CreateEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CreateEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEvent)(nil), "github.CreateEvent")
}

func init() { proto.RegisterFile("create_event.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0x87, 0x55, 0x06, 0x1d, 0xb8, 0x43, 0x9a, 0x7c, 0x0a, 0xbb, 0x30, 0x01, 0x07, 0x4e, 0x3d,
	0x8c, 0x37, 0x00, 0xf1, 0x02, 0x15, 0x9c, 0xab, 0xb6, 0x33, 0x34, 0x07, 0x92, 0xc8, 0x71, 0x91,
	0xf6, 0xee, 0x1c, 0xc8, 0x1c, 0xfe, 0xf4, 0xd4, 0xfa, 0xfb, 0x7d, 0x8e, 0x13, 0x03, 0x0e, 0x4c,
	0x9d, 0x50, 0x4b, 0x9f, 0xe4, 0xa4, 0x0e, 0xec, 0xc5, 0x63, 0xf9, 0x6e, 0x65, 0x9c, 0xfa, 0x0d,
	0x4c, 0x91, 0x38, 0xb3, 0xcd, 0x9a, 0x29, 0xf8, 0x68, 0xc5, 0xf3, 0x21, 0x93, 0x9b, 0xaf, 0x02,
	0xaa, 0x27, 0x6d, 0x7e, 0x3e, 0xf6, 0xe2, 0x1a, 0x16, 0x4c, 0x6f, 0xa6, 0xd8, 0x16, 0xf7, 0x17,
	0xcd, 0xf1, 0x17, 0xaf, 0xe0, 0x3c, 0x7d, 0x5a, 0x39, 0x04, 0x32, 0x27, 0x8a, 0x97, 0xa9, 0x7e,
	0x49, 0x25, 0xde, 0xc2, 0xe5, 0x47, 0x17, 0x85, 0xb8, 0xed, 0xb9, 0x73, 0xc3, 0x68, 0x16, 0x9a,
	0xaf, 0x32, 0x7c, 0x54, 0x86, 0x5b, 0xa8, 0xf6, 0x14, 0x07, 0xb6, 0x41, 0xac, 0x77, 0xe6, 0x54,
	0x95, 0x39, 0xc2, 0x6b, 0xa8, 0xc2, 0x14, 0xc7, 0x74, 0x8c, 0x0e, 0x39, 0x53, 0x03, 0x32, 0xd2,
	0x39, 0x3b, 0x80, 0xff, 0x8b, 0x9b, 0x32, 0xe5, 0xd5, 0x0e, 0xeb, 0xfc, 0xbe, 0xba, 0xf9, 0x4b,
	0x9a, 0x99, 0x85, 0x77, 0x50, 0x46, 0x72, 0x7b, 0x62, 0xb3, 0x54, 0x7f, 0xf5, 0xeb, 0xbf, 0xa6,
	0x75, 0x34, 0x3f, 0x59, 0x5f, 0xea, 0x16, 0x1e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x83,
	0x33, 0xb5, 0x41, 0x01, 0x00, 0x00,
}