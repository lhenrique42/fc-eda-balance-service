package gateway

import "balance-service/internal/entity"

type BalanceGateway interface {
	Get(accountID string) (*entity.Balance, error)
	Upsert(balance *entity.Balance) error
}
