# Quotes Service
REST API сервис для хранения и управления цитатами, написанный на
Go.
## Требования
- Go 1.16 или выше
- Только стандартная библиотека Go (без внешних зависимостей)
## Установка и запуск
1. Установка зависимостей:
   - Убедитесь, что у вас установлен Go (версия 1.16+):
     ```bash
     go version
     ```
   - Клонируйте репозиторий:
     ```bash
     git clone https://github.com/DeepnDrag/Quote-book
     cd Quote-book
     ```
2. Запустите сервис:
```bash
go run cmd/app/main.go
```
Сервис будет запущен на порту 8080.
## API Endpoints
### 1. Добавление новой цитаты
```bash
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Confucius", "quote":"Life is simple, but we
insist on making it complicated."}'
```
**Request:**
- Method: `POST`
- URL: `/quotes`
- Headers: `Content-Type: application/json`
- Body:
```json
{
"author": "string",
"quote": "string"
}
```
**Response:**
- Status: `201 Created`
- Body:
```json
{
"id": 1,
"author": "Confucius",
"quote": "Life is simple, but we insist on making it
complicated."
}
```
### 2. Получение всех цитат
```bash
curl http://localhost:8080/quotes
```
**Request:**
- Method: `GET`
- URL: `/quotes`
**Response:**
- Status: `200 OK`
- Body:
```json
[
{
"id": 1,
"author": "Confucius",
"quote": "Life is simple, but we insist on making it
complicated."
}
]
```
### 3. Получение случайной цитаты
```bash
curl http://localhost:8080/quotes/random
```
**Request:**
- Method: `GET`
- URL: `/quotes/random`
**Response:**
- Status: `200 OK`
- Body:
```json
{
"id": 1,
"author": "Confucius",
"quote": "Life is simple, but we insist on making it
complicated."
}
```
### 4. Фильтрация по автору
```bash
curl http://localhost:8080/quotes?author=Confucius
```
**Request:**
- Method: `GET`
- URL: `/quotes?author={author}`
**Response:**
- Status: `200 OK`
- Body:
```json
[
{
"id": 1,
"author": "Confucius",
"quote": "Life is simple, but we insist on making it
complicated."
}
]
```
### 5. Удаление цитаты по ID
```bash
curl -X DELETE http://localhost:8080/quotes/1
```
**Request:**
- Method: `DELETE`
- URL: `/quotes/{id}`
**Response:**
- Status: `204 No Content` (успешное удаление)
- Status: `404 Not Found` (цитата не найдена)
## Тестирование
Запустить unit-тесты:
```bash
go test -v
```
## Особенности реализации
- Данные хранятся в памяти (используется map с mutex для
потокобезопасности)
- Используется только стандартная библиотека Go
- ID цитат генерируются автоматически (автоинкремент)
- Все операции потокобезопасны (sync.RWMutex)
- Поддерживается конкурентный доступ к API
## Примеры использования
### Добавление нескольких цитат
```bash
# Цитата Конфуция
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Confucius", "quote":"Life is simple, but we
insist on making it complicated."}'
# Цитата Лао-цзы
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Lao Tzu", "quote":"The journey of a thousand
miles begins with one step."}'
# Еще одна цитата Конфуция
curl -X POST http://localhost:8080/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Confucius", "quote":"It does not matter how
slowly you go as long as you do not stop."}'
```
### Получение всех цитат Конфуция
```bash
curl http://localhost:8080/quotes?author=Confucius
```
### Получение случайной цитаты
```bash
curl http://localhost:8080/quotes/random
```
## Обработка ошибок
- `400 Bad Request` - неверный формат запроса или отсутствуют
обязательные поля
- `404 Not Found` - цитата не найдена (при удалении или когда нет
цитат для случайного выбора)
- `405 Method Not Allowed` - неподдерживаемый HTTP метод
