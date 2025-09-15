// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default() // create gin router with default middlewares(logger+recovery)
// 	r.GET("/get", getValues)
// 	r.Run(":8080")
// }

// var url = "https://api.agify.io?name=meelad"

// // Struct to hold the response from the API
// type AgifyResponse struct {
// 	Count int    `json:"count"`
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// }

// func getValues(c *gin.Context) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "something went wrong",
// 		})
// 		return
// 	}
// 	defer resp.Body.Close()//prevent from data breach

// 	// Read response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "error reading response",
// 		})
// 		return
// 	}

// 	// Unmarshal JSON
// 	var data AgifyResponse
// 	if err := json.Unmarshal(body, &data); err != nil {//byte of array
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// Send JSON response to client
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": data,
// 	})
// }

// use map instead of strct
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", getValues)
	r.Run(":8080")
}

var url = "https://api.agify.io?name=meelad"

func getValues(c *gin.Context) {
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error reading response",
		})
		return
	}

	// Use map instead of struct
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error parsing json",
		})
		return
	}

	// Send JSON response back
	c.JSON(http.StatusOK, data)
}
