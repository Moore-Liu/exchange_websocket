package bitmex_websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

type Bitmex interface {
	// ping
	Ping()
	// ws connect
	WsConnect()
	// subscribe
	Subscribe()
	// reade message
	ReadMessage()
	// trade
	BmTradeWebsocket()
	// depth
	BmDepthWebsocket()
	// kline
	BmKlineWebsocket()
}

type bitmex struct {
	Url      string
	Ws       *websocket.Conn
	Channels []*BmWebsocketRequest
}

// 初始化
func BmWebsocketInit() *bitmex {
	bm := new(bitmex)
	bm.Url = os.Getenv("BITMEX_URL")
	bm.WsConnect()
	return bm
}

// 连接websocket
func (o *bitmex) WsConnect() {
	dailer := new(websocket.Dialer)
	ws, _, err := dailer.Dial(o.Url, nil)
	if err != nil {
		fmt.Println("websocket connect error:", err)
	}
	o.Ws = ws
}

// ping
func (o *bitmex) Ping() {
	err := o.Ws.WriteMessage(websocket.TextMessage, []byte("ping"))
}

type BmWebsocketRequest struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}
