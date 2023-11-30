package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
)

func (b *createBiz) ListProductAccordingCateidBiz(c context.Context, flag bool, cateid int, paging common.Paging, options []string) *[]entities.ProductGet {
	// get products
	productGets := b.store.ListProductAccordingCateIDStorage(c, flag, cateid, paging, options)
	// return
	return productGets
}
