package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Quotes/internal/handler"
	"Quotes/internal/store"
)

func TestHandleAddQuote(t *testing.T) {
	store := store.NewQuoteStore()
	handler := handler.NewHandler(store)

	body := `{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var quote struct {
		ID     int    `json:"id"`
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	if err := json.NewDecoder(w.Body).Decode(&quote); err != nil {
		t.Fatal("Failed to decode response:", err)
	}
	if quote.Author != "Confucius" {
		t.Errorf("Expected author 'Confucius', got '%s'", quote.Author)
	}
}

func TestHandleGetQuotes(t *testing.T) {
	store := store.NewQuoteStore()
	handler := handler.NewHandler(store)

	store.Add("Confucius", "Life is simple, but we insist on making it complicated.")
	store.Add("Lao Tzu", "The journey of a thousand miles begins with one step.")

	req := httptest.NewRequest(http.MethodGet, "/quotes", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var quotes []struct {
		ID     int    `json:"id"`
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	if err := json.NewDecoder(w.Body).Decode(&quotes); err != nil {
		t.Fatal("Failed to decode response:", err)
	}

	if len(quotes) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(quotes))
	}
}

func TestHandleDeleteQuote(t *testing.T) {
	store := store.NewQuoteStore()
	handler := handler.NewHandler(store)

	quote := store.Add("Confucius", "Life is simple, but we insist on making it complicated.")

	req := httptest.NewRequest(http.MethodDelete, "/quotes/1", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	if store.Delete(quote.ID) {
		t.Error("Expected quote to be already deleted")
	}
}

func TestHandleInvalidRequests(t *testing.T) {
	store := store.NewQuoteStore()
	handler := handler.NewHandler(store)

	// Invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid JSON, got %d", w.Code)
	}

	// Missing fields
	body := `{"author":"Confucius"}`
	req = httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for missing quote field, got %d", w.Code)
	}

	// Test delete with invalid ID
	req = httptest.NewRequest(http.MethodDelete, "/quotes/invalid", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid ID, got %d", w.Code)
	}

	// Test delete non-existent quote
	req = httptest.NewRequest(http.MethodDelete, "/quotes/999", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent quote, got %d", w.Code)
	}

	// Test unsupported method
	req = httptest.NewRequest(http.MethodPut, "/quotes", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 for unsupported method, got %d", w.Code)
	}
}
