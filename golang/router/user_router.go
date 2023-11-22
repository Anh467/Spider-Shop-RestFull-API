package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/user/entities"
	"SpiderShop-Restfull-API/module/user/transport"

	"github.com/gin-gonic/gin"
)

func getUserRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	user := v1.Group("/users")
	{
		user.GET("/",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.ListUserTrasnport(aptx))
	}
}
