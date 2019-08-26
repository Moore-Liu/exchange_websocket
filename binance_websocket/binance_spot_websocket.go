package binance_websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

type Binance interface {
	// connect
	WsConnect()
	// read message
	ReadMessage()
	// tick
	BATickWebsocket()
	// depth
	BADepthWebsocket()
	// trade
	BATradeWebsocket()
	// kline
	BAKlineUsdtWebsocket()
	BAKlineBtcWebsocket()
	BAKlineEthWebsocket()
}

type binance struct {
	Url string
	Ws  *websocket.Conn
}

// 初始化
func BinanceWebsocketInit() *binance {
	ba := new(binance)
	ba.Url = os.Getenv("BINANCE_URL")
	return ba
}

// websocket connect
func (o *binance) WsConnect() {
	dialer := new(websocket.Dialer)
	ws, _, err := dialer.Dial(o.Url, nil)
	if err != nil {
		fmt.Println("websocket connect error:", err)
		return
	}
	o.Ws = ws
}

// read message
func (o *binance) ReadMessage() {
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
func (o *binance) BATickWebsocket() {
	baSymbols := NewBinanceSymbol()
	for _, symbol := range baSymbols.BinanceSymbols {
		o.Url += symbol + "@miniTicker/"
	}
}

// trade
func (o *binance) BATradeWebsocket() {
	baSymbols := NewBinanceSymbol()
	for _, symbol := range baSymbols.BinanceSymbols {
		o.Url += symbol + "@trade/"
	}
}

// depth
func (o *binance) BADepthWebsocket() {
	baSymbols := NewBinanceSymbol()
	for _, symbol := range baSymbols.BinanceSymbols {
		o.Url += symbol + "@depth10/"
	}
}

/*****************************************/
/****    参数长度限制，需将三个币种分开   ****/
/*****************************************/
// usdt kline
func (o *binance) BAUsdtKlineWebsocket() {
	baSymbols := NewBinanceSymbol()
	baCycles := NewBinanceCycle()
	for _, symbol := range baSymbols.BinanceUsdtSymbol {
		for _, cycle := range baCycles.BinanceCycles {
			o.Url += symbol + "@kline_" + cycle + "/"
		}
	}
}

// btc kline
func (o *binance) BABtcKlineWebsocket() {
	baSymbols := NewBinanceSymbol()
	baCycles := NewBinanceCycle()
	for _, symbol := range baSymbols.BinanceBtcSymbol {
		for _, cycle := range baCycles.BinanceCycles {
			o.Url += symbol + "@kline_" + cycle + "/"
		}
	}
}

// eth kline
func (o *binance) BAEthKlineWebsocket() {
	baSymbols := NewBinanceSymbol()
	baCycles := NewBinanceCycle()
	for _, symbol := range baSymbols.BinanceEthSymbol {
		for _, cycle := range baCycles.BinanceCycles {
			o.Url += symbol + "@kline_" + cycle + "/"
		}
	}
}
