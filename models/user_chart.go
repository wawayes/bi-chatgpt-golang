package models

type UserChart struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      int    `json:"userId" gorm:"column:userId"`
	UserAccount string `json:"userAccount" gorm:"column:userAccount"`
	UserAvatar  string `json:"userAvatar" gorm:"column:userAvatar"`
	Token       int    `json:"token" gorm:"column:token"`
	FreeCount   int    `json:"freeCount" gorm:"column:freeCount"`
}

func (userChart *UserChart) TableName() string {
	return "user_chart"
}
