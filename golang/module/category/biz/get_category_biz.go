package biz

import (
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

func (s *createBiz) GetCategoryStorage(c context.Context, cateid int) entities.CateGet {
	// declare
	var category entities.CateGet
	// deleting
	category = s.store.GetCategoryStorage(c, cateid)
	// return
	return category
}
