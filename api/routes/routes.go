package routes

import (
	"sre-goapi/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/student/:id", handlers.GetStudent)
		v1.GET("/student", handlers.GetAllStudents)
		v1.POST("/student/:id", handlers.UpdateStudentById)
		v1.DELETE("/student/:id", handlers.DeleteStudentById)
	}
}
