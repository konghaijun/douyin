package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/repository"
	"github.com/KumaJie/douyin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type CommentService struct {
	commentDao *repository.CommentDao
}

// /douyin/comment/action/
func (s *CommentService) CommentAction(c *gin.Context) (resp models.CommentResponse, err error) {
	userIdStr, exists := c.Get("user_id")
	var userId int64
	var ok bool
	resp.StatusCode = 1
	resp.StatusMsg = "fail"
	if exists {

		if userId, ok = userIdStr.(int64); ok {
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

	commentService := CommentService{
		commentDao: &repository.CommentDao{},
	}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	videoService := &VideoService{
		videoDAO: &repository.VideoDAO{}}

	tx := utils.DB.Begin()

	//1发布评论 2删除评论
	if atype == "1" {
		content := c.Query("comment_text")
		fmt.Println(content)

		timeNow := time.Now()
		commment := repository.Comment{
			CommentId:  0,
			CreateTime: timeNow,
			VideoId:    videoId,
			UserId:     userId,
			Content:    content,
		}

		cid, err := commentService.commentDao.AddComment(commment)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}

		user, err := userService.userDAO.GetUserById(userId)
		if err != nil {
			fmt.Println(err)
			return resp, err
		}

		timeStr := timeNow.Format("01-02")

		resp.Comment.Content = content
		resp.Comment.ID = cid
		resp.Comment.User = user
		resp.Comment.CreateDate = timeStr

	} else if atype == "2" {
		cidStr := c.Query("comment_id")

		cid, err := strconv.ParseInt(cidStr, 10, 64)
		if err != nil {
			// 处理转换失败的情况
			fmt.Println("Failed to convert user_id to int64:", err)
			log.Println(err)
			return resp, err
			// 返回错误信息或采取其他操作
		}

		err = commentService.commentDao.DelComment(cid)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}

	}

	err = videoService.videoDAO.AddComment(videoId, atype)
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

func (s *CommentService) CommentList(c *gin.Context) (resp models.CommentListResponse, err error) {

	resp.StatusCode = 1

	videoIdStr := c.Query("video_id")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return resp, err
		// 返回错误信息或采取其他操作
	}
	commentService := CommentService{
		commentDao: &repository.CommentDao{},
	}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	commentList, err := commentService.commentDao.ListComment(videoId)

	var newVideo = make([]models.CommentContent, len(commentList))

	//遍历修改
	for i, comment := range commentList {
		u, err := userService.userDAO.GetUserById(comment.UserId)
		if err != nil {
			fmt.Println(err)
			return resp, err
		}

		newVideo[i] = models.CommentContent{
			ID:         comment.CommentId,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("01-02"),
			User:       u,
		}

	}
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.Comment = newVideo

	return resp, nil
}
