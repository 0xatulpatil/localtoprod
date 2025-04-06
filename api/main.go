package main

import (
	"fmt"
	"net/http"
	"sre-goapi/config"
	"sre-goapi/controllers"
	"sre-goapi/db"
	"sre-goapi/handlers"
	"sre-goapi/models"
	"sre-goapi/routes"
	logger "sre-goapi/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	appDB := db.GetAppDB()
	router := gin.Default()

	err := appDB.AutoMigrate(&models.Student{})
	if err != nil {
		logger.Error("Database Migration Failed")
		fmt.Println("ERROR: Error migrating database")
	}

	studentController := controllers.NewStudentController()
	studentHandler := handlers.NewStudentHandler(studentController)

	routes.RegisterRoutes(router, studentHandler)
	router.GET("/healthcheck", func(c *gin.Context) {
		logger.Info("/healthcheck")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":" + config.Port)
	logger.Info("Server up and running")
}
