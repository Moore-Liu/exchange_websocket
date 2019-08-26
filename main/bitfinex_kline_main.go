package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFKlineWebsocket()
	for true {
		bf.WsConnect()
		go func() {
			bf.Ping()
		}()
		bf.Subscribe("kline")
		bf.ReadMessage()
	}
}
