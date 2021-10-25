package handler

import (
	"encoding/json"
	"fmt"
	"im/engine"
	"im/online"
	"log"
	"time"
)

type SingleChat struct {
	Id string
	Payload string
	To string
	From string
}

func (sc *SingleChat) ChannelRead(ctx engine.Context)  {
	err := json.Unmarshal(ctx.Body, sc)
	if err!=nil {
		return
	}
	log.Printf("single chat %v \n" ,sc)
	if err!=nil {
		log.Printf("err %v \n",err.Error())
		return
	}
	content:=fmt.Sprintf("hello : %v",time.Now())
	//ctx.SendById([]byte(content))

	online.SendById(sc.From,content)
}