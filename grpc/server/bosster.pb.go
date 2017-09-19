// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bosster.proto

/*
Package server is a generated protocol buffer package.

server

It is generated from these files:
	bosster.proto

It has these top-level messages:
	Contact
	Post
	PostRequest
	PostJob
	PostReply
*/
package server

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
	Type  Contact_Type `protobuf:"varint,2,opt,name=type,enum=server.Contact_Type" json:"type,omitempty"`
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
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Post *Post  `protobuf:"bytes,2,opt,name=post" json:"post,omitempty"`
	// if async == false, server will send response instantly and enqueue
	// post request to background
	// if async == true, server will be wait to post article to each soc network
	Async    bool       `protobuf:"varint,3,opt,name=async" json:"async,omitempty"`
	Contacts []*Contact `protobuf:"bytes,4,rep,name=contacts" json:"contacts,omitempty"`
	// TODO: maybe rename this
	Targets []SocialType `protobuf:"varint,5,rep,packed,name=targets,enum=server.SocialType" json:"targets,omitempty"`
	// id группы для постинга
	SocialId       string `protobuf:"bytes,6,opt,name=social_id,json=socialId" json:"social_id,omitempty"`
	SocialToken    string `protobuf:"bytes,7,opt,name=social_token,json=socialToken" json:"social_token,omitempty"`
	SocialLogin    string `protobuf:"bytes,8,opt,name=social_login,json=socialLogin" json:"social_login,omitempty"`
	SocialPassword string `protobuf:"bytes,9,opt,name=social_password,json=socialPassword" json:"social_password,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PostRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *PostRequest) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *PostRequest) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *PostRequest) GetTargets() []SocialType {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *PostRequest) GetSocialId() string {
	if m != nil {
		return m.SocialId
	}
	return ""
}

func (m *PostRequest) GetSocialToken() string {
	if m != nil {
		return m.SocialToken
	}
	return ""
}

func (m *PostRequest) GetSocialLogin() string {
	if m != nil {
		return m.SocialLogin
	}
	return ""
}

func (m *PostRequest) GetSocialPassword() string {
	if m != nil {
		return m.SocialPassword
	}
	return ""
}

// для каждой соцсети свой PostJob
type PostJob struct {
	// нужно для запроса состояния,
	// если процедура будет выполняться асинхронно
	PostRequestId string `protobuf:"bytes,1,opt,name=post_request_id,json=postRequestId" json:"post_request_id,omitempty"`
	// id поста
	SocialId string `protobuf:"bytes,2,opt,name=social_id,json=socialId" json:"social_id,omitempty"`
	// на чью стену опубликовали
	SocialOwnerId  string     `protobuf:"bytes,3,opt,name=social_owner_id,json=socialOwnerId" json:"social_owner_id,omitempty"`
	SocialProvider SocialType `protobuf:"varint,4,opt,name=social_provider,json=socialProvider,enum=server.SocialType" json:"social_provider,omitempty"`
	Status         STATUS     `protobuf:"varint,5,opt,name=status,enum=server.STATUS" json:"status,omitempty"`
	// optional
	Error string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *PostJob) Reset()                    { *m = PostJob{} }
func (m *PostJob) String() string            { return proto.CompactTextString(m) }
func (*PostJob) ProtoMessage()               {}
func (*PostJob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PostJob) GetPostRequestId() string {
	if m != nil {
		return m.PostRequestId
	}
	return ""
}

func (m *PostJob) GetSocialId() string {
	if m != nil {
		return m.SocialId
	}
	return ""
}

func (m *PostJob) GetSocialOwnerId() string {
	if m != nil {
		return m.SocialOwnerId
	}
	return ""
}

func (m *PostJob) GetSocialProvider() SocialType {
	if m != nil {
		return m.SocialProvider
	}
	return SocialType_NONE
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
	proto.RegisterType((*Contact)(nil), "server.Contact")
	proto.RegisterType((*Post)(nil), "server.Post")
	proto.RegisterType((*PostRequest)(nil), "server.PostRequest")
	proto.RegisterType((*PostJob)(nil), "server.PostJob")
	proto.RegisterType((*PostReply)(nil), "server.PostReply")
	proto.RegisterEnum("server.SocialType", SocialType_name, SocialType_value)
	proto.RegisterEnum("server.STATUS", STATUS_name, STATUS_value)
	proto.RegisterEnum("server.Contact_Type", Contact_Type_name, Contact_Type_value)
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
	err := grpc.Invoke(ctx, "/server.Poster/Post", in, out, c.cc, opts...)
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
		FullMethod: "/server.Poster/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Poster",
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
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x7f, 0x62, 0xc7, 0x93, 0x26, 0xf5, 0xb7, 0x5f, 0x2f, 0x2c, 0x10, 0x52, 0x30, 0x52,
	0x89, 0x0a, 0x8a, 0xaa, 0xc0, 0x15, 0x20, 0xa1, 0xd0, 0xba, 0x60, 0x9a, 0x3a, 0xc1, 0xb1, 0xdb,
	0xcb, 0xc8, 0x89, 0x57, 0x91, 0xc1, 0xcd, 0x9a, 0x5d, 0xa7, 0x55, 0x9e, 0x83, 0xd7, 0xe1, 0xa5,
	0x78, 0x03, 0xb4, 0xbb, 0x4e, 0xd2, 0x54, 0xdc, 0x79, 0xce, 0x9c, 0xd9, 0x3d, 0x73, 0x76, 0xc6,
	0xd0, 0x9a, 0x11, 0xc6, 0x4a, 0x4c, 0x7b, 0x05, 0x25, 0x25, 0x41, 0x06, 0xc3, 0xf4, 0x0e, 0x53,
	0xf7, 0x97, 0x02, 0xe6, 0x19, 0x59, 0x96, 0xc9, 0xbc, 0x44, 0x08, 0xf4, 0x65, 0x72, 0x8b, 0x1d,
	0xa5, 0xa3, 0x74, 0xad, 0x50, 0x7c, 0xa3, 0x2e, 0xe8, 0xe5, 0xba, 0xc0, 0x8e, 0xda, 0x51, 0xba,
	0xed, 0xfe, 0x51, 0x4f, 0x96, 0xf5, 0xaa, 0x92, 0x5e, 0xb4, 0x2e, 0x70, 0x28, 0x18, 0xe8, 0x08,
	0xea, 0x77, 0x49, 0xbe, 0xc2, 0x8e, 0x26, 0xca, 0x65, 0xe0, 0xbe, 0x05, 0x9d, 0x73, 0x50, 0x0b,
	0xac, 0x38, 0xf0, 0xaf, 0xbd, 0x70, 0x32, 0x18, 0xda, 0x35, 0x64, 0x41, 0x7d, 0xfc, 0x65, 0x14,
	0x78, 0xb6, 0x82, 0x4c, 0xd0, 0xe2, 0x70, 0x68, 0xab, 0x1c, 0xf3, 0xae, 0x06, 0xfe, 0xd0, 0xd6,
	0xdc, 0x8f, 0xa0, 0x8f, 0x09, 0x2b, 0x91, 0x03, 0xe6, 0x2d, 0x66, 0x2c, 0x59, 0x6c, 0x44, 0x6d,
	0x42, 0xf4, 0x0c, 0x20, 0xbb, 0x4d, 0x16, 0x78, 0xba, 0xa2, 0x39, 0x73, 0xd4, 0x8e, 0xd6, 0xb5,
	0x42, 0x4b, 0x20, 0x31, 0xcd, 0x99, 0xfb, 0x5b, 0x85, 0x26, 0x3f, 0x21, 0xc4, 0x3f, 0x57, 0x98,
	0x95, 0xa8, 0x0d, 0x6a, 0x96, 0x56, 0x67, 0xa8, 0x59, 0x8a, 0x3a, 0xa0, 0x17, 0x84, 0x95, 0xa2,
	0xad, 0x66, 0xff, 0x60, 0xd3, 0x96, 0x28, 0x11, 0x19, 0xde, 0x4e, 0xc2, 0xd6, 0xcb, 0xb9, 0x68,
	0xa7, 0x11, 0xca, 0x00, 0xbd, 0x82, 0xc6, 0x5c, 0xb6, 0xce, 0x1c, 0xbd, 0xa3, 0x75, 0x9b, 0xfd,
	0xc3, 0x47, 0x96, 0x84, 0x5b, 0x02, 0x7a, 0x0d, 0x66, 0x99, 0xd0, 0x05, 0x2e, 0x99, 0x53, 0xef,
	0x68, 0xdd, 0x76, 0x1f, 0x6d, 0xb8, 0x13, 0x32, 0xcf, 0x92, 0x5c, 0x98, 0xb7, 0xa1, 0xa0, 0xa7,
	0x60, 0x31, 0x01, 0x4f, 0xb3, 0xd4, 0x31, 0x84, 0xd2, 0x86, 0x04, 0xfc, 0x14, 0x3d, 0x87, 0x83,
	0x2a, 0x59, 0x92, 0x1f, 0x78, 0xe9, 0x98, 0x22, 0xdf, 0x94, 0x58, 0xc4, 0xa1, 0x07, 0x94, 0x9c,
	0x2c, 0xb2, 0xa5, 0xd3, 0x78, 0x48, 0x19, 0x72, 0x08, 0xbd, 0x84, 0xc3, 0x8a, 0x52, 0x24, 0x8c,
	0xdd, 0x13, 0x9a, 0x3a, 0x96, 0x60, 0xb5, 0x25, 0x3c, 0xae, 0x50, 0xf7, 0x8f, 0x02, 0x26, 0xf7,
	0xe2, 0x2b, 0x99, 0xa1, 0x63, 0x38, 0xe4, 0x86, 0x4c, 0xa9, 0xb4, 0x72, 0xba, 0xf5, 0xb1, 0x55,
	0xec, 0x0c, 0xf6, 0xd3, 0x7d, 0xfd, 0xea, 0x23, 0xfd, 0xc7, 0xdb, 0x9b, 0xc9, 0xfd, 0x12, 0x53,
	0x4e, 0x91, 0x63, 0xd2, 0x92, 0xf0, 0x88, 0xa3, 0x7e, 0x8a, 0xde, 0xef, 0x14, 0x52, 0x72, 0x97,
	0xa5, 0x98, 0x3a, 0xba, 0x98, 0xbc, 0x7f, 0x59, 0xb7, 0x51, 0x5d, 0x31, 0xd1, 0x31, 0x18, 0xac,
	0x4c, 0xca, 0x15, 0xb7, 0x9b, 0xd7, 0xb4, 0xb7, 0x35, 0xd1, 0x20, 0x8a, 0x27, 0x61, 0x95, 0xe5,
	0x4f, 0x8b, 0x29, 0x25, 0xb4, 0x72, 0x59, 0x06, 0xee, 0x29, 0x58, 0x72, 0x62, 0x8a, 0x7c, 0x8d,
	0x5e, 0x80, 0xfe, 0x9d, 0xcc, 0x98, 0xa3, 0xec, 0xbf, 0x71, 0xe5, 0x49, 0x28, 0x92, 0x27, 0xd7,
	0x00, 0x3b, 0x35, 0xa8, 0x01, 0x7a, 0xc0, 0x27, 0xba, 0x86, 0x0c, 0x50, 0xaf, 0x2f, 0x6d, 0x05,
	0x1d, 0x40, 0x23, 0xf2, 0x86, 0xde, 0xe7, 0x70, 0x70, 0x65, 0xab, 0x7c, 0x03, 0xfc, 0x60, 0x12,
	0x0d, 0x44, 0xa8, 0xa1, 0x26, 0x98, 0xd1, 0x8d, 0x1f, 0x45, 0x5e, 0x68, 0xeb, 0x9c, 0x79, 0x31,
	0x38, 0xf3, 0x3e, 0x8d, 0x46, 0x97, 0x76, 0xfd, 0xe4, 0x03, 0x18, 0x52, 0x31, 0x27, 0xc5, 0xc1,
	0x65, 0x30, 0xba, 0x09, 0xec, 0x1a, 0x27, 0x79, 0xc1, 0xb7, 0xd8, 0x8b, 0xbd, 0x73, 0x5b, 0xe1,
	0xd7, 0x5d, 0xf0, 0x65, 0x11, 0x07, 0x9f, 0x8d, 0xae, 0xc6, 0x43, 0x2f, 0xf2, 0xce, 0x6d, 0xad,
	0xff, 0x0e, 0x0c, 0x2e, 0x13, 0x53, 0x74, 0x5a, 0x6d, 0xd1, 0xff, 0x7b, 0xe3, 0x2d, 0x1f, 0xec,
	0xc9, 0x7f, 0xfb, 0x60, 0x91, 0xaf, 0xdd, 0xda, 0xcc, 0x10, 0x3f, 0x87, 0x37, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x42, 0xfc, 0x27, 0x2d, 0x04, 0x00, 0x00,
}
