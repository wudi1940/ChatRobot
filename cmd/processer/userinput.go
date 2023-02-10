package processer

import (
	"ChatRobot/cmd/client"
)

func TextInput(user *client.UserClient, msg *client.ChatMessage) {

	respMsg := &client.RespMessage{
		RespType:  2,
		Message:   "这个是自动回复," + msg.Message,
		MessageId: "",
	}

	user.RespChan <- respMsg
}
