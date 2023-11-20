package user

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/transport"

	"github.com/gin-gonic/gin"
)

func getProductRouters(api *gin.RouterGroup, aptx common.AppConext) {
	user := api.Group("/users")
	{
		user.POST("/", transport.CreateUserTransport(&aptx))
	}
}
