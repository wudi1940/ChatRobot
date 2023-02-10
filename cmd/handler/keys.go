package handler

import (
	"ChatRobot/cmd/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type KeysRequest struct {
	OpenAIKey   string `json:"open_ai_key,omitempty"`
	AzureKey    string `json:"azure_key,omitempty"`
	AzureRegion string `json:"azure_region"`
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
	name := c.ClientIP()
	userClient.Name = name
	userClient.AzureClient = client.InitAzureClient(req.AzureKey, req.AzureRegion)
	userClient.RespChan = make(chan *client.RespMessage, 3)
	userClient.SendChan = make(chan *client.Message, 3)
	// 如果用户列表中没有该用户
	if userClients[name] == nil {
		userClients[name] = userClient
	}

	c.JSON(http.StatusOK, Response{
		Code: 1,
		Msg:  "success",
		Data: req,
	})
}
