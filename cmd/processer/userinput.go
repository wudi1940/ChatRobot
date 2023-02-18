package processer

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func UserIutput(user *client.UserClient) {
	logger := client.GetLogClient()
	logger.WithField("uid", user.Uid)
	for {
		_, bytes, err := user.Conn.ReadMessage()
		if err != nil {
			logger.WithError(err).Error("接受用户ws数据错误")
			return
		}

		data := &client.ChatMessage{}
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			logger.WithError(err).Error("用户ws数据转json失败")
			return
		}

		logger.Debug(data)

		switch data.MessageType {
		case client.MessageType_Text:
			textInput(user, data)
		case client.MessageType_Speech:
			speechInput(user, data)
		default:
			fmt.Println("消息类型错误")
		}

	}
}

func textInput(user *client.UserClient, msg *client.ChatMessage) {
	logger := client.GetLogClient()
	logger.Info(msg)

	var respMsg = &client.RespMessage{}
	msgId := user.Uid + "_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	respMsg.MessageId = msgId
	respMsg.RespType = client.RespType_Err
	logger.WithField("uid", user.Uid).WithField("msgId", msgId)

	rpMsg, err := user.OpenAIClient.AskAI(msg.Message)
	if err != nil {
		logger.WithError(err).Error("ai回复错误")
		respMsg.Message = "ai回复错误"
		user.RespChan <- respMsg
		return
	}

	filePath := config.ProjectPath + "/docs/audio/" + msgId + ".wav"
	err = user.AzureClient.TextToSpeech(rpMsg, filePath, msg.LanguageSelection)
	if err != nil {
		logger.WithError(err).Error("text to speech err")
		respMsg.Message = "speech to text err" + err.Error()
		user.RespChan <- respMsg
		return
	}

	respMsg.RespType = client.RespType_Ai
	respMsg.Message = rpMsg

	user.RespChan <- respMsg
}

func speechInput(user *client.UserClient, msg *client.ChatMessage) {
	logger := client.GetLogClient()

	var respMsg = &client.RespMessage{}
	respMsg.MessageId = msg.MessageId
	respMsg.RespType = client.RespType_Err
	logger.WithField("uid", user.Uid).WithField("msgId", msg.MessageId)
	// 语音输入保存成文件
	content := msg.Content
	index := strings.Index(content, ",")
	decodeString, err := base64.StdEncoding.DecodeString(content[index+1:])
	if err != nil {
		logger.WithError(err).Error("decode error")
		respMsg.Message = "decode error"
		user.RespChan <- respMsg
		return
	}

	fileName := config.ProjectPath + "/docs/audio/" + msg.MessageId + ".wav"
	err = os.WriteFile(fileName, decodeString, 0666)
	if err != nil {
		logger.WithError(err).Error("语音文件存储错误")
		respMsg.Message = "语音文件存储错误"
		user.RespChan <- respMsg
		return
	}

	// 语音识别
	rpMsg, err := user.AzureClient.SpeechToTextFromFile(fileName, msg.LanguageSelection)
	if err != nil {
		logger.WithError(err).Error("speech to text err")
		respMsg.RespType = client.RespType_Err
		respMsg.Message = "speech to text err" + err.Error()
		user.RespChan <- respMsg
		return
	}
	respMsg.RespType = client.RespType_Self
	respMsg.Message = rpMsg
	// 语音转文本回显
	user.RespChan <- respMsg

	logger.Info(respMsg)

	// 询问ai并输出
	aiMsg, err := user.OpenAIClient.AskAI(rpMsg)
	if err != nil {
		logger.WithError(err).Error("ai回复错误")
		respMsg.RespType = client.RespType_Err
		respMsg.Message = "ai回复错误"
		user.RespChan <- respMsg
		return
	}

	// ai语音合成：text转语音
	var respAIMsg = &client.RespMessage{}
	msgId := user.Uid + "_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	respAIMsg.RespType = client.RespType_Err
	respAIMsg.MessageId = msgId
	filePath := config.ProjectPath + "/docs/audio/" + msgId + ".wav"
	err = user.AzureClient.TextToSpeech(aiMsg, filePath, msg.LanguageSelection)
	if err != nil {
		logger.WithError(err).Error("text to speech err")
		respMsg.Message = "speech to text err" + err.Error()
		user.RespChan <- respMsg
		return
	}

	respAIMsg.RespType = client.RespType_Ai
	respAIMsg.Message = aiMsg

	user.RespChan <- respAIMsg
}
