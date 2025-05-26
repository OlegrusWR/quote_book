package server

import (
	"github.com/OlegrusWR/quote_book/handlers"
	"github.com/OlegrusWR/quote_book/storage"
	"github.com/gorilla/mux"
)

type Server struct {
    Storage *storage.Storage
    Router  *mux.Router
}

func NewServer() *Server {
    s := &Server{
        Storage: storage.NewStorage(),
        Router:  mux.NewRouter(),
    }
    s.configureRouter()
    return s
}

func (s *Server) configureRouter() {
    handlers := &handlers.QuoteHandlers{Storage: s.Storage}
    s.Router.HandleFunc("/quotes", handlers.QuoteCreate).Methods("POST")
	s.Router.HandleFunc("/quotes", handlers.GetQuotes).Methods("GET")
	s.Router.HandleFunc("/quotes/id/{id}", handlers.GetById).Methods("GET")
	s.Router.HandleFunc("/quotes/random", handlers.GetRandom).Methods("GET")
	s.Router.HandleFunc("/quotes/{id}", handlers.DeleteById).Methods("DELETE")
}
