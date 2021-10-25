package engine

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Client struct {
	App  *app
	Conn *websocket.Conn
	Send chan []byte
	context.Context
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *Client) read() {
	defer func() {
		log.Println("read end")
		c.Conn.Close()
	}()
	//c.Conn.SetReadLimit(maxMsgSize)
	//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.Conn.SetPongHandler(func(appData string) error {
	//	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	//	return nil
	//})
	for {
		mt, msg, err := c.Conn.ReadMessage()

		//ctx.
		log.Printf("read message %v %v \n",mt,time.Now())
		if err != nil {
			log.Println("read err", err.Error())
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		ctx :=NewContext(c)
		ctx.SetBody(msg)
		dispatch(ctx)
		//log.Printf("msg %v \n", msg)
		//c.App.Channel <- ctx
	}
}

func (c *Client) Reply(msg []byte)  {

}
func (c *Client) write() {
	timer := time.NewTicker(pingPeriod)
	defer func() {
		log.Printf("write defer")
		timer.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			log.Printf("msg %v \n",msg)
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			writer, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("write err %v",err.Error())
				return
			}
			writer.Write(msg)
			n := len(c.Send)
			for i := 0; i < n; i++ {
				writer.Write(newLine)
				writer.Write(<-c.Send)
			}
			if err := writer.Close(); err != nil {
				return
			}
		case <-timer.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}
	}
}


func dispatch(ctx Context) {
	for i := range channelHandlers {
		_, ok := channelHandlers[i].(ChannelHandler)
		if !ok {
			log.Println("no handler")
			return
		}
		channelHandlers[i].ChannelRead(ctx)
		select {
		case _ = <-ctx.Done():
			log.Println("done")
			//c.Conn.Close()
			return
		}
	}
}