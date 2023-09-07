package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct {
	videoService *service.VideoService
}

// 支持所有用户刷抖音视频
func (ctrl *VideoController) DouyinFeedHandler(c *gin.Context) {

	// 调用 service.GetDouyinFeed 获取 Douyin Feed 数据
	response, err := ctrl.videoService.GetDouyinFeed(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 将服务层返回的数据作为 JSON 响应发送
	c.JSON(http.StatusOK, response)
}

// 视频投稿接口处理函数
func (ctrl *VideoController) DouyinPublishActionHandler(c *gin.Context) {

	resp, err := ctrl.videoService.CreateVideo(c)
	// 调用服务层创建视频
	if err != nil {
		c.JSON(500, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ctrl *VideoController) DHandler(c *gin.Context) {
	response := map[string]interface{}{
		"status_code": 0,
		"status_msg":  "success",
		"next_time":   1693835887,
		"video_list": []map[string]interface{}{
			{
				"id": 39,
				"author": map[string]interface{}{
					"id":               0,
					"name":             "dad",
					"follow_count":     1,
					"follower_count":   1,
					"is_follow":        true,
					"avatar":           "http://localhost:8080/static/picture/1_1111_1693835886033.png",
					"background_image": "http://localhost:8080/static/picture/1_1111_1693835886033.png",
					"signature":        "dd",
					"total_favorited":  1,
					"work_count":       1,
					"favorite_count":   1,
				},
				"play_url":       "http://localhost:8080/static/video/1_1111_1693835886033.mp4",
				"cover_url":      "http://localhost:8080/static/picture/1_1111_1693835886033.png",
				"favorite_count": 0,
				"comment_count":  0,
				"title":          "4",
				"is_favorite":    true,
			},
			{
				"id": 38,
				"author": map[string]interface{}{
					"id":               0,
					"name":             "dadddd",
					"follow_count":     1,
					"follower_count":   1,
					"is_follow":        true,
					"avatar":           "http://localhost:8080/static/picture/1_1111_1693835886033.png",
					"background_image": "http://localhost:8080/static/picture/1_1111_1693835886033.png",
					"signature":        "dd",
					"total_favorited":  1,
					"work_count":       1,
					"favorite_count":   1,
				},
				"play_url":       "http://localhost:8080/static/video/1_1111_1693835886033.mp4",
				"cover_url":      "http://localhost:8080/static/picture/1_1111_1693835886033.png",
				"favorite_count": 0,
				"comment_count":  0,
				"title":          "44",
				"is_favorite":    true,
			},
		},
	}
	c.JSON(http.StatusOK, response)
}

// 查看用户发布过的视频/douyin/publish/list/
func (ctrl *VideoController) DouyinPublishList(c *gin.Context) {

	resp, err := ctrl.videoService.GetUserVideo(c)
	// 调用服务层创建视频
	if err != nil {
		c.JSON(500, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ctrl *VideoController) GetUserFavoriteHandler(c *gin.Context) {
	resp, err := ctrl.videoService.GetUserFavorite(c)
	// 调用服务层创建视频
	if err != nil {
		c.JSON(500, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
