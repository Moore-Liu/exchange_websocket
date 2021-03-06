package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFTickWebsocket()
	for true {
		bf.WsConnect()
		go func() {
			bf.Ping()
		}()
		bf.Subscribe("tick")
		bf.ReadMessage()
	}
}
