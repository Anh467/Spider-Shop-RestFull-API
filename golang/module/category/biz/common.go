package biz

import (
	"SpiderShop-Restfull-API/common"
	"net/http"

	"github.com/dlclark/regexp2"
)

// Password
func CheckValidName(name string) {
	re := regexp2.MustCompile(CATE_PATTERN_Name, regexp2.None)
	if isMatch, _ := re.MatchString(name); !isMatch {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: CATE_ERR_PATTERN_Name,
		})
	}
}
