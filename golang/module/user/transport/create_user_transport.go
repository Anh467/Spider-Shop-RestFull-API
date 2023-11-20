package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	entity "SpiderShop-Restfull-API/module/user/entity"
	"SpiderShop-Restfull-API/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare variable
		var userGet entity.UserCreate
		// get body data
		if err := c.ShouldBind(&userGet); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Binding data from database failed",
			})
		}
		// Dependency
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// creating
		userJWTModel := business.CreateUserBiz(c, userGet)
		// response
		c.JSON(http.StatusCreated, gin.H{
			"user": userJWTModel,
		})
	}
}
