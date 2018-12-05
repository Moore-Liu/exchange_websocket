package huobi_websocket

import . "exchange_websocket/common"

type HuobiCylce struct {
	HuobiCycles   []string
	HuobiCycleMap map[string]int
}

func NewHuobiCylce() *HuobiCylce {
	hbCycle := new(HuobiCylce)
	return hbCycle.HuobiCylceInit()
}

func (o *HuobiCylce) HuobiCylceInit() *HuobiCylce {
	o.HuobiCycles = []string{"1min", "5min", "15min", "30min", "60min", "1day", "1week", "1mon"}
	o.HuobiCycleMap = map[string]int{
		"1min":  int(KLine1Min),
		"5min":  int(KLine5Min),
		"15min": int(KLine15Min),
		"30min": int(KLine30Min),
		"60min": int(KLine1hour),
		"1day":  int(KLineDay),
		"1week": int(KLineWeek),
		"1mon":  int(KLineMonth),
	}
	return o
}

func (o *HuobiCylce) HuobiCylceTransfer(cycle string) int {
	isExist, _ := Contain(cycle, o.HuobiCycles)
	if isExist {
		return o.HuobiCycleMap[cycle]
	}
	return 0
}
