package engine

import (
	"context"
)

type Context struct {
	context.Context
	Cancel context.CancelFunc
	*Client
	Body []byte
}

func NewContext(client *Client) Context {
	ctx, cancel := context.WithCancel(context.Background())
	return Context{
		ctx,
		cancel,
		client,
		make([]byte,0),
	}
}

//func (ctx *Context) Body() []byte{
//	messageType, msg, err := ctx.Conn.ReadMessage()
//	log.Println("message type %v\n" ,messageType)
//	if err!=nil {
//		log.Println("readmessage err %v\n",err.Error())
//	}
//	return msg
//}
func (ctx *Context) Send(b []byte) {
	ctx.Client.Send <-b
}
func (ctx *Context) SetBody(b []byte) {
	ctx.Body =b
}