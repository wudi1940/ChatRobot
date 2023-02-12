package processer

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/config"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func TextInput(user *client.UserClient, msg *client.ChatMessage) {

	// todo: askAi

	rpMsg := user.OpenAIClient.AskAI(msg.Message)
	msgId := user.Uid + "_" + strconv.FormatInt(time.Now().UnixMicro(), 10)

	respMsg := &client.RespMessage{
		RespType:  2,
		Message:   rpMsg,
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

	fileName := config.ProjectPath + "/docs/audio/" + msg.MessageId + ".wav"
	err = os.WriteFile(fileName, decodeString, 0666)
	if err != nil {
		fmt.Println("语音文件存储错误", err)
		return
	}

	// todo 语音识别
	rpMsg := user.AzureClient.SpeechToTextFromFile(fileName)

	respMsg := &client.RespMessage{
		RespType:  1,
		Message:   rpMsg,
		MessageId: msg.MessageId,
	}
	// 文本发送
	user.RespChan <- respMsg

	// todo 询问ai并输出
	aiMsg := user.OpenAIClient.AskAI(rpMsg)
	// ai 语音识别
	msgId := user.Uid + "_" + strconv.FormatInt(time.Now().UnixMicro(), 10)
	filePath := config.ProjectPath + "/docs/audio/" + msgId + ".wav"
	user.AzureClient.TextToSpeech(aiMsg, filePath)

	respAIMsg := &client.RespMessage{
		RespType:  2,
		Message:   aiMsg,
		MessageId: msgId,
	}

	user.RespChan <- respAIMsg

}
