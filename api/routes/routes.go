package routes

import (
	"sre-goapi/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, studentHandler *handlers.StudentHandler) {
	v1 := r.Group("/v1")
	{
		v1.GET("/student/:id", studentHandler.GetStudentById)
		v1.GET("/student", studentHandler.GetAllStudents)
		v1.POST("/student/:id", studentHandler.UpdateStudentById)
		v1.DELETE("/student/:id", studentHandler.DeleteStudentById)
	}
}
