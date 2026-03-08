package blog

import (
	"context"
	"demotest/base-service/global"
	"demotest/base-service/model/blog"
	"demotest/base-service/proto/blog"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type BlogServe struct{}

// PublishBlog 博客发布功能
func (*BlogServe) PublishBlog(ctx context.Context, in *blog.PublishBlogRequest) (*blog.PublishBlogResponse, error) {
	// 创建博客文章
	article := &blog.Blog{
		UserID:   in.UserId,
		Title:    in.Title,
		Content:  in.Content,
		Tag:      in.Tag,
		Excerpt:  in.Excerpt,
	}

	// 如果没有提供摘要,自动生成
	if article.Excerpt == "" {
		if len(article.Content) > 200 {
			article.Excerpt = article.Content[:200]
		} else {
			article.Excerpt = article.Content
		}
	}

	// 保存博客文章
	result := global.DB.Create(article)
	if result.Error != nil {
		return nil, errors.New("博客发布失败: " + result.Error.Error())
	}

	// 如果指定了分类,添加分类关联
	if in.CategoryId > 0 {
		blogCategory := &blog.BlogCategory{
			BlogID:     article.BlogID,
			CategoryID: in.CategoryId,
		}
		global.DB.Create(blogCategory)
	}

	return &blog.PublishBlogResponse{
		StatusCode: 0,
		StatusMsg:  "博客发布成功",
		BlogId:     article.BlogID,
	}, nil
}

// GetBlogList 获取博客列表
func (*BlogServe) GetBlogList(ctx context.Context, in *blog.BlogListRequest) (*blog.BlogListResponse, error) {
	var blogs []blog.Blog
	query := global.DB.Model(&blog.Blog{}).Where("is_deleted = ?", false)

	// 筛选条件
	if in.UserId > 0 {
		query = query.Where("user_id = ?", in.UserId)
	}

	// 按时间倒序
	query = query.Order("create_time DESC")

	// 获取列表
	result := query.Find(&blogs)
	if result.Error != nil {
		return nil, errors.New("获取博客列表失败: " + result.Error.Error())
	}

	// 处理博客列表,添加作者信息
	infos := make([]*blog.BlogInfo, len(blogs))
	for i, b := range blogs {
		// 获取作者信息
		author := &blog.User{}
		global.DB.Model(&blog.User{}).Where("id = ?", b.UserID).First(&author)

		// 转换时间
		creatTime, _ := ptypes.TimestampProto(b.CreateTime)
		updatTime, _ := ptypes.TimestampProto(b.UpdateTime)

		infos[i] = &blog.BlogInfo{
			BlogId:       b.BlogID,
			UserId:       b.UserID,
			Title:        b.Title,
			Excerpt:      b.Excerpt,
			Views:        b.Views,
			LikeCount:    b.LikeCount,
			CommentCount: b.CommentCount,
			Tag:          b.Tag,
			AuthorName:   author.Name,
			AuthorAvatar: author.Avatar,
			CreateTime:   fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
				creatTime.GetYear(), creatTime.GetMonth(), creatTime.GetDay(),
				creatTime.GetHours(), creatTime.GetMinutes(), creatTime.GetSeconds()),
			UpdateTime: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
				updatTime.GetYear(), updatTime.GetMonth(), updatTime.GetDay(),
				updatTime.GetHours(), updatTime.GetMinutes(), updatTime.GetSeconds()),
		}
	}

	return &blog.BlogListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		BlogList:   infos,
	}, nil
}

// GetBlogDetail 获取博客详情
func (*BlogServe) GetBlogDetail(ctx context.Context, in *blog.BlogDetailRequest) (*blog.BlogDetailResponse, error) {
	var blogItem blog.Blog
	result := global.DB.Model(&blog.Blog{}).Where("blog_id = ? AND is_deleted = ?", in.BlogId, false).First(&blogItem)
	if result.Error != nil {
		return nil, errors.New("博客不存在或已被删除: " + result.Error.Error())
	}

	// 增加阅读量
	// 使用Redis缓存阅读量,定时同步到数据库
	viewsKey := fmt.Sprintf("blog:views:%d", blogItem.BlogID)
	global.RS.Incr(viewsKey)

	// 获取作者信息
	author := &blog.User{}
	global.DB.Model(&blog.User{}).Where("id = ?", blogItem.UserID).First(&author)

	// 转换时间
	creatTime, _ := ptypes.TimestampProto(blogItem.CreateTime)
	updatTime, _ := ptypes.TimestampProto(blogItem.UpdateTime)

	blogInfo := &blog.BlogInfo{
		BlogId:       blogItem.BlogID,
		UserId:       blogItem.UserID,
		Title:        blogItem.Title,
		Content:      blogItem.Content,
		Excerpt:      blogItem.Excerpt,
		Views:        blogItem.Views,
		LikeCount:    blogItem.LikeCount,
		CommentCount: blogItem.CommentCount,
		Tag:          blogItem.Tag,
		IsPublished:  blogItem.IsPublished,
		AuthorName:   author.Name,
		AuthorAvatar: author.Avatar,
		CreateTime:   fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			creatTime.GetYear(), creatTime.GetMonth(), creatTime.GetDay(),
			creatTime.GetHours(), creatTime.GetMinutes(), creatTime.GetSeconds()),
		UpdateTime: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			updatTime.GetYear(), updatTime.GetMonth(), updatTime.GetDay(),
			updatTime.GetHours(), updatTime.GetMinutes(), updatTime.GetSeconds()),
	}

	return &blog.BlogDetailResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		BlogInfo:   blogInfo,
	}, nil
}

