package storage

import (
	"SpiderShop-Restfull-API/common"
	entities_category "SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (s *mySQLStore) CreateProductStorage(ctx context.Context, productCreate entities.ProductCreate) *entities.ProductGet {
	// declare
	var count int64
	productGet := &entities.ProductGet{
		CateID: productCreate.CateID,
		Name:   productCreate.Name,
		Desc:   productCreate.Desc,
		Image:  productCreate.Image,
		Status: productCreate.Status,
	}
	// check the existence of cateid
	if err := s.aptx.GormDB.
		Table(entities_category.CATE_TABLE).
		Where("CateID = ?", productGet.CateID).
		First(nil).
		Count(&count).
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
		Create(&productGet).
		Preload(entities_category.CATE_TABLE, func(db *gorm.DB) *gorm.DB {
			return db.Select(entities_category.USER_TABLE_CateID,
				entities_category.USER_TABLE_Name,
				entities_category.USER_TABLE_Desc,
				entities_category.USER_TABLE_Status)
		}).Error; err != nil {
		panic(err)
	}
	// return
	return productGet
}
