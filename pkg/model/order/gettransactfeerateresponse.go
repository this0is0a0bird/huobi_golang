package order

type GetTransactFeeRateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []*GetTransactFeeRateItem `json:"data"`
}

type GetTransactFeeRateItem struct {
	Symbol          string `json:"symbol"`
	MakerFeeRate    string `json:"makerFeeRate"`
	TakerFeeRate    string `json:"takerFeeRate"`
	ActualMakerRate string `json:"actualMakerRate"`
	ActualTakerRate string `json:"actualTakerRate"`
}
