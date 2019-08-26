package bitfinex_websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

type Bitfinex interface {
	//ping
	Ping()
	// connect
	WsConnect()
	// subscribe
	Subscribe()
	// read message
	ReadMessage()
	// tick
	BFTickWebsocket()
	// depth
	BFDepthWebsocket()
	// trade
	BFTradeWebsocket()
	// kline
	BFKlineWebsocket()
}

type bitfinex struct {
	Url           string
	Ws            *websocket.Conn
	Channels      []*BitfinexWebsocketRequest
	DepthChannels []*BitfinexDepthWebsocketRequest
	KlineChannels []*BitfinexKlineWebsocketRequest
}

// 初始化
func BitfinexWebsocketInit() *bitfinex {
	bf := new(bitfinex)
	bf.Url = os.Getenv("BITFINEX_URL")
	return bf
}

// websocket connect
func (o *bitfinex) WsConnect() {
	dialer := new(websocket.Dialer)
	ws, _, err := dialer.Dial(o.Url, nil)
	if err != nil {
		fmt.Println("websocket connect error:", err)
		return
	}
	o.Ws = ws
}

// ping 保持连接
func (o *bitfinex) Ping() {
	pingMsg := ping{Event: "ping"}
	done := make(chan struct{})
	// 5s定时
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C: // ping消息
			err := o.Ws.WriteJSON(pingMsg)
			if err != nil {
				fmt.Println("ping error: ", err)
				return
			}

		}
	}
}

// subscribe
func (o *bitfinex) Subscribe(channel string) {
	switch channel {
	case "depth":
		if o.DepthChannels == nil {
			fmt.Println("no channels")
			return
		}
		for _, channel := range o.DepthChannels {
			err := o.Ws.WriteJSON(*channel)
			if err != nil {
				fmt.Println("subscribe error: ", err)
				return
			}
		}
	case "kline":
		if o.KlineChannels == nil {
			fmt.Println("no channels")
			return
		}
		for _, channel := range o.KlineChannels {
			err := o.Ws.WriteJSON(*channel)
			if err != nil {
				fmt.Println("subscribe error: ", err)
				return
			}
		}
	default:
		if o.Channels == nil {
			fmt.Println("no channels")
			return
		}
		for _, channel := range o.Channels {
			err := o.Ws.WriteJSON(*channel)
			if err != nil {
				fmt.Println("subscribe error: ", err)
				return
			}
		}
	}

}

// read message
func (o *bitfinex) ReadMessage() {
	for true {
		msgType, msg, err := o.Ws.ReadMessage()
		if err != nil {
			fmt.Println("read message error: ", err)
			break
		}
		if msgType == websocket.TextMessage {
			data := string(msg)
			fmt.Println("message is:", msgType, data)
		} else if msgType == websocket.CloseMessage {
			// 重新连接
			err := o.Ws.Close()
			if err != nil {
				fmt.Println("error:", err)
			}
			break
		}
	}
}

// tick
func (o *bitfinex) BFTickWebsocket() {
	bfSymbols := NewBitfinexSymbol()
	for _, symbol := range bfSymbols.BitfinexSymbols {
		channel := BitfinexWebsocketRequest{
			"subscribe",
			"ticker",
			symbol,
		}
		o.Channels = append(o.Channels, &channel)
	}
}

// trade
func (o *bitfinex) BFTradeWebsocket() {
	bfSymbols := NewBitfinexSymbol()
	for _, symbol := range bfSymbols.BitfinexSymbols {
		channel := BitfinexWebsocketRequest{
			"subscribe",
			"trades",
			symbol,
		}
		o.Channels = append(o.Channels, &channel)
	}
}

// depth
func (o *bitfinex) BFDepthWebsocket() {
	bfSymbols := NewBitfinexSymbol()
	for _, symbol := range bfSymbols.BitfinexSymbols {
		channel := BitfinexDepthWebsocketRequest{
			"subscribe",
			"book",
			"P0",
			symbol,
			"25",
		}
		o.DepthChannels = append(o.DepthChannels, &channel)
	}
}

// kline
func (o *bitfinex) BFKlineWebsocket() {
	bfSymbols := NewBitfinexSymbol()
	bfCylces := NewBitfinexCycle()
	for _, cycle := range bfCylces.BitfinexCycles {
		for _, symbol := range bfSymbols.BitfinexSymbols {
			channel := BitfinexKlineWebsocketRequest{
				"subscribe",
				"candles",
				"trade:" + cycle + ":t" + symbol,
			}
			o.KlineChannels = append(o.KlineChannels, &channel)
		}
	}

}

type ping struct {
	Event string `json:"event"`
}

type BitfinexWebsocketRequest struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Pair    string `json:"pair"`
}

type BitfinexDepthWebsocketRequest struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Prec    string `json:"prec"`
	Symbol  string `json:"symbol"`
	Len     string `json:"len"`
}

type BitfinexKlineWebsocketRequest struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Key     string `json:"key"`
}
