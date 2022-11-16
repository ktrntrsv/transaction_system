# transaction_system

**Quick start**

1. Clone the repo
2. Run `docker compose up`
3. Import `balance-service.postman_collection.json` to Postman app
4. Enjoy
---

Старалась придерживаться чистой архитектуры.
Для реализации "очереди" для одной строки в базочке данных использовала SELECT ... FOR UPDATE в рамках одной транзакции.


Таблица со счетами (accounts) наполнена сущностями для удобства тестирования, их ids:
1. `123e4567-e89b-12d3-a456-426614174000` 
2. `123e4567-e89b-12d3-a456-426614174001`
3. `123e4567-e89b-12d3-a456-426614174002`
4. `123e4567-e89b-12d3-a456-426614174003`
