package binance_websocket

import "exchange_websocket/common"

type BinanceCycle struct {
	BinanceCycles   []string
	BinanceCycleMap map[string]int
}

func NewBinanceCycle() *BinanceCycle {
	bfCycle := new(BinanceCycle)
	return bfCycle.BinanceCycleInit()
}

func (o *BinanceCycle) BinanceCycleInit() *BinanceCycle {
	o.BinanceCycles = []string{"1m", "5m", "15m", "30m", "1h", "4h", "1d", "1w", "1M"}
	o.BinanceCycleMap = map[string]int{
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

func (o *BinanceCycle) BinanceCycleTransfer(cycle string) int {
	isExist, _ := common.Contain(cycle, o.BinanceCycles)
	if isExist {
		return o.BinanceCycleMap[cycle]
	}
	return 0
}
