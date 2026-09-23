package entity

import "time"

type Balance struct {
	AccountID string
	Amount    float64
	UpdatedAt time.Time
}

func NewBalance(accountID string, amount float64) *Balance {
	return &Balance{
		AccountID: accountID,
		Amount:    amount,
		UpdatedAt: time.Now(),
	}
}
