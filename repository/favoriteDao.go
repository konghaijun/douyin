package repository

import (
	"errors"
	"fmt"
	"github.com/KumaJie/douyin/utils"
	"gorm.io/gorm"
	"log"
)

type Favorite struct {
	FavoriteId int64 `json:"favorite_id"`
	VideoId    int64 `json:"video_id"`
	UserId     int64 `json:"user_id"`
}

var (
	favoriteDao *FavoriteDao
)

type FavoriteDao struct {
}

func (Favorite) TableName() string {
	return "favorite"
}

// 用户点赞
func (*FavoriteDao) FavoriteActionAdd(uid int64, vid int64) error {

	favorite := Favorite{
		FavoriteId: 0,
		VideoId:    vid,
		UserId:     uid,
	}
	result := utils.DB.Create(&favorite)
	if result.Error != nil {
		// 发生错误时回滚事务
		fmt.Println(result.Error)
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

// FavoriteActionDel 取消点赞
func (*FavoriteDao) FavoriteActionDel(uid int64, vid int64) error {

	result := db.Delete(&Favorite{}, "user_id = ? AND video_id = ?", uid, vid)

	if result.Error != nil {
		fmt.Println(result.Error)
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

// 查询用户是否有点赞视频
func (*FavoriteDao) CheckFavorite(uid int64, vid int64) (bool, error) {
	var fa Favorite
	result := utils.DB.Where("user_id = ? AND video_id = ?", uid, vid).First(&fa)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 数据不存在
			return false, nil
		} else {

			// 查询出错
			fmt.Println("Failed to query from DB:", result.Error)
			return false, result.Error
		}
	} else {
		// 数据存在
		return true, nil
	}
}

// 查询用户所有点赞过的视频
func (*FavoriteDao) GetUserFavorite(uid int64) ([]int64, error) {
	var videos []int64
	result := utils.DB.Table("favorite").Select("video_id").Where("user_id = ?", uid).Find(&videos)
	if result.Error != nil {
		fmt.Println(result.Error)
		log.Println(result.Error)
		return videos, result.Error
	}
	return videos, nil
}
