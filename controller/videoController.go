package controller

import (
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *VideoController) DouyinFeedHandler(c *gin.Context) {
	// 调用 service.GetDouyinFeed 获取 Douyin Feed 数据
	response, err := ctrl.videoService.GetDouyinFeed()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 设置 HTTP 响应的 Content-Type 为 application/json
	c.Header("Content-Type", "application/json")

	// 将服务层返回的数据作为 JSON 响应发送
	c.JSON(http.StatusOK, response)
}

type VideoController struct {
	videoService *service.VideoService
}

// 视频投稿接口处理函数
func (ctrl *VideoController) DouyinPublishActionHandler(c *gin.Context) {
	var createVideoRequest models.CreateVideoRequest
	// 解析请求参数
	if err := c.ShouldBindJSON(createVideoRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层创建视频
	if err := ctrl.videoService.CreateVideo(createVideoRequest); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video created successfully"})
}
