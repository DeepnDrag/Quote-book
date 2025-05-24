package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"Quotes/internal/model"
	"Quotes/internal/store"
)

type Handler struct {
	store *store.QuoteStore
}

func NewHandler(store *store.QuoteStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case path == "/quotes" && r.Method == http.MethodPost:
		h.handleAddQuote(w, r)
	case path == "/quotes" && r.Method == http.MethodGet:
		h.handleGetQuotes(w, r)
	case path == "/quotes/random" && r.Method == http.MethodGet:
		h.handleGetRandomQuote(w, r)
	case strings.HasPrefix(path, "/quotes/") && r.Method == http.MethodDelete:
		h.handleDeleteQuote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleAddQuote(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.Author == "" || req.Quote == "" {
		http.Error(w, "Author and quote are required", http.StatusBadRequest)
		return
	}
	quote := h.store.Add(req.Author, req.Quote)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (h *Handler) handleGetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []*model.Quote
	if author != "" {
		quotes = h.store.GetByAuthor(author)
	} else {
		quotes = h.store.GetAll()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *Handler) handleGetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote := h.store.GetRandom()
	if quote == nil {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *Handler) handleDeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}
	if h.store.Delete(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Quote not found", http.StatusNotFound)
	}
}
