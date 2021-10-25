package dto

import "im/model/payload"

type Message struct {
	From string
	To string
	ConversationType string
	Type string
	Payload payload.Payload
}