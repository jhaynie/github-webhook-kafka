// Code generated by protoc-gen-go.
// source: commit_user.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommitUser struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *CommitUser) Reset()                    { *m = CommitUser{} }
func (m *CommitUser) String() string            { return proto.CompactTextString(m) }
func (*CommitUser) ProtoMessage()               {}
func (*CommitUser) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*CommitUser)(nil), "github.CommitUser")
}

func init() { proto.RegisterFile("commit_user.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0x2c, 0x89, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x52, 0x0a, 0xe2, 0xe2, 0x72, 0x06, 0x4b, 0x86, 0x02, 0xe5, 0x84,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x23, 0x24,
	0xc5, 0xc5, 0x01, 0x32, 0x0d, 0xac, 0x9a, 0x19, 0x2c, 0x01, 0xe7, 0x27, 0xb1, 0x81, 0xad, 0x30,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xb3, 0x07, 0xc8, 0x77, 0x00, 0x00, 0x00,
}
