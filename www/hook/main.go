package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {
	// execCommand()
	serverStart()
}

func execCommand() {
	cmd := exec.Command("/bin/bash", "git_pull.sh")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())

}

func serverStart() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("hook", func(c *gin.Context) {
		var jsonp map[string]interface{}

		var gitStruct GitHubWebHook

		err := c.BindJSON(&gitStruct)
		if err == nil {
			resp := gitStruct.Repository.GitURL
			fmt.Printf("gitUrl: %s \n", resp)
		}

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

// GitHubWebHook github的json结构
type GitHubWebHook struct {
	Zen    string
	HookID int `json:"hook_id"`
	Hook   struct {
		Type   string
		ID     int
		Name   string
		Active bool
		Events []string
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret, URL string
		}
		UpdatedAt             string `json:"updated_at"`
		CreatedAt             string `json:"created_at"`
		URL, TestURL, PingURL string
		LastResponse          struct {
			Code    interface{}
			Status  string
			Message interface{}
		}
	}
	Repository struct {
		ID                     int
		Name, FullName, NodeID string
		Private                bool
		Owner                  map[string]interface{}
		GitURL                 string `json:"git_url"`
		SSHURL                 string `json:"ssh_url"`
		CloneURL               string `json:"clone_url"`
	}
	Sender map[string]interface{}
}
