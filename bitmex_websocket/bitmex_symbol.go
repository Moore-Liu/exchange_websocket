package bitmex_websocket

import . "exchange_websocket/common"

// bitmex symbols
type BitmexSymbol struct {
	BitmexSymbols []string
}

func NewBitmexSymbol() *BitmexSymbol {
	bm := new(BitmexSymbol)
	return bm.bitmexSymbolInit()
}

func (o *BitmexSymbol) bitmexSymbolInit() *BitmexSymbol {

	o.BitmexSymbols = BitmexYx
	return o
}

func NewBitmexFuturesSymbol() *BitmexSymbol {
	bm := new(BitmexSymbol)
	return bm.bitmexFuturesSymbolInit()
}

func (o *BitmexSymbol) bitmexFuturesSymbolInit() *BitmexSymbol {
	o.BitmexSymbols = BitmexDq
	return o
}
