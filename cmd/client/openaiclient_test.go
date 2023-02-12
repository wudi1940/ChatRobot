package client

import (
	"fmt"
	"testing"
)

func TestOpenAIClientStruct_AskAI(t *testing.T) {

	c := InitOpenAIClient("")
	res := c.AskAI("你能说中文吗？")
	fmt.Println(res)

}
