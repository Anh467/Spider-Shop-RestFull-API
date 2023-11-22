package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
)

func (b createBiz) ListUserBiz(c context.Context, paging common.Paging) []entities.UserGet {
	// declare
	var userList []entities.UserGet
	// get
	userList = b.store.ListUserStorage(c, paging)
	// return users
	return userList
}
