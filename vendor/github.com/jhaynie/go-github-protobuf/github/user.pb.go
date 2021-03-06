// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Id                int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Login             string `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
	AvatarUrl         string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
	GravatarId        string `protobuf:"bytes,4,opt,name=gravatar_id,json=gravatarId" json:"gravatar_id,omitempty"`
	Url               string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	HtmlUrl           string `protobuf:"bytes,6,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	FollowersUrl      string `protobuf:"bytes,7,opt,name=followers_url,json=followersUrl" json:"followers_url,omitempty"`
	FollowingUrl      string `protobuf:"bytes,8,opt,name=following_url,json=followingUrl" json:"following_url,omitempty"`
	GistsUrl          string `protobuf:"bytes,9,opt,name=gists_url,json=gistsUrl" json:"gists_url,omitempty"`
	StarredUrl        string `protobuf:"bytes,10,opt,name=starred_url,json=starredUrl" json:"starred_url,omitempty"`
	SubscriptionsUrl  string `protobuf:"bytes,11,opt,name=subscriptions_url,json=subscriptionsUrl" json:"subscriptions_url,omitempty"`
	OrganizationsUrl  string `protobuf:"bytes,12,opt,name=organizations_url,json=organizationsUrl" json:"organizations_url,omitempty"`
	ReposUrl          string `protobuf:"bytes,13,opt,name=repos_url,json=reposUrl" json:"repos_url,omitempty"`
	EventsUrl         string `protobuf:"bytes,14,opt,name=events_url,json=eventsUrl" json:"events_url,omitempty"`
	ReceivedEventsUrl string `protobuf:"bytes,15,opt,name=received_events_url,json=receivedEventsUrl" json:"received_events_url,omitempty"`
	Type              string `protobuf:"bytes,16,opt,name=type" json:"type,omitempty"`
	SiteAdmin         bool   `protobuf:"varint,17,opt,name=site_admin,json=siteAdmin" json:"site_admin,omitempty"`
	Name              string `protobuf:"bytes,18,opt,name=name" json:"name,omitempty"`
	Company           string `protobuf:"bytes,19,opt,name=company" json:"company,omitempty"`
	Blog              string `protobuf:"bytes,20,opt,name=blog" json:"blog,omitempty"`
	Location          string `protobuf:"bytes,21,opt,name=location" json:"location,omitempty"`
	Email             string `protobuf:"bytes,22,opt,name=email" json:"email,omitempty"`
	Hireable          bool   `protobuf:"varint,23,opt,name=hireable" json:"hireable,omitempty"`
	Bio               string `protobuf:"bytes,24,opt,name=bio" json:"bio,omitempty"`
	PublicRepos       int32  `protobuf:"varint,25,opt,name=public_repos,json=publicRepos" json:"public_repos,omitempty"`
	PublicGists       int32  `protobuf:"varint,26,opt,name=public_gists,json=publicGists" json:"public_gists,omitempty"`
	Followers         int32  `protobuf:"varint,27,opt,name=followers" json:"followers,omitempty"`
	Following         int32  `protobuf:"varint,28,opt,name=following" json:"following,omitempty"`
	CreatedAt         string `protobuf:"bytes,29,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt         string `protobuf:"bytes,30,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

