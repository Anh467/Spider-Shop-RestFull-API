package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/biz"
	"SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/category/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategory godoc
// @Summary      get a category
// @Description  user can able get a specific category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param 		 cateid path int required "Enter appropriate cateid"
// @Success      200  {object} entities.CateGet
// @Failure      400  {string} http.StatusBadRequest
// @Failure      500  {string} common.GLOBAL_UNDEFIND_ERR
// @Router       /categories/:cateid [GET]
func GetCategoryTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var cateid int
		var cateStr string
		var err error
		var cateGet entities.CateGet
		// get data param header
		cateStr, _ = c.Params.Get("cateid")
		if cateid, err = strconv.Atoi(cateStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		cateGet = business.GetCategoryStorage(c, cateid)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"category": cateGet,
		})
	}
}
