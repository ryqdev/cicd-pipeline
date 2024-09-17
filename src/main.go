package main

import (
	"io"
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
		// Read Headers
		headers := make(map[string]string)
		for key, values := range c.Request.Header {
			headers[key] = values[0] // assuming single value headers
		}

		// Log headers
		golog.Info("Headers: ", headers)

		// Read Body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			golog.Error("Error reading body: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read request body"})
			return
		}

		// Log raw body
		golog.Info("Raw Body: ", string(bodyBytes))

		// Respond with JSON including headers and parsed body
		c.JSON(http.StatusOK, gin.H{
			"header": headers,
			"body":   bodyBytes,
		})
	})
	r.Run()
}
