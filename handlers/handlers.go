package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OlegrusWR/quote_book/models"
	"github.com/OlegrusWR/quote_book/storage"
	"github.com/gorilla/mux"
)

type QuoteHandlers struct {
    Storage *storage.Storage
}

func (h *QuoteHandlers) QuoteCreate(w http.ResponseWriter, r *http.Request) {
    var quote models.Quote
    if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
        responseWithError(w, http.StatusBadRequest, "invalid data format")
        return
    }
    id := h.Storage.Add(quote)
    responseWithJSON(w, http.StatusCreated, map[string]int{"id": id})
}

func (h *QuoteHandlers) GetQuotes(w http.ResponseWriter, r *http.Request){
	author := r.URL.Query().Get("author")
	if author != "" {
		quotes, err := h.Storage.GetByAuthor(author)
		if err != nil{
			responseWithError(w, http.StatusNotFound, "Такой автор не найден")
			return
		}
		responseWithJSON(w, http.StatusOK, quotes)
		return
	}

	quote := h.Storage.GetAll()
	responseWithJSON(w, http.StatusOK, quote)
}

func (h *QuoteHandlers) GetById(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok{
		responseWithError(w, http.StatusBadRequest, "требуется указать айди")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil{
		responseWithError(w, http.StatusBadRequest, "неверный формат айди")
	}

	quote, err := h.Storage.GetById(id)
	if err != nil{
		responseWithError(w, http.StatusNotFound, err.Error())
	}
	responseWithJSON(w, http.StatusOK, quote)
}

func (h *QuoteHandlers) GetByAuthor(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	author, ok := vars["author"]
	if !ok{
		responseWithError(w, http.StatusBadRequest, "требуется указать автора")
	}
	quote, err := h.Storage.GetByAuthor(author)
	if err != nil{
		responseWithError(w, http.StatusNotFound, err.Error())
	}
	responseWithJSON(w, http.StatusOK, quote)
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
    responseWithJSON(w, code, map[string]string{"error": msg})
}
