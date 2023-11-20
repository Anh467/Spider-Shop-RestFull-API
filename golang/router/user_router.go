package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/transport"

	"github.com/gin-gonic/gin"
)

func getUserRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	user := v1.Group("/users")
	{
		user.POST("/", transport.CreateUserTransport(aptx))
	}
}
