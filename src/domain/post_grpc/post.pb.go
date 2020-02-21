// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/post/controllers/post_grpc/post.proto

package post_grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UserId               int64                `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Post) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Post) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateReq struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{1}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UpdateReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{2}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdateReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ListReq struct {
	Datetime             *timestamp.Timestamp `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Num                  int64                `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListReq) Reset()         { *m = ListReq{} }
func (m *ListReq) String() string { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()    {}
func (*ListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{3}
}

func (m *ListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReq.Unmarshal(m, b)
}
func (m *ListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReq.Marshal(b, m, deterministic)
}
func (m *ListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReq.Merge(m, src)
}
func (m *ListReq) XXX_Size() int {
	return xxx_messageInfo_ListReq.Size(m)
}
func (m *ListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListReq proto.InternalMessageInfo

func (m *ListReq) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

func (m *ListReq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type DeleteReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{4}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{5}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListPost struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPost) Reset()         { *m = ListPost{} }
func (m *ListPost) String() string { return proto.CompactTextString(m) }
func (*ListPost) ProtoMessage()    {}
func (*ListPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3bb529b00bb5e1e, []int{6}
}

func (m *ListPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPost.Unmarshal(m, b)
}
func (m *ListPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPost.Marshal(b, m, deterministic)
}
func (m *ListPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPost.Merge(m, src)
}
func (m *ListPost) XXX_Size() int {
	return xxx_messageInfo_ListPost.Size(m)
}
func (m *ListPost) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPost.DiscardUnknown(m)
}

var xxx_messageInfo_ListPost proto.InternalMessageInfo

func (m *ListPost) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "post_grpc.Post")
	proto.RegisterType((*CreateReq)(nil), "post_grpc.CreateReq")
	proto.RegisterType((*UpdateReq)(nil), "post_grpc.UpdateReq")
	proto.RegisterType((*ListReq)(nil), "post_grpc.ListReq")
	proto.RegisterType((*DeleteReq)(nil), "post_grpc.DeleteReq")
	proto.RegisterType((*ID)(nil), "post_grpc.ID")
	proto.RegisterType((*ListPost)(nil), "post_grpc.ListPost")
}

func init() {
	proto.RegisterFile("src/post/controllers/post_grpc/post.proto", fileDescriptor_f3bb529b00bb5e1e)
}

var fileDescriptor_f3bb529b00bb5e1e = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x8e, 0x93, 0x40,
	0x14, 0xc7, 0x33, 0x40, 0xa1, 0x9c, 0x8d, 0xee, 0x66, 0x6c, 0xd2, 0x09, 0xc6, 0x6d, 0x43, 0x34,
	0x61, 0x6f, 0x20, 0xd6, 0x2b, 0xbd, 0x72, 0xc7, 0x46, 0xd3, 0xc4, 0x0b, 0x83, 0x1f, 0xb7, 0x0d,
	0x5b, 0xc6, 0x4a, 0x42, 0x3b, 0x38, 0x0c, 0x6b, 0x7c, 0x05, 0x13, 0x5f, 0xc6, 0x27, 0xf1, 0xc6,
	0x97, 0xe9, 0x95, 0x99, 0xa1, 0x40, 0x3f, 0xd6, 0x5d, 0xef, 0x38, 0xe7, 0xfc, 0x0f, 0xe7, 0xf7,
	0x3f, 0x1c, 0xe0, 0xa2, 0x14, 0x8b, 0xa8, 0xe0, 0xa5, 0x8c, 0x16, 0x7c, 0x2d, 0x05, 0xcf, 0x73,
	0x26, 0x4a, 0x9d, 0x98, 0x2f, 0x45, 0x51, 0x97, 0xc2, 0x42, 0x70, 0xc9, 0xb1, 0xdb, 0x66, 0xbd,
	0xe1, 0x75, 0x92, 0x67, 0x69, 0x22, 0x59, 0xd4, 0x3c, 0xd4, 0x1a, 0x6f, 0xb4, 0xe4, 0x7c, 0x99,
	0xb3, 0x48, 0x47, 0x57, 0xd5, 0xe7, 0x48, 0x66, 0x2b, 0x56, 0xca, 0x64, 0x55, 0x6c, 0x05, 0xe7,
	0x87, 0x82, 0x6f, 0x22, 0x29, 0x0a, 0x26, 0xca, 0xba, 0xee, 0xff, 0x41, 0x60, 0xbd, 0xe3, 0xa5,
	0xc4, 0xf7, 0xc1, 0xc8, 0x52, 0x82, 0xc6, 0x28, 0x30, 0x63, 0x23, 0x4b, 0xf1, 0x00, 0x7a, 0x32,
	0x93, 0x39, 0x23, 0xc6, 0x18, 0x05, 0x6e, 0x5c, 0x07, 0x98, 0x80, 0xa3, 0xb8, 0xd9, 0x5a, 0x12,
	0x53, 0xe7, 0x9b, 0x10, 0x3f, 0x07, 0x58, 0x08, 0x96, 0x48, 0x96, 0xce, 0x13, 0x49, 0xac, 0x31,
	0x0a, 0x4e, 0x26, 0x5e, 0x58, 0x4f, 0x0f, 0x9b, 0xe9, 0xe1, 0x87, 0x06, 0x2f, 0x76, 0xb7, 0xea,
	0x4b, 0xdd, 0x5a, 0x15, 0x69, 0xd3, 0xda, 0xbb, 0xbb, 0x75, 0xab, 0xbe, 0x94, 0x78, 0x08, 0x4e,
	0x55, 0x32, 0x31, 0xcf, 0x52, 0x62, 0x6b, 0x74, 0x5b, 0x85, 0xb3, 0xd4, 0x97, 0xe0, 0xbe, 0xd2,
	0x03, 0x62, 0xf6, 0x15, 0x8f, 0x1a, 0x2f, 0xca, 0x9e, 0x4b, 0xdd, 0x0d, 0xb5, 0x85, 0x75, 0x86,
	0xc8, 0xa0, 0xb1, 0xf5, 0xb8, 0xb3, 0xa5, 0xed, 0x52, 0xd8, 0x50, 0x47, 0xf4, 0xce, 0x10, 0xf9,
	0x7d, 0xda, 0x59, 0x1c, 0x77, 0xc3, 0x94, 0x79, 0x93, 0x3a, 0x1b, 0x6a, 0xf9, 0x46, 0x80, 0xda,
	0xa9, 0x3f, 0x11, 0xb8, 0x1f, 0x35, 0x9c, 0x1a, 0x3b, 0xec, 0x56, 0xda, 0x49, 0xd5, 0x6e, 0x47,
	0x7b, 0xbb, 0xbd, 0x9d, 0xc7, 0xfc, 0x2f, 0x1e, 0xeb, 0x66, 0x9e, 0x2f, 0xe0, 0xbc, 0xcd, 0x4a,
	0xa9, 0x60, 0x5e, 0x42, 0x5f, 0x71, 0xa9, 0xfb, 0xd0, 0x48, 0xb7, 0xae, 0x98, 0xf6, 0x37, 0xb4,
	0xf7, 0x0b, 0x19, 0x7d, 0x14, 0xb7, 0x5d, 0xf8, 0x21, 0x98, 0xeb, 0x6a, 0xa5, 0x99, 0x4d, 0xcd,
	0xec, 0x5b, 0xe4, 0x3c, 0x40, 0xb1, 0xca, 0xfa, 0xaf, 0xc1, 0x9d, 0xb2, 0x9c, 0xdd, 0x61, 0x7c,
	0x87, 0xd8, 0xb8, 0x99, 0xf8, 0x11, 0x18, 0xb3, 0xe9, 0x3f, 0x5f, 0xe0, 0x3f, 0x85, 0xbe, 0x32,
	0xa4, 0x2f, 0xf6, 0x09, 0xf4, 0xd4, 0x1f, 0x52, 0x12, 0x34, 0x36, 0x83, 0x93, 0xc9, 0x69, 0xd8,
	0xfe, 0x2f, 0xa1, 0xaa, 0xc7, 0x75, 0x75, 0xf2, 0xc3, 0x80, 0x13, 0x15, 0xbf, 0x67, 0xe2, 0x3a,
	0x5b, 0x30, 0x1c, 0x81, 0x5d, 0x5f, 0x06, 0x1e, 0xec, 0x74, 0xb4, 0xc7, 0xe2, 0x1d, 0xbe, 0x07,
	0x5f, 0x80, 0xf3, 0x86, 0x49, 0xfa, 0x7d, 0x36, 0xc5, 0xf7, 0x76, 0x6a, 0xb3, 0xe9, 0xb1, 0x74,
	0xa2, 0xa5, 0x8a, 0x10, 0xe3, 0x9d, 0xda, 0xf6, 0x1b, 0x78, 0x0f, 0x0e, 0x72, 0xba, 0x27, 0x02,
	0xbb, 0x3e, 0x99, 0x3d, 0x9e, 0xf6, 0x8a, 0x8e, 0x87, 0xbc, 0x00, 0xbb, 0x5e, 0xf5, 0x5e, 0x43,
	0xbb, 0x7d, 0xef, 0xf8, 0xbb, 0x52, 0xce, 0xf3, 0x4f, 0x49, 0x5e, 0xb1, 0x2b, 0x5b, 0xe7, 0x9e,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x49, 0x15, 0x0f, 0x87, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Post, error)
	GetByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Post, error)
	GetList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListPost, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Post, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/post_grpc.PostService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/post_grpc.PostService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListPost, error) {
	out := new(ListPost)
	err := c.cc.Invoke(ctx, "/post_grpc.PostService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/post_grpc.PostService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/post_grpc.PostService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	Create(context.Context, *CreateReq) (*Post, error)
	GetByID(context.Context, *ID) (*Post, error)
	GetList(context.Context, *ListReq) (*ListPost, error)
	Update(context.Context, *UpdateReq) (*Post, error)
	Delete(context.Context, *DeleteReq) (*wrappers.BoolValue, error)
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) Create(ctx context.Context, req *CreateReq) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPostServiceServer) GetByID(ctx context.Context, req *ID) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (*UnimplementedPostServiceServer) GetList(ctx context.Context, req *ListReq) (*ListPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (*UnimplementedPostServiceServer) Update(ctx context.Context, req *UpdateReq) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPostServiceServer) Delete(ctx context.Context, req *DeleteReq) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_grpc.PostService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_grpc.PostService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_grpc.PostService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_grpc.PostService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_grpc.PostService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post_grpc.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PostService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _PostService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _PostService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PostService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PostService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/post/controllers/post_grpc/post.proto",
}
