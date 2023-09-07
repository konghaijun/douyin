package videoutil

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// GetVideoUrl 生成视频地址
func GetVideoUrl(videoName string) string {
	url := fmt.Sprintf("http://%s:%d/static/video/%s",
		viper.Get("oss.server.ip"), viper.Get("oss.server.port"), videoName)
	return url
}

// GetPictureUrl 生成图片地址
func GetPictureUrl(pictureName string) string {
	url := fmt.Sprintf("http://%s:%d/static/picture/%s",
		viper.Get("oss.server.ip"), viper.Get("oss.server.port"), pictureName)
	return url
}

// GenerateUniqueFilename 生成唯一文件名
func GenerateUniqueFilename(file string, userID int64) string {
	extension := filepath.Ext(file)
	filenameWithoutExtension := strings.TrimSuffix(file, extension)
	timestamp := time.Now().UnixNano() / int64(time.Millisecond) // 获取当前时间戳（毫秒级）
	filename := filenameWithoutExtension + "_" + strconv.FormatInt(userID, 10) + "_" + strconv.FormatInt(timestamp, 10) + extension
	return filename
}

// ExtractFilenameWithoutExtension 去掉扩展名
func ExtractFilenameWithoutExtension(file string) string {
	fmt.Println(file)
	extension := filepath.Ext(file)
	filename := strings.TrimSuffix(file, extension)
	fmt.Println(filename)
	return filename
}
