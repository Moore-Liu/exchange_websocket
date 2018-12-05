package main

import . "exchange_websocket/okex_websocket"

func main() {
	for true {
		okex := OkexWebsocketInit()
		okex.OkexDepthWebsocket()
		okex.Subscribe()
		okex.ReadMessage()
	}

}
