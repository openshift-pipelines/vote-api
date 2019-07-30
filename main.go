package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var inMemoryStore = make(map[string]int)
var redirectURL = "http://0.0.0.0:9000"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/vote", func(c *gin.Context) {
		payload := gin.H{}
		for k, v := range inMemoryStore {
			payload[k] = v
		}
		c.JSON(http.StatusOK, payload)
	})

	r.POST("/vote", func(c *gin.Context) {
		name := c.PostForm("vote")
		inMemoryStore[name]++
		c.Redirect(301, redirectURL)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
