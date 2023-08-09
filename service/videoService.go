package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/repository"
	"log"
	"time"
)

type VideoService struct {
	videoDAO *repository.VideoDAO
}

// GetDouyinFeed 根据请求参数获取 Douyin Feed 数据
func (s *VideoService) GetDouyinFeed() (*models.DouyinFeedResponse, error) {

	var response = &models.DouyinFeedResponse{}

	videos, err := s.videoDAO.GetVideoList()
	if err != nil {
		log.Println(err)
		response.StatusCode = 1
		response.StatusMsg = "fail"
		return response, err
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = videos

	if len(videos) > 0 {
		response.NextTime = videos[len(videos)-1].CreateTime.Unix()
	}

	return response, nil
}

func (s *VideoService) CreateVideo(req models.CreateVideoRequest) error {
	saveVideoToFile(req.Data, req.Title)
	vid := saveVideoToAli(req.Title)
	time.Sleep(3 * time.Second)
	v, err := GetPlayInfo(vid)
	if err != nil {
		fmt.Println(err)
	}
	v1 := repository.Video{
		VideoID:    v.VideoID,
		UserID:     v.UserID,
		PlayURL:    v.PlayURL,
		CoverURL:   v.CoverURL,
		Title:      v.Title,
		CreateTime: v.CreateTime,
	}

	fmt.Println(v.PlayURL)
	err = s.videoDAO.InsertVideo(v1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
