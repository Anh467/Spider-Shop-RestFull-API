package biz

import (
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
)

type createStorage interface {
	//storage function
	CreateUserStorage(c context.Context, userCreate entities.UserCreate) entities.UserJWTModel
}

type createBiz struct {
	store createStorage
}

func NewCreateBiz(store createStorage) *createBiz {
	return &createBiz{store: store}
}
