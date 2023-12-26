package storage

import (
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
)

func (s *mySQLStore) CreatePriceStorage(ctx context.Context, priceCreate entities.PriceCreate) *entities.PriceGet {
	var priceGet *entities.PriceGet
	// create
	if err := s.aptx.GormDB.
		Create(&priceCreate).
		Error; err != nil {
	}
	// get price which just created
	if err := s.aptx.GormDB.
		Where("PriceID = ?", priceCreate.PriceID).
		First(&priceGet).
		Error; err != nil {
		panic(err)
	}
	// return
	return priceGet
}
