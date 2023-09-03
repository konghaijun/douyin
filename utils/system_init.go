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
	viper.SetConfigName("app")
	//添加配置文件的路径为 "config"
	viper.AddConfigPath("config")

	fmt.Println("当前 Viper 配置路径:", viper.ConfigFileUsed())

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
	//user:password@(localhost:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	url := viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@(" +
		viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + ")/" +
		viper.GetString("mysql.dbname") + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(url)
	DB, _ = gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	// 其他操作...
}
