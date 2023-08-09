package service

import (
	"fmt"
	"github.com/KumaJie/douyin/models"
	"github.com/KumaJie/douyin/utils"
	"os"
	"testing"
)

func Test(t *testing.T) {

	utils.InitConfig()
	utils.InitMysql()

	fmt.Println("视频：")
	filePath := "..//upload/77cdbeed6b92c020a4ceb5c96c724e74.mp4"
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	v := models.CreateVideoRequest{Title: "to", Data: data, Token: "t"}

	videoService := &VideoService{} // 创建 VideoService 实例

	err = videoService.CreateVideo(v)
	if err != nil {
		fmt.Println(err)
	}
}
