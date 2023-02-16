package handler

import (
	"ChatRobot/cmd/client"
	"ChatRobot/cmd/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeysRequest struct {
	OpenAIKey   string `json:"open_ai_key,omitempty"`
	AzureKey    string `json:"azure_key,omitempty"`
	AzureRegion string `json:"azure_region"`
	Uid         string `json:"uid"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetKeys(c *gin.Context) {
	req := KeysRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code: 3,
			Msg:  "bad request param",
			Data: req,
		})
		return
	}

	userClient := &client.UserClient{}
	name := req.Uid
	userClient.Uid = req.Uid
	userClient.AzureClient = client.InitAzureClient(config.AzureClientKey, config.AzureClientRegion) // fixme 替换
	if req.OpenAIKey == "" {
		userClient.OpenAIClient = client.InitOpenAIClient(config.OpenAIClientKey)
	} else {
		userClient.OpenAIClient = client.InitOpenAIClient(req.OpenAIKey)
	}
	userClient.RespChan = make(chan *client.RespMessage, 3)
	userClient.SendChan = make(chan *client.Message, 3)
	// 如果用户列表中没有该用户
	userClients.LoadOrStore(name, userClient)

	c.JSON(http.StatusOK, Response{
		Code: 1,
		Msg:  "success",
		Data: req,
	})
}
