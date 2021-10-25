package model

type TextMessage struct {
	*Message
	text string
}

func TextMsgBuilder(text string) *TextMessage {
	return &TextMessage{
		text: text,
		Message:&Message{
			Type: "MSG_TEXT",
		},
	}
}