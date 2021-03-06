// Code generated by protoc-gen-go.
// source: push_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Commit struct {
	Id        string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TreeId    string      `protobuf:"bytes,2,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	Distinct  bool        `protobuf:"varint,3,opt,name=distinct" json:"distinct,omitempty"`
	Message   string      `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	Timestamp string      `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Url       string      `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	Author    *CommitUser `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Committer *CommitUser `protobuf:"bytes,8,opt,name=committer" json:"committer,omitempty"`
	Added     []string    `protobuf:"bytes,9,rep,name=added" json:"added,omitempty"`
	Removed   []string    `protobuf:"bytes,10,rep,name=removed" json:"removed,omitempty"`
	Modified  []string    `protobuf:"bytes,11,rep,name=modified" json:"modified,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *Commit) GetAuthor() *CommitUser {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Commit) GetCommitter() *CommitUser {
	if m != nil {
		return m.Committer
	}
	return nil
}

type PushEvent struct {
	Ref        string      `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Before     string      `protobuf:"bytes,2,opt,name=before" json:"before,omitempty"`
	After      string      `protobuf:"bytes,3,opt,name=after" json:"after,omitempty"`
	Created    bool        `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Deleted    bool        `protobuf:"varint,5,opt,name=deleted" json:"deleted,omitempty"`
	Forced     bool        `protobuf:"varint,6,opt,name=forced" json:"forced,omitempty"`
	BaseRef    string      `protobuf:"bytes,7,opt,name=base_ref,json=baseRef" json:"base_ref,omitempty"`
	Compare    string      `protobuf:"bytes,8,opt,name=compare" json:"compare,omitempty"`
	Commits    []*Commit   `protobuf:"bytes,9,rep,name=commits" json:"commits,omitempty"`
	HeadCommit *Commit     `protobuf:"bytes,10,opt,name=head_commit,json=headCommit" json:"head_commit,omitempty"`
	Repository *Repository `protobuf:"bytes,11,opt,name=repository" json:"repository,omitempty"`
	Pusher     *CommitUser `protobuf:"bytes,12,opt,name=pusher" json:"pusher,omitempty"`
	Sender     *User       `protobuf:"bytes,13,opt,name=sender" json:"sender,omitempty"`
}

func (m *PushEvent) Reset()                    { *m = PushEvent{} }
func (m *PushEvent) String() string            { return proto.CompactTextString(m) }
func (*PushEvent) ProtoMessage()               {}
func (*PushEvent) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

func (m *PushEvent) GetCommits() []*Commit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *PushEvent) GetHeadCommit() *Commit {
	if m != nil {
		return m.HeadCommit
	}
	return nil
}

func (m *PushEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PushEvent) GetPusher() *CommitUser {
	if m != nil {
		return m.Pusher
	}
	return nil
}

func (m *PushEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*Commit)(nil), "github.Commit")
	proto.RegisterType((*PushEvent)(nil), "github.PushEvent")
}

