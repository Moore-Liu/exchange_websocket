package okex_websocket

import . "exchange_websocket/common"

type OkexCylce struct {
	OkexCycles   []string
	OkexCycleMap map[string]int
}

func NewOkexCycle() *OkexCylce {
	okCycle := new(OkexCylce)
	return okCycle.OkexCycleInit()
}

func (o *OkexCylce) OkexCycleInit() *OkexCylce {
	o.OkexCycles = []string{"1min", "3min", "5min", "15min", "30min", "1hour", "4hour", "day", "week"}
	o.OkexCycleMap = map[string]int{
		"1min":  int(KLine1Min),
		"3min":  int(KLine3Min),
		"5min":  int(KLine5Min),
		"15min": int(KLine15Min),
		"30min": int(KLine30Min),
		"1hour": int(KLine1hour),
		"4hour": int(KLine4hour),
		"day":   int(KLineDay),
		"week":  int(KLineWeek),
	}
	return o
}

func (o *OkexCylce) OkexCycleTransfer(cycle string) int {
	isExist, _ := Contain(cycle, o.OkexCycles)
	if isExist {
		return o.OkexCycleMap[cycle]
	}
	return 0
}
