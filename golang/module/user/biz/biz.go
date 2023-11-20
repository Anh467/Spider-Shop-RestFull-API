package biz

import (
	"SpiderShop-Restfull-API/module/user/entity"
	"context"
)

type createStorage interface {
	//storage function
	CreateUserStorage(c context.Context, userCreate entity.UserCreate) entity.UserJWTModel
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
