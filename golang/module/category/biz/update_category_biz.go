package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
	"net/http"
	"strings"
)

func (s *createBiz) UpdateCategoryStorage(c context.Context, cateUpdate entities.CateUpdate, cateid int) entities.CateGet {
	// declare
	var category entities.CateGet
	// chekc status
	cateUpdate.Status = strings.Trim(cateUpdate.Status, " ")
	if cateUpdate.Status != "" {
		flag := false
		for _, status := range entities.CATE_STATUS {
			if cateUpdate.Status == status {
				flag = true
			}
		}
		if !flag {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: CATE_ERR_Status_EXIST,
			})
		}
	}
	// updating
	category = *s.store.UpdateCategoryStorage(c, cateUpdate, cateid)
	// return
	return category
}
