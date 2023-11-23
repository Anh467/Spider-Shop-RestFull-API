package storage

import (
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *mySQLStore) CreateCategoryStorage(c context.Context, category entities.CateCreate) entities.CateGet {
	// declare
	cateGet := &entities.CateGet{
		Name:  category.Name,
		Desc:  category.Desc,
		Image: category.Image,
	}
	// get
	if err := s.aptx.GormDB.Create(&cateGet).Error; err != nil {
		panic(err)
	}
	// return
	return *cateGet
}
