package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sre-goapi/controllers"
	"sre-goapi/handlers"
	"sre-goapi/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDB() {
	var err error
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	testDB.AutoMigrate(&models.Student{}) // Migrate schema
}

// Function to initialize a test router
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	studentController := controllers.NewStudentController(testDB)
	studentHandler := handlers.NewStudentHandler(studentController)
	RegisterRoutes(router, studentHandler)

	return router
}

func TestStudent(t *testing.T) {
	setupTestDB() // Initialize in-memory DB
	router := setupRouter()

	t.Run("CreateStudent", func(t *testing.T) {
		student := models.Student{Name: "John Doe", RollNo: 21, ID: 12}
		body, _ := json.Marshal(student)

		// Create a request
		req, _ := http.NewRequest("POST", "/v1/student", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code) // Expect 201 Created
		var resp models.Student
		json.Unmarshal(w.Body.Bytes(), &resp)
	})
	t.Run("GetAllStudent", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/student", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp []models.Student
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, len(resp), 1)
		assert.Equal(t, resp[0].Name, "John Doe")

		t.Cleanup(func() {
			testDB.Exec("DELETE FROM students")
		})
	})

	t.Run("DeleteStudent", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/v1/student/12", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("UpdateStudent", func(t *testing.T) {
		student := models.Student{Name: "John Doe", RollNo: 21, ID: 13}

		res := testDB.Create(&student)
		if res.Error != nil {
			fmt.Print("ERROR: Student Creation failed", res.Error)
		}
		student.Name = "Dohn Joe"
		body, _ := json.Marshal(&student)

		// Create a request
		req, _ := http.NewRequest("PUT", "/v1/student/13", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code) // Expect 201 Created
		var resp models.Student
		json.Unmarshal(w.Body.Bytes(), &resp)

		var updateStudent models.Student
		err := testDB.First(&updateStudent, "id = ?", 13).Error
		assert.NoError(t, err)
		assert.Equal(t, 21, updateStudent.RollNo)
	})

}
