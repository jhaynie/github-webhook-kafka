// Code generated by protoc-gen-go.
// source: comment.proto
// DO NOT EDIT!

/*
Package github is a generated protocol buffer package.

It is generated from these files:
	comment.proto
	commit_comment_event.proto
	commit_user.proto
	create_event.proto
	delete_event.proto
	deployment.proto
	deployment_event.proto
	deployment_status_event.proto
	download_event.proto
	follow_event.proto
	fork_apply_event.proto
	fork_event.proto
	gist_event.proto
	gollum_event.proto
	issue.proto
	issue_comment_event.proto
	issues_event.proto
	member_event.proto
	membership_event.proto
	milestone.proto
	page_build_event.proto
	public_event.proto
	pull_request_event.proto
	pull_request_review_comment_event.proto
	pull_request_review_event.proto
	pullrequest.proto
	push_event.proto
	release.proto
	release_event.proto
	repository.proto
	repository_event.proto
	status_event.proto
	team.proto
	team_add.proto
	user.proto
	watch_event.proto

It has these top-level messages:
	Comment
	CommitCommentEvent
	CommitUser
	CreateEvent
	DeleteEvent
	Deployment
	DeploymentEvent
	DeploymentStatus
	DeploymentStatusEvent
	Download
	DownloadEvent
	FollowEvent
	ForkApplyEvent
	ForkEvent
	GistFile
	GistFork
	GistChangeStatus
	GistHistory
	Gist
	GistEvent
	GollumPage
	GollumEvent
	Label
	Issue
	IssueCommentEvent
	IssuesEvent
	MemberEvent
	MembershipEvent
	Milestone
	BuildError
	Build
	PageBuildEvent
	PublicEvent
	PullRequestEvent
	PullRequestReviewCommentLink
	PullRequestReviewCommentLinks
	PullRequestReviewComment
	PullRequestReviewCommentEvent
	PullRequestReviewLink
	PullRequestReviewLinks
	PullRequestReview
	PullRequestReviewEvent
	RepositoryRef
	LinkHref
	Links
	PullRequest
	Commit
	PushEvent
	Asset
	Release
	ReleaseEvent
	RepositoryPermission
	Repository
	RepositoryEvent
	CommitDetailTree
	CommitBranch
	CommitDetail
	CommitUpdate
	StatusEvent
	Team
	TeamAddEvent
	User
	WatchEvent
*/
package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Comment struct {
	Url       string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	HtmlUrl   string `protobuf:"bytes,2,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	Id        int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	User      *User  `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Position  int32  `protobuf:"varint,5,opt,name=position" json:"position,omitempty"`
	Line      int32  `protobuf:"varint,6,opt,name=line" json:"line,omitempty"`
	Path      string `protobuf:"bytes,7,opt,name=path" json:"path,omitempty"`
	CommitId  string `protobuf:"bytes,8,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Body      string `protobuf:"bytes,11,opt,name=body" json:"body,omitempty"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Comment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Comment)(nil), "github.Comment")
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xd5, 0x34, 0xcd, 0x9f, 0x2b, 0x20, 0x74, 0x93, 0x29, 0x42, 0xaa, 0x98, 0x98, 0x32,
	0xc0, 0x13, 0x20, 0x26, 0xd6, 0x48, 0x9d, 0xa3, 0xa4, 0xb6, 0x88, 0xa5, 0x24, 0x8e, 0xd2, 0xcb,
	0xc0, 0xb3, 0xf0, 0xb2, 0xdc, 0x9d, 0x0b, 0xdb, 0x77, 0xbf, 0x9f, 0xed, 0x4f, 0x3e, 0xb8, 0x3d,
	0x87, 0x71, 0x74, 0x13, 0x55, 0xf3, 0x12, 0x28, 0x60, 0xf6, 0xe5, 0xa9, 0x5f, 0xbb, 0x03, 0xac,
	0x17, 0xb7, 0x44, 0xf6, 0xfc, 0x93, 0x40, 0xfe, 0x11, 0x4f, 0xe1, 0x3d, 0x6c, 0xd7, 0x65, 0x30,
	0x9b, 0xe3, 0xe6, 0xa5, 0xac, 0x25, 0xe2, 0x03, 0x14, 0x3d, 0x8d, 0x43, 0x23, 0x38, 0x51, 0x9c,
	0xcb, 0x7c, 0x62, 0x75, 0x07, 0x89, 0xb7, 0x66, 0xcb, 0x70, 0x57, 0x73, 0xc2, 0x23, 0xa4, 0xf2,
	0xac, 0x49, 0x99, 0xec, 0x5f, 0x6f, 0xaa, 0xd8, 0x55, 0x9d, 0x98, 0xd5, 0x6a, 0xf0, 0x00, 0xc5,
	0x1c, 0x2e, 0x9e, 0x7c, 0x98, 0xcc, 0x4e, 0xef, 0xfd, 0xcf, 0x88, 0x90, 0x0e, 0x7e, 0x72, 0x26,
	0x53, 0xae, 0x59, 0xd8, 0xdc, 0x52, 0x6f, 0x72, 0x2d, 0xd6, 0x8c, 0x8f, 0x50, 0xca, 0x9f, 0x3c,
	0x35, 0x5c, 0x5e, 0xa8, 0x28, 0x22, 0xf8, 0xb4, 0xf8, 0x04, 0x70, 0x5e, 0x5c, 0x4b, 0xce, 0x36,
	0x2d, 0x99, 0x52, 0x6d, 0x79, 0x25, 0xef, 0x24, 0x7a, 0x9d, 0xed, 0x9f, 0x86, 0xa8, 0xaf, 0x84,
	0x35, 0xd7, 0x75, 0xc1, 0x7e, 0x9b, 0x7d, 0xac, 0x93, 0xdc, 0x65, 0xba, 0xa4, 0xb7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x12, 0x1a, 0x94, 0x49, 0x01, 0x00, 0x00,
}
