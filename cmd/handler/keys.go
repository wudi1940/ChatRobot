package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type KeysRequest struct {
	OpenAIKey string `json:"open_ai_key,omitempty"`
	AzureKey  string `json:"azure_key,omitempty"`
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

	c.JSON(http.StatusOK, Response{
		Code: 1,
		Msg:  "success",
		Data: req,
	})
}
