package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/entities"
	"context"
)

type createStorage interface {
	//storage function
	CreateCategoryStorage(c context.Context, category entities.CateCreate) entities.CateGet
	DeleteCategoryStorage(c context.Context, cateid string) *entities.CateGet
	GetCategoryStorage(c context.Context, cateid int) entities.CateGet
	ListCategoryStorage(c context.Context, paging common.Paging, options ...string) []entities.CateGet
	UpdateCategoryStorage(c context.Context, cateUpdate entities.CateUpdate, cateid int) *entities.CateGet
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
