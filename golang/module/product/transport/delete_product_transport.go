package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"SpiderShop-Restfull-API/module/product/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProductTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var productid int
		var err error
		var productGet *entities.ProductGet

		// get query parameters url
		productidStr, _ := c.Params.Get("productid")
		if productid, err = strconv.Atoi(productidStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}

		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// deleting
		productGet = business.DeleteProductBiz(c, productid)

		// response
		if productGet == nil {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"product": productGet,
		})
	}
}
