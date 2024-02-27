package domain

type MessageSender interface {
	Send(content string, contact string)
}
