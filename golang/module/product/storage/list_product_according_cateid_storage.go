package storage

import (
	"SpiderShop-Restfull-API/common"
	biz_category "SpiderShop-Restfull-API/module/category/biz"
	entities_category "SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
	"strings"
)

func (s *mySQLStore) ListProductAccordingCateIDStorage(c context.Context, flag bool, cateid int, paging common.Paging, options []string) *[]entities.ProductGet {
	// declare product
	var productGets *[]entities.ProductGet
	var query = s.aptx.GormDB
	var cateStatus string
	// get list deleted cateidsDeleted
	if !flag {
		if err := s.aptx.GormDB.
			Table(entities_category.CATE_TABLE).
			Where("CateID = ?", cateid).
			Select("Status").
			Scan(&cateStatus).
			Error; err != nil {
			panic(err)
		}

		if strings.EqualFold(cateStatus, entities_category.USER_TABLE_Status_Deleted) {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadGateway,
				ErrorMessage: biz.PRODUCT_ERR_Temporarily_deleted,
			})
		}
		if cateStatus == "" {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz_category.CATE_ERR_CateID_EXIST,
			})
		}
	}
	// init options
	for _, option := range options {
		query = query.Where(option)
	}

	// add conditions user have a rights to access deleted products or not
	if !flag {
		query = query.
			Not("Status = ?", entities.PRODUCT_TABLE_Status_Deleted)
	}
	// get products
	if err := query.
		Preload("Cate").
		Where("CateID = ?", cateid).
		Find(&productGets).
		Error; err != nil {
		panic(err)
	}
	// return
	return productGets
}
