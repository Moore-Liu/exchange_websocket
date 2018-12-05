package bitfinex_websocket

import (
	. "exchange_websocket/common"
	"strings"
)

// bitfinex symbol
type BitfinexSymbol struct {
	BitfinexUsdtSymbol []string
	BitfinexBtcSymbol  []string
	BitfinexEthSymbol  []string
	BitfinexSymbols    []string
}

func NewBitfinexSymbol() *BitfinexSymbol {
	bf := new(BitfinexSymbol)
	return bf.bitfinexSymbolInit()
}

func (o *BitfinexSymbol) bitfinexSymbolInit() *BitfinexSymbol {
	for _, symbol := range append(CommonUsdt, BitfinexUsdt...) {
		// return [symbol.replace('_', '').rsplit('T', 1)[0] for symbol in self.USDT]
		o.BitfinexUsdtSymbol = append(o.BitfinexUsdtSymbol, strings.Replace(symbol[0:len(symbol)-1], "_", "", -1))
	}

	for _, symbol := range append(CommonBtc, BitfinexBtc...) {
		o.BitfinexBtcSymbol = append(o.BitfinexBtcSymbol, strings.Replace(symbol, "_", "", -1))
	}

	for _, symbol := range append(CommonEth, BitfinexEth...) {
		o.BitfinexEthSymbol = append(o.BitfinexEthSymbol, strings.Replace(symbol, "_", "", -1))
	}

	o.BitfinexSymbols = append(o.BitfinexUsdtSymbol, append(o.BitfinexBtcSymbol, o.BitfinexEthSymbol...)...)
	return o
}

func (o *BitfinexSymbol) BitfinexSymbolTransfer(symbol string) string {

	isExist1, _ := Contain(symbol, o.BitfinexUsdtSymbol)
	isExist2, _ := Contain(symbol, o.BitfinexBtcSymbol)
	isExist3, _ := Contain(symbol, o.BitfinexEthSymbol)
	if isExist1 {
		symbol = strings.Replace(symbol, "USD", "_USD", -1)
	} else if isExist2 {
		symbol = strings.Replace(symbol, "BTC", "_BTC", -1)
	} else if isExist3 {
		symbol = strings.Replace(symbol, "ETH", "_ETH", -1)
	} else {
		return ""
	}

	if strings.HasPrefix(symbol, "IOS") {
		return strings.Replace(symbol, "IOS", "IOST", -1)
	}
	if strings.HasPrefix(symbol, "DSH") {
		return strings.Replace(symbol, "DSH", "DASH", -1)
	}
	if strings.HasPrefix(symbol, "IOT") {
		return strings.Replace(symbol, "IOT", "IOTA", -1)
	}

	return symbol

}
