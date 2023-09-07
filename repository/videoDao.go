package repository

import (
	"fmt"
	"github.com/KumaJie/douyin/utils"
	"sync"
	"time"
)

// Video 模型结构体
type Video struct {
	ID            int64     `gorm:"column:id"`
	UserInfoID    int64     `gorm:"column:user_info_id"`
	PlayURL       string    `gorm:"column:play_url"`
	CoverURL      string    `gorm:"column:cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	Title         string    `gorm:"column:title"`
	UploadTime    time.Time `gorm:"column:upload_time"`
}

var (
	videoDao *VideoDAO
	postOnce sync.Once
)

func (Video) TableName() string {
	return "videos"
}

type VideoDAO struct {
}

// 获取按投稿时间倒序的视频列表
func (*VideoDAO) GetVideoList(time time.Time) ([]Video, error) {
	var videos []Video
	result := utils.DB.Where("upload_time < ?", time).Order("upload_time desc").Limit(2).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

// 上传视频
func (*VideoDAO) InsertVideo(video Video) error {
	result := utils.DB.Create(&video)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

// 根据id获取用户投稿视频
func (*VideoDAO) GetUserVideo(uid int64) ([]Video, error) {
	var videos []Video
	result := utils.DB.Where("user_info_id", uid).Find(&videos)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return videos, nil
}

// 点赞数量或者取消点赞
func (*VideoDAO) AddFavorite(vid int64, atype string) error {
	var video Video
	result := utils.DB.Where("id = ?", vid).First(&video)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
		return result.Error
	}

	if atype == "1" {
		video.FavoriteCount++
	} else if atype == "2" {
		video.FavoriteCount--
	}
	utils.DB.Save(&video)
	return nil
}

// 根据视频id获取视频
func (*VideoDAO) GetVideoById(vid int64) (Video, error) {
	var video Video
	result := utils.DB.Where("id = ?", vid).First(&video)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
		return video, result.Error
	}
	return video, nil
}

// 评论或者删除评论
func (*VideoDAO) AddComment(vid int64, atype string) error {
	var video Video
	result := utils.DB.Where("id = ?", vid).First(&video)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
		return result.Error
	}

	if atype == "1" {
		video.CommentCount++
	} else if atype == "2" {
		video.CommentCount--
	}
	utils.DB.Save(&video)
	return nil
}
