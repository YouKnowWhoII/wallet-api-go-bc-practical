package models

type Wallet struct {
	ID      string        `json:"id" validate:"required"`
	Name    string        `json:"name" validate:"required"`
	Balance float64       `json:"balance" validate:"gte=0"`
	Txns    []Transaction `json:"transactions"`
}

type Transaction struct {
	Type   string  `json:"type" validate:"required,oneof=credit debit"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
