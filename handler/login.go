package handler

import (
	"im/engine"
	"log"
)

type Login struct {

}

func (login *Login) ChannelRead(ctx *engine.Context) {
	log.Println("login")
	//ctx.Cancel()
}