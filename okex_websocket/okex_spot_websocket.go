package okex_websocket

import (
	"fmt"
	"github.com/go-log/log"
	"github.com/gorilla/websocket"
	"time"
)

type Okex interface {
	//ping
	Ping() error
	// connect
	WsConnect() error
	// subscribe
	Subscribe()
	// read message
	ReadMessage()
	// tick
	OkexTickWebsocket()
	// depth
	OkexDepthWebsocket()
	// trade
	OkexTradeWebsocket()
	// kline
	OkexKlineWebsocket()
}

type okex struct {
	Url      string
	Ws       *websocket.Conn
	Channels []*OkexWebsocketRequest
}

// 初始化
func OkexWebsocketInit() *okex {
	okex := new(okex)
	okex.Url = "wss://real.okex.com:10441/websocket?compress=true"
	return okex
}

// 连接websocket
func (o *okex) WsConnect() {

	dialer := new(websocket.Dialer)

	ws, _, err := dialer.Dial(o.Url, nil)
	if err != nil {
		panic(err)
	}
	o.Ws = ws

}

//ping 保持连接
func (o *okex) Ping() {
	pingMsg := ping{Event: "ping"}
	for true {
		err := o.Ws.WriteJSON(pingMsg)
		if err != nil {
			log.Log("error message:", err)
		}
		time.Sleep(25 * time.Second)
	}

}

type ping struct {
	Event string `json:"event"`
}

// 订阅channel
func (o *okex) Subscribe() {
	if o.Channels == nil {
		log.Log("no channels")
		return
	}
	fmt.Println(o.Channels)
	for _, channel := range o.Channels {
		fmt.Println(*channel)
		err := o.Ws.WriteJSON(*channel)
		if err != nil {
			fmt.Println("subscribe error: ", err)
		}
	}

}

//读取消息
func (o *okex) ReadMessage() {
	for true {
		msgType, msg, err := o.Ws.ReadMessage()
		if err != nil {
			fmt.Println("read message error: ", err)
			break
		}
		if msgType == websocket.TextMessage {
			data := string(msg)
			fmt.Println("message is:", msgType, data)
		} else if msgType == websocket.BinaryMessage {
			message, err := GzipDecode(msg)
			if err != nil {
				log.Log("error:", err)
				continue
			}
			data := string(message)
			fmt.Println("message is:", msgType, data)
		} else if msgType == websocket.CloseMessage {
			// 重新连接
			err := o.Ws.Close()
			if err != nil {
				log.Log("error:", err)
			}
			break
		}
	}
}

// tick
func (o *okex) OkexTickWebsocket() {
	okSymbols := NewOkexSymbol()
	for _, symbol := range okSymbols.OkexSymbols {
		channel := OkexWebsocketRequest{"addChannel", "ok_sub_spot_" + symbol + "_ticker"}
		o.Channels = append(o.Channels, &channel)
	}
}

// depth
func (o *okex) OkexDepthWebsocket() {
	okSymbols := NewOkexSymbol()
	for _, symbol := range okSymbols.OkexSymbols {
		channel := OkexWebsocketRequest{"addChannel", "ok_sub_spot_" + symbol + "_depth_10"}
		o.Channels = append(o.Channels, &channel)
	}
}

// trade
func (o *okex) OkexTradeWebsocket() {
	okSymbols := NewOkexSymbol()
	for _, symbol := range okSymbols.OkexSymbols {
		channel := OkexWebsocketRequest{"addChannel", "ok_sub_spot_" + symbol + "_deals"}
		o.Channels = append(o.Channels, &channel)
	}
}

// kline
func (o *okex) OkexKlineWebsocket() {
	okSymbols := NewOkexSymbol()
	okCycles := NewOkexCycle()
	for _, symbol := range okSymbols.OkexSymbols {
		for _, cycle := range okCycles.OkexCycles {
			channel := OkexWebsocketRequest{"addChannel", "ok_sub_spot_" + symbol + "_kline_" + cycle}
			o.Channels = append(o.Channels, &channel)
		}
	}
}

/************* okex websocket request **********/
type OkexWebsocketRequest struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

/************ okex tick *****************/
type OkexTickWebsocketResponse struct {
	Channel string `json:"channel"`
	Data    OkexTickData
}

type OkexTickData struct {
	High      string `json:"high"`
	Vol       string `json:"vol"`
	Last      string `json:"last"`
	Low       string `json:"low"`
	Buy       string `json:"buy"`
	Change    string `json:"change"`
	Sell      string `json:"sell"`
	DayLow    string `json:"dayLow"`
	DayHigh   string `json:"dayHigh"`
	Timestamp int64  `json:"timestamp"`
}

/************** okex depth **************/
type OkexDepthWebsocketResponse struct {
	Channel string `json:"channel"`
	Data    OkexDepthData
}

type BidsAsksData struct {
	Price  string
	Volume string
}

type OkexDepthData struct {
	Asks      []*BidsAsksData `json:"asks"`
	Bids      []*BidsAsksData `json:"bids"`
	Timestamp int64           `json:"timestamp"`
}

/************* okex trade **************/
type OkexTradeWebsocketResponse struct {
	Channel string `json:"channel"`
	Data    []*OkexTradeData
}

type OkexTradeData struct {
	Id     string `json:"id"`
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Time   string `json:"time"`
	Type   string `json:"type"`
}

/*************** okex kline ******************/
type OkexKlineWebsocketResponse struct {
	Channel string `json:"channel"`
	Data    [][]*OkexKlinehData
}

type OkexKlinehData struct {
	Time   string `json:"time"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
}
