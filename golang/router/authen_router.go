package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/transport"

	"github.com/gin-gonic/gin"
)

func getAuthenRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	authen := v1.Group("/authen")
	{
		authen.POST("/register", transport.CreateUserTransport(aptx))
		authen.POST("/signin", transport.SigninUserTransport(aptx))
	}
}
