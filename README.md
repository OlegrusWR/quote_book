# Quote Book

Простой сервер на Go для хранения и работы с цитатами.

---

## Что умеет

- Добавлять новую цитату  
- Получать все цитаты или искать по автору  
- Получать цитату по ID  
- Удалять цитату по ID  
- Показывать случайную цитату  

---

## Технологии

- Язык Go  
- Gorilla Mux для маршрутов  
- Хранилище в памяти (без базы данных)  

---

## Структура проекта
```
quote_book/
├── cmd/ — запуск программы (main.go)
├── handlers/ — обработчики HTTP-запросов
├── models/ — модель данных 
├── storage/ — логика хранения данных
├── server/ — настройка сервера и маршрутов
├── unit_tests/ — тесты (Моки делать не стал, тесты только на storage)
├── go.mod
└── README.md
```
## Как запустить

### Склонировать репозиторий:

```bash
git clone https://github.com/OlegrusWR/quote_book.git
cd quote_book
```
### Собрать и запустить сервер:

```bash
go run cmd/main.go
```
Сервер запустится на http://localhost:8080

## Примеры использования API
#### В новом терминале:

### Добавить цитату

```bash
curl -X POST http://localhost:8080/quotes \ -H "Content-Type: application/json" \
-d
'{"author":"Confucius",
"quote":"Life is simple, but we insist on making it complicated."}'
```

#### Ответ:

```json
{
  "id": 1
}
```

### Получить все цитаты

```bash
curl http://localhost:8080/quotes
```

### Получить цитаты по автору

```bash
curl http://localhost:8080/quotes?author=Confucius
```

### Получить цитату по ID

```bash
curl http://localhost:8080/quotes/id/1
```

### Удалить цитату по ID
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

### Получить случайную цитату
``` bash
curl http://localhost:8080/quotes/random
```

## Тестирование

### Для запусков тестов:
*В корне проекта:*
``` bash
go test -v  -coverprofile=coverage.out ./storage ./unit_tests
```
### Для просмотра покрытия 
``` bash
go tool cover -html=coverage.out
```





