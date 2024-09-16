package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryqdev/golog"
)

func main() {
	r := gin.Default()
	golog.ShowDetail(true)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello CICD",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook", func(c *gin.Context) {
		// get and print header and body
		header := c.Request.Header
		body := c.Request.Body
		golog.Info("header: ", header)
		golog.Info("body: ", body)
		c.JSON(http.StatusOK, gin.H{
			"header": header,
			"body":   body,
		})
	})
	r.Run()
}
