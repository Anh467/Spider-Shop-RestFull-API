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

func GetProductTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var productid int
		var err error
		var productGet *entities.ProductGet
		var flag bool

		// get flag to check permissions
		flag = c.GetBool("flag")

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

		// get product
		productGet = business.GetProductBiz(c, flag, productid)

		// res
		c.JSON(http.StatusOK, gin.H{
			"product": productGet,
		})
	}
}
