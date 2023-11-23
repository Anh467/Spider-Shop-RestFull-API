package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
	"net/http"
	"strings"
)

func (s *createBiz) CreateCategoryBiz(c context.Context, category entities.CateCreate) entities.CateGet {
	// declare
	var cateGet entities.CateGet

	// name
	category.Name = strings.Trim(category.Name, " ")
	//check blank
	if len(category.Name) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: CATE_ERR_Name_Not_Blank,
		})
	}
	// regex
	CheckValidName(category.Name)

	// desc
	category.Desc = strings.Trim(category.Desc, " ")

	// image
	category.Image = strings.Trim(category.Image, " ")

	// creating
	cateGet = s.store.CreateCategoryStorage(c, category)
	return cateGet
}
