package web_controller

import (
	"github.com/gin-gonic/gin"
	"grocery/handlers/web_handler"
)


func Route(r *gin.Engine) *gin.Engine {
	r.GET("shopping-list/show/:id", web_handler.ShowHandler)
	r.GET("shopping-list/new/", web_handler.NewHandler)
	r.POST("shopping-list/", web_handler.CreateHandler)
	r.GET("shopping-list/edit/:id", web_handler.EditHandler)
	r.POST("shopping-list/update/:id", web_handler.UpdateHandler)
	r.GET("shopping-list/", web_handler.IndexHandler)
	r.GET("shopping-list/delete/:id", web_handler.DeleteHandler)
	return r
}
