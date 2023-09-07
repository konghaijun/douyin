package controller

import (
	"github.com/KumaJie/douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FavoriteController struct {
	favoriteService *service.FavoriteService
}

func (ctrl *FavoriteController) DouyinFavoriteAction(c *gin.Context) {
	resp, err := ctrl.favoriteService.FavoriteAction(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
	}
	c.JSON(http.StatusOK, resp)
}
