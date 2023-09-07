package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/utils/videoutil"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

type VideoService struct {
	videoDAO *repository.VideoDAO
}

// GetDouyinFeed 根据请求参数获取 Douyin Feed 数据
func (s *VideoService) GetDouyinFeed(c *gin.Context) (*models.DouyinFeedResponse, error) {

	//设置响应信息
	var response = &models.DouyinFeedResponse{}
	response.StatusCode = 1
	response.StatusMsg = "fail"

	latestTimeStr := c.Query("latest_time")
	var latestTime time.Time
	fmt.Println(latestTime)

	if latestTimeStr != "" {
		latestTimestamp, err := strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			c.String(400, "Invalid latest_time format")
			return response, err
		}

		// 将时间戳除以 1000，得到以秒为单位的时间戳
		latestTimestamp /= 1000

		fmt.Println(latestTimestamp)
		location := time.FixedZone("CST", 28800) // 创建一个代表东八区的时区对象

		latestTime = time.Unix(latestTimestamp, 0).In(location)

		fmt.Println(latestTime)

	} else {
		latestTime = time.Now()
	}

	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{},
	}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	favoriteService := &FavoriteService{
		favoriteDAO: &repository.FavoriteDao{}}

	//获取视频数组
	videos, err := videoService.videoDAO.GetVideoList(latestTime)
	if err != nil {
		log.Println(err)
		return response, err
	}

	//创建一个相同大小的数组
	var newVideo = make([]models.FeedVideo, len(videos))

	//遍历修改
	for i, video := range videos {
		fmt.Println(video)
		//根据uid查user对象
		u, err := userService.userDAO.GetUserById(video.UserInfoID)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return response, err
		}

		f, err := favoriteService.favoriteDAO.CheckFavorite(u.ID, video.ID)

		newVideo[i] = models.FeedVideo{
			ID:            video.ID,
			Author:        u,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			//待修改 方法1获取所有用户点赞过的视频再去比较
			IsFavorite: f,
		}

	}

	//设置为成功响应的信息
	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = newVideo
	if len(videos) > 0 {
		response.NextTime = videos[len(videos)-1].UploadTime.Unix()
	}

	return response, nil
}

func (s *VideoService) CreateVideo(c *gin.Context) (resp models.BaseResponse, err error) {
	var req models.DouyinPublishActionRequest

	// 解析请求主体内容
	if err := c.ShouldBind(&req); err != nil {
		// 请求主体内容解析失败
		fmt.Println(err)
		return resp, err
	}

	file, err := c.FormFile("data")
	if err != nil {
		return resp, err
	}

	//拿到用户id
	//	userId := c.MustGet("user_id").(int64)

	//生成唯一文件名字
	filename := videoutil.GenerateUniqueFilename(file.Filename, 1111)

	// 保存文件到本地
	savePath := "./static/video/" + filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return resp, err
	}

	playUrl := videoutil.GetVideoUrl(filename)

	//去掉视频的扩展名
	filename = videoutil.ExtractFilenameWithoutExtension(filename)

	//根据视频生成封面
	coverUrl, err := videoutil.GetVideoPicture(savePath, "./static/picture/"+filename, 1)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	coverUrl = videoutil.GetPictureUrl(filename + ".png")

	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{}}

	video := repository.Video{
		ID: 0, UserInfoID: 0, CoverURL: coverUrl, PlayURL: playUrl, FavoriteCount: 0, CommentCount: 0, Title: req.Title, UploadTime: time.Now()}
	//插入视频
	err = videoService.videoDAO.InsertVideo(video)
	if err != nil {
		log.Println(err)
		err := os.Remove(savePath)
		if err != nil {
			log.Println(err)
			return resp, err
		}
		return resp, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = "上传成功"
	// 返回响应
	return resp, nil
}

func (s *VideoService) GetUserVideo(c *gin.Context) (resp models.PublishResponse, err error) {

	user_id_str := c.Query("user_id")
	user_id, err := strconv.ParseInt(user_id_str, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return
		// 返回错误信息或采取其他操作
	}

	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{}}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	//根据用户id获取用户信息
	u, err := userService.userDAO.GetUserById(user_id)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return resp, err
	}

	//从数据库获取用户投稿视频
	videos, err := videoService.videoDAO.GetUserVideo(user_id)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	//创建一个相同大小的数组
	var newVideo = make([]models.FeedVideo, len(videos))

	//遍历修改
	for i, video := range videos {
		fmt.Println(video)
		//根据uid查user对象

		newVideo[i] = models.FeedVideo{
			ID:            video.ID,
			Author:        u,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			//待修改 方法1获取所有用户点赞过的视频再去比较
			IsFavorite: false,
		}

	}

	resp.VideoList = newVideo

	return resp, nil
}

func (s *VideoService) GetUserFavorite(c *gin.Context) (resp models.PublishResponse, err error) {
	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{}}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	favoriteService := &FavoriteService{
		favoriteDAO: &repository.FavoriteDao{}}

	user_id_str := c.Query("user_id")
	user_id, err := strconv.ParseInt(user_id_str, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return
		// 返回错误信息或采取其他操作
	}

	//拿到用户点赞的视频数组
	videos, err := favoriteService.favoriteDAO.GetUserFavorite(user_id)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	var newVideo = make([]models.FeedVideo, len(videos))

	for i, videoF := range videos {

		video, err := videoService.videoDAO.GetVideoById(videoF)
		u, err := userService.userDAO.GetUserById(user_id)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return resp, err
		}

		newVideo[i] = models.FeedVideo{
			ID:            video.ID,
			Author:        u,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			//待修改 方法1获取所有用户点赞过的视频再去比较
			IsFavorite: true,
		}

	}

	resp.VideoList = newVideo
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return resp, nil
}
