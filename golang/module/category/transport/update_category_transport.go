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

// UpdateCategory godoc
// @Summary      Update a Category
// @Description  user can able Update a category based one cateid and you must have rights admin to Update
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param 		 cateid path int required "Enter appropriate cateid"
// @Param        user body entities.CateUpdate  true "Category's information which needs to be add"
// @Success      200  {object} entities.CateGet
// @Success      204  {string}
// @Failure      400  {string} common.GLOBAL_BINDING_DATA_FAIL
// @Failure      400  {string} common.GLOBAL_WRONG_FORMAT_INTERGER
// @Failure      500  {string} common.GLOBAL_UNDEFIND_ERR
// @Router       /categories/:userid [PUT]
func UpdateCategoryTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var cateid int
		var cateStr string
		var err error
		var cateGet entities.CateGet
		var cateUpdate entities.CateUpdate
		// get data param header
		cateStr, _ = c.Params.Get("cateid")
		if cateid, err = strconv.Atoi(cateStr); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_WRONG_FORMAT_INTERGER,
			})
		}
		// get data body
		if err := c.ShouldBind(&cateUpdate); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_BINDING_DATA_FAIL,
			})
		}
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		cateGet = business.UpdateCategoryStorage(c, cateUpdate, cateid)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"category": cateGet,
		})
	}
}
