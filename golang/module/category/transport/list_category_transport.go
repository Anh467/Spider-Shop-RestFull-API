package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/biz"
	"SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/category/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListCategory godoc
// @Summary      list categories
// @Description  user can get list of categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit the number of users to return"
// @Param        offset query int false "Offset the list of users"
// @Success      200  {object} entities.CateGet
// @Failure      400  {string}
// @Failure      500  {string} common.GLOBAL_UNDEFIND_ERR
// @Router       /categories [GET]
func GetCategoryTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var paging common.Paging
		var offsetStr, limitStr string
		var cateListGet []entities.CateGet
		// get data param header
		offsetStr = c.Query("offset")
		limitStr = c.Query("limit")
		// paging int
		paging.Process(limitStr, offsetStr)
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		cateListGet = business.ListCategoryStorage(c, paging, nil)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"categories": cateListGet,
		})
	}
}
