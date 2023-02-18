package client

import (
	"context"
	"errors"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var openAIClient *openAIClientStruct

type openAIClientStruct struct {
	Key string
}

func GetOpenAIClient() *openAIClientStruct {
	if openAIClient == nil {
		panic(errors.New("openAIClient未初始化"))
	}
	return openAIClient
}

func InitOpenAIClient(key string) *openAIClientStruct {
	openAIClient = &openAIClientStruct{
		Key: key,
	}
	return openAIClient
}

func (r *openAIClientStruct) AskAI(prompt string) (string, error) {
	c := gogpt.NewClient(r.Key)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        2500,
		Prompt:           prompt,
		Temperature:      0.9,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
		Stop:             []string{"HUMAN:", "AI:"},
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Text, nil
}
