package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (b *createBiz) DeleteProductBiz(ctx context.Context, productid int) *entities.ProductGet {
	// Product int
	if productid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRODUCT_ERR_PRODUCT_ID_POSITIVE,
		})
	}

	// deleting
	productGet := b.store.DeleteProductStorage(ctx, productid)

	// return
	return productGet
}
