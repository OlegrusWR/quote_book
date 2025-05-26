package main

import (
	"net/http"

	"github.com/OlegrusWR/quote_book/server"
)


func main(){
	s := server.NewServer()
	http.ListenAndServe(":8080", s.Router)
}