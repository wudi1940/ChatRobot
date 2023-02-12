package client

import (
	"ChatRobot/cmd/config"
	"fmt"
	"testing"
)

func TestAzureClientStruct_TextToSpeech(t *testing.T) {
	InitAzureClient("9e41080c590946229ec766b6d9ea6a6c", "japanwest")
	//client := GetAzureClient()

	//client.TextToSpeech("Hello World, This is a test message for Synthesizing Audio")

}

func TestAzureClientStruct_SpeechToTextFromFile(t *testing.T) {
	InitAzureClient("9e41080c590946229ec766b6d9ea6a6c", "japanwest")
	client := GetAzureClient()

	filePath := config.ProjectPath + "/docs/audio/chinese.wav"
	text := client.SpeechToTextFromFile(filePath)
	fmt.Println(text)
}

func TestIsChinese(t *testing.T) {
	text := "zh-CN-YunxiNeural,语音文件存储错误"
	res := IsChinese(text)
	fmt.Println(res)
}
