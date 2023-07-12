package models

import "gorm.io/gorm"

type Chart struct {
	Model
	Goal      string `json:"goal" gorm:"column:goal"`
	Name      string `json:"name" gorm:"column:name"`
	Data      string `json:"data" gorm:"column:chartData"`
	ChartType string `json:"chartType" gorm:"column:chartType"`
	Token     int    `json:"token" gorm:"column:token"`
	GenChart  string `json:"genChart" gorm:"column:genChart"`
	GenResult string `json:"genResult" gorm:"column:genResult"`
	UserId    int    `json:"userId" gorm:"column:userId"`
}

func (chart *Chart) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "chart"
}

//func (chart *Chart) AfterCreate(tx *gorm.DB) (err error) {
//	var userChart UserChart
//	tx.Model(&UserChart{}).Where("userId = ?", chart.UserId).First(&userChart)
//
//	var user User
//	tx.Model(&User{}).Where("userId = ?", chart.UserId).First(&user).Update("freeCount", user.FreeCount-1)
//	freeCount := user.FreeCount
//
//	finalToken := chart.Token + userChart.Token
//	tx.Model(&UserChart{}).Where("userId = ?", chart.UserId).UpdateColumn("token", finalToken).UpdateColumn("freeCount", freeCount)
//	return nil
//}

func (chart *Chart) AfterCreate(tx *gorm.DB) (err error) {
	var userChart UserChart
	tx.Model(&UserChart{}).Where("userId = ?", chart.UserId).First(&userChart)

	var user User
	tx.Preload("UserChart").Where("userId = ?", chart.UserId).First(&user)

	finalToken := chart.Token + userChart.Token
	tx.Model(&UserChart{}).Where("userId = ?", chart.UserId).Updates(UserChart{Token: finalToken})

	tx.Model(&User{}).Where("userId = ?", chart.UserId).Updates(User{FreeCount: user.FreeCount - 1})

	tx.Model(&UserChart{}).Where("userId = ?", chart.UserId).UpdateColumn("freeCount", gorm.Expr("freeCount - ?", 1))

	return nil
}