func init() { proto.RegisterFile("push_event.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x9b, 0xad, 0x9b, 0x4c, 0x97, 0x55, 0xb1, 0x10, 0x98, 0x8a, 0xc3, 0x6a, 0xc5, 0x61,
	0xc5, 0xa1, 0xa0, 0xf2, 0x08, 0x88, 0x03, 0x37, 0x14, 0x89, 0x73, 0xe4, 0xd6, 0xd3, 0xad, 0xa5,
	0x4d, 0x1d, 0xd9, 0xce, 0x4a, 0x3c, 0x14, 0x0f, 0xc1, 0x9b, 0x31, 0x63, 0x27, 0xed, 0xae, 0x44,
	0x6f, 0xfe, 0x7e, 0x12, 0x7f, 0xf3, 0x79, 0x60, 0xd9, 0xf5, 0xe1, 0xd0, 0xe0, 0x13, 0x1e, 0xe3,
	0xba, 0xf3, 0x2e, 0x3a, 0x29, 0x1e, 0x6c, 0x3c, 0xf4, 0xdb, 0x15, 0xf4, 0x01, 0x7d, 0xe6, 0x56,
	0xaf, 0x77, 0xae, 0x6d, 0x6d, 0x6c, 0x9e, 0x51, 0x4b, 0x8f, 0x9d, 0x0b, 0x36, 0x3a, 0xff, 0x3b,
	0x33, 0x77, 0x7f, 0xa7, 0x20, 0xbe, 0x25, 0x9f, 0xbc, 0x81, 0xa9, 0x35, 0x6a, 0x72, 0x3b, 0xb9,
	0xaf, 0x6a, 0x3a, 0xc9, 0x77, 0x30, 0x8f, 0x1e, 0xb1, 0x21, 0x72, 0x9a, 0x48, 0xc1, 0xf0, 0x87,
	0x91, 0x2b, 0x28, 0x8d, 0x0d, 0xd1, 0x1e, 0x77, 0x51, 0x15, 0xa4, 0x94, 0xf5, 0x09, 0x4b, 0x05,
	0xf3, 0x16, 0x43, 0xd0, 0x0f, 0xa8, 0xae, 0xd2, 0x47, 0x23, 0x94, 0x1f, 0xa0, 0x8a, 0x96, 0x40,
	0xd4, 0x6d, 0xa7, 0x66, 0x49, 0x3b, 0x13, 0x72, 0x09, 0x45, 0xef, 0x1f, 0x95, 0x48, 0x3c, 0x1f,
	0xe5, 0x27, 0x10, 0xba, 0x8f, 0x07, 0xe7, 0xd5, 0x9c, 0xc8, 0xc5, 0x46, 0xae, 0xf3, 0x8c, 0xeb,
	0x1c, 0xf7, 0x17, 0x4d, 0x55, 0x0f, 0x0e, 0xf9, 0x05, 0xaa, 0x3c, 0x6c, 0x44, 0xaf, 0xca, 0x8b,
	0xf6, 0xb3, 0x49, 0xbe, 0x81, 0x99, 0x36, 0x06, 0x8d, 0xaa, 0x6e, 0x0b, 0xba, 0x31, 0x03, 0x4e,
	0xef, 0xb1, 0x75, 0x4f, 0xc4, 0x43, 0xe2, 0x47, 0xc8, 0x33, 0xb7, 0xce, 0xd8, 0xbd, 0x25, 0x69,
	0x91, 0xa4, 0x13, 0xbe, 0xfb, 0x53, 0x40, 0xf5, 0x93, 0x5e, 0xe4, 0x3b, 0x3f, 0x08, 0x4f, 0xe2,
	0x71, 0x3f, 0xf4, 0xc8, 0x47, 0xf9, 0x16, 0xc4, 0x16, 0xf7, 0xce, 0xe3, 0xd8, 0x63, 0x46, 0x29,
	0xc3, 0x9e, 0x13, 0x17, 0x89, 0xce, 0x80, 0x33, 0xec, 0x3c, 0xea, 0x48, 0x17, 0x5d, 0xa5, 0x72,
	0x47, 0xc8, 0x8a, 0xc1, 0x47, 0x64, 0x65, 0x96, 0x95, 0x01, 0xf2, 0x0d, 0xf4, 0xc7, 0x1d, 0x09,
	0x22, 0x09, 0x03, 0x92, 0xef, 0xa1, 0xdc, 0xea, 0x80, 0x0d, 0x07, 0x9a, 0xe7, 0xe7, 0x60, 0x5c,
	0x53, 0x28, 0xbe, 0xc6, 0xb5, 0x9d, 0xa6, 0x54, 0x65, 0x56, 0x06, 0x28, 0xef, 0x93, 0x42, 0x3d,
	0x85, 0x54, 0xce, 0x62, 0x73, 0xf3, 0xb2, 0xca, 0x7a, 0x94, 0xe5, 0x67, 0x58, 0x1c, 0x50, 0x9b,
	0x26, 0x63, 0xaa, 0x6c, 0xf2, 0x1f, 0x37, 0xb0, 0x65, 0x58, 0xb1, 0x0d, 0xc0, 0x79, 0x03, 0xa9,
	0xc7, 0x17, 0x0f, 0x55, 0x9f, 0x94, 0xfa, 0x99, 0x8b, 0xf7, 0x80, 0xd7, 0x9d, 0x6a, 0xba, 0xbe,
	0xbc, 0x07, 0xd9, 0x21, 0x3f, 0x82, 0x08, 0x78, 0x34, 0xe4, 0x7d, 0x95, 0xbc, 0xd7, 0xa3, 0x37,
	0xbb, 0xb2, 0xb6, 0x15, 0x69, 0xf5, 0xbf, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x37, 0x94, 0xcc,
	0x1d, 0x47, 0x03, 0x00, 0x00,
}
