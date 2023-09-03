package controller

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func (ctrl *UserController) DouyinUserHandler(c *gin.Context) {
	var userRequest models.DouyinUserRequest

	user_id_str := c.Query("user_id")
	user_id, err := strconv.ParseInt(user_id_str, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return
		// 返回错误信息或采取其他操作
	}
	userRequest.UserID = user_id

	userRequest.Token = c.Query("token")

	resp, err := ctrl.userService.GetUserById(userRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "fail"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// /douyin/user/register/
func (ctrl *UserController) DouyinUserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	resp, err := ctrl.userService.Register(username, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *UserController) DouyinUserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	resp, err := ctrl.userService.Login(username, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *UserController) DdddHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"user": gin.H{
			"id":               8,
			"name":             "123456",
			"follow_count":     0,
			"follower_count":   0,
			"is_follow":        true,
			"avatar":           "string",
			"background_image": "string",
			"signature":        "string",
			"total_favorited":  1,
			"work_count":       0,
			"favorite_count":   0,
		},
	})
}
func (ctrl *UserController) DdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"user_id":     8,
		"token":       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM2NDY0NTIsImlhdCI6MTY5MzY0NjQ1MiwiaWQiOjh9.mvv8IRK_eLG8fGyPfMBmDAJCNBgWHG_9t-5WDHF_q5g",
	})
}
