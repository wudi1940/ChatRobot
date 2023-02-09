package processer

//
//func TextToSpeech() {
//	// This example requires environment variables named "SPEECH_KEY" and "SPEECH_REGION"
//	speechKey := os.Getenv("SPEECH_KEY")
//	speechRegion := os.Getenv("SPEECH_REGION")
//
//	audioConfig, err := audio.NewAudioConfigFromDefaultSpeakerOutput()
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer audioConfig.Close()
//	speechConfig, err := speech.NewSpeechConfigFromSubscription(speechKey, speechRegion)
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer speechConfig.Close()
//
//	speechConfig.SetSpeechSynthesisVoiceName("en-US-JennyNeural")
//
//	speechSynthesizer, err := speech.NewSpeechSynthesizerFromConfig(speechConfig, audioConfig)
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer speechSynthesizer.Close()
//
//	speechSynthesizer.SynthesisStarted(synthesizeStartedHandler)
//	speechSynthesizer.Synthesizing(synthesizingHandler)
//	speechSynthesizer.SynthesisCompleted(synthesizedHandler)
//	speechSynthesizer.SynthesisCanceled(cancelledHandler)
//
//	for {
//		fmt.Printf("Enter some text that you want to speak, or enter empty text to exit.\n> ")
//		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//		text = strings.TrimSuffix(text, "\n")
//		if len(text) == 0 {
//			break
//		}
//
//		task := speechSynthesizer.SpeakTextAsync(text)
//		var outcome speech.SpeechSynthesisOutcome
//		select {
//		case outcome = <-task:
//		case <-time.After(60 * time.Second):
//			fmt.Println("Timed out")
//			return
//		}
//		defer outcome.Close()
//		if outcome.Error != nil {
//			fmt.Println("Got an error: ", outcome.Error)
//			return
//		}
//
//		if outcome.Result.Reason == common.SynthesizingAudioCompleted {
//			fmt.Printf("Speech synthesized to speaker for text [%s].\n", text)
//		} else {
//			cancellation, _ := speech.NewCancellationDetailsFromSpeechSynthesisResult(outcome.Result)
//			fmt.Printf("CANCELED: Reason=%d.\n", cancellation.Reason)
//
//			if cancellation.Reason == common.Error {
//				fmt.Printf("CANCELED: ErrorCode=%d\nCANCELED: ErrorDetails=[%s]\nCANCELED: Did you set the speech resource key and region values?\n",
//					cancellation.ErrorCode,
//					cancellation.ErrorDetails)
//			}
//		}
//	}
//}
//
//func synthesizeStartedHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("Synthesis started.")
//}
//
//func synthesizingHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Printf("Synthesizing, audio chunk size %d.\n", len(event.Result.AudioData))
//}
//
//func synthesizedHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Printf("Synthesized, audio length %d.\n", len(event.Result.AudioData))
//}
//
//func cancelledHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("Received a cancellation.")
//}
