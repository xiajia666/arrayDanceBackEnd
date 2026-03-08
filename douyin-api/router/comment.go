package router

import (
	"arrayDanceBackEnddouyin-api/api"
	"arrayDanceBackEnddouyin-api/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("/douyin/comment/")
	zap.S().Info("配置评论相关的url")
	{
		CommentRouter.POST("action/", util.CommentMiddleware(), api.CommentAction)
		CommentRouter.GET("list/", api.CommentList)
	}
}
