package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"Quotes/internal/handler"
	"Quotes/internal/store"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	quoteStore := store.NewQuoteStore()
	quoteHandler := handler.NewHandler(quoteStore)

	fmt.Println("Starting quotes service on :8080...")
	if err := http.ListenAndServe(":8080", quoteHandler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
