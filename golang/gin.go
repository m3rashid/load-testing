package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

type HealthCheck struct {
	uptime    string
	message   string
	timestamp string
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		healthCheck := HealthCheck{
			uptime:    uptime().String(),
			message:   "OK",
			timestamp: time.Now().Format(time.RFC3339),
		}

		fmt.Printf("HealthCheck: %+v", healthCheck)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":4000")
}
