package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func (s *sqlserverStore) GetUserStorage(c context.Context, userid string) entities.UserGet {
	// declare
	var user *entities.UserGet
	// get user
	if err := s.aptx.GormDB.Where("UserID = ?", userid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&common.ErrorHandler{
				ErrorMessage: biz.USER_ERR_CANNOT_GET,
				ErrorCode:    http.StatusNotFound,
			})
		}
		panic(err)
	}
	// return user
	return *user
}
