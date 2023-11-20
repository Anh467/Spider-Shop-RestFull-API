package module

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func V1Routers(r *gin.Engine, aptx *common.AppConext) {
	r.Use(middleware.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "pinghhhhhhhhhhhhh",
		})
	})
	// // router API
	// api := r.Group("/api")

	// // router V1
	// v1 := api.Group("/v1")

}
