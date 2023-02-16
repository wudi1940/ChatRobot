package handler

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func AudioRecord(c *gin.Context) {
	client, err := up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级websocket错误", err)
		panic(err)
	}
	defer client.Close()

	// 接受文件
	_, content, err := client.ReadMessage()
	if err != nil {
		fmt.Println("消息接受错误", err)
	} else {
		fmt.Println("收到文件")
		err = os.WriteFile("docs/audio/temp.wav", content, 0666)
		if err != nil {
			fmt.Println("语音文件存储错误", err)
		}
	}
}
