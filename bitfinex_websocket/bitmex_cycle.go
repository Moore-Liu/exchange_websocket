package bitfinex_websocket

import "exchange_websocket/common"

type BitfinexCycle struct {
	BitfinexCycles   []string
	BitfinexCycleMap map[string]int
}

func NewBitfinexCycle() *BitfinexCycle {
	bfCycle := new(BitfinexCycle)
	return bfCycle.BitfinexCycleInit()
}

func (o *BitfinexCycle) BitfinexCycleInit() *BitfinexCycle {
	o.BitfinexCycles = []string{"1m", "5m", "15m", "30m", "1h", "1D", "7D", "1M"}
	o.BitfinexCycleMap = map[string]int{
		"1m":  int(common.KLine1Min),
		"5m":  int(common.KLine5Min),
		"15m": int(common.KLine15Min),
		"30m": int(common.KLine30Min),
		"1h":  int(common.KLine1hour),
		"1D":  int(common.KLineDay),
		"7D":  int(common.KLineWeek),
		"1M":  int(common.KLineMonth),
	}
	return o
}

func (o *BitfinexCycle) BitfinexCycleTransfer(cycle string) int {
	isExist, _ := common.Contain(cycle, o.BitfinexCycles)
	if isExist {
		return o.BitfinexCycleMap[cycle]
	}
	return 0
}
