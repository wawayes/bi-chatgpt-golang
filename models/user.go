package models

// User 与数据库对齐的model
type User struct {
	Model

	UserAccount  string `json:"userAccount" gorm:"column:userAccount"`
	UserPassword string `json:"userPassword" gorm:"column:userPassword"`
	UserName     string `json:"userName" gorm:"column:userName"`
	UserAvatar   string `json:"userAvatar" gorm:"column:userAvatar"`
	UserRole     string `json:"userRole" gorm:"column:userRole"`
	FreeCount    string `json:"freeCount" gorm:"column:freeCount"`
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "user"
}
