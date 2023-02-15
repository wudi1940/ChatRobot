package main

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/handler"
	"net/http"
	"sync"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	Init()

	r := gin.Default()
	r.Use(client.LoggerToFile())

	r.LoadHTMLGlob("frontend/templates/*")
	r.Use(static.ServeRoot("/", "frontend/templates"))
	r.StaticFS("docs", http.Dir("docs"))

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/keys", handler.GetKeys)

	r.GET("/chat", handler.Chat)
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}

func Init() {
	var once sync.Once

	once.Do(client.InitLogClient)
}
