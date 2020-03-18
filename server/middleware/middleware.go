package middleware

import (
	"github.com/gin-gonic/gin"

	"hostname-chainer/types"
)

func SetHostInfo(currentHost types.CurrentHost) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("currentHost", currentHost)
	}
}
