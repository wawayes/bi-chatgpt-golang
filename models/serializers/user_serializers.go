package serializers

type CurrentUser struct {
	ID          int    `json:"id" gorm:"primary_key"` // 主键ID
	UserAccount string `json:"userAccount" gorm:"column:userAccount"`
	UserName    string `json:"userName" gorm:"column:userName"`
	UserAvatar  string `json:"userAvatar" gorm:"column:userAvatar"`
	UserRole    string `json:"userRole" gorm:"column:userRole"`
}
