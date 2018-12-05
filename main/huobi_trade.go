package main

import "exchange_websocket/huobi_websocket"

func main() {
	for true {
		hb := huobi_websocket.HbWebsocketInit()
		hb.HbTradeWebsocket()
		hb.Subscribe()
		hb.ReadMessage()
	}
}
