package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()//Creates a Gin router with default middlewares (logger + recovery).
	r.GET("/ping", test) // <-- added /// for path /ping perfrom test function
	r.Run(":9090")
}

func test(c *gin.Context) {//handler function for path /ping
	c.JSON(201, gin.H{
		"message": "pong",
	})
}
