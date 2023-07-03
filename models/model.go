package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Model 共有字段
type Model struct {
	ID          int64                 `gorm:"primary_key"` // 主键ID
	CreateTime  time.Time             // 创建时间
	UpdatedTime time.Time             // 更新时间
	IsDelete    soft_delete.DeletedAt `gorm:"column:isDelete;softDelete:flag" json:"-"` // 删除时间
}
