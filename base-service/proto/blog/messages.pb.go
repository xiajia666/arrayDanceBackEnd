// BlogListRequest and other messages

type BlogListRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId           int64    `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	TagId                int64    `protobuf:"varint,3,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Keyword              string   `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page                 int32    `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	LatestTime           int64    `protobuf:"varint,7,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogListRequest) Reset()         { *m = BlogListRequest{} }
func (m *BlogListRequest) String() string { return proto.CompactTextString(m) }
func (*BlogListRequest) ProtoMessage()    {}
func (*BlogListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{4}
}

func (m *BlogListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogListRequest.Unmarshal(m, b)
}
func (m *BlogListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogListRequest.Marshal(b, m, deterministic)
}
func (m *BlogListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogListRequest.Merge(m, src)
}
func (m *BlogListRequest) XXX_Size() int {
	return xxx_messageInfo_BlogListRequest.Size(m)
}
func (m *BlogListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogListRequest proto.InternalMessageInfo

func (m *BlogListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BlogListRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *BlogListRequest) GetTagId() int64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *BlogListRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *BlogListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BlogListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *BlogListRequest) GetLatestTime() int64 {
	if m != nil {
		return m.LatestTime
	}
	return 0
}

type BlogListResponse struct {
	StatusCode           int64       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	BlogList             []*BlogInfo `protobuf:"bytes,3,rep,name=blog_list,json=blogList,proto3" json:"blog_list,omitempty"`
	TotalCount           int64       `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	LatestTime           int64       `protobuf:"varint,5,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlogListResponse) Reset()         { *m = BlogListResponse{} }
func (m *BlogListResponse) String() string { return proto.CompactTextString(m) }
func (*BlogListResponse) ProtoMessage()    {}
func (*BlogListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{5}
}

func (m *BlogListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogListResponse.Unmarshal(m, b)
}
func (m *BlogListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogListResponse.Marshal(b, m, deterministic)
}
func (m *BlogListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogListResponse.Merge(m, src)
}
func (m *BlogListResponse) XXX_Size() int {
	return xxx_messageInfo_BlogListResponse.Size(m)
}
func (m *BlogListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogListResponse proto.InternalMessageInfo

func (m *BlogListResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *BlogListResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *BlogListResponse) GetBlogList() []*BlogInfo {
	if m != nil {
		return m.BlogList
	}
	return nil
}

func (m *BlogListResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *BlogListResponse) GetLatestTime() int64 {
	if m != nil {
		return m.LatestTime
	}
	return 0
}

type BlogDetailRequest struct {
	BlogId               int64    `protobuf:"varint,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogDetailRequest) Reset()         { *m = BlogDetailRequest{} }
func (m *BlogDetailRequest) String() string { return proto.CompactTextString(m) }
func (*BlogDetailRequest) ProtoMessage()    {}
func (*BlogDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{6}
}

func (m *BlogDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDetailRequest.Unmarshal(m, b)
}
func (m *BlogDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDetailRequest.Marshal(b, m, deterministic)
}
func (m *BlogDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDetailRequest.Merge(m, src)
}
func (m *BlogDetailRequest) XXX_Size() int {
	return xxx_messageInfo_BlogDetailRequest.Size(m)
}
func (m *BlogDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDetailRequest proto.InternalMessageInfo

func (m *BlogDetailRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

func (m *BlogDetailRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type BlogDetailResponse struct {
	StatusCode           int64       `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	BlogInfo             *BlogInfo   `protobuf:"bytes,3,opt,name=blog_info,json=blogInfo,proto3" json:"blog_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlogDetailResponse) Reset()         { *m = BlogDetailResponse{} }
func (m *BlogDetailResponse) String() string { return proto.CompactTextString(m) }
func (*BlogDetailResponse) ProtoMessage()    {}
func (*BlogDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{7}
}

func (m *BlogDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDetailResponse.Unmarshal(m, b)
}
func (m *BlogDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDetailResponse.Marshal(b, m, deterministic)
}
func (m *BlogDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDetailResponse.Merge(m, src)
}
func (m *BlogDetailResponse) XXX_Size() int {
	return xxx_messageInfo_BlogDetailResponse.Size(m)
}
func (m *BlogDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDetailResponse proto.InternalMessageInfo

func (m *BlogDetailResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *BlogDetailResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *BlogDetailResponse) GetBlogInfo() *BlogInfo {
	if m != nil {
		return m.BlogInfo
	}
	return nil
}

type DeleteBlogRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BlogId               int64    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRequest) Reset()         { *m = DeleteBlogRequest{} }
func (m *DeleteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRequest) ProtoMessage()    {}
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{8}
}

func (m *DeleteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRequest.Unmarshal(m, b)
}
func (m *DeleteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRequest.Merge(m, src)
}
func (m *DeleteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRequest.Size(m)
}
func (m *DeleteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRequest proto.InternalMessageInfo

func (m *DeleteBlogRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DeleteBlogRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

type DeleteBlogResponse struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogResponse) Reset()         { *m = DeleteBlogResponse{} }
func (m *DeleteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogResponse) ProtoMessage()    {}
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{9}
}

func (m *DeleteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogResponse.Unmarshal(m, b)
}
func (m *DeleteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogResponse.Merge(m, src)
}
func (m *DeleteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogResponse.Size(m)
}
func (m *DeleteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogResponse proto.InternalMessageInfo

func (m *DeleteBlogResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DeleteBlogResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

type UpdateBlogRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BlogId               int64    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Tag                  string   `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	CategoryId           int64    `protobuf:"varint,6,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Excerpt              string   `protobuf:"bytes,7,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRequest) Reset()         { *m = UpdateBlogRequest{} }
func (m *UpdateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRequest) ProtoMessage()    {}
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{10}
}

func (m *UpdateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRequest.Unmarshal(m, b)
}
func (m *UpdateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRequest.Merge(m, src)
}
func (m *UpdateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRequest.Size(m)
}
func (m *UpdateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRequest proto.InternalMessageInfo

func (m *UpdateBlogRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateBlogRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

func (m *UpdateBlogRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateBlogRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdateBlogRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *UpdateBlogRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *UpdateBlogRequest) GetExcerpt() string {
	if m != nil {
		return m.Excerpt
	}
	return ""
}

type UpdateBlogResponse struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogResponse) Reset()         { *m = UpdateBlogResponse{} }
func (m *UpdateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogResponse) ProtoMessage()    {}
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{11}
}

func (m *UpdateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogResponse.Unmarshal(m, b)
}
func (m *UpdateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogResponse.Marshal(b, m, deterministic)
}
func (m *UpdateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogResponse.Merge(m, src)
}
func (m *UpdateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogResponse.Size(m)
}
func (m *UpdateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogResponse proto.InternalMessageInfo

func (m *UpdateBlogResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UpdateBlogResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

type CategoryListResponse struct {
	StatusCode           int64           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	CategoryList         []*CategoryInfo `protobuf:"bytes,3,rep,name=category_list,json=categoryList,proto3" json:"category_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CategoryListResponse) Reset()         { *m = CategoryListResponse{} }
func (m *CategoryListResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryListResponse) ProtoMessage()    {}
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{12}
}

func (m *CategoryListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryListResponse.Unmarshal(m, b)
}
func (m *CategoryListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryListResponse.Marshal(b, m, deterministic)
}
func (m *CategoryListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryListResponse.Merge(m, src)
}
func (m *CategoryListResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryListResponse.Size(m)
}
func (m *CategoryListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryListResponse proto.InternalMessageInfo

func (m *CategoryListResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CategoryListResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *CategoryListResponse) GetCategoryList() []*CategoryInfo {
	if m != nil {
		return m.CategoryList
	}
	return nil
}

type CategoryInfo struct {
	CategoryId           int64    `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ArticleCount         int64    `protobuf:"varint,3,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	DisplayOrder         int64    `protobuf:"varint,4,opt,name=display_order,json=displayOrder,proto3" json:"display_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryInfo) Reset()         { *m = CategoryInfo{} }
func (m *CategoryInfo) String() string { return proto.CompactTextString(m) }
func (*CategoryInfo) ProtoMessage()    {}
func (*CategoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{13}
}

func (m *CategoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryInfo.Unmarshal(m, b)
}
func (m *CategoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryInfo.Marshal(b, m, deterministic)
}
func (m *CategoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryInfo.Merge(m, src)
}
func (m *CategoryInfo) XXX_Size() int {
	return xxx_messageInfo_CategoryInfo.Size(m)
}
func (m *CategoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryInfo proto.InternalMessageInfo

func (m *CategoryInfo) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *CategoryInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryInfo) GetArticleCount() int64 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}

func (m *CategoryInfo) GetDisplayOrder() int64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

type TagListResponse struct {
	StatusCode           int64        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg            string       `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	TagList              []*TagInfo   `protobuf:"bytes,3,rep,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TagListResponse) Reset()         { *m = TagListResponse{} }
func (m *TagListResponse) String() string { return proto.CompactTextString(m) }
func (*TagListResponse) ProtoMessage()    {}
func (*TagListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{14}
}

func (m *TagListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagListResponse.Unmarshal(m, b)
}
func (m *TagListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagListResponse.Marshal(b, m, deterministic)
}
func (m *TagListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagListResponse.Merge(m, src)
}
func (m *TagListResponse) XXX_Size() int {
	return xxx_messageInfo_TagListResponse.Size(m)
}
func (m *TagListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TagListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TagListResponse proto.InternalMessageInfo

func (m *TagListResponse) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *TagListResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *TagListResponse) GetTagList() []*TagInfo {
	if m != nil {
		return m.TagList
	}
	return nil
}

type TagInfo struct {
	TagId                int64    `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ArticleCount         int64    `protobuf:"varint,3,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagInfo) Reset()         { *m = TagInfo{} }
func (m *TagInfo) String() string { return proto.CompactTextString(m) }
func (*TagInfo) ProtoMessage()    {}
func (*TagInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f9c95c3c9e8b0a5, []int{15}
}

func (m *TagInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagInfo.Unmarshal(m, b)
}
func (m *TagInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagInfo.Marshal(b, m, deterministic)
}
func (m *TagInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagInfo.Merge(m, src)
}
func (m *TagInfo) XXX_Size() int {
	return xxx_messageInfo_TagInfo.Size(m)
}
func (m *TagInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TagInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TagInfo proto.InternalMessageInfo

func (m *TagInfo) GetTagId() int64 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *TagInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagInfo) GetArticleCount() int64 {
	if m != nil {
		return m.ArticleCount
	}
	return 0
}
