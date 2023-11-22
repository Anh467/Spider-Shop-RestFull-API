package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	"SpiderShop-Restfull-API/module/user/entities"
	"SpiderShop-Restfull-API/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers godoc
// @Summary      get users
// @Description  this function must have permission administrator
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit the number of users to return"
// @Param        offset query int false "Offset the list of users"
// @Success      200  {object}   []entities.UserGet
// @Failure      404  {string}   http.StatusNotFound
// @Router       /users [get]
func ListUserTrasnport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var offsetStr, limitStr string
		var paging common.Paging
		var userList []entities.UserGet
		// get param
		offsetStr = c.Query("offset")
		limitStr = c.Query("limit")
		// init paging
		paging.Process(offsetStr, limitStr)
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// get users
		userList = business.ListUserBiz(c, paging)
		// response
		c.JSON(http.StatusOK, gin.H{
			"users": userList,
		})
	}
}
