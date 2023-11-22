package transport

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/user/biz"
	"SpiderShop-Restfull-API/module/user/entities"
	"SpiderShop-Restfull-API/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signin godoc
// @Summary      Signin
// @Description  Guess can sign in to have a user's right
// @Tags         authen
// @Accept       json
// @Produce      json
// @Param        user body  entities.UserCredential  true "Data of user's credentials"
// @Success      200  {object}   []entities.UserJWTModel
// @Failure      400  {string}   http.StatusBadRequest
// @Failure      401  {string}   http.StatusUnauthorized
// @Failure      500  {string}   http.StatusInternalServerError
// @Router       /authen/signin [post]
func SigninUserTransport(aptx *common.AppConext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// declare
		var userCredential entities.UserCredential
		var userJWTModel entities.UserJWTModel
		// get body
		if err := c.ShouldBind(&userCredential); err != nil {
			panic(&common.ErrorHandler{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: common.AUTHEN_ERR_SIGN_REQUEST_BODY,
			})
		}
		// dependencies
		store := storage.NewMySQLStorage(aptx)
		business := biz.NewCreateBiz(store)
		// logging in
		userJWTModel = business.SigninUserBiz(c, userCredential)
		// response
		c.JSON(http.StatusOK, gin.H{
			"user": userJWTModel,
		})
	}
}
