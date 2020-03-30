package router

import (
	"github.com/gin-gonic/gin"

	"hostname-chainer/server/controllers"
)

// routing
func Build(r *gin.Engine) {
	r.
		POST("/", controllers.Chain)
}
