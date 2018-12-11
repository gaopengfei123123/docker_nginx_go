package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("hook", func(c *gin.Context) {
		var jsonp map[string]interface{}

		if c.BindJSON(&jsonp) == nil {
			saveToFile(jsonp)
			c.JSON(http.StatusOK, jsonp)
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "false"})
		}
	})

	r.Run(":9090")
}

func saveToFile(jsonp map[string]interface{}) {
	jsonByte, err := json.Marshal(jsonp)

	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("./hook.json", jsonByte, 0755)

	if err != nil {
		fmt.Println("error:", err)
	}
}
