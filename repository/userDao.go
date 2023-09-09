package repository

import (
	"fmt"
	"github.com/KumaJie/douyin/utils"
	"log"
)

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

/*id 是用户的唯一标识符，使用 bigint(20) 存储；
name 是用户名称，使用 varchar(255) 存储；
follow_count 是关注总数，使用 bigint(20) 存储；
follower_count 是粉丝总数，使用 bigint(20) 存储；
is_follow 表示是否已关注，使用 tinyint(1) 存储，0 表示未关注，1 表示已关注；
avatar 是用户头像，使用 varchar(255) 存储；
background_image 是用户个人页顶部大图，使用 varchar(255) 存储；
signature 是个人简介，使用 varchar(255) 存储；
total_favorited 是获赞数量，使用 bigint(20) 存储；
work_count 是作品数量，使用 bigint(20) 存储；
favorite_count 是点赞数量，使用 bigint(20) 存储。*/

var (
	userDao *UserDAO
)

func (User) TableName() string {
	return "user"
}

type UserDAO struct {
}

// 根据用户id查询用户信息
func (*UserDAO) GetUserById(id int64) (User, error) {
	var user User
	utils.DB.First(&user, id)

	if utils.DB.Error != nil {
		log.Println(utils.DB.Error)
		return user, utils.DB.Error
	}
	return user, nil
}

// 修改关注数量
func (*UserDAO) Modify(atype string, uid, toUid int64) error {
	user := User{}
	toUser := User{}

	result := utils.DB.First(&user, uid)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	result = utils.DB.First(&toUser, toUid)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if atype == "1" {
		user.FollowCount++
		toUser.FollowerCount++
	} else if atype == "2" {
		user.FollowCount--
		toUser.FollowerCount--
	}

	utils.DB.Save(&user)
	utils.DB.Save(&toUser)
	return nil

}
