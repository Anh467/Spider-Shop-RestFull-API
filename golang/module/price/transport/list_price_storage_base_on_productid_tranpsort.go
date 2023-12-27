package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/biz"
	"SpiderShop-Restfull-API/module/price/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPriceStorageBaseOnProductIDTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare variable
		var productid int
		var err error
		// get price
		// get query parameters url
		productidStr, _ := c.Params.Get("productid")
		if productid, err = strconv.Atoi(productidStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}

		// Dependency
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// creating
		priceList := business.ListPriceStorageBaseOnProductIDBiz(c, productid)

		// response
		c.JSON(http.StatusOK, gin.H{
			"price": priceList,
		})
	}
}