// DeleteBlog 删除博客
func (*BlogServe) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.DeleteBlogResponse, error) {
	// 检查博客是否存在且属于该用户
	var blogItem blog.Blog
	result := global.DB.Model(&blog.Blog{}).Where("blog_id = ? AND user_id = ?", in.BlogId, in.UserId).First(&blogItem)
	if result.Error != nil {
		return nil, errors.New("博客不存在或无权删除: " + result.Error.Error())
	}

	// 软删除
	// global.DB.Model(&blog.Blog{}).Where("blog_id = ?", in.BlogId).Update("is_deleted", true)
	blogItem.IsDeleted = true
	global.DB.Save(&blogItem)

	return &blog.DeleteBlogResponse{
		StatusCode: 0,
		StatusMsg:  "博客删除成功",
	}, nil
}

// UpdateBlog 更新博客
func (*BlogServe) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.UpdateBlogResponse, error) {
	// 检查博客是否存在且属于该用户
	var blogItem blog.Blog
	result := global.DB.Model(&blog.Blog{}).Where("blog_id = ? AND user_id = ?", in.BlogId, in.UserId).First(&blogItem)
	if result.Error != nil {
		return nil, errors.New("博客不存在或无权修改: " + result.Error.Error())
	}

	// 更新博客信息
	updates := make(map[string]interface{})
	if in.Title != "" {
		updates["title"] = in.Title
	}
	if in.Content != "" {
		updates["content"] = in.Content
		// 更新摘要
		updates["excerpt"] = in.Excerpt
		if in.Excerpt == "" {
			if len(in.Content) > 200 {
				updates["excerpt"] = in.Content[:200]
			} else {
				updates["excerpt"] = in.Content
			}
		}
	}
	if in.Tag != "" {
		updates["tag"] = in.Tag
	}

	if len(updates) > 0 {
		global.DB.Model(&blog.Blog{}).Where("blog_id = ?", in.BlogId).Updates(updates)
	}

	return &blog.UpdateBlogResponse{
		StatusCode: 0,
		StatusMsg:  "博客更新成功",
	}, nil
}

// GetCategoryList 获取分类列表
func (*BlogServe) GetCategoryList(ctx context.Context, in *blog.BlogListRequest) (*blog.CategoryListResponse, error) {
	var categories []blog.Category
	result := global.DB.Model(&blog.Category{}).Order("display_order ASC, create_time DESC").Find(&categories)
	if result.Error != nil {
		return nil, errors.New("获取分类列表失败: " + result.Error.Error())
	}

	infos := make([]*blog.CategoryInfo, len(categories))
	for i, c := range categories {
		infos[i] = &blog.CategoryInfo{
			CategoryId:   c.CategoryID,
			Name:         c.Name,
			ArticleCount: c.ArticleCount,
			DisplayOrder: c.DisplayOrder,
		}
	}

	return &blog.CategoryListResponse{
		StatusCode:    0,
		StatusMsg:     "success",
		CategoryList:  infos,
	}, nil
}

// GetTagList 获取标签列表
func (*BlogServe) GetTagList(ctx context.Context, in *blog.BlogListRequest) (*blog.TagListResponse, error) {
	var tags []blog.Tag
	result := global.DB.Model(&blog.Tag{}).Order("article_count DESC").Limit(20).Find(&tags)
	if result.Error != nil {
		return nil, errors.New("获取标签列表失败: " + result.Error.Error())
	}

	infos := make([]*blog.TagInfo, len(tags))
	for i, t := range tags {
		infos[i] = &blog.TagInfo{
			TagId:         t.TagID,
			Name:          t.Name,
			ArticleCount:  t.ArticleCount,
		}
	}

	return &blog.TagListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		TagList:    infos,
	}, nil
}

// GetUser 引用现有的用户模型
type User struct {
	Id              int64
	Name            string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
	Password        string
}
