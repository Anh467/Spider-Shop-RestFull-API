package router

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping (testing api working or not)
// @Tags Testing
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func V1Router(r *gin.Engine, aptx *common.AppConext) {
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
			// user
			getUserRouters(v1, aptx)
		}
	}
}
