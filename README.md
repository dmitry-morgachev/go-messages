## API

### Сохранение сообщения

- **URL:** `/message/save`
- **Метод:** `POST`
- **Описание:** Сохраняет сообщение в БД и отправляет его в Kafka.
- **Параметры запроса:**
  - `content` - строка, содержимое сообщения.

#### Пример запроса:

```bash
curl -X POST http://localhost:8080/message/save -H "Content-Type: application/json" -d '{"content":"Hello!"}'
```

### Получение статистики

- **URL:** `/stats`
- **Метод:** `GET`
- **Описание:** Возвращает статистику по обработанным сообщениям.

#### Пример запроса:

```bash
curl -X GET http://localhost:8080/stats
```