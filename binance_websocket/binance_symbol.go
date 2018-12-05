package binance_websocket

import (
	. "exchange_websocket/common"
	"strings"
)

// Binance symbols
type BinanceSymbol struct {
	BinanceUsdtSymbol []string
	BinanceBtcSymbol  []string
	BinanceEthSymbol  []string
	BinanceSymbols    []string
}

func NewBinanceSymbol() *BinanceSymbol {
	ba := new(BinanceSymbol)
	return ba.binanceSymbolInit()
}

func (o *BinanceSymbol) binanceSymbolInit() *BinanceSymbol {
	for _, symbol := range append(CommonUsdt, BinanceUsdt...) {
		//[symbol.replace('_', '').lower() for symbol in self.USDT]
		o.BinanceUsdtSymbol = append(o.BinanceUsdtSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonBtc, BinanceBtc...) {
		o.BinanceBtcSymbol = append(o.BinanceBtcSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonEth, BinanceEth...) {
		o.BinanceEthSymbol = append(o.BinanceEthSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	o.BinanceSymbols = append(o.BinanceUsdtSymbol, append(o.BinanceBtcSymbol, o.BinanceEthSymbol...)...)
	return o
}

func (o *BinanceSymbol) BinanceSymbolTransfer(symbol string) string {
	isExist1, _ := Contain(symbol, o.BinanceUsdtSymbol)
	if isExist1 {
		return strings.ToUpper(strings.Replace(symbol, "usdt", "_usdt", -1))
	}

	isExist2, _ := Contain(symbol, o.BinanceBtcSymbol)
	if isExist2 {
		return strings.ToUpper(strings.Replace(symbol, "btc", "_btc", -1))
	}

	isExist3, _ := Contain(symbol, o.BinanceEthSymbol)
	if isExist3 {
		return strings.ToUpper(strings.Replace(symbol, "eth", "_eth", -1))
	}

	return ""
}
