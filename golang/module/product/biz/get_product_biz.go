package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
)

func (b *createBiz) GetProductBiz(c context.Context, flag bool, productid int) *entities.ProductGet {
	// check productid must be positive
	if productid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRODUCT_ERR_PRODUCT_ID_POSITIVE,
		})
	}
	// get product
	productGet := b.store.GetProductStorage(c, flag, productid)
	// return
	return productGet
}
