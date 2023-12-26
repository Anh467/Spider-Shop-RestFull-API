package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/entities"
	productEntities "SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) GetPriceStorageBaseOnProductID(ctx context.Context, productID int) []*entities.PriceGet {
	// declare
	var count int64
	var prices []*entities.PriceGet
	// check exist productid
	if err := s.aptx.GormDB.
		Table(productEntities.PRODUCT_TABLE).
		Where("ProductID = ?", productID).
		Count(&count).
		Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRICE_ERR_PRODUCTID_NOT_FOUND,
		})
	}

	// get list of prices base on productid
	if err := s.aptx.GormDB.
		Where("ProductID = ?", productID).
		Find(&prices).
		Error; err != nil {
		panic(err)
	}
	// return
	return prices
}
