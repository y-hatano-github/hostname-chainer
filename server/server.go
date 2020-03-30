package server

import (
	"log"

	"hostname-chainer/config"
	"hostname-chainer/server/middleware"
	"hostname-chainer/server/router"
	"hostname-chainer/types"

	"github.com/gin-gonic/gin"
)

// start server
func Run() error {

	var conf config.Config
	if err := conf.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	currentHostInfo := types.NewHost(conf)

	r := gin.Default()
	r.Use(middleware.SetHostInfo(currentHostInfo))
	router.Build(r)

	return r.Run(":" + conf.ListenPort)

}
