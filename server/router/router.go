package router

import (
	"github.com/gin-gonic/gin"

	"recursive-echo/config"
	"recursive-echo/server/controllers"
)

func Build(r *gin.Engine, config config.Config) {
	r.
		GET("/", controllers.Echo)
}
