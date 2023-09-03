package models

import "github.com/KumaJie/douyin/repository"

type CreateVideoRequest struct {
	Token string `json:"token" binding:"required"`
	Data  []byte `json:"data" binding:"required"`
	Title string `json:"title" binding:"required"`
}

// DouyinFeedRequest 是 Douyin Feed 请求结构体
type DouyinFeedRequest struct {
	LatestTime int64  `json:"latest_time"`
	Token      string `json:"token"`
}

// DouyinFeedResponse 是 Douyin Feed 响应结构体
type DouyinFeedResponse struct {
	StatusCode int32              `json:"status_code"`
	StatusMsg  string             `json:"status_msg"`
	VideoList  []repository.Video `json:"video_list"`
	NextTime   int64              `json:"next_time"`
}

// DouyinUserRequest 用户信息请求结构体定义
type DouyinUserRequest struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// DouyinUserResponse 用户信息响应结构体定义
type DouyinUserResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	User       repository.User `json:"user"`
}

type DouyinUserRegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}
