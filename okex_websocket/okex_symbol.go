package okex_websocket

import (
	. "exchange_websocket/common"
	"strings"
)

// okex symbols
type OkexSymbol struct {
	OkexUsdtSymbol []string
	OkexBtcSymbol  []string
	OkexEthSymbol  []string
	OkexSymbols    []string
}

func NewOkexSymbol() *OkexSymbol {
	okex := new(OkexSymbol)
	return okex.okexSymbolInit()
}

func (o *OkexSymbol) okexSymbolInit() *OkexSymbol {
	for _, symbol := range append(CommonUsdt, OkexUsdt...) {
		o.OkexUsdtSymbol = append(o.OkexUsdtSymbol, strings.ToLower(symbol))
	}

	for _, symbol := range append(CommonBtc, OkexBtc...) {
		o.OkexBtcSymbol = append(o.OkexBtcSymbol, strings.ToLower(symbol))
	}

	for _, symbol := range append(CommonEth, OkexEth...) {
		o.OkexEthSymbol = append(o.OkexEthSymbol, strings.ToLower(symbol))
	}

	o.OkexSymbols = append(o.OkexUsdtSymbol, append(o.OkexBtcSymbol, o.OkexEthSymbol...)...)
	return o
}

func (o *OkexSymbol) OkexSymbolTransfer(symbol string) string {
	isExist, _ := Contain(symbol, o.OkexSymbols)
	if isExist {
		return strings.ToUpper(symbol)
	}
	return ""
}
