package main

import "exchange_websocket/binance_websocket"

func main() {
	ba := binance_websocket.BinanceWebsocketInit()
	ba.BABtcKlineWebsocket()
	for true {
		ba.WsConnect()
		ba.ReadMessage()
	}
}
