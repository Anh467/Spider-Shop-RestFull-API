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

// DeleteCategory godoc
// @Summary      delete a Category
// @Description  user can able delete a category based one cateid and you must have rights admin to delete
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param 		 cateid path int required "Enter appropriate cateid"
// @Success      201  {object} entities.CateGet
// @Success      204  {string}
// @Failure      400  {string} common.GLOBAL_BINDING_DATA_FAIL
// @Failure      400  {string} common.GLOBAL_WRONG_FORMAT_INTERGER
// @Failure      500  {string} common.GLOBAL_UNDEFIND_ERR
// @Router       /categories/:userid [DELETE]
func DeleteCategoryTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var cateid int
		var cateStr string
		var err error
		var cateGet *entities.CateGet
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
		cateGet = business.DeleteCategoryStorage(c, cateid)
		// response
		if cateGet == nil {
			c.JSON(http.StatusNoContent, gin.H{
				"status": "dpne",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"category": cateGet,
		})
	}
}
