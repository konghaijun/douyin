package repository

import (
	"github.com/KumaJie/douyin/utils"
	"sync"
	"time"
)

// Video 模型结构体
type Video struct {
	VideoID    int       `gorm:"primaryKey" json:"video_id"`
	UserID     int       `json:"user_id"`
	PlayURL    string    `json:"play_url"`
	CoverURL   string    `json:"cover_url"`
	Title      string    `json:"title"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

var (
	videoDao *VideoDAO
	postOnce sync.Once
)

func (Video) TableName() string {
	return "video"
}

type VideoDAO struct {
}

// 获取按投稿时间倒序的视频列表
func (*VideoDAO) GetVideoList() ([]Video, error) {
	var videos []Video
	result := utils.DB.Order("create_time desc").Limit(30).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func (*VideoDAO) InsertVideo(video Video) error {
	result := utils.DB.Create(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
