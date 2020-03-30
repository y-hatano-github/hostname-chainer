package middleware

import (
	"github.com/gin-gonic/gin"

	"hostname-chainer/types"
)

// currentHost interface set to middleware
func SetHostInfo(currentHost types.CurrentHost) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("currentHost", currentHost)
	}
}
