package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
)

type createStorage interface {
	//storage function
	CreateProductStorage(ctx context.Context, productCreate entities.ProductCreate) *entities.ProductGet
	DeleteProductStorage(ctx context.Context, productid int) *entities.ProductGet
	GetProductStorage(c context.Context, flag bool, productid int) *entities.ProductGet
	ListProductStorage(c context.Context, flag bool, paging common.Paging, options []string) *[]entities.ProductGet
	UpdateProductStorage(ctx context.Context, productUpdate entities.ProductUpdate, productid int) *entities.ProductGet
	ListProductAccordingCateIDStorage(c context.Context, flag bool, cateid int, paging common.Paging, options []string) *[]entities.ProductGet
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
