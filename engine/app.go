package engine

import (
	"log"
	"net/http"
)


type app struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Channel    chan  Context
}

var channelHandlers []ChannelHandler

type middleWareHandler func() func(ctx Context)

var middleWares  = make([]middleWareHandler,0)
func NewApp() *app {
	app := &app{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Channel:    make(chan Context),
	}
	return app
}
func (app *app) AddHandler(h ChannelHandler) *app {
	//once := sync.Once{}
	//once.Do(newChannel)
	if channelHandlers==nil {
		newChannel()
	}
	channelHandlers = append(channelHandlers, h)
	return app
}
func newChannel()  {
	channelHandlers =make([]ChannelHandler,0)
}

func (app *app) Run()  {
	for {
		select {
		case client := <-app.Register:
			log.Printf("register %v\n",client)
			app.Clients[client] = true
		case client := <-app.Unregister:
			log.Printf("unregister %v\n",client)
			if _, ok := app.Clients[client]; ok {
				close(client.Send)
				delete(NewApp().Clients, client)
			}
		case ctx:=<-app.Channel:
			dispatch(ctx)
			//log.Println("app msg %v\n",string(msg))

		}
	}
}
func (app *app) Use(fn middleWareHandler){
	middleWares = append(middleWares,fn)
}
func (app *app) handle(ctx Context){
	for i := range middleWares {
		handler := middleWares[i]
		handler()(ctx)
	}
}
func (app *app)open(func()) func(ctx Context) {
	return func(ctx Context) {}
}
func (app *app) NewServer(writer http.ResponseWriter, req *http.Request, header http.Header) *Client {
	conn, _ := upGrader.Upgrade(writer, req, header)
	log.Println("upgrade")
	//if err != nil {
	//	log.Println("upgrade err", err.Error())
	//	return
	//}
	client := &Client{
		Conn: conn,
		App:  app,
		Send: make(chan []byte, 256),
	}
	context := NewContext(client)
	//app.handle(context)
	select {
	case _ =<-context.Done():
		conn.Close()
		log.Println("before done")
		return client

	default:
		client.App.Register <- client
		go client.read()
		go client.write()
	}
	return client
}

