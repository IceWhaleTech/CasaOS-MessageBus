package service

type MessageBus struct{}

func NewMessageBus() *MessageBus {
	return &MessageBus{}
}
