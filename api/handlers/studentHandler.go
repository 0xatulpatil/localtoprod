package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"sre-goapi/models"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	student := models.Student{Name: "Atul Patil", RollNo: 23, ID: 123}
	c.IndentedJSON(http.StatusOK, student)
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	students = append(students, models.Student{Name: "Atul Patil", RollNo: 23, ID: 123})

	c.IndentedJSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	student := models.Student{Name: "Atul Patil", RollNo: 23, ID: 123}
	c.IndentedJSON(http.StatusOK, student)
}

func UpdateStudentById(c *gin.Context) {
	students := []models.Student{
		{Name: "Atul", RollNo: 12, ID: 123},
		{Name: "Ani", RollNo: 13, ID: 124},
		{Name: "Him", RollNo: 14, ID: 125},
	}
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("ERROR: error converting id to integer", err)
	}
	var updateStudent models.Student

	if err := c.ShouldBindBodyWithJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i := range students {
		if students[i].ID == studentId {
			students[i] = updateStudent
			c.JSON(http.StatusOK, gin.H{"message": "Student Updated", "student": updateStudent})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Student not found"})
}

func DeleteStudentById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
