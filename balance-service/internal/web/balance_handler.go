package web

import (
	"balance-service/internal/usecase/get_balance"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
	GetBalanceUseCase get_balance.GetBalanceUseCase
}

func NewWebBalanceHandler(ue get_balance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{GetBalanceUseCase: ue}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountId")

	output, err := h.GetBalanceUseCase.Execute(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
