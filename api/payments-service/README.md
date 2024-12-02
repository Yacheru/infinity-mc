### Данный сервис является часть проекта [infinity-mc](https://github.com/Yacheru/infinity-mc)

Сервис создает платеж в агрегаторе платежей YooKassa, а после ждет подтверждения его оплаты на /payments/accept. 
После получения подтверждения отправляет сообщение в Kafka для дальнейшей обработки...

### Переменные окружения
| Ключ              | Значение       | Описание                   |
|-------------------|----------------|----------------------------|
| `API_PORT`        | 80             | Порт http api              |
| `API_ENVIRONMENT` | development    | Контекст выполнения api    |
| `API_DEBUG`       | true           | Работать ли в debug-режиме |
| `YKASSA_ID`       | youridhere     | ID Вашего магазина         |
| `YKASSA_PASS`     | yourpasshere   | Пароль вашего магазина     |
| `KAFKA_BROKER`    | localhost:9095 | Адрес брокера сообщений    |
| `KAFKA_TOPIC`     | KafkaTopic     | Название топика сообщений  |

### Запуск осуществляется через команду:
- [`make docker-up`](Makefile)