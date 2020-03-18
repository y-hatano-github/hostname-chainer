package router

import (
	"github.com/gin-gonic/gin"

	"hostname-chainer/server/controllers"
)

func Build(r *gin.Engine) {
	r.
		POST("/", controllers.Chain)
}
