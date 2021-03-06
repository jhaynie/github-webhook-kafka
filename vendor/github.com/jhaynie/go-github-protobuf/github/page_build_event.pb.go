// Code generated by protoc-gen-go.
// source: page_build_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuildError struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *BuildError) Reset()                    { *m = BuildError{} }
func (m *BuildError) String() string            { return proto.CompactTextString(m) }
func (*BuildError) ProtoMessage()               {}
func (*BuildError) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

type Build struct {
	Url       string      `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Status    string      `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Error     *BuildError `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Pusher    *User       `protobuf:"bytes,4,opt,name=pusher" json:"pusher,omitempty"`
	Commit    string      `protobuf:"bytes,5,opt,name=commit" json:"commit,omitempty"`
	Duration  int32       `protobuf:"varint,6,opt,name=duration" json:"duration,omitempty"`
	CreatedAt string      `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

func (m *Build) GetError() *BuildError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Build) GetPusher() *User {
	if m != nil {
		return m.Pusher
	}
	return nil
}

type PageBuildEvent struct {
	Id         int32       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Build      *Build      `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	Repository *Repository `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
}

func (m *PageBuildEvent) Reset()                    { *m = PageBuildEvent{} }
func (m *PageBuildEvent) String() string            { return proto.CompactTextString(m) }
func (*PageBuildEvent) ProtoMessage()               {}
func (*PageBuildEvent) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

func (m *PageBuildEvent) GetBuild() *Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *PageBuildEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PageBuildEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildError)(nil), "github.BuildError")
	proto.RegisterType((*Build)(nil), "github.Build")
	proto.RegisterType((*PageBuildEvent)(nil), "github.PageBuildEvent")
}

func init() { proto.RegisterFile("page_build_event.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6a, 0x02, 0x31,
	0x10, 0xc6, 0x59, 0x6d, 0x56, 0x1d, 0x5b, 0x91, 0x39, 0x48, 0x10, 0x0a, 0xc5, 0x96, 0xe2, 0xc9,
	0x83, 0x7d, 0x02, 0x0b, 0xbd, 0x97, 0x40, 0xcf, 0xb2, 0x9a, 0xa0, 0x01, 0x35, 0x4b, 0x32, 0x29,
	0xf4, 0x69, 0xfa, 0x90, 0x7d, 0x81, 0xe6, 0xcf, 0xae, 0xb6, 0x85, 0xde, 0x32, 0xdf, 0xef, 0x4b,
	0x66, 0xe6, 0x0b, 0x4c, 0xea, 0x6a, 0xa7, 0xd6, 0x1b, 0xaf, 0x0f, 0x72, 0xad, 0xde, 0xd5, 0x89,
	0x16, 0xb5, 0x35, 0x64, 0xb0, 0xdc, 0x69, 0xda, 0xfb, 0xcd, 0x14, 0xbc, 0x53, 0x36, 0x6b, 0xd3,
	0xb1, 0x55, 0xb5, 0x71, 0x9a, 0x8c, 0xfd, 0xc8, 0xca, 0xec, 0x11, 0xe0, 0x39, 0x5e, 0x7d, 0xb1,
	0xd6, 0x58, 0xe4, 0xd0, 0x3b, 0x2a, 0xe7, 0xc2, 0x83, 0xbc, 0xb8, 0x2b, 0xe6, 0x03, 0xd1, 0x96,
	0xb3, 0xaf, 0x02, 0x58, 0x32, 0xe2, 0x18, 0xba, 0xde, 0x1e, 0x1a, 0x1e, 0x8f, 0x38, 0x81, 0xd2,
	0x51, 0x45, 0xde, 0xf1, 0x4e, 0x12, 0x9b, 0x0a, 0xe7, 0xc0, 0x54, 0x7c, 0x96, 0x77, 0x83, 0x3c,
	0x5c, 0xe2, 0x22, 0x4f, 0xb4, 0xb8, 0x34, 0x14, 0xd9, 0x80, 0x0f, 0x50, 0xd6, 0xde, 0xed, 0x95,
	0xe5, 0x57, 0xc9, 0x7a, 0xdd, 0x5a, 0xdf, 0xc2, 0xec, 0xa2, 0x61, 0xb1, 0xcf, 0xd6, 0x1c, 0x8f,
	0x9a, 0x38, 0xcb, 0x7d, 0x72, 0x85, 0x53, 0xe8, 0x4b, 0x6f, 0x2b, 0xd2, 0xe6, 0xc4, 0xcb, 0x40,
	0x98, 0x38, 0xd7, 0x78, 0x0b, 0xb0, 0xb5, 0xaa, 0x22, 0x25, 0xd7, 0x15, 0xf1, 0x5e, 0xba, 0x37,
	0x68, 0x94, 0x15, 0x45, 0xec, 0x6b, 0xd9, 0xe2, 0x7e, 0xc6, 0x8d, 0xb2, 0xa2, 0xd9, 0x67, 0x01,
	0xa3, 0xd7, 0xb0, 0x7e, 0x9e, 0x38, 0x86, 0x8b, 0x23, 0xe8, 0x68, 0x99, 0xb6, 0x67, 0x22, 0x9c,
	0xf0, 0x1e, 0x58, 0xca, 0x3e, 0xed, 0x3e, 0x5c, 0xde, 0xfc, 0x5a, 0x52, 0x64, 0x86, 0x4b, 0x80,
	0x4b, 0xf2, 0x7f, 0xe3, 0x10, 0x67, 0x22, 0x7e, 0xb8, 0x62, 0x26, 0x4e, 0x9d, 0xe4, 0x7f, 0x99,
	0x64, 0xb6, 0x29, 0xd3, 0x37, 0x3e, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x32, 0xed, 0x36, 0x86,
	0x06, 0x02, 0x00, 0x00,
}
