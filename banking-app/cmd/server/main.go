package main

import (
	"banking-app/internal/api"
	"banking-app/internal/bank"
	"log"
	"net/http"
)

func main() {
	log.Println("Setup in process....")
	b := bank.NewBank()
	h := &api.Handler{Bank: b}
	b.StartWorkers(4)
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", h.Router())
}
