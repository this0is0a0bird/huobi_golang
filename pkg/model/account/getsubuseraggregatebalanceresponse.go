package account

type GetSubUserAggregateBalanceResponse struct {
	Status string           `json:"status"`
	Data   []AccountBalance `json:"data"`
}
