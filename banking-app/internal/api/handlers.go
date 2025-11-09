package api

import (
	"banking-app/internal/bank"
	"banking-app/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	data := map[string]interface{}{
		"message": fmt.Sprintf("Account created successfully with ID: %s", body.ID),
		"account": acc,
	}
	utils.WriteJSON(w, http.StatusOK, data, nil)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	acc, err := h.Bank.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"message": fmt.Sprintf("Balance retrieved for account: %s", id),
		"balance": acc.Balance,
	}
	utils.WriteJSON(w, http.StatusOK, data, nil)
}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Id     string  `json:"id"`
		Amount float64 `json:"amount"`
	}
	var body req
	json.NewDecoder(r.Body).Decode(&body)

	err := h.Bank.Deposit(body.Id, body.Amount)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "", err)
		return
	}

	acc, _ := h.Bank.GetAccount(body.Id)
	data := map[string]interface{}{
		"message":    fmt.Sprintf("Successfully deposited %f to account %s", body.Amount, body.Id),
		"accountId":  body.Id,
		"amount":     body.Amount,
		"newBalance": acc.Balance,
	}
	utils.WriteJSON(w, http.StatusOK, data, nil)
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Id     string  `json:"id"`
		Amount float64 `json:"amount"`
	}
	var body req
	json.NewDecoder(r.Body).Decode(&body)

	err := h.Bank.Withdraw(body.Id, body.Amount)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "", err)
		return
	}

	acc, _ := h.Bank.GetAccount(body.Id)
	data := map[string]interface{}{
		"message":    fmt.Sprintf("Successfully withdrew %f from account %s", body.Amount, body.Id),
		"accountId":  body.Id,
		"amount":     body.Amount,
		"newBalance": acc.Balance,
	}
	utils.WriteJSON(w, http.StatusOK, data, nil)
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	tx := bank.Transaction{
		ID:        uuid.New().String(),
		From:      req.From,
		To:        req.To,
		Amount:    req.Amount,
		Status:    bank.Pending,
		CreatedAt: time.Now(),
	}

	// store as pending
	h.Bank.Mu.Lock()
	h.Bank.TxStore[tx.ID] = &tx
	h.Bank.Mu.Unlock()

	// send to worker
	h.Bank.Transactions <- tx

	json.NewEncoder(w).Encode(tx)
}

func (h *Handler) GetTransactionStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	h.Bank.Mu.RLock()
	tx, ok := h.Bank.TxStore[id]
	h.Bank.Mu.RUnlock()

	if !ok {
		http.Error(w, "transaction not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tx)
}
