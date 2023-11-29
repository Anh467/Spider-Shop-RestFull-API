package middleware

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	"SpiderShop-Restfull-API/module/user/storage"

	"github.com/gin-gonic/gin"
)

func CheckRole(aptx *common.AppConext, role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// declare
		var roleUser string
		var flag bool = false
		// get header
		token := c.GetHeader("token")
		// parse to claims
		claims, err := common.PraseToken(token, aptx.SuperSecretKey)
		if err != nil {
			// check permission
			c.Set("flag", flag)
			// go next middleware
			c.Next()
			return
		}
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// get role according account
		roleUser = business.GetUserRoleBiz(c.Request.Context(), claims.Account)
		// check in
		for _, ele := range role {
			if ele == roleUser {
				flag = true
				break
			}
		}
		// check permission
		c.Set("flag", flag)
		// go next middleware
		c.Next()
	}
}
