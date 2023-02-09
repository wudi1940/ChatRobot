package client

import (
	"fmt"
	"testing"
)

func TestOpenAIClientStruct_AskAI(t *testing.T) {

	c := InitOpenAIClient("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")
	res := c.AskAI("你能说中文吗？")
	fmt.Println(res)

}
