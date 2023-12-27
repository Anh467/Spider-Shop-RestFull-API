package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	entities "SpiderShop-Restfull-API/module/price/entities"
	"SpiderShop-Restfull-API/module/price/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePriceTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare price
		var price entities.PriceUpdate
		var priceid int
		var err error

		// get price
		// get query parameters url
		priceidStr, _ := c.Params.Get("priceid")
		if priceid, err = strconv.Atoi(priceidStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}

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

		// updating
		business.UpdatePriceBiz(c, price, priceid)

		// response
		c.JSON(http.StatusCreated, nil)
	}
}
