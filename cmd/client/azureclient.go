package client

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
)

var azureClient *azureClientStruct

func GetAzureClient() *azureClientStruct {
	if azureClient == nil {
		panic(errors.New("azureClient未初始化"))
	}
	return azureClient
}

func InitAzureClient(key, region string) *azureClientStruct {
	azureClient = &azureClientStruct{
		Key:    key,
		Region: region,
	}
	return azureClient
}

// 9e41080c590946229ec766b6d9ea6a6c
// japanwest
type azureClientStruct struct {
	Key    string
	Region string
}

func (c *azureClientStruct) SpeechToTextFromFile(filePath string) string {

	speechKey := c.Key
	speechRegion := c.Region

	audioConfig, err := audio.NewAudioConfigFromWavFileInput(filePath)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return ""
	}
	defer audioConfig.Close()
	config, err := speech.NewSpeechConfigFromSubscription(speechKey, speechRegion)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return ""
	}
	defer config.Close()
	speechRecognizer, err := speech.NewSpeechRecognizerFromConfig(config, audioConfig)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return ""
	}
	defer speechRecognizer.Close()
	speechRecognizer.SessionStarted(func(event speech.SessionEventArgs) {
		defer event.Close()
		fmt.Println("Session Started (ID=", event.SessionID, ")")
	})
	speechRecognizer.SessionStopped(func(event speech.SessionEventArgs) {
		defer event.Close()
		fmt.Println("Session Stopped (ID=", event.SessionID, ")")
	})

	task := speechRecognizer.RecognizeOnceAsync()
	var outcome speech.SpeechRecognitionOutcome
	select {
	case outcome = <-task:
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out")
		return ""
	}
	defer outcome.Close()
	if outcome.Error != nil {
		fmt.Println("Got an error: ", outcome.Error)
	}
	fmt.Println("Got a recognition!")
	return outcome.Result.Text
}

func (c *azureClientStruct) TextToSpeech(text, filePath string) {
	speechKey := c.Key
	speechRegion := c.Region

	audioConfig, err := audio.NewAudioConfigFromDefaultSpeakerOutput()
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer audioConfig.Close()
	speechConfig, err := speech.NewSpeechConfigFromSubscription(speechKey, speechRegion)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer speechConfig.Close()

	speechConfig.SetSpeechSynthesisVoiceName("en-US-JennyNeural")
	//speechConfig.SetSpeechSynthesisVoiceName("zh-CN-YunxiNeural")

	speechSynthesizer, err := speech.NewSpeechSynthesizerFromConfig(speechConfig, audioConfig)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer speechSynthesizer.Close()

	speechSynthesizer.SynthesisStarted(synthesizeStartedHandler)
	speechSynthesizer.Synthesizing(synthesizingHandler)
	speechSynthesizer.SynthesisCompleted(synthesizedHandler)
	speechSynthesizer.SynthesisCanceled(cancelledSynthesisHandler)

	//
	if len(text) == 0 {
		return
	}

	task := speechSynthesizer.SpeakTextAsync(text)
	var outcome speech.SpeechSynthesisOutcome
	select {
	case outcome = <-task:
	case <-time.After(60 * time.Second):
		fmt.Println("Timed out")
		return
	}
	defer outcome.Close()
	if outcome.Error != nil {
		fmt.Println("Got an error: ", outcome.Error)
		return
	}

	if outcome.Result.Reason == common.SynthesizingAudioCompleted {
		fmt.Printf("Speech synthesized to speaker for text [%s].\n", text)

		err = os.WriteFile(filePath, outcome.Result.AudioData, 0666)
		if err != nil {
			fmt.Println("语音文件存储错误", err)
		}

	} else {
		cancellation, _ := speech.NewCancellationDetailsFromSpeechSynthesisResult(outcome.Result)
		fmt.Printf("CANCELED: Reason=%d.\n", cancellation.Reason)

		if cancellation.Reason == common.Error {
			fmt.Printf("CANCELED: ErrorCode=%d\nCANCELED: ErrorDetails=[%s]\nCANCELED: Did you set the speech resource key and region values?\n",
				cancellation.ErrorCode,
				cancellation.ErrorDetails)
		}
	}

}

func synthesizeStartedHandler(event speech.SpeechSynthesisEventArgs) {
	defer event.Close()
	fmt.Println("Synthesis started.")
}

func synthesizingHandler(event speech.SpeechSynthesisEventArgs) {
	defer event.Close()
	fmt.Printf("Synthesizing, audio chunk size %d.\n", len(event.Result.AudioData))
}

func synthesizedHandler(event speech.SpeechSynthesisEventArgs) {
	defer event.Close()
	fmt.Printf("Synthesized, audio length %d.\n", len(event.Result.AudioData))
}

func cancelledSynthesisHandler(event speech.SpeechSynthesisEventArgs) {
	defer event.Close()
	fmt.Println("Received a cancellation.")
}
