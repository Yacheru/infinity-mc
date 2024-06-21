include .env

docker-network:
	docker network create traefik

docker-local:
	docker-compose -f docker-compose-local.yml up -d --remove-orphans --build

docker-prod:
	docker-compose -f docker-compose.yml up -d --remove-orphans --build

docker-prune:
	docker system prune -fa

migrate-up:
	migrate -path ./backend/schema -database $(POSTGRES_URL) up

migrate-down:
	migrate -path ./backend/schema -database $(POSTGRES_URL) down