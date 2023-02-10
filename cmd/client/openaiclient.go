package client

import (
	"context"
	"errors"
	"fmt"

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

func (r *openAIClientStruct) AskAI(prompt string) string {
	c := gogpt.NewClient(r.Key)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 5,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return resp.Choices[0].Text
}
