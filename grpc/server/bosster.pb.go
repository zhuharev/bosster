// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bosster.proto

/*
Package bosster is a generated protocol buffer package.

It is generated from these files:
	bosster.proto

It has these top-level messages:
	Contact
	Post
	PostRequest
	PostJob
	PostReply
*/
package bosster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SocialType int32

const (
	// enum must start from zero
	// we will return error if ST == none
	SocialType_NONE      SocialType = 0
	SocialType_VK        SocialType = 1
	SocialType_TELEGRAM  SocialType = 2
	SocialType_INSTAGRAM SocialType = 3
	SocialType_TWITTER   SocialType = 4
	SocialType_FACEBOOK  SocialType = 5
)

var SocialType_name = map[int32]string{
	0: "NONE",
	1: "VK",
	2: "TELEGRAM",
	3: "INSTAGRAM",
	4: "TWITTER",
	5: "FACEBOOK",
}
var SocialType_value = map[string]int32{
	"NONE":      0,
	"VK":        1,
	"TELEGRAM":  2,
	"INSTAGRAM": 3,
	"TWITTER":   4,
	"FACEBOOK":  5,
}

func (x SocialType) String() string {
	return proto.EnumName(SocialType_name, int32(x))
}
func (SocialType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type STATUS int32

const (
	STATUS_UNKNOWN   STATUS = 0
	STATUS_ENQUEUED  STATUS = 1
	STATUS_FAIL      STATUS = 2
	STATUS_COMPLETED STATUS = 3
)

var STATUS_name = map[int32]string{
	0: "UNKNOWN",
	1: "ENQUEUED",
	2: "FAIL",
	3: "COMPLETED",
}
var STATUS_value = map[string]int32{
	"UNKNOWN":   0,
	"ENQUEUED":  1,
	"FAIL":      2,
	"COMPLETED": 3,
}

func (x STATUS) String() string {
	return proto.EnumName(STATUS_name, int32(x))
}
func (STATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Contact_Type int32

const (
	Contact_UNIVERSAL Contact_Type = 0
	Contact_PHONE     Contact_Type = 1
	Contact_URL       Contact_Type = 2
	Contact_EMAIL     Contact_Type = 3
)

var Contact_Type_name = map[int32]string{
	0: "UNIVERSAL",
	1: "PHONE",
	2: "URL",
	3: "EMAIL",
}
var Contact_Type_value = map[string]int32{
	"UNIVERSAL": 0,
	"PHONE":     1,
	"URL":       2,
	"EMAIL":     3,
}

func (x Contact_Type) String() string {
	return proto.EnumName(Contact_Type_name, int32(x))
}
func (Contact_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// user contact
// будет приложено внизу поста
type Contact struct {
	Name  string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type  Contact_Type `protobuf:"varint,2,opt,name=type,enum=bosster.Contact_Type" json:"type,omitempty"`
	Value string       `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetType() Contact_Type {
	if m != nil {
		return m.Type
	}
	return Contact_UNIVERSAL
}

func (m *Contact) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// post represent an article with body (text) and images
type Post struct {
	Message   string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ImageUrls []string `protobuf:"bytes,2,rep,name=image_urls,json=imageUrls" json:"image_urls,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Post) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Post) GetImageUrls() []string {
	if m != nil {
		return m.ImageUrls
	}
	return nil
}

// The request message
type PostRequest struct {
	Post *Post `protobuf:"bytes,1,opt,name=post" json:"post,omitempty"`
	// if async == false, server will send response instantly and enqueue
	// post request to background
	// if async == true, server will be wait to post article to each soc network
	Sync     bool       `protobuf:"varint,2,opt,name=sync" json:"sync,omitempty"`
	Contacts []*Contact `protobuf:"bytes,4,rep,name=contacts" json:"contacts,omitempty"`
	// TODO: maybe rename this
	Targets []*PostJob `protobuf:"bytes,5,rep,name=targets" json:"targets,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *PostRequest) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

func (m *PostRequest) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *PostRequest) GetTargets() []*PostJob {
	if m != nil {
		return m.Targets
	}
	return nil
}

// для каждой соцсети свой PostJob
type PostJob struct {
	// for webhook
	ExternalId      string     `protobuf:"bytes,1,opt,name=external_id,json=externalId" json:"external_id,omitempty"`
	Type            SocialType `protobuf:"varint,2,opt,name=type,enum=bosster.SocialType" json:"type,omitempty"`
	SocialId        string     `protobuf:"bytes,3,opt,name=social_id,json=socialId" json:"social_id,omitempty"`
	SocialWallType  string     `protobuf:"bytes,11,opt,name=social_wall_type,json=socialWallType" json:"social_wall_type,omitempty"`
	SocialToken     string     `protobuf:"bytes,4,opt,name=social_token,json=socialToken" json:"social_token,omitempty"`
	SocialLogin     string     `protobuf:"bytes,5,opt,name=social_login,json=socialLogin" json:"social_login,omitempty"`
	SocialPassword  string     `protobuf:"bytes,6,opt,name=social_password,json=socialPassword" json:"social_password,omitempty"`
	SocialAppId     string     `protobuf:"bytes,7,opt,name=social_app_id,json=socialAppId" json:"social_app_id,omitempty"`
	SocialAppSecret string     `protobuf:"bytes,8,opt,name=social_app_secret,json=socialAppSecret" json:"social_app_secret,omitempty"`
	Status          STATUS     `protobuf:"varint,9,opt,name=status,enum=bosster.STATUS" json:"status,omitempty"`
	// optional
	Error string `protobuf:"bytes,10,opt,name=error" json:"error,omitempty"`
}

func (m *PostJob) Reset()                    { *m = PostJob{} }
func (m *PostJob) String() string            { return proto.CompactTextString(m) }
func (*PostJob) ProtoMessage()               {}
func (*PostJob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PostJob) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *PostJob) GetType() SocialType {
	if m != nil {
		return m.Type
	}
	return SocialType_NONE
}

func (m *PostJob) GetSocialId() string {
	if m != nil {
		return m.SocialId
	}
	return ""
}

func (m *PostJob) GetSocialWallType() string {
	if m != nil {
		return m.SocialWallType
	}
	return ""
}

func (m *PostJob) GetSocialToken() string {
	if m != nil {
		return m.SocialToken
	}
	return ""
}

func (m *PostJob) GetSocialLogin() string {
	if m != nil {
		return m.SocialLogin
	}
	return ""
}

func (m *PostJob) GetSocialPassword() string {
	if m != nil {
		return m.SocialPassword
	}
	return ""
}

func (m *PostJob) GetSocialAppId() string {
	if m != nil {
		return m.SocialAppId
	}
	return ""
}

func (m *PostJob) GetSocialAppSecret() string {
	if m != nil {
		return m.SocialAppSecret
	}
	return ""
}

func (m *PostJob) GetStatus() STATUS {
	if m != nil {
		return m.Status
	}
	return STATUS_UNKNOWN
}

func (m *PostJob) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// The response message
// это будет возвращаться
type PostReply struct {
	Jobs []*PostJob `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *PostReply) Reset()                    { *m = PostReply{} }
func (m *PostReply) String() string            { return proto.CompactTextString(m) }
func (*PostReply) ProtoMessage()               {}
func (*PostReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PostReply) GetJobs() []*PostJob {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "bosster.Contact")
	proto.RegisterType((*Post)(nil), "bosster.Post")
	proto.RegisterType((*PostRequest)(nil), "bosster.PostRequest")
	proto.RegisterType((*PostJob)(nil), "bosster.PostJob")
	proto.RegisterType((*PostReply)(nil), "bosster.PostReply")
	proto.RegisterEnum("bosster.SocialType", SocialType_name, SocialType_value)
	proto.RegisterEnum("bosster.STATUS", STATUS_name, STATUS_value)
	proto.RegisterEnum("bosster.Contact_Type", Contact_Type_name, Contact_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Poster service

type PosterClient interface {
	// Sends a greeting
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostReply, error)
}

type posterClient struct {
	cc *grpc.ClientConn
}

func NewPosterClient(cc *grpc.ClientConn) PosterClient {
	return &posterClient{cc}
}

func (c *posterClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostReply, error) {
	out := new(PostReply)
	err := grpc.Invoke(ctx, "/bosster.Poster/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Poster service

type PosterServer interface {
	// Sends a greeting
	Post(context.Context, *PostRequest) (*PostReply, error)
}

func RegisterPosterServer(s *grpc.Server, srv PosterServer) {
	s.RegisterService(&_Poster_serviceDesc, srv)
}

func _Poster_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bosster.Poster/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bosster.Poster",
	HandlerType: (*PosterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Poster_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bosster.proto",
}

func init() { proto.RegisterFile("bosster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x51, 0x6f, 0xda, 0x48,
	0x10, 0xc7, 0x31, 0x36, 0x18, 0x8f, 0x43, 0xe2, 0x9b, 0xcb, 0x49, 0xab, 0x3b, 0x9d, 0x4a, 0xac,
	0x4a, 0xa1, 0xa8, 0x8a, 0x54, 0xda, 0xc7, 0x48, 0x95, 0x9b, 0x38, 0xad, 0x1b, 0x62, 0xa8, 0xb1,
	0x93, 0xc7, 0xc8, 0xc0, 0x0a, 0xd1, 0x3a, 0xac, 0xeb, 0x5d, 0x9a, 0xf2, 0x3d, 0xfa, 0xd0, 0xa7,
	0x7e, 0xd6, 0x6a, 0xd7, 0x86, 0x90, 0x48, 0x7d, 0xf3, 0xfc, 0xf7, 0xb7, 0xff, 0xdd, 0xd9, 0x99,
	0x31, 0xb4, 0x27, 0x8c, 0x73, 0x41, 0x8b, 0x93, 0xbc, 0x60, 0x82, 0xa1, 0x59, 0x85, 0xee, 0x0f,
	0x0d, 0xcc, 0x33, 0xb6, 0x14, 0xe9, 0x54, 0x20, 0x82, 0xb1, 0x4c, 0xef, 0x28, 0xd1, 0x3a, 0x5a,
	0xd7, 0x8a, 0xd4, 0x37, 0xbe, 0x00, 0x43, 0xac, 0x73, 0x4a, 0xea, 0x1d, 0xad, 0xbb, 0xdf, 0xff,
	0xe7, 0x64, 0x63, 0x53, 0xed, 0x39, 0x89, 0xd7, 0x39, 0x8d, 0x14, 0x82, 0x87, 0xd0, 0xf8, 0x96,
	0x66, 0x2b, 0x4a, 0x74, 0xb5, 0xbf, 0x0c, 0xdc, 0x37, 0x60, 0x48, 0x06, 0xdb, 0x60, 0x25, 0x61,
	0x70, 0xed, 0x47, 0x63, 0x6f, 0xe0, 0xd4, 0xd0, 0x82, 0xc6, 0xe8, 0xc3, 0x30, 0xf4, 0x1d, 0x0d,
	0x4d, 0xd0, 0x93, 0x68, 0xe0, 0xd4, 0xa5, 0xe6, 0x5f, 0x79, 0xc1, 0xc0, 0xd1, 0xdd, 0xb7, 0x60,
	0x8c, 0x18, 0x17, 0x48, 0xc0, 0xbc, 0xa3, 0x9c, 0xa7, 0xf3, 0xcd, 0xad, 0x36, 0x21, 0xfe, 0x0f,
	0xb0, 0xb8, 0x4b, 0xe7, 0xf4, 0x76, 0x55, 0x64, 0x9c, 0xd4, 0x3b, 0x7a, 0xd7, 0x8a, 0x2c, 0xa5,
	0x24, 0x45, 0xc6, 0xdd, 0x5f, 0x1a, 0xd8, 0xd2, 0x21, 0xa2, 0x5f, 0x57, 0x94, 0x0b, 0x3c, 0x02,
	0x23, 0x67, 0x5c, 0x28, 0x17, 0xbb, 0xdf, 0xde, 0xe6, 0xa1, 0x18, 0xb5, 0x24, 0xd3, 0xe7, 0xeb,
	0xe5, 0x54, 0xa5, 0xda, 0x8a, 0xd4, 0x37, 0xbe, 0x84, 0xd6, 0xb4, 0xcc, 0x94, 0x13, 0xa3, 0xa3,
	0x77, 0xed, 0xbe, 0xf3, 0xf4, 0x09, 0xa2, 0x2d, 0x81, 0x3d, 0x30, 0x45, 0x5a, 0xcc, 0xa9, 0xe0,
	0xa4, 0xf1, 0x04, 0x96, 0xe7, 0x7c, 0x64, 0x93, 0x68, 0x03, 0xb8, 0x3f, 0x75, 0x30, 0x2b, 0x11,
	0x9f, 0x81, 0x4d, 0xbf, 0x0b, 0x5a, 0x2c, 0xd3, 0xec, 0x76, 0x31, 0xab, 0x32, 0x85, 0x8d, 0x14,
	0xcc, 0xf0, 0xf8, 0x51, 0x15, 0xfe, 0xde, 0xba, 0x8e, 0xd9, 0x74, 0x91, 0x66, 0x3b, 0x35, 0xf8,
	0x0f, 0x2c, 0xae, 0x34, 0xe9, 0x53, 0xd6, 0xa1, 0x55, 0x0a, 0xc1, 0x0c, 0xbb, 0xe0, 0x54, 0x8b,
	0xf7, 0x69, 0x96, 0xdd, 0x2a, 0x47, 0x5b, 0x31, 0xfb, 0xa5, 0x7e, 0x93, 0x66, 0xca, 0x0c, 0x8f,
	0x60, 0xaf, 0x22, 0x05, 0xfb, 0x42, 0x97, 0xc4, 0x50, 0x94, 0x5d, 0x6a, 0xb1, 0x94, 0x76, 0x90,
	0x8c, 0xcd, 0x17, 0x4b, 0xd2, 0xd8, 0x45, 0x06, 0x52, 0xc2, 0x63, 0x38, 0xa8, 0x90, 0x3c, 0xe5,
	0xfc, 0x9e, 0x15, 0x33, 0xd2, 0xdc, 0x3d, 0x6e, 0x54, 0xa9, 0xe8, 0x42, 0xbb, 0x02, 0xd3, 0x3c,
	0x97, 0x37, 0x37, 0x77, 0xcd, 0xbc, 0x3c, 0x0f, 0x66, 0xd8, 0x83, 0xbf, 0x76, 0x18, 0x4e, 0xa7,
	0x05, 0x15, 0xa4, 0xa5, 0xb8, 0x83, 0x2d, 0x37, 0x56, 0x32, 0x1e, 0x43, 0x93, 0x8b, 0x54, 0xac,
	0x38, 0xb1, 0xd4, 0x83, 0x1d, 0x3c, 0x3c, 0x58, 0xec, 0xc5, 0xc9, 0x38, 0xaa, 0x96, 0x65, 0xcb,
	0xd2, 0xa2, 0x60, 0x05, 0x81, 0xb2, 0x65, 0x55, 0xe0, 0xbe, 0x02, 0xab, 0x6c, 0x9d, 0x3c, 0x5b,
	0xe3, 0x73, 0x30, 0x3e, 0xb3, 0x09, 0x27, 0xda, 0x1f, 0x0a, 0xaa, 0x56, 0x7b, 0xd7, 0x00, 0x0f,
	0xb5, 0xc0, 0x16, 0x18, 0xa1, 0xec, 0xed, 0x1a, 0x36, 0xa1, 0x7e, 0x7d, 0xe9, 0x68, 0xb8, 0x07,
	0xad, 0xd8, 0x1f, 0xf8, 0xef, 0x23, 0xef, 0xca, 0xa9, 0xcb, 0x59, 0x08, 0xc2, 0x71, 0xec, 0xa9,
	0x50, 0x47, 0x1b, 0xcc, 0xf8, 0x26, 0x88, 0x63, 0x3f, 0x72, 0x0c, 0x49, 0x5e, 0x78, 0x67, 0xfe,
	0xbb, 0xe1, 0xf0, 0xd2, 0x69, 0xf4, 0x4e, 0xa1, 0x59, 0x5e, 0x59, 0x42, 0x49, 0x78, 0x19, 0x0e,
	0x6f, 0x42, 0xa7, 0x26, 0x21, 0x3f, 0xfc, 0x94, 0xf8, 0x89, 0x7f, 0xee, 0x68, 0xf2, 0xb8, 0x0b,
	0x39, 0x36, 0xca, 0xf8, 0x6c, 0x78, 0x35, 0x1a, 0xf8, 0xb1, 0x7f, 0xee, 0xe8, 0xfd, 0x53, 0x68,
	0xca, 0x6b, 0xd2, 0x02, 0xfb, 0xd5, 0x3c, 0x1d, 0x3e, 0x6e, 0xfc, 0x72, 0x38, 0xfe, 0xc5, 0x27,
	0x6a, 0x9e, 0xad, 0xdd, 0xda, 0xa4, 0xa9, 0x7e, 0x15, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x86, 0xc6, 0x24, 0x3b, 0x04, 0x00, 0x00,
}
