package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	entities "SpiderShop-Restfull-API/module/price/entities"
	"SpiderShop-Restfull-API/module/price/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePriceTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare price
		var price entities.PriceCreate

		// get price information
		if err := c.ShouldBind(&price); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_BINDING_DATA_FAIL,
			})
		}

		// Dependency
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// creating
		priceGet := business.CreatePriceStorage(c, price)

		// response
		c.JSON(http.StatusCreated, gin.H{
			"price": priceGet,
		})
	}
}
