package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *mySQLStore) ListCategoryStorage(c context.Context, paging common.Paging, options []string) []entities.CateGet {
	// declare
	var categories []entities.CateGet
	var query = s.aptx.GormDB
	// options
	for _, option := range options {
		query.Where(option)
	}
	// get category
	if err := query.
		Limit(paging.GetLimit()).
		Offset(paging.GetOffset()).
		Find(&categories).Error; err != nil {
		panic(err)
	}
	// return
	return categories
}
