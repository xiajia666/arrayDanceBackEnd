package router

import (
	"demotest/douyin-api/api/blog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitBlogRouter 博客路由配置
func InitBlogRouter(Router *gin.RouterGroup) {
	BlogRouter := Router.Group("/blog")

	zap.S().Info("配置博客相关的url")

	{
		// 获取博客列表(不需要登录)
		BlogRouter.GET("/list", blog.GetBlogList)

		// 获取博客详情(不需要登录)
		BlogRouter.GET("/detail", blog.GetBlogDetail)

		// 发布博客(需要登录)
		BlogRouter.POST("/publish", blog.PublishBlog)

		// 更新博客(需要登录)
		BlogRouter.POST("/update", blog.UpdateBlog)

		// 删除博客(需要登录)
		BlogRouter.POST("/delete", blog.DeleteBlog)

		// 获取分类列表
		BlogRouter.GET("/categories", blog.GetCategoryList)

		// 获取标签列表
		BlogRouter.GET("/tags", blog.GetTagList)
	}
}
