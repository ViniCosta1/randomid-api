package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"randomid-api/lib" //function to generate random id

	"os"
)

func main() {
	// loading dotenv
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// production mode
	if (os.Getenv("GIN_MODE") == "release") {
		gin.SetMode(gin.ReleaseMode)
	}

	// endpoints
	router := gin.Default()
	router.GET("/id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": lib.GenerateRandomId(),
		})
	})

	router.Run(":"+os.Getenv("PORT"));
}