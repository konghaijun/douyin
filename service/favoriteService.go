package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type FavoriteService struct {
	favoriteDAO *repository.FavoriteDao
}

func (s *FavoriteService) FavoriteAction(c *gin.Context) (models.BaseResponse, error) {
	var resp models.BaseResponse
	var userId int64
	var err error

	resp.StatusCode = 1
	resp.StatusMsg = "fail"

	userIdStr, exists := c.Get("user_id")
	if exists {
		if userId, ok := userIdStr.(int64); ok {
			// userIdStr 是 int64 类型
			fmt.Println("User ID:", userId)
		} else {
			// userIdStr 不是 int64 类型
			fmt.Println("User ID is not an int64")
		}
	} else {
		fmt.Println("User ID not found")
	}

	videoIdStr := c.Query("video_id")
	atype := c.Query("action_type")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return resp, err
		// 返回错误信息或采取其他操作
	}

	favoriteService := &FavoriteService{
		favoriteDAO: &repository.FavoriteDao{}}

	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{},
	}

	// 开启数据库事务
	tx := utils.DB.Begin()
	if atype == "1" {
		err = favoriteService.favoriteDAO.FavoriteActionAdd(userId, videoId)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}
	} else if atype == "2" {
		err = favoriteService.favoriteDAO.FavoriteActionDel(userId, videoId)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}
	}

	err = videoService.videoDAO.AddFavorite(videoId, atype)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return resp, err
	}

	tx.Commit()

	resp.StatusCode = 0
	resp.StatusMsg = "success"

	return resp, nil
}
