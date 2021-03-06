// Code generated by protoc-gen-go.
// source: release.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Asset struct {
	Url                string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	BrowserDownloadUrl string `protobuf:"bytes,2,opt,name=browser_download_url,json=browserDownloadUrl" json:"browser_download_url,omitempty"`
	Id                 int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	Name               string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Label              string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	State              string `protobuf:"bytes,6,opt,name=state" json:"state,omitempty"`
	ContentType        string `protobuf:"bytes,7,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Size               int32  `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
	DownloadCount      int32  `protobuf:"varint,9,opt,name=download_count,json=downloadCount" json:"download_count,omitempty"`
	CreatedAt          string `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Uploader           *User  `protobuf:"bytes,12,opt,name=uploader" json:"uploader,omitempty"`
}

func (m *Asset) Reset()                    { *m = Asset{} }
func (m *Asset) String() string            { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()               {}
func (*Asset) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

func (m *Asset) GetUploader() *User {
	if m != nil {
		return m.Uploader
	}
	return nil
}

type Release struct {
	Url             string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	HtmlUrl         string   `protobuf:"bytes,2,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	AssetsUrl       string   `protobuf:"bytes,3,opt,name=assets_url,json=assetsUrl" json:"assets_url,omitempty"`
	UploadUrl       string   `protobuf:"bytes,4,opt,name=upload_url,json=uploadUrl" json:"upload_url,omitempty"`
	TarballUrl      string   `protobuf:"bytes,5,opt,name=tarball_url,json=tarballUrl" json:"tarball_url,omitempty"`
	ZipballUrl      string   `protobuf:"bytes,6,opt,name=zipball_url,json=zipballUrl" json:"zipball_url,omitempty"`
	Id              int32    `protobuf:"varint,7,opt,name=id" json:"id,omitempty"`
	TagName         string   `protobuf:"bytes,8,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	TargetCommitish string   `protobuf:"bytes,9,opt,name=target_commitish,json=targetCommitish" json:"target_commitish,omitempty"`
	Name            string   `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
	Body            string   `protobuf:"bytes,11,opt,name=body" json:"body,omitempty"`
	Draft           bool     `protobuf:"varint,12,opt,name=draft" json:"draft,omitempty"`
	Prerelease      bool     `protobuf:"varint,13,opt,name=prerelease" json:"prerelease,omitempty"`
	CreatedAt       string   `protobuf:"bytes,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	PublishedAt     string   `protobuf:"bytes,15,opt,name=published_at,json=publishedAt" json:"published_at,omitempty"`
	Author          *User    `protobuf:"bytes,16,opt,name=author" json:"author,omitempty"`
	Assets          []*Asset `protobuf:"bytes,17,rep,name=assets" json:"assets,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

func (m *Release) GetAuthor() *User {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Release) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*Asset)(nil), "github.Asset")
	proto.RegisterType((*Release)(nil), "github.Release")
}

