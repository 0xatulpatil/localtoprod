package handlers

import (
	"net/http"
	"strconv"

	"sre-goapi/controllers"
	"sre-goapi/models"
	logger "sre-goapi/utils"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentController *controllers.StudentController
}

func NewStudentHandler(stdController *controllers.StudentController) *StudentHandler {
	return &StudentHandler{
		studentController: stdController,
	}
}

func (s *StudentHandler) GetAllStudents(c *gin.Context) {
	student, err := s.studentController.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive students"})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (s *StudentHandler) CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		logger.Error("Invalid Request body for creating student")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request body"})
	}

	_, err := s.studentController.CreateStudent(&newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid"})
	}
}

func (s *StudentHandler) GetStudentById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := s.studentController.GetStudentById(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (s *StudentHandler) UpdateStudentById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var updatedStudent models.Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	student, err := s.studentController.UpdateStudent(int(id), &updatedStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (s *StudentHandler) DeleteStudentById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	err = s.studentController.DeleteStudent(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
