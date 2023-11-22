package biz

import (
	"SpiderShop-Restfull-API/common"
	"net/http"
	"regexp"

	"github.com/dlclark/regexp2"
)

// Account
func CheckValidAccount(account string) {
	match, _ := regexp.MatchString(USER_PATTERN_Account, account)
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Account,
		})
	}

}

// Password
func CheckValidPassword(password string) {
	re := regexp2.MustCompile(USER_PATTERN_Password, regexp2.None)
	if isMatch, _ := re.MatchString(password); !isMatch {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Password,
		})
	}
}

// Name
func CheckValidName(name string) {
	//check regex
	match, _ := regexp.MatchString(USER_PATTERN_Name, name)
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Name,
		})
	}
}

// gender
func CheckValidGender(gender string) {
	// if Gender is not empty check its in range gender
	if len(gender) == 0 {
		return
	}
	genderFlag := false
	for _, gender := range USER_GENDER {
		if gender == gender {
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
func CheckValidMail(mail string) {
	match, err := regexp.MatchString(USER_PATTERN_Mail, mail)
	if err != nil {
		panic(err)
	}
	if !match {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: USER_ERR_PATTERN_Mail,
		})
	}
}
