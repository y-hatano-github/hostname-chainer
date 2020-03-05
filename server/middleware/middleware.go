package middleware

import (
	"github.com/gin-gonic/gin"
	"recursive-echo/types"
)

func SetContext(host types.Host) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("host", host)
	}
}
