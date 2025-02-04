package middleware

import (
	"errors"
	"net/http"

	"example.com/my_errors"
	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware handles errors returned by handlers.
func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			var customErr *my_errors.CustomError
			if errors.As(err.Err, &customErr) {
				c.JSON(customErr.Code, gin.H{
					"error": customErr.Message,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "An unexpected error occurred",
				})
			}
		}
	}
}
