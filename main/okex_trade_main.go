package main

import . "exchange_websocket/okex_websocket"

func main() {
	okex := OkexWebsocketInit()
	okex.OkexTradeWebsocket()
	for {
		okex.Subscribe()
		okex.ReadMessage()
	}

}
