package biz

import (
	"SpiderShop-Restfull-API/common"
	"net/http"
	"regexp"
)

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
