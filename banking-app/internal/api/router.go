package api

import "net/http"

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/account", h.CreateAccount)
	mux.HandleFunc("/balance", h.GetBalance)
	mux.HandleFunc("/transfer", h.Transfer)
	mux.HandleFunc("/deposit", h.Deposit)
	mux.HandleFunc("/withdraw", h.Withdraw)
	mux.HandleFunc("/transaction-status", h.GetTransactionStatus)
	return mux
}
