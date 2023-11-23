package storage

import (
	"SpiderShop-Restfull-API/module/category/biz"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *mySQLStore) UpdateCategoryStorage(c context.Context, cateUpdate entities.CateUpdate, cateid int) *entities.CateGet {
	// declare
	cateGet := entities.CateGet{
		CateID: cateid,
		Name:   cateUpdate.Name,
		Desc:   cateUpdate.Desc,
		Image:  cateUpdate.Image,
		Status: cateUpdate.Status,
	}
	// update
	tx := s.aptx.GormDB.Updates(&cateGet)
	// check error
	if tx.Error != nil {
		panic(tx.Error)
	}
	// check its modify or not
	if tx.RowsAffected == 0 {
		panic(biz.CATE_ERR_TABLE_NO_CHANGE)
	}
	// return
	return &cateGet
}
