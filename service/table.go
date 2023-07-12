package service

import (
	"errors"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/models"
)

type TableService struct{}

func (tableService TableService) ListTable(req *requests.Page) ([]models.UserChart, error) {
	pageNum := req.PageNum
	pageSize := req.PageSize
	var list []models.UserChart
	if err := models.BI_DB.Model(&list).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	return list, nil
}
