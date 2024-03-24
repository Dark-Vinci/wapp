package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	gin.ForceConsoleColor()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     "tomiwa",
			"response": "200",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Something bad is about to happen")
	}

	fmt.Printf("THE ACCOUNT SERVICE")
}
