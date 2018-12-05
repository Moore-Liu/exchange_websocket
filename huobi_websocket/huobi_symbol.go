package huobi_websocket

import (
	. "exchange_websocket/common"
	"strings"
)

// huobi symbols
type HuobiSymbol struct {
	HuobiUsdtSymbol []string
	HuobiBtcSymbol  []string
	HuobiEthSymbol  []string
	HuobiSymbols    []string
}

func NewHuobiSymbol() *HuobiSymbol {
	hb := new(HuobiSymbol)
	return hb.huobiSymbolInit()
}

func (o *HuobiSymbol) huobiSymbolInit() *HuobiSymbol {
	for _, symbol := range append(CommonUsdt, HuobiUsdt...) {
		o.HuobiUsdtSymbol = append(o.HuobiUsdtSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonBtc, HuobiBtc...) {
		o.HuobiBtcSymbol = append(o.HuobiBtcSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonEth, HuobiEth...) {
		o.HuobiEthSymbol = append(o.HuobiEthSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	o.HuobiSymbols = append(o.HuobiUsdtSymbol, append(o.HuobiBtcSymbol, o.HuobiBtcSymbol...)...)
	return o
}

func (o *HuobiSymbol) huobiSymbolTransfer(symbol string) string {
	isExist1, _ := Contain(symbol, o.HuobiUsdtSymbol)
	if isExist1 {
		return strings.ToUpper(strings.Replace(symbol, "usdt", "_usdt", -1))
	}

	isExist2, _ := Contain(symbol, o.HuobiBtcSymbol)
	if isExist2 {
		return strings.ToUpper(strings.Replace(symbol, "btc", "_btc", -1))
	}

	isExist3, _ := Contain(symbol, o.HuobiEthSymbol)
	if isExist3 {
		return strings.ToUpper(strings.Replace(symbol, "eth", "_eth", -1))
	}

	return ""
}
