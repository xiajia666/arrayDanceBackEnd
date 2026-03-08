package blog

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlogInfo struct {
	BlogId               int64    `protobuf:"varint,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Excerpt              string   `protobuf:"bytes,5,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	Views                int64    `protobuf:"varint,6,opt,name=views,proto3" json:"views,omitempty"`
	LikeCount            int64    `protobuf:"varint,7,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	CommentCount         int64    `protobuf:"varint,8,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	Tag                  string   `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`
	IsPublished          bool     `protobuf:"varint,10,opt,name=is_published,json=isPublished,proto3" json:"is_published,omitempty"`
	CategoryId           int64    `protobuf:"varint,11,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CategoryName         string   `protobuf:"bytes,12,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	AuthorName           string   `protobuf:"bytes,13,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	AuthorAvatar         string   `protobuf:"bytes,14,opt,name=author_avatar,json=authorAvatar,proto3" json:"author_avatar,omitempty"`
	CreateTime           string   `protobuf:"bytes,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,16,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsRecommended        bool     `protobuf:"varint,17,opt,name=is_recommended,json=isRecommended,proto3" json:"is_recommended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogInfo) Reset()         { *m = BlogInfo{} }
func (m *BlogInfo) String() string { return proto.CompactTextString(m) }
func (*BlogInfo) ProtoMessage()    {}
func (*BlogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{0}
}

func (m *BlogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogInfo.Unmarshal(m, b)
}
func (m *BlogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogInfo.Marshal(b, m, deterministic)
}
func (m *BlogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogInfo.Merge(m, src)
}
func (m *BlogInfo) XXX_Size() int {
	return xxx_messageInfo_BlogInfo.Size(m)
}
func (m *BlogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlogInfo proto.InternalMessageInfo

func (m *BlogInfo) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

func (m *BlogInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BlogInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BlogInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *BlogInfo) GetExcerpt() string {
	if m != nil {
		return m.Excerpt
	}
	return ""
}

func (m *BlogInfo) GetViews() int64 {
	if m != nil {
		return m.Views
	}
	return 0
}

func (m *BlogInfo) GetLikeCount() int64 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *BlogInfo) GetCommentCount() int64 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

func (m *BlogInfo) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BlogInfo) GetIsPublished() bool {
	if m != nil {
		return m.IsPublished
	}
	return false
}

func (m *BlogInfo) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *BlogInfo) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *BlogInfo) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

func (m *BlogInfo) GetAuthorAvatar() string {
	if m != nil {
		return m.AuthorAvatar
	}
	return ""
}

func (m *BlogInfo) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *BlogInfo) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *BlogInfo) GetIsRecommended() bool {
	if m != nil {
		return m.IsRecommended
	}
	return false
}

type UserInfo struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type PublishBlogRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	CategoryId           int64    `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Excerpt              string   `protobuf:"bytes,6,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishBlogRequest) Reset()         { *m = PublishBlogRequest{} }
func (m *PublishBlogRequest) String() string { return proto.CompactTextString(m) }
func (*PublishBlogRequest) ProtoMessage()    {}
func (*PublishBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{2}
}

func (m *PublishBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishBlogRequest.Unmarshal(m, b)
}
func (m *PublishBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishBlogRequest.Marshal(b, m, deterministic)
}
func (m *PublishBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishBlogRequest.Merge(m, src)
}
func (m *PublishBlogRequest) XXX_Size() int {
	return xxx_messageInfo_PublishBlogRequest.Size(m)
}
func (m *PublishBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishBlogRequest proto.InternalMessageInfo

func (m *PublishBlogRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PublishBlogRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PublishBlogRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PublishBlogRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *PublishBlogRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *PublishBlogRequest) GetExcerpt() string {
	if m != nil {
		return m.Excerpt
	}
	return ""
}

type PublishBlogResponse struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	BlogId               int64    `protobuf:"varint,3,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishBlogResponse) Reset()         { *m = PublishBlogResponse{} }
func (m *PublishBlogResponse) String() string { return proto.CompactTextString(m) }
func (*PublishBlogResponse) ProtoMessage()    {}
func (*PublishBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{3}
}

func (m *PublishBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishBlogResponse.Unmarshal(m, b)
}
func (m *PublishBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishBlogResponse.Marshal(b, m, deterministic)
}
func (m *PublishBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishBlogResponse.Merge(m, src)
}
func (m *PublishBlogResponse) XXX_Size() int {
	return xxx_messageInfo_PublishBlogResponse.Size(m)
}
func (m *PublishBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishBlogResponse proto.InternalMessageInfo

func (m *PublishBlogResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *PublishBlogResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *PublishBlogResponse) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

// ... (继续生成其他message和service的代码)
// 由于篇幅限制,这里省略了其他部分的生成代码
// 实际使用时需要使用protoc命令完整生成

func init() { proto.RegisterFile("blog.proto", fileDescriptor_7f9c95c3c9e8b0a5) }

var fileDescriptor_7f9c95c3c9e8b0a5 = []byte{
	// 此处应为完整的proto编译后的字节码
	// 需要使用 protoc --go_out=plugins=grpc:. blog.proto 命令生成
}
