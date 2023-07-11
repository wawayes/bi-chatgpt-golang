package models

type Chart struct {
	Model
	Goal      string `json:"goal" gorm:"column:goal"`
	Name      string `json:"name" gorm:"column:name"`
	Data      string `json:"data" gorm:"column:chartData"`
	ChartType string `json:"chartType" gorm:"column:chartType"`
	GenChart  string `json:"genChart" gorm:"column:genChart"`
	GenResult string `json:"genResult" gorm:"column:genResult"`
	UserId    string `json:"userId" gorm:"column:userId"`
}

func (Chart) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "chart"
}
