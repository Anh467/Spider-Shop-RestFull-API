package biz

import (
	"SpiderShop-Restfull-API/common"
	"context"
	"net/http"
)

func (b *createBiz) DeletePriceStorage(ctx context.Context, priceid int) {
	// check the postive of the priceid
	if priceid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_PriceID_POSITIVE,
		})
	}
	// delete the price
	b.store.DeletePriceStorage(ctx, priceid)
}
