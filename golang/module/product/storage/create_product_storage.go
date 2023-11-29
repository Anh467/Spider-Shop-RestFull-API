package storage

import (
	"SpiderShop-Restfull-API/common"
	entities_category "SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) CreateProductStorage(ctx context.Context, productCreate entities.ProductCreate) *entities.ProductGet {
	// declare
	var count int64
	var productGet *entities.ProductGet
	// check the existence of cateid
	if err := s.aptx.GormDB.
		Table(entities_category.CATE_TABLE).
		Where("CateID = ?", productCreate.CateID).
		Select("CateID").
		Scan(&count).
		Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRODUCT_ERR_CateID_NOT_EXIST,
		})
	}
	// create product
	if err := s.aptx.GormDB.
		Create(&productCreate).
		Error; err != nil {
		panic(err)
	}
	// get product
	if err := s.aptx.GormDB.
		Select(entities.USER_TABLE_ProductID, entities.USER_TABLE_CateID, entities.USER_TABLE_Name,
			entities.USER_TABLE_Desc, entities.USER_TABLE_Image, entities.USER_TABLE_Status,
			entities.USER_TABLE_CreatedAt, entities.USER_TABLE_UpdatedAt).
		Where("ProductID = ?", productCreate.ProductID).
		Preload(entities_category.CATE_TABLE).
		First(&productGet).Error; err != nil {
		panic(err)
	}
	// return
	return productGet
}
