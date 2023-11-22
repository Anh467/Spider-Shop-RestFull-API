package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/transport"

	"github.com/gin-gonic/gin"
)

func getAuthenRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	user := v1.Group("/authen")
	{
		user.POST("/register", transport.CreateUserTransport(aptx))
		user.POST("/signin", transport.SigninUserTransport(aptx))
	}
}
