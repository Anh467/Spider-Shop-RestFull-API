package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"SpiderShop-Restfull-API/module/category/transport"
	"SpiderShop-Restfull-API/module/user/entities"

	"github.com/gin-gonic/gin"
)

func getCategoryRouters(v1 *gin.RouterGroup, aptx *common.AppConext) {
	user := v1.Group("/categories")
	{
		user.GET("/",
			transport.ListCategoryTransport(aptx))
		user.GET("/:cateid",
			transport.GetCategoryTransport(aptx))
		user.DELETE("/:cateid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.DeleteCategoryTransport(aptx))
		user.PUT("/:cateid",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.UpdateCategoryTransport(aptx))
		user.POST("/",
			middleware.CheckRole(aptx, entities.USER_TABLE_Role_ADMIN),
			transport.CreateCategoryTransport(aptx))
	}
}
