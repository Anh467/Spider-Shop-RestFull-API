package biz

import (
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *createBiz) UpdateCategoryStorage(c context.Context, cateUpdate entities.CateUpdate, cateid int) entities.CateGet {
	// declare
	var category entities.CateGet
	// deleting
	category = *s.store.UpdateCategoryStorage(c, cateUpdate, cateid)
	// return
	return category
}
