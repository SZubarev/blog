// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/comment/protobuf/comment.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//------------------------------------------------------------------------------------------------------------------
//  CREATE
//------------------------------------------------------------------------------------------------------------------
type CreateCommentRequest struct {
	ParentId             string   `protobuf:"bytes,1,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	PostId               string   `protobuf:"bytes,3,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCommentRequest) Reset()         { *m = CreateCommentRequest{} }
func (m *CreateCommentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCommentRequest) ProtoMessage()    {}
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{0}
}

func (m *CreateCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCommentRequest.Unmarshal(m, b)
}
func (m *CreateCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCommentRequest.Marshal(b, m, deterministic)
}
func (m *CreateCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCommentRequest.Merge(m, src)
}
func (m *CreateCommentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCommentRequest.Size(m)
}
func (m *CreateCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCommentRequest proto.InternalMessageInfo

func (m *CreateCommentRequest) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *CreateCommentRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateCommentRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *CreateCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateCommentResponse struct {
	Comment              *Comment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCommentResponse) Reset()         { *m = CreateCommentResponse{} }
func (m *CreateCommentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCommentResponse) ProtoMessage()    {}
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{1}
}

func (m *CreateCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCommentResponse.Unmarshal(m, b)
}
func (m *CreateCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCommentResponse.Marshal(b, m, deterministic)
}
func (m *CreateCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCommentResponse.Merge(m, src)
}
func (m *CreateCommentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCommentResponse.Size(m)
}
func (m *CreateCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCommentResponse proto.InternalMessageInfo

func (m *CreateCommentResponse) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

//------------------------------------------------------------------------------------------------------------------
//  UPDATE
//------------------------------------------------------------------------------------------------------------------
type UpdateCommentRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCommentRequest) Reset()         { *m = UpdateCommentRequest{} }
func (m *UpdateCommentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCommentRequest) ProtoMessage()    {}
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{2}
}

