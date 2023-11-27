package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"SpiderShop-Restfull-API/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var offsetStr, limitStr string
		var paging common.Paging
		var productGets []entities.ProductGet
		var flag bool

		// get flag to check permissions
		flag = c.GetBool("flag")

		// get parma
		offsetStr = c.Query("offset")
		limitStr = c.Query("limit")

		// init paging
		paging.Process(limitStr, offsetStr)

		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// get products
		productGets = *business.ListProductBiz(c, flag, paging, nil)

		// return
		c.JSON(http.StatusOK, gin.H{
			"products": productGets,
		})
	}
}
