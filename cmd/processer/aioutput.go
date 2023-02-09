package processer

import (
	"ChatRobot/cmd/client"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func AIOutput(client *client.UserClient) {
	for {
		select {
		case msg := <-client.RespChan:
			data, err := json.Marshal(msg)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = client.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
