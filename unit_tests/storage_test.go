package unit_tests

import (
	"fmt"
	"testing"

	"github.com/OlegrusWR/quote_book/models"
	"github.com/OlegrusWR/quote_book/storage"
)

func TestStorage_Add(t *testing.T) {
	st := storage.NewStorage()

	quote := models.Quote{
		Author: "First Author",
		Quote:  "Лишь тот достоин жизни и свободы, кто каждый день идет за них на бой",
	}
	id := st.Add(quote)

	if id != 1 {
		t.Errorf("ожидаемый ID должен быть равен 1, получено %d", id)
	}

	if len(st.GetAll()) != 1 {
		t.Errorf("ожидаемая длина хранилища должна быть 1, получено %d", len(st.GetAll()))
	}

	storageQuote, err := st.GetById(id)
	if err != nil {
		t.Errorf("Ошибка: %v", err)
	}

	if storageQuote.Author != quote.Author || storageQuote.Quote != quote.Quote {
		t.Errorf("сохраненная цитата не совпадает с введенной цитатой")
	}
}

func TestStorage_GetById_Error(t *testing.T) {
	st := storage.NewStorage()

	_, err := st.GetById(999)

	if err == nil {
		t.Error("expected error for non-existent ID, got nil")
	} else if err.Error() != "цитата с айди  999 не найдена" {
		t.Errorf("unexpected error message: got %v, want %v", err.Error(), "цитата с айди  999 не найдена")
	}
}

func TestStorage_DeleteById(t *testing.T) {
	st := storage.NewStorage()

	quote := models.Quote{
		Author: "Oleg",
		Quote:  "qwe",
	}
	id := st.Add(quote)
	st.DeleteById(id)
	_, err := st.GetById(id)
	if err == nil {
		t.Errorf("ожидалась ошибка при попытке получить удалённую цитату с ID %d, но ошибка не произошла", id)
	} else if err.Error() != fmt.Sprintf("цитата с айди  %d не найдена", id) {
		t.Errorf("неверное сообщение об ошибке: %v, Должно быть %v", err.Error(), fmt.Sprintf("цитата с айди  %d не найдена", id))

	}
}

func TestStorage_GetRandom(t *testing.T) {
	st := storage.NewStorage()

	_, err := st.GetRandom()
	if err == nil {
		t.Errorf("ожидалась ошибка при получении случайной цитаты из пустого хранилища")
	} else if err.Error() != "в цитатнике нет цитат" {
		t.Errorf("неверное сообщение об ошибке: получили %v, ожидалось %v", err.Error(), "в цитатнике нет цитат")
	}

	quotes := []models.Quote{
		{Author: "Amir", Quote: "aaa"},
		{Author: "Oleg", Quote: "bbb"},
		{Author: "Stas", Quote: "ccc"},
	}
	for _, q := range quotes {
		st.Add(q)
	}

	randomQuote, err := st.GetRandom()
	if err != nil {
		t.Errorf("неожиданная ошибка при получении случайной цитаты: %v", err)
	}

	found := false
	for _, q := range quotes {
		if randomQuote.Author == q.Author && randomQuote.Quote == q.Quote {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("случайная цитата не найдена среди добавленных цитат")
	}
}

func TestStorage_GetByAuthor(t *testing.T) {
	st := storage.NewStorage()

	quotes := []models.Quote{
		{Author: "Amir", Quote: "aaa"},
		{Author: "Oleg", Quote: "bbb"},
		{Author: "Amir", Quote: "ccc"},
	}
	for _, q := range quotes {
		st.Add(q)
	}

	author := "Amir"
	results, err := st.GetByAuthor(author)
	if err != nil {
		t.Errorf("неожиданная ошибка при поиске цитат автора %s: %v", author, err)
	}
	if len(results) != 2 {
		t.Errorf("ожидалось 2 цитаты автора %s, получили %d", author, len(results))
	}

	nonExistentAuthor := "Egor"
	_ , err = st.GetByAuthor(nonExistentAuthor)
	if err == nil {
		t.Errorf("ожидалась ошибка при поиске цитат несуществующего автора %s, но ошибка не возникла", nonExistentAuthor)
	} else if err.Error() != fmt.Sprintf("не найдено цитат автора %s", nonExistentAuthor) {
		t.Errorf("неверное сообщение об ошибке: получили %v, ожидалось %v", err.Error(), fmt.Sprintf("не найдено цитат автора %s", nonExistentAuthor))
	}
}

