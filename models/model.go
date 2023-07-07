package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/plugin/soft_delete"
	"time"
)

// Model 共有字段
type Model struct {
	ID          string                `json:"id" gorm:"primary_key"`                           // 主键ID
	CreateTime  LocalTime             `json:"createTime" gorm:"column:createTime"`             // 创建时间
	UpdatedTime LocalTime             `json:"updateTime" gorm:"column:updateTime"`             // 更新时间
	IsDelete    soft_delete.DeletedAt `json:"isDelete" gorm:"column:isDelete;softDelete:flag"` // 删除时间
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
