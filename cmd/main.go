package main

import (
	"net/http"

	"github.com/OlegrusWR/quote_book/storage"
	"github.com/gorilla/mux"
)

type Server struct{
	storage *storage.Storage
	router *mux.Router
}

func newServer() *Server{
	s := &Server{
		storage: storage.NewStorage(),
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter(){
	s.router.HandleFunc("/quotes", s.handlerQuoteCreate).Methods("POST")
}

func main(){
	s := newServer()
	http.ListenAndServe(":8080", s.router)
}