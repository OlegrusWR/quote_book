package storage

import (
	"fmt"
	"sync"

	m "github.com/OlegrusWR/quote_book/models"
)

type Storage struct {
	mu    sync.Mutex
	quote map[int]m.Quote
	lastId int
}

func NewStorage() *Storage{
	return &Storage{
		quote: make(map[int]m.Quote),
	}
}

func (s *Storage) Add(quote m.Quote) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.quote == nil{
		s.quote = make(map[int]m.Quote)
	}
	s.lastId++
	NewId := s.lastId
	quote.Id = NewId
	s.quote[NewId] = quote

	return NewId
}

func (s *Storage) GetAll() []m.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	quoteSlice := make([]m.Quote, 0, len(s.quote))
	for _, q := range s.quote{
		quoteSlice = append(quoteSlice, q)
	}

	return quoteSlice
}

func (s *Storage) GetById(id int) (m.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if quote, exists := s.quote[id]; exists {
		return quote, nil
	}
	return m.Quote{}, fmt.Errorf("цитата с айди %d не найдена", id)
}

func (s *Storage) GetByAuthor(author string) ([]m.Quote, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	var quoteSlice []m.Quote
	for _, q := range s.quote{
		if q.Author == author{
			quoteSlice = append(quoteSlice, q)
		}
	}
	if quoteSlice == nil{
		return nil, fmt.Errorf("не найдено цитат автора %s", author)
	}
	return quoteSlice, nil
}