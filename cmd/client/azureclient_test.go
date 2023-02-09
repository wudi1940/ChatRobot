package client

import "testing"

func TestAzureClientStruct_TextToSpeech(t *testing.T) {
	InitAzureClient("9e41080c590946229ec766b6d9ea6a6c", "japanwest")
	client := GetAzureClient()

	client.TextToSpeech("Hello World, This is a test message for Synthesizing Audio")

}
