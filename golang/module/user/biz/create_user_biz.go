package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/entity"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) CreateUserBiz(c context.Context, userCreate entity.UserCreate) entity.UserJWTModel {
	// reset UserID bc of increase of userid so it have to reset userid to default
	userCreate.UserID = 0
	// checking biz

	// Account
	userCreate.Account = strings.Trim(userCreate.Account, " ")
	CheckValidAccount(userCreate.Account)

	// Password
	CheckValidPassword(userCreate.Password)

	// Name
	userCreate.Name = strings.Trim(userCreate.Name, " ")
	// check blank
	if len(userCreate.Name) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_Name_MINIMUM_1,
		})
	}
	// regex
	CheckValidName(userCreate.Name)

	// Gender
	userCreate.Gender = strings.Trim(userCreate.Gender, " ")
	// regex
	CheckValidGender(userCreate.Gender)

	// Mail
	// regex
	CheckValidMail(userCreate.Mail)

	// Birth

	// creating
	userJWTModel := b.store.CreateUserStorage(c, userCreate)
	// return userJWTModel
	return userJWTModel
}
