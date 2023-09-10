package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FollowController struct {
	followController *service.FollowService
}

func (ctrl *FollowController) DouyinRelationActionHandler(c *gin.Context) {
	resp, err := ctrl.followController.RelationAction(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *FollowController) DouyinRelationFollowListHandler(c *gin.Context) {
	resp, err := ctrl.followController.GetFollowList(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *FollowController) DouyinRelationFollowerListHandler(c *gin.Context) {
	resp, err := ctrl.followController.GetFollowerList(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *FollowController) DouyinRelationFriendListHandler(c *gin.Context) {
	resp, err := ctrl.followController.GetFriendList(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}
