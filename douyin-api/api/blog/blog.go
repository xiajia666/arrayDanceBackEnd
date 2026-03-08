package blog

import (
	"demotest/douyin-api/global"
	"demotest/douyin-api/proto/blog"
	"demotest/douyin-api/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PublishBlog 博客发布接口
func PublishBlog(c *gin.Context) {
	// 获取用户ID
	userId, err := util.GetUserIdByToken(c)
	if err != nil {
		c.JSON(401, gin.H{
			"status_code": 401,
			"status_msg":  "用户未登录",
		})
		return
	}

	// 解析请求参数
	// 这里简化处理,实际应该使用bind绑定
	// 获取请求参数
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")
	categoryId := c.PostForm("category_id")
	excerpt := c.PostForm("excerpt")

	// 验证必填参数
	if title == "" || content == "" {
		c.JSON(400, gin.H{
			"status_code": 400,
			"status_msg":  "标题和内容不能为空",
		})
		return
	}

	var cid int64
	if categoryId != "" {
		fmt.Sscanf(categoryId, "%d", &cid)
	}

	// 调用gRPC服务
	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.PublishBlog(c.Request.Context(), &blog.PublishBlogRequest{
		UserId:   userId,
		Title:    title,
		Content:  content,
		Tag:      tag,
		CategoryId: cid,
		Excerpt:  excerpt,
	})

	if err != nil {
		zap.S().Errorf("发布博客失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "发布博客失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"blog_id":     response.BlogId,
	})
}

// GetBlogList 获取博客列表接口
func GetBlogList(c *gin.Context) {
	// 获取请求参数
	userIdStr := c.Query("user_id")
	categoryIdStr := c.Query("category_id")
	tagIdStr := c.Query("tag_id")
	keyword := c.Query("keyword")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	latestTimeStr := c.Query("latest_time")

	var userId, categoryId, tagId, latestTime int64
	if userIdStr != "" {
		fmt.Sscanf(userIdStr, "%d", &userId)
	}
	if categoryIdStr != "" {
		fmt.Sscanf(categoryIdStr, "%d", &categoryId)
	}
	if tagIdStr != "" {
		fmt.Sscanf(tagIdStr, "%d", &tagId)
	}
	if latestTimeStr != "" {
		fmt.Sscanf(latestTimeStr, "%d", &latestTime)
	}

	page := int32(1)
	pageSize := int32(10)
	if pageStr != "" {
		fmt.Sscanf(pageStr, "%d", &page)
	}
	if pageSizeStr != "" {
		fmt.Sscanf(pageSizeStr, "%d", &pageSize)
	}

	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.GetBlogList(c.Request.Context(), &blog.BlogListRequest{
		UserId:      userId,
		CategoryId:  categoryId,
		TagId:       tagId,
		Keyword:     keyword,
		Page:        page,
		PageSize:    pageSize,
		LatestTime:  latestTime,
	})

	if err != nil {
		zap.S().Errorf("获取博客列表失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "获取博客列表失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"blog_list":   response.BlogList,
	})
}

// GetBlogDetail 获取博客详情接口
func GetBlogDetail(c *gin.Context) {
	// 获取博客ID
	blogIdStr := c.Query("blog_id")
	if blogIdStr == "" {
		c.JSON(400, gin.H{
			"status_code": 400,
			"status_msg":  "博客ID不能为空",
		})
		return
	}

	var blogId int64
	fmt.Sscanf(blogIdStr, "%d", &blogId)

	// 获取用户ID(可选)
	userId, _ := util.GetUserIdByToken(c)

	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.GetBlogDetail(c.Request.Context(), &blog.BlogDetailRequest{
		BlogId: blogId,
		UserId: userId,
	})

	if err != nil {
		zap.S().Errorf("获取博客详情失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "获取博客详情失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"blog_info":   response.BlogInfo,
	})
}

// DeleteBlog 删除博客接口
func DeleteBlog(c *gin.Context) {
	// 获取用户ID
	userId, err := util.GetUserIdByToken(c)
	if err != nil {
		c.JSON(401, gin.H{
			"status_code": 401,
			"status_msg":  "用户未登录",
		})
		return
	}

	// 获取博客ID
	blogIdStr := c.PostForm("blog_id")
	if blogIdStr == "" {
		c.JSON(400, gin.H{
			"status_code": 400,
			"status_msg":  "博客ID不能为空",
		})
		return
	}

	var blogId int64
	fmt.Sscanf(blogIdStr, "%d", &blogId)

	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.DeleteBlog(c.Request.Context(), &blog.DeleteBlogRequest{
		UserId: userId,
		BlogId: blogId,
	})

	if err != nil {
		zap.S().Errorf("删除博客失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "删除博客失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
	})
}

// UpdateBlog 更新博客接口
func UpdateBlog(c *gin.Context) {
	// 获取用户ID
	userId, err := util.GetUserIdByToken(c)
	if err != nil {
		c.JSON(401, gin.H{
			"status_code": 401,
			"status_msg":  "用户未登录",
		})
		return
	}

	// 获取请求参数
	blogIdStr := c.PostForm("blog_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	tag := c.PostForm("tag")
	categoryId := c.PostForm("category_id")
	excerpt := c.PostForm("excerpt")

	if blogIdStr == "" {
		c.JSON(400, gin.H{
			"status_code": 400,
			"status_msg":  "博客ID不能为空",
		})
		return
	}

	var blogId, cid int64
	fmt.Sscanf(blogIdStr, "%d", &blogId)
	if categoryId != "" {
		fmt.Sscanf(categoryId, "%d", &cid)
	}

	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.UpdateBlog(c.Request.Context(), &blog.UpdateBlogRequest{
		UserId:   userId,
		BlogId:   blogId,
		Title:    title,
		Content:  content,
		Tag:      tag,
		CategoryId: cid,
		Excerpt:  excerpt,
	})

	if err != nil {
		zap.S().Errorf("更新博客失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "更新博客失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
	})
}

// GetCategoryList 获取分类列表接口
func GetCategoryList(c *gin.Context) {
	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.GetCategoryList(c.Request.Context(), &blog.BlogListRequest{})

	if err != nil {
		zap.S().Errorf("获取分类列表失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "获取分类列表失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"category_list": response.CategoryList,
	})
}

// GetTagList 获取标签列表接口
func GetTagList(c *gin.Context) {
	conn, err := util.GetGrpcConn()
	if err != nil {
		zap.S().Errorf("获取gRPC连接失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	client := blog.NewBlogServiceClient(conn)
	response, err := client.GetTagList(c.Request.Context(), &blog.BlogListRequest{})

	if err != nil {
		zap.S().Errorf("获取标签列表失败: %v", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"status_msg":  "获取标签列表失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": response.StatusCode,
		"status_msg":  response.StatusMsg,
		"tag_list": response.TagList,
	})
}
