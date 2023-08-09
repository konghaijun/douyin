package router

import (
	"github.com/KumaJie/douyin/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	videoCtrl := &controller.VideoController{}
	router.POST("/douyin/publish/action", videoCtrl.DouyinPublishActionHandler)
	// 定义路由和处理函数
	router.GET("/douyin/feed", videoCtrl.DouyinFeedHandler)

	return router
}
