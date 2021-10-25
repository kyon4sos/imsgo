package online

import ("im/engine"
"sync")

var onlineUser sync.Map

type User struct {
	Id string
	Client *engine.Client
}


func GetAllUser() *sync.Map {
	return &onlineUser
}
func GetClientById(id string) *engine.Client {
	c,ok:= onlineUser.Load(id)
	if ok {
		return c.(*engine.Client)
	}
	return nil
}
func SetOnlineUser(user *User) {
	onlineUser.Store(user.Id,user.Client)
}
func SendById(id string,content string)  {
	client := GetClientById(id)
	if client !=nil{
		client.Send<-[]byte(content)
	}
}
//func RemoveOnlineUser(id string)  {
//	delete(online,id)
//}