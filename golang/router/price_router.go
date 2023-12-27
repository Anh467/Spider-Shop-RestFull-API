package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/price/transport"
	"SpiderShop-Restfull-API/module/user/entities"

	"github.com/gin-gonic/gin"
)

func getPriceRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	price := v1.Group("/prices")
	{
		price.GET("/:priceid", transport.GetPriceTransport(aptx))
		price.GET("/products/:productid", transport.ListPriceStorageBaseOnProductIDTransport(aptx))
		price.POST("/",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.CreatePriceTransport(aptx))
		price.PUT("/:priceid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.UpdatePriceTransport(aptx))
		price.DELETE("/:priceid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.DeletePriceTransport(aptx))
	}
}
