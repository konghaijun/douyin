package repository

import (
	"github.com/KumaJie/douyin/utils"
	"time"
)

type Comment struct {
	CommentId  int64     `json:"comment_id"`
	VideoId    int64     `json:"video_id"`
	UserId     int64     `json:"user_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

var (
	comment *Comment
)

type CommentDao struct {
}

func (Comment) TableName() string {
	return "comment"
}

// 添加评论
func (*CommentDao) AddComment(comment Comment) (int64, error) {
	result := utils.DB.Create(&comment)
	if result.Error != nil {
		// 发生错误时回滚事务
		return 0, result.Error
	}
	return comment.CommentId, nil
}

// 删除评论
func (*CommentDao) DelComment(cid int64) error {
	comment := Comment{CommentId: cid}
	result := utils.DB.Delete(&comment)
	if result.Error != nil {
		// 发生错误时回滚事务
		return result.Error
	}
	return nil
}

// 根据视频ID查评论
func (*CommentDao) ListComment(cid int64) ([]Comment, error) {
	var list []Comment
	result := utils.DB.Where("video_id = ?", cid).Order("create_time desc").Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
