package unit

import (
	"testing"

	"Quotes/internal/store"
)

func TestQuoteStore(t *testing.T) {
	store := store.NewQuoteStore()

	// Test Add
	quote := store.Add("Confucius", "Life is simple, but we insist on making it complicated.")
	if quote.ID != 1 {
		t.Errorf("Expected ID 1, got %d", quote.ID)
	}
	if quote.Author != "Confucius" {
		t.Errorf("Expected author Confucius, got %s", quote.Author)
	}

	// Test GetAll
	quotes := store.GetAll()
	if len(quotes) != 1 {
		t.Errorf("Expected 1 quote, got %d", len(quotes))
	}

	// Add another quote
	store.Add("Lao Tzu", "The journey of a thousand miles begins with one step.")

	// Test GetByAuthor
	confuciusQuotes := store.GetByAuthor("Confucius")
	if len(confuciusQuotes) != 1 {
		t.Errorf("Expected 1 Confucius quote, got %d", len(confuciusQuotes))
	}

	// Test GetRandom
	randomQuote := store.GetRandom()
	if randomQuote == nil {
		t.Error("Expected a random quote, got nil")
	}

	// Test Delete
	deleted := store.Delete(1)
	if !deleted {
		t.Error("Expected quote to be deleted")
	}

	// Verify deletion
	quotes = store.GetAll()
	if len(quotes) != 1 {
		t.Errorf("Expected 1 quote after deletion, got %d", len(quotes))
	}
}
