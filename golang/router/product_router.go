package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/product/transport"
	"SpiderShop-Restfull-API/module/user/entities"

	"github.com/gin-gonic/gin"
)

func getProductRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	product := v1.Group("/products")
	{
		// inner product
		category := product.Group("/categories")
		{
			category.GET("/:cateid",
				middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
				transport.ListProductAccordingCateidTransport(aptx),
			)
		}
		// get a product
		product.GET("/",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.ListProductTransport(aptx))

		// get list products
		product.GET("/:productid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.GetProductTransport(aptx))

		// delete a product
		product.DELETE("/:productid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.DeleteProductTransport(aptx))

		// update a product
		product.PUT("/:productid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.UpdateProductTransport(aptx))

		// create a product
		product.POST("/",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.CreateProductTransport(aptx))
	}
}
