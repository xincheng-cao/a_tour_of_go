package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		reference:
		https://github.com/gin-gonic/gin
	*/

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:14000")
}
