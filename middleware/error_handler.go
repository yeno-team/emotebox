package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yeno-team/emotebox"
)

func ErrorHandler(logger emotebox.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) <= 0 {
			return
		}

		for _, ginErr := range c.Errors {
			logger.Warnf("ERROR: ", ginErr.Err.Error())
		}

		c.JSON(-1, gin.H{
			"errors": c.Errors,
		})
	}
}
