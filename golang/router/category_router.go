package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/category/transport"
	"SpiderShop-Restfull-API/module/user/entities"

	"github.com/gin-gonic/gin"
)

func getCategoryRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	category := v1.Group("/categories")
	{
		category.GET("/",
			transport.ListCategoryTransport(aptx))
		category.GET("/:cateid",
			transport.GetCategoryTransport(aptx))
		category.DELETE("/:cateid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.DeleteCategoryTransport(aptx))
		category.PUT("/:cateid",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.UpdateCategoryTransport(aptx))
		category.POST("/",
			middleware.CheckPermission(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.CreateCategoryTransport(aptx))
	}
}
