package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() {

	//设置配置文件的名称为 "app"
	viper.SetConfigName("App")

	//添加配置文件的路径为 "config"
	viper.AddConfigPath("/config")

	//读取配置文件并将其加载到 viper 中。
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

var DB *gorm.DB

func InitMysql() {
	//自定义日志模板打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.url")), &gorm.Config{Logger: newLogger})
	// 其他操作...
}
