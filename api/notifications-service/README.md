### Данный сервис является часть проекта [infinity-mc](https://github.com/Yacheru/infinity-mc)

Сервис получает из Kafka [сообщение](https://github.com/Yacheru/payments-service/) об оплате и отправляет уведомление об этом в discord используя Webhook

### Переменные окружения
| Ключ                   | Значение           | Описание                               |
|------------------------|--------------------|----------------------------------------|
| `KAFKA_BROKERS`        | localhost:9095     | Адрес брокера сообщений                |
| `KAFKA_CONSUMER_GROUP` | NotificationsGroup | Название консьюмер группы              |
| `KAFKA_TOPIC`          | KafkaTopic         | Название топика для подключения к нему |
| `WEBHOOK_TOKEN`        | yourtokenhere      | Токен вебхука                          |
| `WEBHOOK_ID`           | youridhere         | ID вебхука                             |

### Запуск осуществляется через команду:
- [`make docker-up`](Makefile)