func init() {
	proto.RegisterType((*User)(nil), "github.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor34) }

var fileDescriptor34 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x93, 0xc9, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x91, 0xc4, 0x8b, 0x34, 0x76, 0x52, 0x9b, 0x49, 0x5b, 0x66, 0xeb, 0x7a, 0x29, 0x50,
	0x20, 0x97, 0x3e, 0x81, 0x0f, 0x45, 0xd1, 0xab, 0x81, 0x9c, 0x0d, 0x4a, 0x62, 0x65, 0x02, 0x94,
	0x28, 0x50, 0x94, 0x8b, 0xf4, 0xd9, 0x7b, 0x08, 0x67, 0x46, 0xb6, 0xe5, 0x1b, 0xe7, 0xfb, 0x3f,
	0xd2, 0x23, 0x72, 0x0c, 0xd0, 0xb5, 0xda, 0x3f, 0x35, 0xde, 0x05, 0x27, 0x26, 0xa5, 0x09, 0xdb,
	0x2e, 0xfb, 0xf2, 0x7f, 0x02, 0xa3, 0xe7, 0x88, 0xc5, 0x15, 0x9c, 0x9b, 0x42, 0x9e, 0x7d, 0x3a,
	0xfb, 0x36, 0x5e, 0xc7, 0x95, 0xb8, 0x81, 0xb1, 0x75, 0xa5, 0xa9, 0xe5, 0x79, 0x44, 0xe9, 0x9a,
	0x0b, 0xf1, 0x08, 0xa0, 0x76, 0x2a, 0x28, 0xbf, 0xe9, 0xbc, 0x95, 0x17, 0x14, 0xa5, 0x4c, 0x9e,
	0xbd, 0x15, 0x1f, 0x61, 0x56, 0xfa, 0x5e, 0x88, 0xa7, 0x8d, 0x28, 0x87, 0x3d, 0xfa, 0x5d, 0x88,
	0x05, 0x5c, 0xe0, 0xc6, 0x31, 0x05, 0xb8, 0x14, 0xb7, 0x90, 0x6c, 0x43, 0x65, 0xe9, 0xbc, 0x09,
	0xe1, 0x29, 0xd6, 0x78, 0xda, 0x57, 0xb8, 0xfc, 0xe3, 0xac, 0x75, 0x7f, 0xb5, 0x6f, 0x29, 0x9f,
	0x52, 0x3e, 0x3f, 0xc0, 0x13, 0xc9, 0xd4, 0x25, 0x49, 0xc9, 0x50, 0x8a, 0x10, 0xa5, 0x7b, 0x48,
	0x4b, 0xd3, 0x06, 0x3e, 0x25, 0x25, 0x21, 0x21, 0xd0, 0x37, 0xdd, 0xc6, 0xee, 0xbc, 0x2e, 0x28,
	0x06, 0x6e, 0xba, 0x47, 0x28, 0x7c, 0x87, 0x65, 0xdb, 0x65, 0x6d, 0xee, 0x4d, 0x13, 0x8c, 0xab,
	0xf9, 0x94, 0x19, 0x69, 0x8b, 0x93, 0xa0, 0x97, 0x9d, 0x2f, 0x55, 0x6d, 0xfe, 0xa9, 0xa3, 0x3c,
	0x67, 0xf9, 0x24, 0xe8, 0xfb, 0xf2, 0xba, 0x71, 0x2c, 0x5d, 0x72, 0x5f, 0x04, 0x30, 0x8c, 0x77,
	0xad, 0x77, 0xba, 0xee, 0xbb, 0xbe, 0xe2, 0xbb, 0x66, 0x82, 0xf1, 0x13, 0x5c, 0x7b, 0x9d, 0x6b,
	0xb3, 0x8b, 0x7d, 0x0f, 0xbc, 0x37, 0xe4, 0x2d, 0xf7, 0xd1, 0xcf, 0x83, 0x2f, 0x60, 0x14, 0x5e,
	0x1a, 0x2d, 0x17, 0x24, 0xd0, 0x1a, 0x7f, 0xa2, 0x35, 0x41, 0x6f, 0x54, 0x51, 0xc5, 0x97, 0x5e,
	0xc6, 0x24, 0x59, 0xa7, 0x48, 0x56, 0x08, 0x70, 0x4b, 0xad, 0x2a, 0x2d, 0x05, 0x6f, 0xc1, 0xb5,
	0x90, 0x30, 0xcd, 0x5d, 0xd5, 0xa8, 0xfa, 0x45, 0x5e, 0xf3, 0x73, 0xf5, 0x25, 0xda, 0x59, 0x9c,
	0x12, 0x79, 0xc3, 0x36, 0xae, 0xc5, 0x1d, 0x24, 0xd6, 0xe5, 0xf4, 0xc1, 0xf2, 0x2d, 0x7f, 0xdf,
	0xbe, 0xc6, 0x09, 0xd3, 0x95, 0x32, 0x56, 0xbe, 0xe3, 0x09, 0xa3, 0x02, 0x77, 0x6c, 0x8d, 0xd7,
	0x2a, 0xb3, 0x5a, 0xbe, 0xa7, 0x86, 0x0e, 0x35, 0x4e, 0x4f, 0x66, 0x9c, 0x94, 0x3c, 0x3d, 0x71,
	0x29, 0x3e, 0xc3, 0xbc, 0xe9, 0x32, 0x6b, 0xf2, 0x0d, 0x5d, 0x9b, 0xbc, 0xa5, 0xf9, 0x9d, 0x31,
	0x5b, 0x23, 0x1a, 0x28, 0xf4, 0xe2, 0xf2, 0x6e, 0xa8, 0xfc, 0x42, 0x24, 0x1e, 0x20, 0x3d, 0xcc,
	0x94, 0xbc, 0xa7, 0xfc, 0x08, 0x8e, 0x69, 0x1c, 0x26, 0xf9, 0x30, 0x4c, 0x23, 0xc0, 0x2b, 0xcc,
	0x63, 0x7b, 0x21, 0xbe, 0x82, 0x0a, 0xf2, 0x91, 0x5f, 0xa9, 0x27, 0xab, 0x80, 0x71, 0xd7, 0x14,
	0xfb, 0xf8, 0x03, 0xc7, 0x3d, 0x59, 0x85, 0x6c, 0x42, 0xff, 0xc6, 0x1f, 0xaf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0xd4, 0x60, 0x06, 0x9b, 0x03, 0x00, 0x00,
}
