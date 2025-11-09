package api

import (
	"banking-app/internal/bank"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Bank *bank.Bank
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	type req struct {
		ID string `json:"id"`
	}
	var body req
	json.NewDecoder(r.Body).Decode(&body)

	acc := h.Bank.CreateAccount(body.ID)
	json.NewEncoder(w).Encode(acc)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	acc, err := h.Bank.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{"balance": acc.Balance})
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	type req struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int64  `json:"amount"`
	}
	var body req
	json.NewDecoder(r.Body).Decode(&body)

	err := h.Bank.Transfer(body.From, body.To, body.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Id     string `json:"id"`
		Amount int64  `json:"amount"`
	}
	var body req
	json.NewDecoder(r.Body).Decode(&body)

	fmt.Printf("Depositing %d to account %s\n", body.Amount, body.Id)

	err := h.Bank.Deposit(body.Id, body.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}
