// Code generated by protoc-gen-go.
// source: milestone.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Milestone struct {
	Url          string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	HtmlUrl      string `protobuf:"bytes,2,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	LabelsUrl    string `protobuf:"bytes,3,opt,name=labels_url,json=labelsUrl" json:"labels_url,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	Number       int32  `protobuf:"varint,5,opt,name=number" json:"number,omitempty"`
	State        string `protobuf:"bytes,6,opt,name=state" json:"state,omitempty"`
	Title        string `protobuf:"bytes,7,opt,name=title" json:"title,omitempty"`
	Description  string `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
	Creator      *User  `protobuf:"bytes,9,opt,name=creator" json:"creator,omitempty"`
	OpenIssues   int32  `protobuf:"varint,10,opt,name=open_issues,json=openIssues" json:"open_issues,omitempty"`
	ClosedIssues int32  `protobuf:"varint,11,opt,name=closed_issues,json=closedIssues" json:"closed_issues,omitempty"`
	CreatedAt    string `protobuf:"bytes,12,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	ClosedAt     string `protobuf:"bytes,14,opt,name=closed_at,json=closedAt" json:"closed_at,omitempty"`
	DueOn        string `protobuf:"bytes,15,opt,name=due_on,json=dueOn" json:"due_on,omitempty"`
}

func (m *Milestone) Reset()                    { *m = Milestone{} }
func (m *Milestone) String() string            { return proto.CompactTextString(m) }
func (*Milestone) ProtoMessage()               {}
func (*Milestone) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *Milestone) GetCreator() *User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func init() {
	proto.RegisterType((*Milestone)(nil), "github.Milestone")
}

func init() { proto.RegisterFile("milestone.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x6b, 0xd3, 0x66, 0xd2, 0x0f, 0x59, 0x54, 0x56, 0x45, 0x2c, 0x0a, 0xe2, 0xa9,
	0x07, 0xfd, 0x05, 0x1e, 0x3d, 0x88, 0x50, 0xe8, 0xb9, 0xe4, 0x63, 0xb0, 0x0b, 0x9b, 0x6c, 0xd8,
	0x8f, 0x9f, 0xe5, 0x7f, 0x74, 0x76, 0x36, 0x01, 0x6f, 0xfb, 0x3e, 0xcf, 0xcb, 0x26, 0x33, 0x0b,
	0xdb, 0x56, 0x69, 0x74, 0xde, 0x74, 0xb8, 0xef, 0xad, 0xf1, 0x46, 0x64, 0x3f, 0xca, 0x9f, 0x43,
	0x75, 0x07, 0xc1, 0xa1, 0x4d, 0xec, 0xe9, 0x77, 0x06, 0xf9, 0xd7, 0xd8, 0x13, 0x97, 0x30, 0x0b,
	0x56, 0xcb, 0xc9, 0x6e, 0xf2, 0x9a, 0x1f, 0xe2, 0x51, 0xdc, 0xc2, 0xf2, 0xec, 0x5b, 0x7d, 0x8a,
	0x78, 0xca, 0x78, 0x11, 0xf3, 0x91, 0xd4, 0x03, 0x80, 0x2e, 0x2b, 0xd4, 0x8e, 0xe5, 0x8c, 0x65,
	0x9e, 0x48, 0xd4, 0x1b, 0x98, 0xaa, 0x46, 0x5e, 0x10, 0x9e, 0x1f, 0xe8, 0x24, 0x6e, 0x20, 0xeb,
	0x42, 0x5b, 0xa1, 0x95, 0x73, 0x66, 0x43, 0x12, 0x57, 0x30, 0x77, 0xbe, 0xf4, 0x28, 0x33, 0xbe,
	0x21, 0x85, 0x48, 0xbd, 0xf2, 0x1a, 0xe5, 0x22, 0x51, 0x0e, 0x62, 0x07, 0x45, 0x83, 0xae, 0xb6,
	0xaa, 0xf7, 0xca, 0x74, 0x72, 0xc9, 0xee, 0x3f, 0x12, 0x2f, 0xb0, 0xa8, 0x2d, 0x96, 0xde, 0x58,
	0x99, 0x93, 0x2d, 0xde, 0x56, 0xfb, 0x34, 0xf5, 0xfe, 0x48, 0x43, 0x1f, 0x46, 0x29, 0x1e, 0xa1,
	0x30, 0x3d, 0x76, 0x27, 0xe5, 0x5c, 0x40, 0x27, 0x81, 0x7f, 0x09, 0x22, 0xfa, 0x64, 0x22, 0x9e,
	0x61, 0x5d, 0x6b, 0xe3, 0xb0, 0x19, 0x2b, 0x05, 0x57, 0x56, 0x09, 0x0e, 0x25, 0x5a, 0x01, 0x5f,
	0x48, 0xad, 0xd2, 0xcb, 0x55, 0x5a, 0xc1, 0x40, 0x3e, 0x7c, 0xd4, 0xa1, 0x6f, 0x46, 0xbd, 0x4e,
	0x7a, 0x20, 0xa4, 0xef, 0x21, 0x1f, 0x3e, 0x41, 0x76, 0xc3, 0x76, 0x99, 0x00, 0xc9, 0x6b, 0xc8,
	0x9a, 0x80, 0x27, 0x9a, 0x72, 0x9b, 0x36, 0x40, 0xe9, 0xbb, 0xab, 0x32, 0x7e, 0xb6, 0xf7, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x41, 0x81, 0xe5, 0xdd, 0x01, 0x00, 0x00,
}