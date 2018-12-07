package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFTradeWebsocket()
	for true {
		bf.WsConnect()
		bf.Subscribe("trade")
		bf.ReadMessage()
	}
}
