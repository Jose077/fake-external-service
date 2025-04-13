package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/leads", func(c *gin.Context) {
		if rand.Intn(2) == 0 {
			c.JSON(http.StatusCreated, gin.H{"message": "Lead created"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Bad gateway"})
		}
	})

	r.Run(":8080")
}
