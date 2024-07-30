package consts

type MessageType int

const (
	TextMessage   MessageType = 1
	BinaryMessage             = 2
	CloseMessage              = 8
	PingMessage               = 9
	PongMessage               = 10
)
