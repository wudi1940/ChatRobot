package client

import "github.com/gorilla/websocket"

type UserClient struct {
	Conn         *websocket.Conn // 用户websocket连接
	Uid          string          // 用户名称
	AzureClient  *azureClientStruct
	OpenAIClient *openAIClientStruct
	SendChan     chan *Message
	RespChan     chan *RespMessage
}

type ChatMessage struct {
	MessageType       int32  `json:"message_type,omitempty"` // 1表示text 2表示语音
	Message           string `json:"message,omitempty"`
	MessageId         string `json:"message_id,omitempty"`
	Content           string `json:"content,omitempty"`
	LanguageSelection string `json:"language"`
}

const (
	MessageType_Text   = 1
	MessageType_Speech = 2
)

type RespMessage struct {
	RespType  uint32 `json:"resp_type,omitempty"` // 1表示自己 2表示对方
	Message   string `json:"message,omitempty"`
	MessageId string `json:"message_id,omitempty"`
}

const (
	RespType_Self = 1
	RespType_Ai   = 2
	RespType_Err  = 3
)

type Message struct {
	EventType byte   `json:"type"`       // 0表示有新消息到达
	Name      string `json:"name"`       // 用户名称
	Message   string `json:"message"`    // 消息内容
	MessageId string `json:"message_id"` // 消息id对应语音
}
