## JWT-Auth-Service

1. Заполните ./configs/.env своими данными
2. make run или `docker compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build`

Swagger по умолчанию доступен по localhost:80/auth/swagger/index.html

### Ключевые технологии:
- PostgreSQL
- Golang
- Redis
- JWT
