package router

import (
	"github.com/KumaJie/douyin/controller"
	"github.com/KumaJie/douyin/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	videoCtrl := &controller.VideoController{}
	userCtrl := &controller.UserController{}
	favoriteCtrl := &controller.FavoriteController{}
	commentCtrl := &controller.CommentController{}

	router.Static("/static", "./static")
	//视频投稿
	router.POST("/douyin/publish/action/", utils.JWTMiddleWare(), videoCtrl.DouyinPublishActionHandler)

	//视频feed流
	router.GET("/douyin/feed/", videoCtrl.DouyinFeedHandler)

	//注册
	router.POST("/douyin/user/register/", userCtrl.DouyinUserRegisterHandler)

	//登录
	router.POST("/douyin/user/login/", userCtrl.DouyinUserLoginHandler)

	//查用戶信息
	router.GET("/douyin/user/", utils.JWTMiddleWare(), userCtrl.DouyinUserHandler)

	//查看发布信息
	router.GET("/douyin/publish/list/", utils.JWTMiddleWare(), videoCtrl.DouyinPublishList)

	//点赞
	router.POST("/douyin/favorite/action/", utils.JWTMiddleWare(), favoriteCtrl.DouyinFavoriteAction)

	//喜欢列表
	router.GET("/douyin/favorite/list/", utils.JWTMiddleWare(), videoCtrl.GetUserFavoriteHandler)

	//评论
	router.POST("/douyin/comment/action/", utils.JWTMiddleWare(), commentCtrl.CommentActionHandler)

	return router
}
