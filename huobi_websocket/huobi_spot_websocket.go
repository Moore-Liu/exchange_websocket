package huobi_websocket

import (
	"fmt"
	"github.com/go-log/log"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

type Huobi interface {
	// ws ping pong
	Pong()
	// connect
	WsConnect()
	// Subscribe
	Subscribe()
	// ReadMessage
	ReadMessage()
	// tick
	HbTickWebsocket()
	// depth
	HbDepthWebsocket()
	// trade
	HbTradeWebsocket()
	// Kline
	HbKlineWebsocket()
}

type huobi struct {
	Url      string
	Ws       *websocket.Conn
	Channels []*HbWebsocketRequest
}

// 初始化
func HbWebsocketInit() *huobi {
	huobi := new(huobi)
	huobi.Url = os.Getenv("HUOBI_URL")
	return huobi
}

// 连接websocket
func (o *huobi) WsConnect() {
	dailer := new(websocket.Dialer)
	ws, _, err := dailer.Dial(o.Url, nil)
	if err != nil {
		log.Log("websocket connect error:", err)
	}
	o.Ws = ws

}

// pong
func (o *huobi) Pong(pingData string) {
	ts := pingData[8:21]
	pong := pong{Pong: ts}
	err := o.Ws.WriteJSON(pong)
	if err != nil {
		log.Log("pong error: ", err)
	}
}

type pong struct {
	Pong string `json:"pong"`
}

func (o huobi) Subscribe() {
	if o.Channels == nil {
		log.Log("no channels")
		return
	}
	fmt.Println(o.Channels)

	for _, channel := range o.Channels {
		fmt.Println(*channel)
		err := o.Ws.WriteJSON(*channel)
		if err != nil {
			log.Log("channel subscribe error!")
		}
	}
}

func (o huobi) ReadMessage() {
	for true {
		msgType, msg, err := o.Ws.ReadMessage()
		if err != nil {
			log.Log("read message error:", err)
			break
		}
		if msgType == websocket.TextMessage {
			data := string(msg)
			fmt.Println("message is:", msgType, data)
		} else if msgType == websocket.BinaryMessage {
			message, err := GzipDecompress(msg)
			if err != nil {
				log.Log("message error:", err)
			}
			data := string(message)
			if strings.Contains(data, "ping") {
				o.Pong(data)
			}
			fmt.Println("message is:", msgType, data)
		} else if msgType == websocket.CloseMessage {
			//重连
			err := o.Ws.Close()
			if err != nil {
				log.Log("websocket close error:", err)
			} else {
				log.Log("websocket close")
			}
			break
		}

	}
}

// tick
func (o *huobi) HbTickWebsocket() {
	hbSymbols := NewHuobiSymbol()
	for _, symbol := range hbSymbols.HuobiSymbols {
		channel := HbWebsocketRequest{
			"market." + symbol + ".detail",
			"id12",
		}
		o.Channels = append(o.Channels, &channel)
	}
}

//depth
func (o *huobi) HbDepthWebsocket() {
	hbSymbols := NewHuobiSymbol()
	var step string
	if os.Getenv("HUOBI_STEP") != "" {
		step = os.Getenv("HUOBI_STEP")
	} else {
		step = "step0"
	}
	for _, symbol := range hbSymbols.HuobiSymbols {
		channel := HbWebsocketRequest{
			"market." + symbol + ".depth." + step,
			"id1",
		}
		o.Channels = append(o.Channels, &channel)
	}
}

//trade
func (o *huobi) HbTradeWebsocket() {
	hbSymbols := NewHuobiSymbol()
	for _, symbol := range hbSymbols.HuobiSymbols {
		channel := HbWebsocketRequest{
			"market." + symbol + ".trade.detail",
			"id1",
		}
		o.Channels = append(o.Channels, &channel)
	}
}

//kline
func (o *huobi) HbKlineWebsocket() {
	hbSymbols := NewHuobiSymbol()
	hbCycles := NewHuobiCylce()
	for _, symbol := range hbSymbols.HuobiSymbols {
		for _, cycle := range hbCycles.HuobiCycles {
			channel := HbWebsocketRequest{
				"market." + symbol + ".kline." + cycle,
				"id10",
			}
			o.Channels = append(o.Channels, &channel)
		}
	}
}

/********** huobi websocket request *************/
type HbWebsocketRequest struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

/*************** huobi tick *********************/
type tick struct {
	Amount float32 `json:"amount"`
	Open   float32 `json:"open"`
	Close  float32 `json:"close"`
	High   float32 `json:"high"`
	Ts     int64   `json:"ts"`
	Id     int64   `json:"id"`
	Count  int32   `json:"count"`
	Low    float32 `json:"low"`
	Vol    float32 `json:"volume"`
}

type HbTickData struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Data tick   `json:"data"`
}

/**************** huobi depth ********************/
type depthTick struct {
	Bids [][2]float32 `json:"bids"`
	Asks [][2]float32 `json:"asks"`
}

type HbDepthData struct {
	Ch   string    `json:"ch"`
	Ts   int64     `json:"ts"`
	Tick depthTick `json:"tick"`
}

/**************** huobi trade ********************/
type tradeData struct {
	Amount    float32 `json:"amount"`
	Ts        int64   `json:"ts"`
	Id        int64   `json:"id"`
	Price     float32 `json:"price"`
	Direction string  `json:"direction"`
}

type tradeTick struct {
	Id   int64       `json:"id"`
	Ts   int64       `json:"ts"`
	Data []tradeData `json:"data"`
}

type HbTradeTick struct {
	Ch   string    `json:"ch"`
	Ts   int64     `json:"ts"`
	Tick tradeTick `json:"tick"`
}

/*************** huobi kline ******************/
type klineTick struct {
	Amount float32 `json:"amount"`
	Open   float32 `json:"open"`
	Close  float32 `json:"close"`
	High   float32 `json:"high"`
	Id     int64   `json:"id"`
	Count  int32   `json:"count"`
	Low    float32 `json:"low"`
	Vol    float32 `json:"volume"`
}

type HbKlineData struct {
	Ch   string    `json:"ch"`
	Ts   int64     `json:"ts"`
	Tick klineTick `json:"tick"`
}
