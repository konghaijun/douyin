package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func (ctrl *UserController) DouyinUserHandler(c *gin.Context) {
	resp, err := ctrl.userService.GetUserById(c)
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
