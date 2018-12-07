package main

import "exchange_websocket/bitfinex_websocket"

func main() {
	bf := bitfinex_websocket.BitfinexWebsocketInit()
	bf.BFTradeWebsocket()
	for true {
		bf.WsConnect()
		go func() {
			bf.Ping()
		}()
		bf.Subscribe("trade")
		bf.ReadMessage()
	}
}
