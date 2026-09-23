package handler

import (
	"balance-service/internal/entity"
	"balance-service/internal/gateway"
	"encoding/json"
	"fmt"
)

type balanceUpdatedMessage struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom        string  `json:"account_id_from"`
		AccountIDTo          string  `json:"account_id_to"`
		BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
		BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
	} `json:"Payload"`
}

type BalanceUpdatedHandler struct {
	BalanceGateway gateway.BalanceGateway
}

func NewBalanceUpdatedHandler(g gateway.BalanceGateway) *BalanceUpdatedHandler {
	return &BalanceUpdatedHandler{BalanceGateway: g}
}

func (h *BalanceUpdatedHandler) Handle(rawMessage []byte) error {
	var msg balanceUpdatedMessage
	if err := json.Unmarshal(rawMessage, &msg); err != nil {
		return err
	}

	if err := h.BalanceGateway.Upsert(entity.NewBalance(msg.Payload.AccountIDFrom, msg.Payload.BalanceAccountIDFrom)); err != nil {
		return err
	}

	if err := h.BalanceGateway.Upsert(entity.NewBalance(msg.Payload.AccountIDTo, msg.Payload.BalanceAccountIDTo)); err != nil {
		return err
	}

	fmt.Println("balances updated for", msg.Payload.AccountIDFrom, msg.Payload.AccountIDTo)
	return nil
}
