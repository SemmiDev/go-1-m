package web

import (
	"github.com/gin-gonic/gin"
	"grocery/handlers/html_handler"
)


func Route(r *gin.Engine) *gin.Engine {
	r.GET("shopping-list/show/:id", html_handler.ShowHandler)
	r.GET("shopping-list/new/", html_handler.NewHandler)
	r.POST("shopping-list/", html_handler.CreateHandler)
	r.GET("shopping-list/edit/:id", html_handler.EditHandler)
	r.POST("shopping-list/update/:id", html_handler.UpdateHandler)
	r.GET("shopping-list/", html_handler.IndexHandler)
	r.GET("shopping-list/delete/:id", html_handler.DeleteHandler)
	return r
}
