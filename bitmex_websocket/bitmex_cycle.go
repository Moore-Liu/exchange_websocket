package bitmex_websocket

import "exchange_websocket/common"

type BitmexCycle struct {
	BitmexCycles   []string
	BitmexCycleMap map[string]int
}

func NewBitmexCycle() *BitmexCycle {
	bmCycle := new(BitmexCycle)
	return bmCycle.BitmexCycleInit()
}

func (o *BitmexCycle) BitmexCycleInit() *BitmexCycle {
	o.BitmexCycles = []string{"1m", "5m", "1h", "1d"}
	o.BitmexCycleMap = map[string]int{
		"1m": int(common.KLine1Min),
		"5m": int(common.KLine5Min),
		"1h": int(common.KLine1hour),
		"1d": int(common.KLineDay),
	}
	return o
}

func (o *BitmexCycle) BitmexCycleTransfer(cycle string) int {
	isExist, _ := common.Contain(cycle, o.BitmexCycles)
	if isExist {
		return o.BitmexCycleMap[cycle]
	}
	return 0
}
