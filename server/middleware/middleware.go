package middleware

import (
	"github.com/gin-gonic/gin"

	"recursive-echo/config"
)

func SetContext(c config.Config) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("host", c.HostName)
	}
}
