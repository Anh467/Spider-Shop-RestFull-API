package storage

import (
	"SpiderShop-Restfull-API/common"
	entities_category "SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) UpdateProductStorage(ctx context.Context, productUpdate entities.ProductUpdate, productid int) *entities.ProductGet {
	// delcare
	var productGet *entities.ProductGet
	var count int64
	// check the exsistence of the productid

	if err := s.aptx.GormDB.
		Table(entities.PRODUCT_TABLE).
		Where("ProductID = ?", productid).
		Select("ProductID").
		Scan(&count).
		Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
		})
	}

	// it will be checked if catid != 0 (==0 means that we won't update cateid)
	if productUpdate.CateID != 0 {
		// check the exsistence of the cateid
		if err := s.aptx.GormDB.Table(entities_category.CATE_TABLE).Where("CateID=?", productUpdate.CateID).Count(&count).Error; err != nil {
			panic(err)
		}
		if count == 0 {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz.PRODUCT_ERR_CateID_NOT_EXIST,
			})
		}
	}
	// update the product
	if tx := s.aptx.GormDB.
		Where("ProductID=?", productid).
		Updates(&productUpdate); tx != nil {
		// check there was an error
		if tx.Error != nil {
			panic(tx.Error)
		}
		// check there was an modifications
		if tx.RowsAffected == 0 {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: biz.PRODUCT_ERR_TABLE_NO_CHANGE,
			})
		}
	}
	// getting the product
	if err := s.aptx.GormDB.
		Where("ProductID=?", productid).
		Preload(entities_category.CATE_TABLE).
		First(&productGet).
		Error; err != nil {
		panic(err)
	}
	// return
	return productGet
}
