package blog

import (
	"time"
)

// 博客文章表
type Blog struct {
	BlogID       int64     `gorm:"column:blog_id;primarykey;autoIncrement" json:"blog_id"`
	UserID       int64     `gorm:"column:user_id;not null;index" json:"user_id"` // 作者ID
	Title        string    `gorm:"column:title;type:varchar(255);not null" json:"title"` // 标题
	Content      string    `gorm:"column:content;type:text;not null" json:"content"` // 内容
	Excerpt      string    `gorm:"column:excerpt;type:varchar(500)" json:"excerpt"` // 摘要
	Views        int64     `gorm:"column:views;default:0" json:"views"` // 阅读量
	LikeCount    int64     `gorm:"column:like_count;default:0" json:"like_count"` // 点赞数
	CommentCount int64     `gorm:"column:comment_count;default:0" json:"comment_count"` // 评论数
	Tag          string    `gorm:"column:tag;type:varchar(100)" json:"tag"` // 标签,多个标签用逗号分隔
	IsPublished  bool      `gorm:"column:is_published;default:true" json:"is_published"` // 是否发布
	IsDeleted    bool      `gorm:"column:is_deleted;default:false" json:"is_deleted"` // 软删除
	IsRecommended bool     `gorm:"column:is_recommended;default:false" json:"is_recommended"` // 是否推荐
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

// 博客分类表
type Category struct {
	CategoryID   int64     `gorm:"column:category_id;primarykey;autoIncrement" json:"category_id"`
	Name         string    `gorm:"column:name;type:varchar(50);not null;unique" json:"name"` // 分类名称
	DisplayOrder int64     `gorm:"column:display_order;default:0" json:"display_order"` // 显示顺序
	ArticleCount int64     `gorm:"column:article_count;default:0" json:"article_count"` // 文章数量
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

// 博客标签表
type Tag struct {
	TagID        int64     `gorm:"column:tag_id;primarykey;autoIncrement" json:"tag_id"`
	Name         string    `gorm:"column:name;type:varchar(50);not null;unique" json:"name"` // 标签名称
	ArticleCount int64     `gorm:"column:article_count;default:0" json:"article_count"` // 文章数量
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

// 博客文章与标签关联表
type BlogTag struct {
	ID        int64     `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	BlogID    int64     `gorm:"column:blog_id;not null;index" json:"blog_id"`
	TagID     int64     `gorm:"column:tag_id;not null;index" json:"tag_id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

// 博客文章与分类关联表
type BlogCategory struct {
	ID          int64     `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	BlogID      int64     `gorm:"column:blog_id;not null;index" json:"blog_id"`
	CategoryID  int64     `gorm:"column:category_id;not null;index" json:"category_id"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

func (Blog) TableName() string {
	return "blogs"
}

func (Category) TableName() string {
	return "categories"
}

func (Tag) TableName() string {
	return "tags"
}

func (BlogTag) TableName() string {
	return "blog_tags"
}

func (BlogCategory) TableName() string {
	return "blog_categories"
}
