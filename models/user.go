package models

// User 与数据库对齐的model
type User struct {
	Model

	UserAccount  string `json:"userAccount" gorm:"userAccount"`
	UserPassword string `json:"userPassword" gorm:"userPassword"`
	UserName     string `json:"userName" gorm:"userName"`
	UserAvatar   string `json:"userAvatar" gorm:"userAvatar"`
	UserRole     string `json:"UserRole" gorm:"userRole"`
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "user"
}
