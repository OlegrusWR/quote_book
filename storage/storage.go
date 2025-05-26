package storage

import (
	"math/rand"
	"fmt"
	"sync"

	"github.com/OlegrusWR/quote_book/models"
)

type Storage struct {
    mu     sync.Mutex
    quote  map[int]models.Quote
    lastId int
}

func NewStorage() *Storage {
    return &Storage{
        quote: make(map[int]models.Quote),
    }
}

func (s *Storage) Add(quote models.Quote) int {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.lastId++
    quote.Id = s.lastId
    s.quote[s.lastId] = quote

    return s.lastId
}

func (s *Storage) GetAll() []models.Quote {
    s.mu.Lock()
    defer s.mu.Unlock()

    quoteSlice := make([]models.Quote, 0, len(s.quote))
    for _, q := range s.quote {
        quoteSlice = append(quoteSlice, q)
    }

    return quoteSlice
}

func (s *Storage) GetById(id int) (models.Quote, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if quote, exists := s.quote[id]; exists {
        return quote, nil
    }
    return models.Quote{}, fmt.Errorf("цитата с айди  %d не найдена", id)
}

func (s *Storage) DeleteById(id int){
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.quote, id)
}

func (s *Storage) GetByAuthor(author string) ([]models.Quote, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    var quoteSlice []models.Quote
    for _, q := range s.quote {
        if q.Author == author {
            quoteSlice = append(quoteSlice, q)
        }
    }
    if len(quoteSlice) == 0 {
        return nil, fmt.Errorf("не найдено цитат автора %s", author)
    }
    return quoteSlice, nil
}

func (s *Storage) GetRandom() (models.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.quote) == 0{
		return models.Quote{}, fmt.Errorf("в цитатнике нет цитат")
	}

	keys := make([]int, 0, len(s.quote))
	for key := range s.quote{
		keys = append(keys, key)
	}

	randomIndex := rand.Intn(len(keys))
	randomKeys := randomIndex
	if randomIndex == 0{
		randomKeys++
	}
	
	return s.quote[randomKeys], nil

}

