package wallet

import "github.com/shopspring/decimal"

type QueryDepositWithdrawResponse struct {
	Status string            `json:"status"`
	Data   []DepositWithdraw `json:"data"`
}
type DepositWithdraw struct {
	Id         int64           `json:"id"`
	Type       string          `json:"type"`
	Currency   string          `json:"currency"`
	TxHash     string          `json:"tx-hash"`
	Chain      string          `json:"chain"`
	Amount     decimal.Decimal `json:"amount"`
	Address    string          `json:"address"`
	AddressTag string          `json:"address-tag"`
	Fee        decimal.Decimal `json:"fee"`
	State      string          `json:"state"`
	CreatedAt  int64           `json:"created-at"`
	UpdatedAt  int64           `json:"updated-at"`
}

type GetWithdrawByClientOrderIdResponse struct {
	Status string            `json:"status"`
	Data   *GetWithdrawByClientOrderId `json:"data"`
}

type GetWithdrawByClientOrderId struct {
	Id            int             `json:"id"`
	ClientOrderId string          `json:"client-order-id"`
	Type          string          `json:"type"`
	SubType       string          `json:"sub-type"`
	Currency      string          `json:"currency"`
	Chain         string          `json:"chain"`
	TxHash        string          `json:"tx-hash"`
	Amount        decimal.Decimal `json:"amount"`
	FromAddrTag   string          `json:"from-addr-tag"`
	Address       string          `json:"address"`
	AddressTag    string          `json:"address-tag"`
	Fee           decimal.Decimal `json:"fee"`
	State         string          `json:"state"`
	CreatedAt     int64           `json:"created-at"`
	UpdatedAt     int64           `json:"updated-at"`
}
