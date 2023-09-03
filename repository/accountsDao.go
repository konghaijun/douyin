package repository

import (
	"github.com/KumaJie/douyin/utils"
	"log"
)

type Accounts struct {
	ID       int    `json:"id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Accounts) TableName() string {
	return "accounts"
}

// 查询用户名是否唯一
func (*UserDAO) IsUsernameUnique(username string) (bool, error) {
	var count int64
	utils.DB.Model(&Accounts{}).Where("username = ?", username).Count(&count)
	if utils.DB.Error != nil {
		log.Println(utils.DB.Error)
		return false, utils.DB.Error
	}
	return count == 0, nil
}

// 插入用户
func (*UserDAO) CreateNewAccounts(username string, password string) (int64, error) {

	// 开启数据库事务
	tx := utils.DB.Begin()

	// 插入 user 记录
	user := User{
		ID: 0,
	}
	result := tx.Create(&user)
	if result.Error != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return 0, result.Error
	}

	// 插入 accounts 记录
	accounts := Accounts{
		UserId:   user.ID,
		Password: password,
		Username: username,
	}
	result = tx.Create(&accounts)
	if result.Error != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		tx.Delete(&user)
		return 0, result.Error
	}

	// 如果没有发生错误，则提交事务
	tx.Commit()

	// 返回成功结果
	return user.ID, nil

}

func (*UserDAO) UserLogin(username, password string) (Accounts, error) {
	var accounts Accounts

	err := utils.DB.Where("username = ?", username).First(&accounts).Error
	if err != nil {
		log.Println(err)
		return accounts, err
	}
	return accounts, nil
}
