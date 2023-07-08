package models

type GenChartRequest struct {
	Model
	Goal      string `json:"goal" gorm:"column:goal"`
	Name      string `json:"name" gorm:"column:name"`
	Data      string `json:"data" gorm:"column:chartData"`
	ChartType string `json:"chartType" gorm:"column:chartType"`
	GenChart  string `json:"genChart" gorm:"column:genChart"`
	GenResult string `json:"genResult" gorm:"column:genResult"`
	UserId    string `json:"userId" gorm:"column:userId"`
}
