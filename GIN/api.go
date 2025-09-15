package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var data api

func main() {
	r := gin.Default()                //create gin router with default middlewares(logger+recovery)
	r.GET("/get", getValues)          //get data from somewhre
	r.POST("/post", postValues)       //post data to somewhere
	r.PUT("/put", putValues)          //update prev data
	r.DELETE("/delete", deleteValues) //delete data
	r.Run(":8080")
}
func getValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func postValues(c *gin.Context) {
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "data added successfully",
		"data":    data,
	})
}

func putValues(c *gin.Context) {
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
