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

func (s *mySQLStore) CreateUserStorage(c context.Context, userCreate entities.UserCreate) entities.UserJWTModel {
	// declare variables
	userJWTModel := &entities.UserJWTModel{}
	var token string
	var err error
	var count int64
	// check existing account
	s.aptx.GormDB.Table(entities.USER_TABLE).Where("Account = ?", userCreate.Account).
		Count(&count)
	if count > 0 {
		panic(&common.ErrorHandler{
			ErrorMessage: biz.USER_ERR_Account_UNIQUE,
			ErrorCode:    http.StatusBadRequest,
		})
	}
	// create
	if err := s.aptx.GormDB.Table(entities.USER_TABLE).Create(&userCreate).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidField) {
			panic(&common.ErrorHandler{
				ErrorMessage: biz.USER_ERR_CANNOT_CREATE,
				ErrorCode:    http.StatusBadRequest,
			})
		}
		panic(&common.ErrorHandler{
			ErrorMessage: biz.USER_ERR_CANNOT_CREATE,
			ErrorCode:    http.StatusConflict,
		})
	}
	// transfer token
	if token, err = common.CreateToken(&userCreate, s.aptx.SuperSecretKey); err != nil {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusNotImplemented,
			ErrorMessage: err.Error(),
		})
	}
	// set data to entity.UserJWTModel
	userJWTModel.UserID = userCreate.UserID
	userJWTModel.Account = userCreate.Account
	userJWTModel.Name = userCreate.Name
	userJWTModel.Gender = userCreate.Gender
	userJWTModel.Mail = userCreate.Mail
	userJWTModel.Birth = userCreate.Birth
	userJWTModel.Token = token
	// return user data which was created
	return *userJWTModel
}
