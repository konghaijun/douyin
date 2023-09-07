package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CommentController struct {
	commentService *service.CommentService
}

func (ctrl *CommentController) CommentActionHandler(c *gin.Context) {
	resp, err := ctrl.commentService.CommentAction(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *CommentController) CommentListHandler(c *gin.Context) {
	resp, err := ctrl.commentService.CommentList(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}
