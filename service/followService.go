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

type FollowService struct {
	followDao *repository.FollowDao
}

// /douyin/relation/action/
func (s *FollowService) RelationAction(c *gin.Context) (models.BaseResponse, error) {
	var resp models.BaseResponse
	resp.StatusCode = 1
	resp.StatusMsg = "fail"

	var userId int64
	var ok bool
	userIdStr, exists := c.Get("user_id")
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

	followService := &FollowService{
		followDao: &repository.FollowDao{},
	}

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	toUserIdStr := c.Query("to_user_id")

	atype := c.Query("action_type")

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		return resp, err
		// 返回错误信息或采取其他操作
	}

	follow := repository.Follow{FromUserId: userId,
		ToUserId: toUserId}

	tx := utils.DB.Begin()

	if atype == "1" {
		err = followService.followDao.AddFollow(follow)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}
	} else if atype == "2" {
		err = followService.followDao.DelFollow(follow)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return resp, err
		}
	}

	err = userService.userDAO.Modify(atype, userId, toUserId)
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

// 获取关注
func (s *FollowService) GetFollowList(c *gin.Context) (models.FollowListResponse, error) {
	var resp models.FollowListResponse
	resp.StatusCode = 1
	resp.StatusMsg = "fail"

	userIdStr := c.Query("user_id")

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	followService := &FollowService{
		followDao: &repository.FollowDao{},
	}

	UserID, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	userList, err := followService.followDao.GetFollowList(UserID)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	var newUserList []repository.User

	for _, user := range userList {
		u, err := userService.userDAO.GetUserById(user)
		if err != nil {
			// 处理转换失败的情况
			fmt.Println("Failed to convert user_id to int64:", err)
			log.Println(err)
			return resp, err
		}

		u.IsFollow = true
		newUserList = append(newUserList, u)

	}

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.UserList = newUserList
	return resp, nil
}

// 获取粉丝列表
func (s *FollowService) GetFollowerList(c *gin.Context) (models.FollowListResponse, error) {
	var resp models.FollowListResponse
	resp.StatusCode = 1
	resp.StatusMsg = "fail"

	userIdStr := c.Query("user_id")

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	followService := &FollowService{
		followDao: &repository.FollowDao{},
	}

	UserID, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	userList, err := followService.followDao.GetFollowerList(UserID)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	var newUserList []repository.User

	for _, user := range userList {
		u, err := userService.userDAO.GetUserById(user)
		if err != nil {
			// 处理转换失败的情况
			fmt.Println("Failed to convert user_id to int64:", err)
			log.Println(err)
			return resp, err
		}

		u.IsFollow = true
		newUserList = append(newUserList, u)

	}

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.UserList = newUserList
	return resp, nil
}

// 获取好友
func (s *FollowService) GetFriendList(c *gin.Context) (models.FollowListResponse, error) {
	var resp models.FollowListResponse
	resp.StatusCode = 1
	resp.StatusMsg = "fail"

	userIdStr := c.Query("user_id")

	userService := &UserService{
		userDAO: &repository.UserDAO{},
	}

	followService := &FollowService{
		followDao: &repository.FollowDao{},
	}

	UserID, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	//取到我的关注列表
	userList, err := followService.followDao.GetFollowList(UserID)
	if err != nil {
		// 处理转换失败的情况
		fmt.Println("Failed to convert user_id to int64:", err)
		log.Println(err)
		// 返回错误信息或采取其他操作
	}

	var newUserList []repository.User

	for _, user := range userList {

		//判断对方是否关注我
		f, err := followService.followDao.CheckFollow(user, UserID)
		if err != nil {
			// 处理转换失败的情况
			fmt.Println("Failed to convert user_id to int64:", err)
			log.Println(err)
			return resp, err
		}

		if f {
			u, err := userService.userDAO.GetUserById(user)
			if err != nil {
				// 处理转换失败的情况
				fmt.Println("Failed to convert user_id to int64:", err)
				log.Println(err)
				return resp, err
			}

			u.IsFollow = true
			newUserList = append(newUserList, u)
		}

	}

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.UserList = newUserList
	return resp, nil
}
