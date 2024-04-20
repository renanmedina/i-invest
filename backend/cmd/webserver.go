package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	webserver := gin.Default()

	webserver.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	webserver.Run() // listen and serve on 0.0.0.0:8080
}
