package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) DeleteProductStorage(ctx context.Context, productid int) *entities.ProductGet {
	// Declare
	var status string
	productGet := entities.ProductGet{}
	// get status of that category
	if err := s.aptx.GormDB.
		Table(entities.PRODUCT_TABLE).
		Where("ProductID = ?", productid).
		Pluck(entities.USER_TABLE_Status, &status).Error; err != nil {
		panic(err)
	}
	// if cateid is exist
	if status == "" {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRODUCT_ERR_PRODUCT_ID_NOT_EXIST,
		})
	}
	// check if status is Hot or Normal, turn it into Deleted (soft deletes)
	if status == entities.PRODUCT_TABLE_Status_Normal || status == entities.PRODUCT_TABLE_Status_HOT {
		// set model to Deleted
		productGet.Status = entities.PRODUCT_TABLE_Status_Deleted
		// get deleted information
		tx := s.aptx.GormDB.Where("ProductID = ?", productid).Updates(&productGet)
		// check there is errors
		if tx.Error != nil {
			panic(tx.Error)
		}
		if tx.RowsAffected == 0 {
			panic(biz.PRODUCT_ERR_TABLE_NO_CHANGE)
		}
		return &productGet
	}

	// check if is Deleted, delete it in table (hard deletes)
	if status == entities.PRODUCT_TABLE_Status_Deleted {
		// get deleted information
		tx := s.aptx.GormDB.Table(entities.PRODUCT_TABLE).Where("ProductID = ?", productid).Delete(nil)
		// check there is errors
		if tx.Error != nil {
			panic(tx.Error)
		}
		if tx.RowsAffected == 0 {
			panic(biz.PRODUCT_ERR_TABLE_NO_CHANGE)
		}
		return nil
	}
	return nil
}
