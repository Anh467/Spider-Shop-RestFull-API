package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
	"net/http"
)

func (s *sqlserverStore) ListUserStorage(c context.Context, paging common.Paging) []entities.UserGet {
	// declare
	var userList []entities.UserGet
	// get user
	if err := s.aptx.GormDB.Offset(paging.GetOffset()).
		Limit(paging.GetLimit()).
		Find(&userList); err != nil {
		panic(err)
	}
	// check list exist
	if len(userList) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: biz.USER_ERR_CANNOT_GET,
		})
	}
	// response
	return userList
}
