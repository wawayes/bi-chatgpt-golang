package service

import (
	"fmt"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/logx"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

var data string

func File2Data(file multipart.File) (string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		logx.Warning(err.Error())
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, row := range rows {
		for _, colCell := range row {
			data += colCell + "\t"
		}
		data += "\n"
	}
	return data, nil
}
