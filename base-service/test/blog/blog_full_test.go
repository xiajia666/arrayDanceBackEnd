package blog

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:8889"
)

// 测试用例说明
// 1. 启动博客gRPC服务: go run base-service/cmd/blog/main.go
// 2. 运行测试: go test -v ./base-service/test/blog/

// TestBlogCRUD 测试博客CRUD操作
func TestBlogCRUD(t *testing.T) {
	// 连接gRPC服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	// 1. 发布博客
	fmt.Println("=== 测试发布博客 ===")
	publishReq := &blog.PublishBlogRequest{
		UserId:  1,
		Title:   "Go语言最佳实践",
		Content: "# Go语言最佳实践\n\n这是一篇关于Go语言最佳实践的博客...",
		Tag:     "Go,后端开发,最佳实践",
	}
	publishRes, err := c.PublishBlog(context.Background(), publishReq)
	if err != nil {
		t.Fatalf("发布博客失败: %v", err)
	}
	fmt.Printf("博客发布成功, ID: %d\n", publishRes.BlogId)

	// 2. 获取博客列表
	fmt.Println("\n=== 测试获取博客列表 ===")
	listReq := &blog.BlogListRequest{
		Page:     1,
		PageSize: 10,
	}
	listRes, err := c.GetBlogList(context.Background(), listReq)
	if err != nil {
		t.Fatalf("获取博客列表失败: %v", err)
	}
	fmt.Printf("博客列表共 %d 篇\n", len(listRes.BlogList))
	for i, b := range listRes.BlogList {
		fmt.Printf("%d. [%s] %s - 作者: %s, 阅读: %d\n",
			i+1, b.CategoryName, b.Title, b.AuthorName, b.Views)
	}

	// 3. 获取博客详情
	fmt.Println("\n=== 测试获取博客详情 ===")
	detailReq := &blog.BlogDetailRequest{
		BlogId: publishRes.BlogId,
		UserId: 1,
	}
	detailRes, err := c.GetBlogDetail(context.Background(), detailReq)
	if err != nil {
		t.Fatalf("获取博客详情失败: %v", err)
	}
	fmt.Printf("博客详情:\n标题: %s\n内容: %s\n作者: %s\n阅读量: %d\n",
		detailRes.BlogInfo.Title,
		detailRes.BlogInfo.Content,
		detailRes.BlogInfo.AuthorName,
		detailRes.BlogInfo.Views)

	// 4. 更新博客
	fmt.Println("\n=== 测试更新博客 ===")
	updateReq := &blog.UpdateBlogRequest{
		UserId:  1,
		BlogId:  publishRes.BlogId,
		Title:   "Go语言最佳实践(更新版)",
		Content: "# Go语言最佳实践(更新版)\n\n这是一篇更新后的博客...",
		Tag:     "Go,后端开发,最佳实践,更新",
	}
	_, err = c.UpdateBlog(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("更新博客失败: %v", err)
	}
	fmt.Println("博客更新成功")

	// 5. 获取分类列表
	fmt.Println("\n=== 测试获取分类列表 ===")
	_, err = c.GetCategoryList(context.Background(), &blog.BlogListRequest{})
	if err != nil {
		t.Fatalf("获取分类列表失败: %v", err)
	}
	fmt.Println("分类列表获取成功")

	// 6. 获取标签列表
	fmt.Println("\n=== 测试获取标签列表 ===")
	_, err = c.GetTagList(context.Background(), &blog.BlogListRequest{})
	if err != nil {
		t.Fatalf("获取标签列表失败: %v", err)
	}
	fmt.Println("标签列表获取成功")

	// 7. 删除博客
	// 注意: 这里注释掉删除操作,保留测试数据
	/*
		fmt.Println("\n=== 测试删除博客 ===")
		deleteReq := &blog.DeleteBlogRequest{
			UserId: 1,
			BlogId: publishRes.BlogId,
		}
		_, err = c.DeleteBlog(context.Background(), deleteReq)
		if err != nil {
			t.Fatalf("删除博客失败: %v", err)
		}
		fmt.Println("博客删除成功")
	*/
}

// BenchmarkBlogPublish 基准测试 - 博客发布
func BenchmarkBlogPublish(b *testing.B) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		b.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := &blog.PublishBlogRequest{
			UserId:  int64(i % 100),
			Title:   fmt.Sprintf("测试博客 %d", i),
			Content: fmt.Sprintf("这是第 %d 篇测试博客", i),
			Tag:     "测试",
		}
		_, err := c.PublishBlog(context.Background(), req)
		if err != nil {
			b.Fatalf("发布博客失败: %v", err)
		}
	}
}
