package storage

import (
	"SpiderShop-Restfull-API/common"
	entities_category "SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
)

func (s *mySQLStore) ListProductStorage(c context.Context, flag bool, paging common.Paging, options []string) *[]entities.ProductGet {
	// declare product
	var cateidsDeleted []string
	var productGets *[]entities.ProductGet
	var query = s.aptx.GormDB
	// get list deleted cateidsDeleted
	if !flag {
		if err := s.aptx.GormDB.
			Table(entities_category.CATE_TABLE).
			Where("Status = ?", entities_category.USER_TABLE_Status_Deleted).
			Pluck("CateID", &cateidsDeleted).
			Error; err != nil {
			panic(err)
		}
	}
	// init options
	for _, option := range options {
		query = query.Where(option)
	}
	// add conditions user have a rights to access deleted products or not
	if !flag {
		query = query.
			Not("CateID in ?", cateidsDeleted).
			Not("Status = ?", entities.PRODUCT_TABLE_Status_Deleted)
	}
	// get products
	if err := query.Preload("Cate").Find(&productGets).Error; err != nil {
		panic(err)
	}
	// return
	return productGets
}
