package handler

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/processer"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var userClients sync.Map // 用户组映射

var up = &websocket.Upgrader{
	// 校验请求来源，这里不校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chat(c *gin.Context) {

	logger := client.GetLogClient()

	// 检验http头中upgrader属性，若为websocket，则将http协议升级为websocket协议
	conn, err := up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.WithError(err).Error("升级websocket错误")
	}

	// 如果用户列表中没有该用户
	name, _ := c.GetQuery("uid")
	user, ok := userClients.Load(name)
	if !ok {
		logger.Error("用户列表中没有该用户")
		return
	}
	userClient := user.(*client.UserClient)
	userClient.Conn = conn
	go processer.AIOutput(userClient)
	processer.UserIutput(userClient)

	// 当函数返回时，将该用户加入退出通道，并断开用户连接
	defer func() {
		userClients.Delete(name)
		userClient.Conn.Close()
	}()

}
