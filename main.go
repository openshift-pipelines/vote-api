package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/vote", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"votes": 100,
		})
	})

	r.POST("/vote", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"votes": 100,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
