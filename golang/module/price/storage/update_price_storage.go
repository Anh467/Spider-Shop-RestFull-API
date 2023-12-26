package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) UpdatePriceStorage(ctx context.Context, priceUpdate entities.PriceUpdate, priceid int) {
	// declare
	var count int64
	// check existing priceid
	if err := s.aptx.GormDB.
		Table(entities.Price_TABLE).
		Where("PriceID = ?", priceid).
		First(nil).
		Count(&count).
		Error; err != nil {
		panic(err)
	}
	if count == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRICE_ERR_PriceID_NOT_FOUND,
		})
	}
	// update price
	if err := s.aptx.GormDB.
		Table(entities.Price_TABLE).
		Where("PriceID = ?", priceid).
		Updates(&priceUpdate).
		Error; err != nil {
		panic(err)
	}
}
