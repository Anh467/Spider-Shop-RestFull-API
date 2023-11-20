package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func V1Router(r *gin.Engine, aptx *common.AppConext) {
	r.Use(middleware.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "pinghhhhhhhhhhhhh",
		})
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// user
			getUserRouters(v1, aptx)
		}
	}
}
