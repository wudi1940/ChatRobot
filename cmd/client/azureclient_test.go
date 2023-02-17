package client

import (
	"fmt"
	"testing"
)

func TestAzureClientStruct_TextToSpeech(t *testing.T) {
	InitAzureClient("", "")
	//client := GetAzureClient()

	//client.TextToSpeech("Hello World, This is a test message for Synthesizing Audio")

}

func TestAzureClientStruct_SpeechToTextFromFile(t *testing.T) {
	//InitAzureClient("", "")
	//client := GetAzureClient()
	//
	//filePath := config.ProjectPath + "/docs/audio/chinese.wav"
	//text := client.SpeechToTextFromFile(filePath)
	//fmt.Println(text)
}

func TestIsChinese(t *testing.T) {
	text := "zh-CN-YunxiNeural,语音文件存储错误"
	res := IsChinese(text)
	fmt.Println(res)
}

func TestDetectLanguage(t *testing.T) {
	language := DetectLanguage("这是啥")
	fmt.Println(language)
}
