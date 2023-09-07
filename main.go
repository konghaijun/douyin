package main

import (
	"github.com/KumaJie/douyin/router"
	"github.com/KumaJie/douyin/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMysql()
	appRouter := router.SetupRouter()

	appRouter.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
