// Code generated by protoc-gen-go.
// source: team.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Team struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id              int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Slug            string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	Description     string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Permission      string `protobuf:"bytes,5,opt,name=permission" json:"permission,omitempty"`
	Url             string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	MembersUrl      string `protobuf:"bytes,7,opt,name=members_url,json=membersUrl" json:"members_url,omitempty"`
	RepositoriesUrl string `protobuf:"bytes,8,opt,name=repositories_url,json=repositoriesUrl" json:"repositories_url,omitempty"`
}

func (m *Team) Reset()                    { *m = Team{} }
func (m *Team) String() string            { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()               {}
func (*Team) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

func init() {
	proto.RegisterType((*Team)(nil), "github.Team")
}

func init() { proto.RegisterFile("team.proto", fileDescriptor32) }

var fileDescriptor32 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x8e, 0xc2, 0x30,
	0x10, 0x46, 0x95, 0xdf, 0xdd, 0x9d, 0x48, 0x4b, 0x34, 0x95, 0x2b, 0x88, 0xa8, 0xa0, 0xa1, 0xe1,
	0x26, 0x11, 0xd4, 0x28, 0x21, 0xa3, 0x60, 0x29, 0x8e, 0xad, 0xb1, 0x73, 0x5d, 0xce, 0x82, 0xed,
	0x50, 0xa4, 0x7b, 0x7a, 0xdf, 0x9b, 0x62, 0x00, 0x1c, 0x75, 0xea, 0x62, 0x58, 0x3b, 0x8d, 0xe5,
	0x28, 0xdd, 0x6b, 0xe9, 0x8f, 0xef, 0x04, 0xf2, 0x9b, 0xd7, 0x88, 0x90, 0xcf, 0x9d, 0x22, 0x91,
	0x34, 0xc9, 0xe9, 0xaf, 0x8d, 0x8c, 0xff, 0x90, 0xca, 0x41, 0xa4, 0xde, 0x14, 0xad, 0xa7, 0xd0,
	0xd8, 0x69, 0x19, 0x45, 0xb6, 0x36, 0x81, 0xb1, 0x81, 0x6a, 0x20, 0xfb, 0x64, 0x69, 0x9c, 0xd4,
	0xb3, 0xc8, 0xe3, 0xb4, 0x55, 0xb8, 0x07, 0x30, 0xc4, 0x4a, 0x5a, 0x1b, 0x82, 0x22, 0x06, 0x1b,
	0x83, 0x35, 0x64, 0x0b, 0x4f, 0xa2, 0x8c, 0x43, 0x40, 0x3c, 0x40, 0xa5, 0x48, 0xf5, 0xc4, 0xf6,
	0x11, 0x96, 0x9f, 0xf5, 0xe4, 0xab, 0xee, 0x3e, 0x38, 0x43, 0xcd, 0x64, 0xb4, 0x95, 0x4e, 0xb3,
	0xa4, 0xb5, 0xfa, 0x8d, 0xd5, 0x6e, 0xeb, 0x7d, 0xda, 0x97, 0xf1, 0xdf, 0xeb, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x12, 0xae, 0x55, 0xbe, 0xfd, 0x00, 0x00, 0x00,
}