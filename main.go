package main

import (
	"fmt"
	"nicolascastro/go/isMutant/controllers"
	"nicolascastro/go/isMutant/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := database.ConnectDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database!")
		return
	}

	r.POST("/mutant/", controllers.CheckMutant)
	r.GET("/stats/", controllers.GetStats)

	r.Run(":8080")
}
