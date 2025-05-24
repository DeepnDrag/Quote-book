package store

import (
	"math/rand"
	"sync"

	"Quotes/internal/model"
)

type QuoteStore struct {
	mu     sync.RWMutex
	quotes map[int]*model.Quote
	nextID int
}

func NewQuoteStore() *QuoteStore {
	return &QuoteStore{
		quotes: make(map[int]*model.Quote),
		nextID: 1,
	}
}

func (s *QuoteStore) Add(author, quote string) *model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()
	q := &model.Quote{
		ID:     s.nextID,
		Author: author,
		Quote:  quote,
	}
	s.quotes[s.nextID] = q
	s.nextID++
	return q
}

func (s *QuoteStore) GetAll() []*model.Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()
	quotes := make([]*model.Quote, 0, len(s.quotes))
	for _, q := range s.quotes {
		quotes = append(quotes, q)
	}
	return quotes
}

func (s *QuoteStore) GetByAuthor(author string) []*model.Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []*model.Quote
	for _, q := range s.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result
}

func (s *QuoteStore) GetRandom() *model.Quote {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.quotes) == 0 {
		return nil
	}

	ids := make([]int, 0, len(s.quotes))
	for id := range s.quotes {
		ids = append(ids, id)
	}
	randomID := ids[rand.Intn(len(ids))]
	return s.quotes[randomID]
}

func (s *QuoteStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.quotes[id]; exists {
		delete(s.quotes, id)
		return true
	}
	return false
}
