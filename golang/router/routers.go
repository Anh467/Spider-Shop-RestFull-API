package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func V1Router(r *gin.Engine, aptx *common.AppConext) {
	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// add hanlder error
	r.Use(middleware.Recovery())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//test
			v1.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"test": "pinghhhhhhhhhhhhh",
				})
			})

			// authen
			getAuthenRouters(v1, aptx)

			// users
			getUserRouters(v1, aptx)

			// categories
			getCategoryRouters(v1, aptx)

			// products
			getProductRouters(v1, aptx)

			// prices
			getPriceRouters(v1, aptx)
		}
	}
}
