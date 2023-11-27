package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
)

func (b *createBiz) ListProductBiz(c context.Context, flag bool, paging common.Paging, options []string) *[]entities.ProductGet {
	// get products
	productGets := b.store.ListProductStorage(c, flag, paging, options)
	// return
	return productGets
}
