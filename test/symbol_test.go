package test

import (
	. "exchange_websocket/okex_websocket"
	"fmt"
	"testing"
)

func TestSymbol(t *testing.T) {
	okex := NewOkexSymbol()
	fmt.Println(okex)
}
