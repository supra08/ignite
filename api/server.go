package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/weaveworks/ignite/cmd/ignite/run"
	"os/exec"
)

func psCmd(c *gin.Context) {
	out, err := exec.Command("./ignite", "ps", "-a").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	c.JSON(200, gin.H{
		"success": true,
		"data":    string(out),
	})
}

func runCmd(c *gin.Context) {
	var data map[string]interface{}
	c.BindJSON(&data)
	args := []string{"run", data["image"].(string), "--cpus", data["cpus"].(string), "--memory", data["memory"].(string) + "GB", "--ssh", "--name", data["name"].(string)}
	go exec.Command("./ignite", args...).Output()
	c.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func rmCmd(c *gin.Context) {
	var data map[string]interface{}
	c.BindJSON(&data)
	go exec.Command("./ignite", "rm", data["name"].(string)).Output()
	c.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func stopCmd(c *gin.Context) {
	var data map[string]interface{}
	c.BindJSON(&data)
	go exec.Command("./ignite", "stop", data["name"].(string)).Output()
	c.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func startCmd(c *gin.Context) {
	var data map[string]interface{}
	c.BindJSON(&data)
	go exec.Command("./ignite", "start", data["name"].(string)).Output()
	c.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ps", psCmd)
		v1.POST("/run", runCmd)
		v1.POST("/rm", rmCmd)
		v1.POST("/stop", stopCmd)
		v1.POST("/start", startCmd)
	}
	router.Run()
}
