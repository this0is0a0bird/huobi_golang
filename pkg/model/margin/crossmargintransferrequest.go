package margin

type CrossMarginTransferRequest struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
