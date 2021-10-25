package handler

import (
	"im/command"
	"im/engine"
	"im/online"
	"log"
)

type GroupChat struct {
	command.Command
}

//func (g *GroupChat) Save()  {
//	id := uuid.NewString()
//	db.GetDb().Model()
//}

func (g *GroupChat) ChannelRead(ctx *engine.Context) {
	log.Printf("group chat")
	allUser := online.GetAllUser()
	allUser.Range(func(key, value interface{}) bool {
		client := value.(*engine.Client)
		client.Send <- ctx.Body
		//if user,ok := value.(online.OnlineUser);ok {
		//	user.Client.SendById<-ctx.Body
		//}
		return true
	})
}


