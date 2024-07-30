package types

import "davinci-chat/consts"

type Message struct {
	User        string          `json:"user"`
	Message     string          `json:"message"`
	Time        string          `json:"time"`
	MessageType int             `json:"messageType"`
	UserType    consts.UserType `json:"userType"`
}
