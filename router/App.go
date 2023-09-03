package router

import (
	"github.com/KumaJie/douyin/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	videoCtrl := &controller.VideoController{}
	userCtrl := &controller.UserController{}
	router.POST("/douyin/publish/action/", videoCtrl.DouyinPublishActionHandler)

	router.GET("/douyin/feed/", videoCtrl.DouyinFeedHandler)

	//注册
	router.POST("/douyin/user/register/", userCtrl.DouyinUserRegisterHandler)

	//登录
	router.POST("/douyin/user/login/", userCtrl.DouyinUserLoginHandler)

	//查用戶信息
	router.GET("/douyin/user/", userCtrl.DouyinUserHandler)

	return router
}
