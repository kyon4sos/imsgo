package engine
//
//import (
//	"im/online"
//)
//
//type Hub struct {
//	Clients    map[string]*Client
//	Register   chan *online.OnlineUser
//	Unregister chan *online.OnlineUser
//	broadcast  chan []byte
//}
//
//func NewHub() *Hub {
//	return &Hub{
//		Clients:    make(map[string]*Client),
//		Register:   make(chan *online.OnlineUser),
//		Unregister: make(chan *online.OnlineUser),
//		broadcast:  make(chan []byte),
//	}
//}
//
////func (h *Hub) Run() {
////	for {
////		select {
////		case online := <-h.Register:
////			h.Clients[online.UUID] = online.Client
////		case online := <-h.Unregister:
////			if _, ok := h.Clients[online.UUID]; ok {
////				close(online.Client.Send)
////				delete(h.Clients, online.UUID)
////			}
////		case msg := <-h.broadcast:
////			log.Printf("channelHandlers %v\n", string(msg))
////			var message dto.Message
////			for idx, ch := range channelHandlers {
////				log.Println("channelHandlers idx",idx)
////				ch.ChannelRead(nil,msg)
////			}
////			json.Unmarshal(msg,&message)
////			log.Printf("message unmarshal  %v\n", message)
////			//if message.Type == "text" {
////			//		json.Unmarshal()
////			//}
////			//for client := range h.Clients {
////			//	select {
////			//	case client.Send <- msg:
////			//	default:
////			//		close(client.Send)
////			//		delete(h.Clients, client)
////			//	}
////			//}
////		}
////	}
////}
