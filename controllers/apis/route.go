package apis

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"grocery/handlers/api_handler"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1ShoppingList := r.Group("/v1/shopping-list")
	{
		v1ShoppingList.POST("/", api_handler.CreateHandler)
		v1ShoppingList.GET("/:id", api_handler.ShowHandler)
		v1ShoppingList.PUT("/:id", api_handler.PutHandler)
		v1ShoppingList.DELETE("/:id", api_handler.DeleteHandler)
	}

	v1Student := r.Group("/v1/student")
	{
		v1Student.POST("/", api_handler.CreateStudentHandler)
		v1Student.GET("/:id", api_handler.ShowStudentHandler)
		v1Student.GET("/join", api_handler.ShowStudentJoinDOsenHandler)
		v1Student.PUT("/:id", api_handler.PutStudentHandler)
		v1Student.DELETE("/:id", api_handler.DeleteStudentHandler)
	}

	v1Dosen := r.Group("/v1/dosen")
	{
		v1Dosen.POST("/", api_handler.CreateDosenPAHandler)
		v1Dosen.GET("/:id", api_handler.ShowDosenPAHandler)
		v1Dosen.PUT("/:id", api_handler.PutDosenPAHandler)
		v1Dosen.DELETE("/:id", api_handler.DeleteDosenPAHandler)
	}

	return r
}