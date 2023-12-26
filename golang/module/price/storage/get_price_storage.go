package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func (s *mySQLStore) GetPriceStorage(ctx context.Context, priceid int) *entities.PriceGet {
	// declare
	var price *entities.PriceGet = nil

	// get the price base on priceid
	if err := s.aptx.GormDB.
		Where("PriceID = ?", priceid).
		First(&price).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz.PRICE_ERR_PriceID_NOT_FOUND,
			})
		} else {
			panic(err)
		}
	}
	// check get successful
	if price == nil {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.PRICE_ERR_PriceID_NOT_FOUND,
		})
	}
	// return
	return price
}
