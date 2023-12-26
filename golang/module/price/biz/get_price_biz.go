package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
)

func (b *createBiz) GetPriceBiz(ctx context.Context, priceid int) *entities.PriceGet {
	// declare
	var priceGet *entities.PriceGet
	// check the postive of the priceid
	if priceid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_PriceID_POSITIVE,
		})
	}
	// delete the price
	priceGet = b.store.GetPriceStorage(ctx, priceid)
	// return the price
	return priceGet
}
