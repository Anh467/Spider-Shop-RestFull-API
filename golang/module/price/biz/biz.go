package biz

import (
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
)

type createStorage interface {
	//storage function
	CreatePriceStorage(ctx context.Context, priceCreate entities.PriceCreate) *entities.PriceGet
	DeletePriceStorage(ctx context.Context, priceid int)
	GetPriceStorage(ctx context.Context, priceid int) *entities.PriceGet
	ListPriceStorageBaseOnProductIDStorage(ctx context.Context, productID int) []*entities.PriceGet
	UpdatePriceStorage(ctx context.Context, priceUpdate entities.PriceUpdate, priceid int)
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
