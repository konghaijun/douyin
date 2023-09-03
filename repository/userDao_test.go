package repository

import (
	"github.com/KumaJie/douyin/utils"
	"testing"
)

func TestUser(t *testing.T) {
	utils.InitConfig()
	utils.InitMysql()
	videoService := &UserDAO{} // 创建 VideoService 实例
	videoService.GetUserById(1)

}