func init() { proto.RegisterFile("release.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xd5, 0x38, 0x71, 0x9c, 0xc9, 0x9f, 0x86, 0x55, 0x0f, 0x0b, 0x12, 0x50, 0x2a, 0x90,
	0xca, 0x25, 0x42, 0xf0, 0x04, 0x55, 0x39, 0x73, 0xb0, 0xe8, 0xd9, 0x5a, 0xc7, 0x4b, 0x62, 0xc9,
	0xf6, 0x5a, 0xeb, 0xb5, 0xaa, 0xf6, 0x6d, 0x79, 0x0c, 0x6e, 0xcc, 0xce, 0xac, 0xdd, 0xa8, 0xca,
	0x6d, 0xe7, 0xf7, 0x8d, 0x77, 0x47, 0xdf, 0x37, 0x86, 0xb5, 0xd5, 0x95, 0x56, 0x9d, 0xde, 0xb5,
	0xd6, 0x38, 0x23, 0xe2, 0x43, 0xe9, 0x8e, 0x7d, 0xfe, 0x0e, 0xfa, 0x4e, 0x5b, 0x66, 0x37, 0x7f,
	0x27, 0x30, 0xbb, 0xeb, 0x3a, 0xed, 0xc4, 0x16, 0xa2, 0xde, 0x56, 0xf2, 0xe2, 0xfa, 0xe2, 0x76,
	0x91, 0xfa, 0xa3, 0xf8, 0x06, 0x57, 0xb9, 0x35, 0x8f, 0xd8, 0x9c, 0x15, 0xe6, 0xb1, 0xa9, 0x8c,
	0x2a, 0x32, 0xdf, 0x32, 0xa1, 0x16, 0x11, 0xb4, 0x9f, 0x41, 0x7a, 0xc0, 0x2f, 0x36, 0x30, 0x29,
	0x0b, 0x19, 0xa1, 0x3e, 0x4b, 0xf1, 0x24, 0x04, 0x4c, 0x1b, 0x55, 0x6b, 0x39, 0xa5, 0x2f, 0xe8,
	0x2c, 0xae, 0x60, 0x56, 0xa9, 0x5c, 0x57, 0x72, 0x46, 0x90, 0x0b, 0x4f, 0x3b, 0xa7, 0x9c, 0x96,
	0x31, 0x53, 0x2a, 0xc4, 0x27, 0x58, 0xed, 0x4d, 0xe3, 0x74, 0xe3, 0x32, 0xf7, 0xd4, 0x6a, 0x39,
	0x27, 0x71, 0x19, 0xd8, 0x6f, 0x44, 0xfe, 0x89, 0xae, 0x7c, 0xd6, 0x32, 0xa1, 0x47, 0xe9, 0x2c,
	0xbe, 0xc0, 0x66, 0x1c, 0x78, 0x6f, 0xfa, 0xc6, 0xc9, 0x05, 0xa9, 0xeb, 0x81, 0xde, 0x7b, 0x28,
	0xde, 0x03, 0xec, 0xad, 0xc6, 0x77, 0x8a, 0x4c, 0x39, 0x09, 0x74, 0xf7, 0x22, 0x90, 0x3b, 0x92,
	0xfb, 0xb6, 0x18, 0xe4, 0x25, 0xcb, 0x81, 0xa0, 0x7c, 0x0b, 0x49, 0xdf, 0xfa, 0xcb, 0xb4, 0x95,
	0x2b, 0x14, 0x97, 0xdf, 0x57, 0x3b, 0x36, 0x78, 0xf7, 0x80, 0xb6, 0xa4, 0xa3, 0x7a, 0xf3, 0x2f,
	0x82, 0x79, 0xca, 0x49, 0x9c, 0x71, 0xf9, 0x2d, 0x24, 0x47, 0x57, 0x57, 0x27, 0xce, 0xce, 0x7d,
	0xed, 0xed, 0xc4, 0x09, 0x94, 0xcf, 0xa6, 0x23, 0x31, 0xe2, 0x09, 0x98, 0x04, 0x99, 0xdf, 0x20,
	0x79, 0x3a, 0x0c, 0x38, 0x84, 0xf1, 0x11, 0x96, 0x4e, 0xd9, 0x5c, 0x55, 0x7c, 0x37, 0xdb, 0x0d,
	0x01, 0x85, 0x86, 0xe7, 0xb2, 0x1d, 0x1b, 0xd8, 0x79, 0x08, 0xe8, 0x25, 0xce, 0xf9, 0x18, 0x27,
	0x8e, 0xea, 0xd4, 0x21, 0xa3, 0x48, 0x13, 0x1e, 0x15, 0xeb, 0x5f, 0x3e, 0xd5, 0xaf, 0xb0, 0xc5,
	0x9b, 0x0f, 0xda, 0xa1, 0xe1, 0x75, 0x5d, 0xba, 0xb2, 0x3b, 0x92, 0xe9, 0x8b, 0xf4, 0x92, 0xf9,
	0xfd, 0x80, 0xc7, 0xa5, 0x80, 0x93, 0xa5, 0x40, 0x96, 0x9b, 0xe2, 0x29, 0xb8, 0x4c, 0x67, 0xbf,
	0x12, 0x85, 0x55, 0x7f, 0x1c, 0xb9, 0x9b, 0xa4, 0x5c, 0x88, 0x0f, 0x00, 0xad, 0xd5, 0x61, 0xb1,
	0xe5, 0x9a, 0xa4, 0x13, 0xf2, 0x2a, 0xd4, 0xcd, 0xeb, 0x50, 0x71, 0xa3, 0xda, 0x3e, 0xaf, 0x70,
	0x0e, 0x6e, 0xb8, 0xe4, 0x8d, 0x1a, 0x19, 0xb6, 0x7c, 0x86, 0x58, 0xf5, 0xee, 0x68, 0xac, 0xdc,
	0x9e, 0x89, 0x35, 0x68, 0xb8, 0x63, 0x31, 0x27, 0x21, 0xdf, 0x5c, 0x47, 0xd8, 0xb5, 0x1e, 0xba,
	0xe8, 0x6f, 0x4a, 0x83, 0x98, 0xc7, 0xf4, 0x9b, 0xfd, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x98,
	0xb2, 0x2a, 0x17, 0x8b, 0x03, 0x00, 0x00,
}
