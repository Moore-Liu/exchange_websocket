package main

import "exchange_websocket/bitmex_websocket"

func main() {
	bm := bitmex_websocket.BmWebsocketInit()
	bm.BmFuturesInstrumentQuetoWebsocket()
	for true {
		bm.WsConnect()
		go func() {
			bm.Ping()
		}()
		bm.Subscribe()
		bm.ReadMessage()
	}
}
