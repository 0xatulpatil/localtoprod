package main

import (
	"fmt"
	"net/http"
	"sre-goapi/config"
	"sre-goapi/db"
	"sre-goapi/models"
	"sre-goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	appDB := db.GetAppDB()
	router := gin.Default()

	err := appDB.AutoMigrate(&models.Student{})
	if err != nil {
		fmt.Println("ERROR: Error migrating database")
	}

	routes.RegisterRoutes(router)
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":" + config.Port)
}
