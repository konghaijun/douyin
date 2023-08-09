package main

import (
	"github.com/KumaJie/douyin/router"
	"github.com/KumaJie/douyin/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMysql()
	router.SetupRouter()

	/*	r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})*/

	/*err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}*/

	router := router.SetupRouter()

	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
