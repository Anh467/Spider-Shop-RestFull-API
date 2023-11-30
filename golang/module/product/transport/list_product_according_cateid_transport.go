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

func ListProductAccordingCateidTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var offsetStr, limitStr string
		var paging common.Paging
		var productGets []entities.ProductGet
		var flag bool
		var cateid int
		var err error

		// get flag to check permissions
		flag = c.GetBool("flag")

		// get query parameters url
		cateidStr, _ := c.Params.Get("cateid")
		if cateid, err = strconv.Atoi(cateidStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}

		// get Query
		offsetStr = c.Query("offset")
		limitStr = c.Query("limit")

		// init paging
		paging.Process(limitStr, offsetStr)

		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)

		// get products
		productGets = *business.ListProductAccordingCateidBiz(c, flag, cateid, paging, nil)

		// return
		c.JSON(http.StatusOK, gin.H{
			"products": productGets,
		})
	}
}
