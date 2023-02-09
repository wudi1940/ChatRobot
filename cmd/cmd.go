package main

import (
	"ChatRobot/cmd/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("frontend/templates/*")

	v1 := r.Group("v1")
	{
		v1.POST("/input/text", handler.InputText)
	}

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/websocket", handler.AudioRecord)

	r.POST("/keys", handler.GetKeys)

	r.POST("/form", func(context *gin.Context) {
		context.HTML(http.StatusOK, "record.html", gin.H{})

		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	//i := &processer.AudioInput{}
	//i.Record()
}
