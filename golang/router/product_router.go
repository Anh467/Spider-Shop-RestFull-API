package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/product/transport"
	"SpiderShop-Restfull-API/module/user/entities"

	"github.com/gin-gonic/gin"
)

func getProductRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	user := v1.Group("/products")
	{
		// get a product
		user.GET("/",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.ListProductTransport(aptx))

		// get list products
		user.GET("/:cateid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.GetProductTransport(aptx))

		// delete a product
		user.DELETE("/:cateid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.DeleteProductTransport(aptx))

		// update a product
		user.PUT("/:cateid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.UpdateProductTransport(aptx))

		// create a product
		user.POST("/",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.CreateProductTransport(aptx))
	}
}
