package storage

import (
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *mySQLStore) GetCategoryStorage(c context.Context, cateid int) entities.CateGet {
	// declare
	var cateGet *entities.CateGet
	// get first category
	if err := s.aptx.GormDB.
		Where("CateID = ?", cateid).
		First(&cateGet).Error; err != nil {
		panic(err)
	}
	// return
	return *cateGet
}
