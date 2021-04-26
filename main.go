package main

import (
	"github.com/gin-gonic/gin"
	"grocery/controllers/api_controller"
	"grocery/controllers/web_controller"
	"log"
)

func main() {
	log.Println("starting go with templates")

	r := gin.Default()
	api_controller.Route(r)

	r.LoadHTMLGlob("templates/**/*")
	web_controller.Route(r)

	_ = r.Run()
}