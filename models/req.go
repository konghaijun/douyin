package models

import (
	"github.com/KumaJie/douyin/repository"
	"mime/multipart"
)

type BaseResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// /douyin/publish/action/
type DouyinPublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data" binding:"required"`
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
}

// DouyinFeedResponse 是 Douyin Feed 响应结构体
type DouyinFeedResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []FeedVideo `json:"video_list"`
	NextTime   int64       `json:"next_time"`
}

// PublishResponse 是 Douyin Feed 响应结构体
type PublishResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []FeedVideo `json:"video_list"`
}

// Video 模型结构体
type FeedVideo struct {
	ID            int64           `json:"id"`
	Author        repository.User `json:"author"`
	PlayURL       string          `json:"play_url"`
	CoverURL      string          `json:"cover_url"`
	FavoriteCount int64           `json:"favorite_count"`
	CommentCount  int64           `json:"comment_count"`
	Title         string          `json:"title"`
	IsFavorite    bool            `json:"is_favorite"`
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

// 评论相关响应
type CommentResponse struct {
	StatusCode int            `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
	Comment    CommentContent `json:"comment"`
}

type CommentContent struct {
	ID         int64           `json:"id"`
	User       repository.User `json:"user"`
	Content    string          `json:"content"`
	CreateDate string          `json:"create_date"`
}

// 评论相关响应
type CommentListResponse struct {
	StatusCode int              `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	Comment    []CommentContent `json:"comment_list"`
}

type FollowListResponse struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	UserList   []repository.User `json:"user_list"`
}
