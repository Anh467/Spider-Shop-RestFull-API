package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
)

func (b *createBiz) ListPriceStorageBaseOnProductIDBiz(ctx context.Context, productID int) []*entities.PriceGet {
	var priceGets []*entities.PriceGet
	// ProductID int     `json:"ProductID" gorm:"column:ProductID"`
	if productID <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_PRODUCTID_POSITIVE,
		})
	}

	// list prices of the product
	priceGets = b.store.ListPriceStorageBaseOnProductIDStorage(ctx, productID)

	// return prices
	return priceGets
}
