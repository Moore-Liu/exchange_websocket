package main

import "exchange_websocket/huobi_websocket"

func main() {
	hb := huobi_websocket.HbWebsocketInit()
	hb.HbTickWebsocket()
	for {
		hb.WsConnect()
		hb.Subscribe()
		hb.ReadMessage()
	}
}
