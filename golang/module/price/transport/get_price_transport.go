package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPriceTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare variable
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
		// Dependency
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		userJWTModel := business.GetPriceBiz(c, priceid)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"price": userJWTModel,
		})
	}
}
