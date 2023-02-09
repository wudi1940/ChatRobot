package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InputTextRequest struct {
	Text string `json:"text,omitempty"`
}

func InputText(c *gin.Context) {
	req := InputTextRequest{}
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
