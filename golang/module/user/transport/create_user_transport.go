package transport

import (
	common "SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	entities "SpiderShop-Restfull-API/module/user/entities"
	"SpiderShop-Restfull-API/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUsers godoc
// @Summary      post users
// @Description  user can able create new account
// @Tags         authen
// @Accept       json
// @Produce      json
// @Param        user body  entities.UserCreate  true "User object that needs to be added"
// @Success      201  {object}   entities.UserJWTModel
// @Failure      400  {string} biz.USER_ERR_CANNOT_CREATE
// @Failure      409  {string} biz.USER_ERR_CANNOT_CREATE
// @Failure      501  {string} biz.USER_ERR_CANNOT_CREATE
// @Router       /register [POST]
func CreateUserTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare variable
		var userGet entities.UserCreate
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
