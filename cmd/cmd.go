package main

import (
	"ChatRobot/cmd/handler"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("frontend/templates/*")
	r.Use(static.ServeRoot("/", "frontend/templates"))
	r.StaticFS("docs", http.Dir("docs"))

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/keys", handler.GetKeys)

	r.GET("/chat", handler.Chat) //r.POST("/chat", handler.Core)

	//r.GET("/websocket", handler.AudioRecord)

	r.Run(":9999") // listen and serve on 0.0.0.0:8080

	//i := &processer.AudioInput{}
	//i.Record()

	//v1 := r.Group("v1")
	//{
	//	v1.POST("/input/text", handler.InputText)
	//}
}
