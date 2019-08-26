package bitmex_websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
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
	done := make(chan struct{})
	// 5s定时
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C: // ping消息
			err := o.Ws.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				fmt.Println("ping error: ", err)
				return
			}

		}
	}

}

// 订阅subscribe
func (o *bitmex) Subscribe() {
	if o.Channels == nil {
		fmt.Println("no channels")
		return
	}
	for _, channel := range o.Channels {
		err := o.Ws.WriteJSON(*channel)
		if err != nil {
			fmt.Println("subscribe error: ", err)
		}
	}
}

// 读取消息
func (o *bitmex) ReadMessage() {
	for true {
		msgType, msg, err := o.Ws.ReadMessage()
		if err != nil {
			fmt.Println("read message error: ", err)
			break
		}
		if msgType == websocket.TextMessage {
			data := string(msg)
			fmt.Println("message: ", msgType, data)
		} else if msgType == websocket.CloseMessage {
			// 关闭当前websocket
			err := o.Ws.Close()
			if err != nil {
				fmt.Println("close websocket error: ", err)
			}
			break
		}

	}

}

// depth funding Instrument消息
func (o *bitmex) BmDepthFundingInstrumentWebsocket() {
	bmSymbols := NewBitmexSymbol()
	var depthArgs []string
	var fundingArgs []string
	var instrumentArgs []string
	for _, symbol := range bmSymbols.BitmexSymbols {
		depthCh := "orderBook10:" + symbol
		fundingCh := "funding:" + symbol
		instrumentCh := "instrument:" + symbol
		depthArgs = append(depthArgs, depthCh)
		fundingArgs = append(fundingArgs, fundingCh)
		instrumentArgs = append(instrumentArgs, instrumentCh)
	}
	args := append(append(depthArgs, fundingArgs...), instrumentArgs...)
	channel := BmWebsocketRequest{"subscribe", args}
	o.Channels = append(o.Channels, &channel)
}

// trade quote kline消息
func (o *bitmex) BmTradeQuoteKlineWebsocket() {
	bmSymbols := NewBitmexSymbol()
	bmCycles := NewBitmexCycle()
	var tradeArgs []string
	var quoteArgs []string
	var klineArgs []string
	for _, symbol := range bmSymbols.BitmexSymbols {
		tradeCh := "trade:" + symbol
		quoteCh := "quote:" + symbol
		tradeArgs = append(tradeArgs, tradeCh)
		quoteArgs = append(quoteArgs, quoteCh)
		for _, cycle := range bmCycles.BitmexCycles {
			klineCh := "tradeBin" + cycle + ":" + symbol
			klineArgs = append(klineArgs, klineCh)
		}
	}
	args := append(append(tradeArgs, quoteArgs...), klineArgs...)
	channel := BmWebsocketRequest{"subscribe", args}
	o.Channels = append(o.Channels, &channel)
}

func (o *bitmex) BmFuturesTradeDepthWebsocket() {
	bmFuturesSymbols := NewBitmexFuturesSymbol()
	var tradeArgs []string
	var depthArgs []string
	for _, symbol := range bmFuturesSymbols.BitmexSymbols {
		tradeCh := "trade:" + symbol
		depthCh := "orderBook10:" + symbol
		tradeArgs = append(tradeArgs, tradeCh)
		depthArgs = append(depthArgs, depthCh)
	}
	args := append(tradeArgs, depthArgs...)
	channel := BmWebsocketRequest{"subscribe", args}
	o.Channels = append(o.Channels, &channel)
}

func (o *bitmex) BmFuturesInstrumentQuetoWebsocket() {
	bmFuturesSymbols := NewBitmexFuturesSymbol()
	var instrumentArgs []string
	var quetoArgs []string
	for _, symbol := range bmFuturesSymbols.BitmexSymbols {
		instrumentCh := "instrument:" + symbol
		quetoCh := "queto:" + symbol
		instrumentArgs = append(instrumentArgs, instrumentCh)
		quetoArgs = append(quetoArgs, quetoCh)
	}
	args := append(instrumentArgs, quetoArgs...)
	channel := BmWebsocketRequest{"subscribe", args}
	o.Channels = append(o.Channels, &channel)
}

func (o *bitmex) BmFuturesKlineWebsocket() {
	bmFuturesSymbols := NewBitmexFuturesSymbol()
	bmCycles := NewBitmexCycle()
	var klineArgs []string
	for _, symbol := range bmFuturesSymbols.BitmexSymbols {
		for _, cycle := range bmCycles.BitmexCycles[0:2] {
			klineCh := "tradeBin" + cycle + ":" + symbol
			klineArgs = append(klineArgs, klineCh)
		}
	}
	channel := BmWebsocketRequest{"subscribe", klineArgs}
	o.Channels = append(o.Channels, &channel)
}

type BmWebsocketRequest struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}
