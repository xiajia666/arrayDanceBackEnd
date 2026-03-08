package blog

import (
	"context"
	"demotest/base-service/proto/blog"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

const (
	address = "127.0.0.1:8889"
)

// TestPublishBlog 测试发布博客
func TestPublishBlog(t *testing.T) {
	// 连接grpc服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	// 发布博客
	req := &blog.PublishBlogRequest{
		UserId:   1,
		Title:    "我的第一篇博客",
		Content:  "这是一篇测试博客内容...",
		Tag:      "技术,Go",
		Excerpt:  "这是一篇测试博客摘要...",
	}
	res, err := c.PublishBlog(context.Background(), req)
	if err != nil {
		t.Fatalf("could not publish blog: %v", err)
	}

	fmt.Printf("Blog published successfully, ID: %d\n", res.BlogId)
}

// TestGetBlogList 测试获取博客列表
func TestGetBlogList(t *testing.T) {
	// 连接grpc服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	// 获取博客列表
	req := &blog.BlogListRequest{
		Page:     1,
		PageSize: 10,
	}
	res, err := c.GetBlogList(context.Background(), req)
	if err != nil {
		t.Fatalf("could not get blog list: %v", err)
	}

	fmt.Printf("Blog list retrieved, total: %d\n", len(res.BlogList))
	for _, b := range res.BlogList {
		fmt.Printf("Blog ID: %d, Title: %s\n", b.BlogId, b.Title)
	}
}

// TestGetBlogDetail 测试获取博客详情
func TestGetBlogDetail(t *testing.T) {
	// 连接grpc服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	// 获取博客详情
	// 请替换为实际的博客ID
	blogId := int64(1)
	req := &blog.BlogDetailRequest{
		BlogId: blogId,
	}
	res, err := c.GetBlogDetail(context.Background(), req)
	if err != nil {
		t.Fatalf("could not get blog detail: %v", err)
	}

	fmt.Printf("Blog detail: %+v\n", res.BlogInfo)
}

// TestDeleteBlog 测试删除博客
func TestDeleteBlog(t *testing.T) {
	// 连接grpc服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	// 删除博客
	// 请替换为实际的博客ID和用户ID
	blogId := int64(1)
	userId := int64(1)
	req := &blog.DeleteBlogRequest{
		UserId: userId,
		BlogId: blogId,
	}
	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		t.Fatalf("could not delete blog: %v", err)
	}

	fmt.Printf("Blog deleted: %+v\n", res)
}
