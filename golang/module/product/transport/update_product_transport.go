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

func UpdateProductTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var productid int
		var err error
		var productUpdate entities.ProductUpdate
		var productGet entities.ProductGet

		// get query parameters url
		productidStr, _ := c.Params.Get("productid")
		if productid, err = strconv.Atoi(productidStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}

		// get data from body
		if err := c.ShouldBind(&productUpdate); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_BINDING_DATA_FAIL,
			})
		}

		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// creating
		productGet = *business.UpdateProductBiz(c, productUpdate, productid)

		// res
		c.JSON(http.StatusCreated, gin.H{
			"product": productGet,
		})
	}
}
