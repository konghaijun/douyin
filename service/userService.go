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

type UserService struct {
	userDAO *repository.UserDAO
}

func (s *UserService) GetUserById(c *gin.Context) (resp models.DouyinUserResponse, err error) {
	// 判断token
	// 根据id获取user

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	to_user_id_str := c.Query("user_id")

	toUserID, err := strconv.ParseInt(to_user_id_str, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return
		// 返回错误信息或采取其他操作
	}

	user, err := userService.userDAO.GetUserById(toUserID)
	if err != nil {
		log.Println(err)
		resp.StatusCode = -1
		resp.StatusMsg = "fail"
		return resp, err
	}

	resp.User = user
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return resp, nil
}

func (s *UserService) Register(username string, password string) (resp models.DouyinUserRegisterResponse, err error) {
	//查询用户名是否唯一
	//插入用户返回id
	//生成token
	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	f, err := userService.userDAO.IsUsernameUnique(username)
	if !f {
		resp.StatusMsg = "already"
		resp.StatusCode = 1
		return resp, nil
	}

	userid, err := userService.userDAO.CreateNewAccounts(username, password)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	token, err := utils.GetToken(userid)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	resp.UserID = userid
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.Token = token
	return resp, nil
}

func (s *UserService) Login(username, password string) (resp models.DouyinUserRegisterResponse, err error) {
	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	account, err := userService.userDAO.UserLogin(username, password)
	if err != nil {
		// 处理查询错误
		return resp, err
	}

	// 验证密码
	if account.Password == password {
		// 密码正确
		token, err := utils.GetToken(account.UserId)
		if err != nil {
			log.Println(err)
			return resp, err
		}

		resp.UserID = account.UserId
		resp.Token = token
		resp.StatusMsg = "success"
		resp.StatusCode = 0
	} else {
		// 密码错误
		resp.StatusMsg = "fail"
		resp.StatusCode = 1
	}
	return resp, nil
}
