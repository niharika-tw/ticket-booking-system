package app

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {

	r := gin.Default()
	r.POST("/book", func(c *gin.Context) {
		
		c.JSON(201, gin.H{
			"message": "Ticket Created",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server started",
		})
	})
	go r.Run()
	go r.Run()
}
