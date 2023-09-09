package repository

import (
	"errors"
	"fmt"
	"github.com/KumaJie/douyin/utils"
	"gorm.io/gorm"
)

type Follow struct {
	FollowId   int64 `json:"follow_id"`
	FromUserId int64 `json:"from_user_id"`
	ToUserId   int64 `json:"to_user_id"`
}

type FollowDao struct {
}

var (
	followDao *FollowDao
)

func (Follow) TableName() string {
	return "follow"
}

func (*FollowDao) AddFollow(follow Follow) error {
	result := utils.DB.Create(&follow)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func (*FollowDao) DelFollow(follow Follow) error {
	result := utils.DB.Delete(&follow, "from_user_id = ? AND to_user_id = ?", follow.FromUserId, follow.ToUserId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

// 判断用户之间是否有关注
func (*FollowDao) CheckFollow(uid, toUid int64) (bool, error) {
	var fo Follow
	result := utils.DB.Where("from_user_id = ? AND to_user_id = ?", uid, toUid).First(&fo)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			fmt.Println("fail", result.Error)
			return false, result.Error
		}
	} else {
		return true, nil
	}
}

func (*FollowDao) GetFollowList(uid int64) ([]int64, error) {
	var list []int64
	result := utils.DB.Where("from_user_id=?", uid).Table("follow").Select("to_user_id").Find(&list)
	if result.Error != nil {
		fmt.Println(result.Error)
		return list, result.Error
	}
	return list, nil
}
