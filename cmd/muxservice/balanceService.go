package muxservice

// Balance stores total amount for accounts
type Balance struct {
	Amount float32 `json:"amount"`
}

// GetBalance calculates and returns the balance of all accounts
func GetBalance() (Balance, error) {
	var balance Balance

	balance.Amount = 30.32 // Mocked values

	//err := e.marshal(v, encOpts{escapeHTML: true})
	//if err != nil {
	//	return nil, err
	//}
	return balance, nil
}
