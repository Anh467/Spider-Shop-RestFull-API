package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/entity"
	"context"
	"net/http"
	"regexp"
	"strings"
)

func (b *createBiz) CreateUserBiz(c context.Context, userCreate entity.UserCreate) entity.UserJWTModel {
	// declare variables
	var match bool
	var err error
	// reset UserID bc of increase of userid so it have to reset userid to default
	userCreate.UserID = 0
	// checking biz
	// Account
	/*
		userCreate.Account = strings.Trim(userCreate.Account, " ")
		match, _ = regexp.MatchString(USER_PATTERN_Account, userCreate.Account)
		if !match {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: USER_ERR_PATTERN_Account,
			})
		}
	*/
	// Password
	match, err = regexp.MatchString(USER_PATTERN_Password, userCreate.Password)
	if err != nil {
		panic(err)
	}
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Password,
		})
	}
	// Name
	// check blank
	userCreate.Name = strings.Trim(userCreate.Name, " ")
	if len(userCreate.Name) == 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_Name_MINIMUM_1,
		})
	}
	//check regex
	match, _ = regexp.MatchString(USER_PATTERN_Name, userCreate.Name)
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Name,
		})
	}
	// Gender
	userCreate.Gender = strings.Trim(userCreate.Gender, " ")
	// if Gender is not empty check its in range gender
	if len(userCreate.Gender) != 0 {
		genderFlag := false
		for _, gender := range USER_GENDER {
			if gender == userCreate.Gender {
				genderFlag = true
				break
			}
		}
		if !genderFlag {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: USER_ERR_Gender,
			})
		}
	}

	// Mail
	match, _ = regexp.MatchString(USER_ERR_PATTERN_Mail, userCreate.Mail)
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Mail,
		})
	}
	// Birth
	// creating
	userJWTModel := b.store.CreateUserStorage(c, userCreate)
	return userJWTModel
}
