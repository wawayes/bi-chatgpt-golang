package models

import (
	"gorm.io/gorm"
)

// User 与数据库对齐的model
type User struct {
	Model

	UserAccount  string `json:"userAccount" gorm:"column:userAccount"`
	UserPassword string `json:"userPassword" gorm:"column:userPassword"`
	UserName     string `json:"userName" gorm:"column:userName"`
	UserAvatar   string `json:"userAvatar" gorm:"column:userAvatar"`
	UserRole     string `json:"userRole" gorm:"column:userRole"`
	FreeCount    int    `json:"freeCount" gorm:"column:freeCount"`
}

func (user *User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "user"
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&UserChart{}).Where("userId = ?", user.ID).Count(&count)
	if count == 0 {
		userChart := &UserChart{
			UserId:      user.ID,
			UserAccount: user.UserAccount,
		}
		err := tx.Select("userId", "userAccount").Create(&userChart).Error
		return err
	}
	return nil
}
