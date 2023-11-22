package middleware

import (
	common "SpiderShop-Restfull-API/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer recovery
		defer func() {
			// check if recovery is already in progress or not
			if err := recover(); err != nil {
				// cast err to Error HandlerFunc
				if errorHandler, ok := err.(*common.ErrorHandler); ok {
					// response to user error handler
					c.JSON(errorHandler.ErrorCode, gin.H{
						"Error": gin.H{
							"code":    errorHandler.ErrorCode,
							"message": errorHandler.ErrorMessage,
						},
					})
					c.Abort()
					return
					// simultaneously stop middleware and return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": err,
				})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
