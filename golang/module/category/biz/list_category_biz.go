package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *createBiz) ListCategoryStorage(c context.Context, paging common.Paging, options []string) []entities.CateGet {
	// declare
	var category []entities.CateGet
	// deleting
	category = s.store.ListCategoryStorage(c, paging, options)
	// return
	return category
}
