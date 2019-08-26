package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFDepthWebsocket()
	for true {
		bf.WsConnect()
		go func() {
			bf.Ping()
		}()
		bf.Subscribe("depth")
		bf.ReadMessage()
	}
}