func (m *UpdateCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCommentRequest.Unmarshal(m, b)
}
func (m *UpdateCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCommentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCommentRequest.Merge(m, src)
}
func (m *UpdateCommentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCommentRequest.Size(m)
}
func (m *UpdateCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCommentRequest proto.InternalMessageInfo

func (m *UpdateCommentRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *UpdateCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdateCommentRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UpdateCommentResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCommentResponse) Reset()         { *m = UpdateCommentResponse{} }
func (m *UpdateCommentResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCommentResponse) ProtoMessage()    {}
func (*UpdateCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{3}
}

func (m *UpdateCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCommentResponse.Unmarshal(m, b)
}
func (m *UpdateCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCommentResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCommentResponse.Merge(m, src)
}
func (m *UpdateCommentResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCommentResponse.Size(m)
}
func (m *UpdateCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCommentResponse proto.InternalMessageInfo

func (m *UpdateCommentResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

//------------------------------------------------------------------------------------------------------------------
//  DELETE
//------------------------------------------------------------------------------------------------------------------
type DeleteCommentRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCommentRequest) Reset()         { *m = DeleteCommentRequest{} }
func (m *DeleteCommentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCommentRequest) ProtoMessage()    {}
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{4}
}

func (m *DeleteCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCommentRequest.Unmarshal(m, b)
}
func (m *DeleteCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCommentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCommentRequest.Merge(m, src)
}
func (m *DeleteCommentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCommentRequest.Size(m)
}
func (m *DeleteCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCommentRequest proto.InternalMessageInfo

func (m *DeleteCommentRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type DeleteCommentResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCommentResponse) Reset()         { *m = DeleteCommentResponse{} }
func (m *DeleteCommentResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCommentResponse) ProtoMessage()    {}
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{5}
}

func (m *DeleteCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCommentResponse.Unmarshal(m, b)
}
func (m *DeleteCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCommentResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCommentResponse.Merge(m, src)
}
func (m *DeleteCommentResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCommentResponse.Size(m)
}
func (m *DeleteCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCommentResponse proto.InternalMessageInfo

func (m *DeleteCommentResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

//------------------------------------------------------------------------------------------------------------------
//  GET
//------------------------------------------------------------------------------------------------------------------
type GetCommentRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommentRequest) Reset()         { *m = GetCommentRequest{} }
func (m *GetCommentRequest) String() string { return proto.CompactTextString(m) }
func (*GetCommentRequest) ProtoMessage()    {}
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{6}
}

func (m *GetCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentRequest.Unmarshal(m, b)
}
func (m *GetCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentRequest.Marshal(b, m, deterministic)
}
func (m *GetCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentRequest.Merge(m, src)
}
func (m *GetCommentRequest) XXX_Size() int {
	return xxx_messageInfo_GetCommentRequest.Size(m)
}
func (m *GetCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentRequest proto.InternalMessageInfo

func (m *GetCommentRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type GetCommentResponse struct {
	Comment              *Comment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommentResponse) Reset()         { *m = GetCommentResponse{} }
func (m *GetCommentResponse) String() string { return proto.CompactTextString(m) }
func (*GetCommentResponse) ProtoMessage()    {}
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{7}
}

func (m *GetCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommentResponse.Unmarshal(m, b)
}
func (m *GetCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommentResponse.Marshal(b, m, deterministic)
}
func (m *GetCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommentResponse.Merge(m, src)
}
func (m *GetCommentResponse) XXX_Size() int {
	return xxx_messageInfo_GetCommentResponse.Size(m)
}
func (m *GetCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommentResponse proto.InternalMessageInfo

func (m *GetCommentResponse) GetComment() *Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

//------------------------------------------------------------------------------------------------------------------
//  FIND
//------------------------------------------------------------------------------------------------------------------
type FindCommentRequest struct {
	PostId               string   `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindCommentRequest) Reset()         { *m = FindCommentRequest{} }
func (m *FindCommentRequest) String() string { return proto.CompactTextString(m) }
func (*FindCommentRequest) ProtoMessage()    {}
func (*FindCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{8}
}

func (m *FindCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindCommentRequest.Unmarshal(m, b)
}
func (m *FindCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindCommentRequest.Marshal(b, m, deterministic)
}
func (m *FindCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindCommentRequest.Merge(m, src)
}
func (m *FindCommentRequest) XXX_Size() int {
	return xxx_messageInfo_FindCommentRequest.Size(m)
}
func (m *FindCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindCommentRequest proto.InternalMessageInfo

func (m *FindCommentRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type FindCommentResponse struct {
	Comments             []*Comment `protobuf:"bytes,1,rep,name=Comments,proto3" json:"Comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindCommentResponse) Reset()         { *m = FindCommentResponse{} }
func (m *FindCommentResponse) String() string { return proto.CompactTextString(m) }
func (*FindCommentResponse) ProtoMessage()    {}
func (*FindCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{9}
}

func (m *FindCommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindCommentResponse.Unmarshal(m, b)
}
func (m *FindCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindCommentResponse.Marshal(b, m, deterministic)
}
func (m *FindCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindCommentResponse.Merge(m, src)
}
func (m *FindCommentResponse) XXX_Size() int {
	return xxx_messageInfo_FindCommentResponse.Size(m)
}
func (m *FindCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindCommentResponse proto.InternalMessageInfo

func (m *FindCommentResponse) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

//------------------------------------------------------------------------------------------------------------------
//  COMMENT
//------------------------------------------------------------------------------------------------------------------
type Comment struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=Slug,proto3" json:"Slug,omitempty"`
	ParentId             string   `protobuf:"bytes,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	PostId               string   `protobuf:"bytes,4,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b9c4a1f94de3f1, []int{10}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Comment) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *Comment) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Comment) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateCommentRequest)(nil), "protobuf.CreateCommentRequest")
	proto.RegisterType((*CreateCommentResponse)(nil), "protobuf.CreateCommentResponse")
	proto.RegisterType((*UpdateCommentRequest)(nil), "protobuf.UpdateCommentRequest")
	proto.RegisterType((*UpdateCommentResponse)(nil), "protobuf.UpdateCommentResponse")
	proto.RegisterType((*DeleteCommentRequest)(nil), "protobuf.DeleteCommentRequest")
	proto.RegisterType((*DeleteCommentResponse)(nil), "protobuf.DeleteCommentResponse")
	proto.RegisterType((*GetCommentRequest)(nil), "protobuf.GetCommentRequest")
	proto.RegisterType((*GetCommentResponse)(nil), "protobuf.GetCommentResponse")
	proto.RegisterType((*FindCommentRequest)(nil), "protobuf.FindCommentRequest")
	proto.RegisterType((*FindCommentResponse)(nil), "protobuf.FindCommentResponse")
	proto.RegisterType((*Comment)(nil), "protobuf.Comment")
}

func init() {
	proto.RegisterFile("services/comment/protobuf/comment.proto", fileDescriptor_e5b9c4a1f94de3f1)
}

var fileDescriptor_e5b9c4a1f94de3f1 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x49, 0xd3, 0x85, 0xf2, 0x43, 0x02, 0xd5, 0x74, 0x23, 0x0a, 0x63, 0x4c, 0xbe, 0x6c,
	0x1a, 0xd0, 0x88, 0xf1, 0x04, 0xa8, 0x85, 0x69, 0xb7, 0xa9, 0xd5, 0x6e, 0x1c, 0x48, 0xd7, 0x1f,
	0x55, 0xa4, 0xce, 0x2e, 0xb1, 0xdb, 0x0b, 0x70, 0xe1, 0x15, 0x38, 0xf2, 0x58, 0xbc, 0x02, 0x6f,
	0xc0, 0x0b, 0x20, 0xdb, 0x71, 0xeb, 0xa4, 0x8e, 0xa8, 0x76, 0xcb, 0xef, 0xef, 0xe7, 0xeb, 0xf8,
	0x6b, 0x38, 0x11, 0x58, 0xac, 0xf2, 0x1b, 0x14, 0xe9, 0x0d, 0xbf, 0xbd, 0x45, 0x26, 0xd3, 0x45,
	0xc1, 0x25, 0x9f, 0x2c, 0x3f, 0xdb, 0x44, 0x5f, 0x27, 0x48, 0xc7, 0xe6, 0x93, 0xc3, 0x19, 0xe7,
	0xb3, 0x39, 0xa6, 0xd9, 0x22, 0x4f, 0x33, 0xc6, 0xb8, 0xcc, 0x64, 0xce, 0x99, 0x30, 0x7d, 0xf4,
	0x1b, 0xf4, 0x06, 0x05, 0x66, 0x12, 0x07, 0x66, 0x7c, 0x84, 0x5f, 0x96, 0x28, 0x24, 0x49, 0xa0,
	0x73, 0x95, 0x15, 0xc8, 0xe4, 0xe5, 0x34, 0x0e, 0x8e, 0x83, 0xd3, 0x07, 0xa3, 0x75, 0x4c, 0x0e,
	0x20, 0xba, 0x16, 0x58, 0x5c, 0x4e, 0xe3, 0x96, 0xae, 0x94, 0x91, 0xca, 0x5f, 0x71, 0xa1, 0x26,
	0x42, 0x93, 0x37, 0x11, 0x89, 0xe1, 0xfe, 0x80, 0x33, 0x89, 0x4c, 0xc6, 0x6d, 0x5d, 0xb0, 0x21,
	0x1d, 0xc2, 0x7e, 0x8d, 0x2e, 0x16, 0x9c, 0x09, 0x24, 0x2f, 0xd5, 0x88, 0x4e, 0x69, 0xfa, 0xc3,
	0xf3, 0x6e, 0xdf, 0x1e, 0xa8, 0x6f, 0x7b, 0x6d, 0x07, 0xfd, 0x08, 0xbd, 0xeb, 0xc5, 0x74, 0xfb,
	0x0c, 0x04, 0xda, 0xe3, 0xf9, 0x72, 0x56, 0xea, 0xd7, 0xdf, 0xae, 0x96, 0xbd, 0x8a, 0x16, 0xa5,
	0x7e, 0x2c, 0x33, 0xb9, 0x14, 0x71, 0x74, 0x1c, 0x9c, 0xee, 0x8d, 0xca, 0x88, 0xa6, 0xb0, 0x5f,
	0xdb, 0x5e, 0x6a, 0xdc, 0x0c, 0x04, 0x95, 0x81, 0x33, 0xe8, 0x0d, 0x71, 0x8e, 0xbb, 0xc8, 0x51,
	0xcb, 0x6b, 0xbd, 0xff, 0x59, 0x7e, 0x02, 0xdd, 0x0b, 0x94, 0x3b, 0x6c, 0x7e, 0x07, 0xc4, 0x6d,
	0xbc, 0xcb, 0x7f, 0x7d, 0x05, 0xe4, 0x43, 0xce, 0xa6, 0x35, 0xd8, 0xe6, 0x96, 0x03, 0xf7, 0x96,
	0xe9, 0x10, 0x9e, 0x54, 0xba, 0x4b, 0xe2, 0x6b, 0xe8, 0x94, 0x29, 0x75, 0x94, 0xd0, 0x8f, 0x5c,
	0xb7, 0xd0, 0x5f, 0xc1, 0x5a, 0xa1, 0xf7, 0xfe, 0x5c, 0x5f, 0xb6, 0x1a, 0x7d, 0x19, 0x36, 0xf8,
	0xb2, 0xdd, 0xe4, 0xcb, 0xdd, 0xbc, 0x70, 0xfe, 0x37, 0x84, 0x47, 0xa5, 0xba, 0xb1, 0x79, 0x88,
	0x64, 0x02, 0x91, 0xb1, 0x30, 0x39, 0x72, 0xce, 0xe5, 0x79, 0x52, 0xc9, 0x8b, 0xc6, 0xba, 0xf9,
	0x55, 0xf4, 0xe9, 0x8f, 0xdf, 0x7f, 0x7e, 0xb6, 0xba, 0xf4, 0xb1, 0x7e, 0xab, 0xab, 0x37, 0xf6,
	0x49, 0x93, 0x1c, 0x22, 0x63, 0x41, 0x97, 0xe1, 0xb3, 0xbc, 0xcb, 0xf0, 0x9a, 0x96, 0x1e, 0x69,
	0x46, 0x4c, 0x0f, 0x6a, 0x8c, 0xf4, 0xab, 0xfa, 0xbd, 0xdf, 0x15, 0xca, 0x18, 0xd2, 0x45, 0xf9,
	0xec, 0xec, 0xa2, 0xbc, 0x16, 0xb6, 0xa8, 0xb3, 0x26, 0xd4, 0x27, 0x08, 0x2f, 0x50, 0x92, 0x67,
	0x9b, 0x3d, 0x5b, 0xce, 0x4e, 0x0e, 0xfd, 0xc5, 0x2a, 0x81, 0x34, 0x11, 0xde, 0x43, 0x5b, 0x59,
	0x92, 0x38, 0x5b, 0xb6, 0x0d, 0x9d, 0x3c, 0x6f, 0xa8, 0x96, 0x90, 0x7b, 0x93, 0x48, 0xd7, 0xdf,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x7e, 0xea, 0x14, 0x7d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	//Создание записи
	Create(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	//Обновление записи
	Update(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	//Удаление записи
	Delete(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	//Возвращает запись по SLUG
	Get(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	//Писк
	Find(ctx context.Context, in *FindCommentRequest, opts ...grpc.CallOption) (*FindCommentResponse, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) Create(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CommentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) Update(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CommentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) Delete(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CommentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) Get(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CommentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) Find(ctx context.Context, in *FindCommentRequest, opts ...grpc.CallOption) (*FindCommentResponse, error) {
	out := new(FindCommentResponse)
	err := c.cc.Invoke(ctx, "/protobuf.CommentService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	//Создание записи
	Create(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	//Обновление записи
	Update(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	//Удаление записи
	Delete(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	//Возвращает запись по SLUG
	Get(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
	//Писк
	Find(context.Context, *FindCommentRequest) (*FindCommentResponse, error)
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) Create(ctx context.Context, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCommentServiceServer) Update(ctx context.Context, req *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCommentServiceServer) Delete(ctx context.Context, req *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCommentServiceServer) Get(ctx context.Context, req *GetCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCommentServiceServer) Find(ctx context.Context, req *FindCommentRequest) (*FindCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CommentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Create(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CommentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Update(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CommentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Delete(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CommentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Get(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CommentService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Find(ctx, req.(*FindCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CommentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CommentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CommentService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CommentService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _CommentService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/comment/protobuf/comment.proto",
}
