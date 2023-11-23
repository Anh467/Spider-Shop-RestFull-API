package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/biz"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) DeleteCategoryStorage(c context.Context, cateid int) *entities.CateGet {
	// Declare
	var status string
	var cateGet *entities.CateGet
	// get status of that category
	if err := s.aptx.GormDB.
		Table(entities.CATE_TABLE).
		Where("CateID = ?", cateid).
		Pluck(entities.USER_TABLE_Status, &status).Error; err != nil {
		panic(err)
	}
	// if cateid is exist
	if status == "" {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.CATE_ERR_CateID_UNIQUE,
		})
	}
	// check if is Hot or Normal, turn it into Deleted (soft deletes)
	if status == entities.USER_TABLE_Status_Normal || status == entities.USER_TABLE_Status_HOT {
		// set model to Deleted
		cateGet.Status = entities.USER_TABLE_Status_Deleted
		// get deleted information
		tx := s.aptx.GormDB.Where("CateID = ?", cateid).Updates(&cateGet)
		// check there is errors
		if tx.Error != nil {
			panic(tx.Error)
		}
		if tx.RowsAffected == 0 {
			panic(biz.CATE_ERR_TABLE_NO_CHANGE)
		}
		return cateGet
	}

	// check if is Deleted, delete it in table (hard deletes)
	if status == entities.USER_TABLE_Status_Deleted {
		// get deleted information
		tx := s.aptx.GormDB.Table(entities.CATE_TABLE).Where("CateID = ?", cateid).Delete(nil)
		// check there is errors
		if tx.Error != nil {
			panic(tx.Error)
		}
		if tx.RowsAffected == 0 {
			panic(biz.CATE_ERR_TABLE_NO_CHANGE)
		}
		return nil
	}
	return nil
}
