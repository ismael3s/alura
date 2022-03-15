package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ismael3s/alura/go/05/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.DELETE("/students/:id", controllers.RemoveStudent)
	r.POST("/students", controllers.CreateStudent)

	log.Fatal(r.Run())
}
