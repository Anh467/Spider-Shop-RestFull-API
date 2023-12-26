package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
)

func (s *mySQLStore) DeletePriceStorage(ctx context.Context, priceid int) {
	// declare
	var count int64
	var price *entities.PriceGet
	// check existing priceid
	if err := s.aptx.GormDB.
		Table(entities.Price_TABLE).
		Where("PriceID = ?", priceid).
		First(&price).
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
	// soft deleting priceid
	if price.Status != entities.Price_TABLE_Status_Deleted {
		if err := s.aptx.GormDB.
			Table(entities.Price_TABLE).
			Where("PriceID = ?", priceid).
			Update(entities.Price_TABLE_Status, entities.Price_TABLE_Status_Deleted).
			Error; err != nil {
			panic(err)
		}
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusOK,
			ErrorMessage: "Status changed to deleted",
		})
	}
	// deleting priceid
	if err := s.aptx.GormDB.
		Table(entities.Price_TABLE).
		Where("PriceID = ?", priceid).
		Delete(nil).
		Error; err != nil {
		panic(err)
	}
}
