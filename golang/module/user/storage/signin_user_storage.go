package storage

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	entities "SpiderShop-Restfull-API/module/user/entities"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func (s *mySQLStore) SigninUserStorage(c context.Context, userCredential entities.UserCredential) entities.UserJWTModel {
	// declare
	var userJWTModel entities.UserJWTModel
	var err error
	// check user credentials
	if err := s.aptx.GormDB.Unscoped().
		Where(userCredential).
		Select(entities.USER_TABLE_UserID, entities.USER_TABLE_Account,
			entities.USER_TABLE_Name, entities.USER_TABLE_Gender,
			entities.USER_TABLE_Mail, entities.USER_TABLE_Birth).
		First(&userJWTModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: biz.USER_ERR_CREDENTIAL,
			})
		}
		panic(err)
	}
	// encode user token
	if userJWTModel.Token, err = common.CreateToken(&entities.UserCreate{
		UserID:  userJWTModel.UserID,
		Account: userJWTModel.Account,
	}, s.aptx.SuperSecretKey); err != nil {
		panic(err)
	}
	// return
	return userJWTModel
}
