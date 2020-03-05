package server

import (
	"log"
	"recursive-echo/types"

	"github.com/gin-gonic/gin"

	"recursive-echo/config"
	"recursive-echo/server/middleware"
	"recursive-echo/server/router"
)

func Run() error {

	var conf config.Config
	if err := conf.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	host := types.NewHost(conf)

	r := gin.Default()
	r.Use(middleware.SetContext(host))
	router.Build(r, conf)

	return r.Run()

}
