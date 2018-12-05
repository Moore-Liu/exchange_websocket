package main

import . "exchange_websocket/okex_websocket"

func main() {
	for true {
		okex := OkexWebsocketInit()
		okex.OkexTradeWebsocket()
		okex.Subscribe()
		okex.ReadMessage()
	}

}
