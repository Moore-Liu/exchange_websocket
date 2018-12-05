package common

var (
	/************************* BTC交易对 **************************/
	CommonUsdt = []string{"BTC_USDT", "ETH_USDT", "EOS_USDT", "XRP_USDT", "ETC_USDT", "LTC_USDT", "NEO_USDT",
		"TRX_USDT"}

	OkexUsdt = []string{"XLM_USDT", "XMR_USDT", "OKB_USDT", "DASH_USDT", "ZEC_USDT", "IOTA_USDT", "IOST_USDT",
		"OMG_USDT", "ONT_USDT", "ADA_USDT", "BCHABC_USDT", "BCHSV_USDT"}

	HuobiUsdt = []string{"HT_USDT", "DASH_USDT", "ZEC_USDT", "IOTA_USDT", "IOST_USDT", "OMG_USDT", "ONT_USDT",
		"ADA_USDT", "BCH_USDT"}

	BinanceUsdt = []string{"XLM_USDT", "BNB_USDT", "ONT_USDT", "ADA_USDT", "IOTA_USDT", "BCHABC_USDT", "BCHSV_USDT"}

	BitfinexUsdt = []string{"DSH_USDT", "XLM_USDT", "ZEC_USDT", "XMR_USDT", "IOS_USDT", "OMG_USDT", "IOT_USDT"}

	/************************* BTC交易对 **************************/
	CommonBtc = []string{"ETH_BTC", "EOS_BTC", "XRP_BTC", "ETC_BTC", "LTC_BTC", "NEO_BTC", "TRX_BTC",
		"XLM_BTC", "ZEC_BTC", "XMR_BTC", "OMG_BTC"}

	OkexBtc = []string{"OKB_BTC", "ONT_BTC", "ADA_BTC", "IOTA_BTC", "IOST_BTC", "DASH_BTC", "BCHABC_BTC", "BCHSV_BTC"}

	HuobiBtc = []string{"HT_BTC", "ONT_BTC", "ADA_BTC", "IOTA_BTC", "IOST_BTC", "DASH_BTC", "BCH_BTC"}

	BinanceBtc = []string{"BNB_BTC", "ONT_BTC", "ADA_BTC", "IOTA_BTC", "IOST_BTC", "DASH_BTC", "BCHABC_BTC",
		"BCHSV_BTC"}

	BitfinexBtc = []string{"IOT_BTC", "IOS_BTC", "DSH_BTC"}

	/************************* ETH交易对 **************************/
	CommonEth = []string{"EOS_ETH", "TRX_ETH", "XLM_ETH", "OMG_ETH"}

	OkexEth = []string{"ETC_ETH", "LTC_ETH", "NEO_ETH", "XRP_ETH", "DASH_ETH", "ZEC_ETH", "OKB_ETH", "XMR_ETH",
		"ONT_ETH", "ADA_ETH", "IOTA_ETH", "IOST_ETH"}

	HuobiEth = []string{"HT_ETH", "XMR_ETH", "ONT_ETH", "ADA_ETH", "IOTA_ETH", "IOST_ETH"}

	BinanceEth = []string{"XRP_ETH", "ETC_ETH", "LTC_ETH", "NEO_ETH", "DASH_ETH", "ZEC_ETH", "BNB_ETH",
		"XMR_ETH", "ONT_ETH", "ADA_ETH", "IOTA_ETH", "IOST_ETH"}

	BitfinexEth = []string{"NEO_ETH", "IOT_ETH", "IOS_ETH"}

	/************************* bitmex 永续合约 *********************/
	BitmexYx = []string{"XBTUSD"}
)

/************ kline cycle *****************/
type Cycle int

var (
	KLine1Min  = Cycle(1)
	KLine3Min  = Cycle(2)
	KLine5Min  = Cycle(3)
	KLine15Min = Cycle(4)
	KLine30Min = Cycle(5)
	KLine1hour = Cycle(6)
	KLine4hour = Cycle(7)
	KLineDay   = Cycle(8)
	KLineWeek  = Cycle(10)
	KLineMonth = Cycle(11)
)
