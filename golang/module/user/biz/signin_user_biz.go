package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) SigninUserBiz(c context.Context, userCredential entities.UserCredential) entities.UserJWTModel {
	// declare
	var userJWTModel entities.UserJWTModel
	// check biz
	// Account
	// Account
	userCredential.Account = strings.Trim(userCredential.Account, " ")
	if len(userCredential.Account) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_Account_Blank,
		})
	}
	// Password
	userCredential.Password = strings.Trim(userCredential.Password, " ")
	if len(userCredential.Password) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_Password_Blank,
		})
	}
	//	logging in
	userJWTModel = b.store.SigninUserStorage(c, userCredential)
	// return
	return userJWTModel
}
