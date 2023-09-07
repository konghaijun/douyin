package videoutil

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

// GetVideoPicture 调用ffmpeg为视频截取封面
func GetVideoPicture(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	// 创建一个缓冲区来接收ffmpeg的输出
	buf := bytes.NewBuffer(nil)

	// 调用ffmpeg命令行工具截取视频封面，将输出写入缓冲区
	err = ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()

	/*ffmpeg_go.Input(videoPath)：指定要处理的视频文件路径作为输入。

	.Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)})：添加筛选器，通过设置选择条件来选择视频帧。
	"select" 是 ffmpeg 的筛选器名称，"gte(n,frameNum)" 表示选择帧号大于等于 frameNum 的视频帧。

	.Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"})：指定输出参数，将视频帧转换为图像。
	"pipe:" 表示将输出结果传递到管道，"vframes": 1 表示只输出一帧图像，"format": "image2" 表示输出图像格式为 image2，"vcodec": "mjpeg" 表示图像编解码器为 MJPEG。

	.WithOutput(buf, os.Stdout).Run()：将输出结果存储在缓冲区 buf 中，并运行命令。*/

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	// 从缓冲区中解码封面图像
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	// 将图像保存到指定路径
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	// 获取封面图像的文件名
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"

	return
}
