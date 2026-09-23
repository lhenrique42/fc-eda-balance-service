package get_balance

import "balance-service/internal/gateway"

type GetBalanceOutputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(g gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{BalanceGateway: g}
}

func (ue *GetBalanceUseCase) Execute(accountID string) (*GetBalanceOutputDTO, error) {
	balance, err := ue.BalanceGateway.Get(accountID)

	if err != nil {
		return nil, err
	}

	return &GetBalanceOutputDTO{AccountID: accountID, Balance: balance.Amount}, nil
}
