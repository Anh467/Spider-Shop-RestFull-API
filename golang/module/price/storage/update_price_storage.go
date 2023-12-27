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

func (s *mySQLStore) UpdatePriceStorage(ctx context.Context, priceUpdate entities.PriceUpdate, priceid int) {
	// declare
	var count int64
	// check existing priceid
	if err := s.aptx.GormDB.
		Table(entities.Price_TABLE).
		Where("PriceID = ?", priceid).
		First(&entities.PriceGet{}).
		Count(&count).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: biz.PRICE_ERR_PriceID_NOT_FOUND,
			})
		}
		panic(err)
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
