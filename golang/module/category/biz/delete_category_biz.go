package biz

import (
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *createBiz) DeleteCategoryStorage(c context.Context, cateid int) *entities.CateGet {
	// declare
	var category *entities.CateGet
	// deleting
	category = s.store.DeleteCategoryStorage(c, cateid)
	// return
	return category
}
