package server

import (
	"log"

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

	r := gin.Default()
	r.Use(middleware.SetContext(conf))
	router.Build(r, conf)

	return r.Run()

}
