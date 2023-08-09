package service

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("视频：")
	fmt.Println(viper.GetString("aliyun.keyid"))

	GetPlayInfo("299e4a50354171ee80d60764a3fd0102")

	//	videoService := &VideoService{}

}
