package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/KumaJie/douyin/repository"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vod20170321 "github.com/alibabacloud-go/vod-20170321/v3/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/spf13/viper"

	"os"
)

var accessKeyId string = viper.GetString("aliyun.vod.file.keyid")

var accessKeySecret string = viper.GetString("aliyun.vod.file.keysecret")

// 初始化VOD客户端并获取上传地址和凭证
func InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {
	// 点播服务接入区域
	regionId := "cn-shanghai"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

// 方法用于创建视频上传请求。它接受一个 vod 客户端和视频标题作为参数，然后返回一个 vod.CreateUploadVideoResponse 对象和一个错误。
func MyCreateUploadVideo(client *vod.Client, title string) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = title
	//request.Description = "Sample Description"
	request.FileName = title + ".mp4"
	//request.CateId = "-1"
	//Cover URL示例：http://example.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-****.png
	//request.CoverURL = "<your CoverURL>"
	//	request.Tags = "tag1,tag2"
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

// 方法用于初始化阿里云对象存储（OSS）客户端。它接受一个包含上传权限信息的 UploadAuthDTO 对象和包含上传地址信息的 UploadAddressDTO 对象作为参数，然后返回一个 OSS 客户端和一个错误。
func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

// 方法用于将本地文件上传到 OSS
func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// 方法用于创建 vod 客户端。它接受访问密钥 ID 和访问密钥 Secret 作为参数，然后返回一个配置好的 vod 客户端和一个错误。
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *vod20170321.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/vod
	config.Endpoint = tea.String("vod.cn-shanghai.aliyuncs.com")
	_result = &vod20170321.Client{}
	_result, _err = vod20170321.NewClient(config)
	return _result, _err
}

// 方法用于根据视频 ID 获取视频播放信息。
// 它接受视频 ID（字符串类型）作为参数，然后使用 vod 客户端发送请求来获取视频的播放地址、封面地址、标题和创建时间等信息。
// 最后，它将获取到的视频信息填充到 repository.Video 结构体中并返回。
func GetPlayInfo(videoID string) (repository.Video, error) {
	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		return repository.Video{}, fmt.Errorf("failed to create client: %w", err)
	}

	fmt.Println(videoID)
	id := videoID
	getPlayInfoRequest := &vod20170321.GetPlayInfoRequest{
		VideoId: tea.String(id),
	}

	var v repository.Video
	err = func() error {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				err = fmt.Errorf("panic occurred: %v", r)
			}
		}()

		info, err := client.GetPlayInfo(getPlayInfoRequest)
		if err != nil {
			return fmt.Errorf("failed to get play info: %w", err)
		}
		fmt.Println(info)

		title := *info.Body.VideoBase.Title
		playURL := *info.Body.PlayInfoList.PlayInfo[0].PlayURL

		coverURL := *info.Body.VideoBase.CoverURL

		creationTimeString := *info.Body.VideoBase.CreationTime
		//creationTime, err := time.Parse(time.RFC3339, creationTimeString)
		fmt.Println(playURL, coverURL, title, creationTimeString)
		if err != nil {
			// 处理解析错误
			return fmt.Errorf("failed to parse creation time: %w", err)
		}

		//v = repository.Video{VideoID: -1, UserID: -1, PlayURL: playURL, CoverURL: coverURL, Title: title, CreateTime: creationTime}
		return nil
	}()

	if err != nil {
		return repository.Video{}, fmt.Errorf("API error: %w", err)
	}

	return v, nil
}

// 视频保存到本地目录
func saveVideoToFile(data []byte, title string) error {
	// 指定保存视频的文件路径
	filePath := "..//Upload/" + title + ".mp4"

	// 打开文件，创建或截断现有文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入数据到文件
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("Video saved successfully")
	return nil
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}

func saveVideoToAli(title string) string {
	// 您的AccessKeySecret
	var localFile string = "../Upload/" + title + ".mp4" // 需要上传到VOD的本地视频文件的完整路径

	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return "1"
	}
	// 获取上传地址和凭证
	var response, createUploadVideoErr = MyCreateUploadVideo(vodClient, title)
	if createUploadVideoErr != nil {
		fmt.Println("Error:", createUploadVideoErr)
		return "1"
	}

	// 执行成功会返回VideoId、UploadAddress和UploadAuth
	var videoId = response.VideoId
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
	var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)

	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	var ossClient, _ = InitOssClient(uploadAuthDTO, uploadAddressDTO)
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	UploadLocalFile(ossClient, uploadAddressDTO, localFile)
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	return videoId

}
