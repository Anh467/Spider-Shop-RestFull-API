package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/biz"
	"SpiderShop-Restfull-API/module/product/entities"
	"SpiderShop-Restfull-API/module/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProductTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var productCreate entities.ProductCreate
		var productGet entities.ProductGet

		// get data from body
		if err := c.ShouldBind(&productCreate); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_BINDING_DATA_FAIL,
			})
		}

		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// creating
		productGet = *business.CreateProductBiz(c, productCreate)

		// res
		c.JSON(http.StatusCreated, gin.H{
			"product": productGet,
		})
	}
}
