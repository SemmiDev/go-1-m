package main

import (
	"github.com/gin-gonic/gin"
	"grocery/controllers/apis"
	"grocery/controllers/web"
	"log"
)

func main() {
	log.Println("starting go with templates")
	r := gin.Default()

	apis.Route(r)

	r.LoadHTMLGlob("templates/**/*")
	web.Route(r)

	_ = r.Run()
}