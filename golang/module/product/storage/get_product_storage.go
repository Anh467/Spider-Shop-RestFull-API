package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
	"strings"
)

func (s *mySQLStore) GetProductStorage(c context.Context, flag bool, productid int) *entities.ProductGet {
	// declare
	var productget entities.ProductGet
	var count int64
	// find first product
	if err := s.aptx.GormDB.
		Preload("Cate").
		Where("productid = ?", productid).
		First(&productget).
		Count(&count).
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
	if !flag && (strings.EqualFold(productget.Status, entities.PRODUCT_TABLE_Status_Deleted) ||
		strings.EqualFold(productget.Cate.Status, entities.PRODUCT_TABLE_Status_Deleted)) {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: biz.PRODUCT_ERR_Temporarily_deleted,
		})
	}
	// return
	return &productget
}
