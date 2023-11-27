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

func (s *mySQLStore) GetProductStorage(c context.Context, flag bool, productid int) *entities.ProductGet {
	// declare
	var productget entities.ProductGet
	var count int64
	// find first product
	if err := s.aptx.GormDB.Where("ProductID = ?", productid).
		Preload(entities_category.CATE_TABLE, func(db *gorm.DB) *gorm.DB {
			return db.Select(entities_category.USER_TABLE_CateID,
				entities_category.USER_TABLE_Name,
				entities_category.USER_TABLE_Desc,
				entities_category.USER_TABLE_Status)
		}).
		Count(&count).
		First(&productget).
		Error; err != nil {
		panic(err)
	}
	// check product is exist or not
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
		})
	}
	// check status of product and category
	if !flag && (productget.Status == entities.PRODUCT_TABLE_Status_Deleted || productget.Cate.Status == entities.PRODUCT_TABLE_Status_Deleted) {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: biz.PRODUCT_ERR_Temporarily_deleted,
		})
	}
	// return
	return &productget
}
