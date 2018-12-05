package main

import . "exchange_websocket/okex_websocket"

func main() {
	okex := OkexWebsocketInit()
	okex.OkexDepthWebsocket()
	for {
		okex.Subscribe()
		okex.ReadMessage()
	}

}
