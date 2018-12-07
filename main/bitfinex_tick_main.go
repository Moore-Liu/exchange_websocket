package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFTickWebsocket()
	for true {
		bf.WsConnect()
		bf.Subscribe("tick")
		bf.ReadMessage()
	}
}
