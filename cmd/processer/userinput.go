package processer

import (
	"ChatRobot/cmd/client"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func TextInput(user *client.UserClient, msg *client.ChatMessage) {

	// todo: askAi

	// rpMsg := user.OpenAIClient.AskAI(msg.Message)
	msgId := user.Uid + "_" + strconv.FormatInt(time.Now().UnixMicro(), 10)

	respMsg := &client.RespMessage{
		RespType:  2,
		Message:   "这个是自动回复," + msg.Message,
		MessageId: msgId,
	}

	user.RespChan <- respMsg
}

func SpeechInput(user *client.UserClient, msg *client.ChatMessage) {
	content := msg.Content
	index := strings.Index(content, ",")
	decodeString, err := base64.StdEncoding.DecodeString(content[index+1:])
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fileName := "docs/audio/" + msg.MessageId + ".wav"
	err = os.WriteFile(fileName, decodeString, 0666)
	if err != nil {
		fmt.Println("语音文件存储错误", err)
	}

}
