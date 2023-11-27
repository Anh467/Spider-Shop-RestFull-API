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

func (s *mySQLStore) UpdateProductStorage(ctx context.Context, productUpdate entities.ProductUpdate, productid int) *entities.ProductGet {
	// delcare
	productGet := &entities.ProductGet{
		ProductID: productid,
		CateID:    productUpdate.CateID,
		Name:      productUpdate.Name,
		Desc:      productUpdate.Desc,
		Image:     productUpdate.Image,
		Status:    productUpdate.Name,
	}
	var count int64
	// check the exsistence of the productid
	if err := s.aptx.GormDB.Where("ProductID=?", productid).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
		})
	}
	// it will be checked if catid != 0 (==0 means that we won't update cateid)
	if productGet.CateID != 0 {
		// check the exsistence of the cateid
		if err := s.aptx.GormDB.Where("CateID=?", productGet.CateID).Count(&count).Error; err != nil {
			panic(err)
		}
		if count == 0 {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
			})
		}
		if count == 0 {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
			})
		}
	}
	// update the product
	if tx := s.aptx.GormDB.
		Where("ProductID=?", productid).
		Updates(&productGet).
		Preload(entities_category.CATE_TABLE, func(db *gorm.DB) *gorm.DB {
			return db.Select(entities_category.USER_TABLE_CateID,
				entities_category.USER_TABLE_Name,
				entities_category.USER_TABLE_Desc,
				entities_category.USER_TABLE_Status)
		}); tx != nil {
		// check there was an modifications
		if tx.RowsAffected == 0 {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: biz.PRODUCT_ERR_TABLE_NO_CHANGE,
			})
		}
		// check there was an error
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
	// return
	return productGet
}
