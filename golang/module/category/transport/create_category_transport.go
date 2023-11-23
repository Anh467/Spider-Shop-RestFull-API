package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/category/biz"
	"SpiderShop-Restfull-API/module/category/entities"
	"SpiderShop-Restfull-API/module/category/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary      post Categories
// @Description  user can able create new account
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        user body entities.CateCreate  true "Category's information which needs to be add"
// @Success      201  {object}   entities.CateGet
// @Failure      400  {string} common.GLOBAL_BINDING_DATA_FAIL
// @Failure      500  {string} common.GLOBAL_UNDEFIND_ERR
// @Router       /categories [POST]
func CreateCategoryTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var cateGet entities.CateGet
		var cateCreate *entities.CateCreate
		// get data body
		if err := c.ShouldBind(&cateCreate); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.GLOBAL_BINDING_DATA_FAIL,
			})
		}
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		cateGet = business.CreateCategoryBiz(c, *cateCreate)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"category": cateGet,
		})
	}
}
