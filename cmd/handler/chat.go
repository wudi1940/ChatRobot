package handler

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/processer"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var userClients = make(map[string]*client.UserClient) // 用户组映射

var up = &websocket.Upgrader{
	// 校验请求来源，这里不校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chat(c *gin.Context) {
	// 检验http头中upgrader属性，若为websocket，则将http协议升级为websocket协议
	conn, err := up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级websocket错误", err)
		panic(err)
	}
	defer conn.Close()

	// 如果用户列表中没有该用户
	name := c.ClientIP()
	if userClients[name] == nil {
		fmt.Println("未找到连接")
		return
	}
	userClient := userClients[name]
	userClient.Conn = conn
	go processer.AIOutput(userClient)

	// 当函数返回时，将该用户加入退出通道，并断开用户连接
	defer func() {
		userClients[name] = nil
		userClient.Conn.Close()
	}()

	for {
		_, bytes, err := userClient.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		data := &client.ChatMessage{}
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)

		switch data.MessageType {
		case 1:
			processer.TextInput(userClient, data)
		case 2:
			// todo 语音识别
			// 文本发送
		default:
			fmt.Println("消息类型错误")
		}

	}

}
