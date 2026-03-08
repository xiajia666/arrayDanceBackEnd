package router

import (
	"arrayDanceBackEnddouyin-api/api"
	"arrayDanceBackEnddouyin-api/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitFavoriteRouter(Router *gin.RouterGroup) {
	FavoriteRouter := Router.Group("/douyin/favorite/")
	zap.S().Info("配置点赞相关的url")
	{
		FavoriteRouter.POST("action/", util.AuthMiddleware(), api.FavoriteAction)
		FavoriteRouter.GET("list/", util.AuthMiddleware(), api.FavoriteList)
	}
}
