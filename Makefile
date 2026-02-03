.PHONY: build up down logs clean

# Development
dev-build:
	docker compose -f docker-compose.yml build

dev-up:
	docker compose -f docker-compose.yml up

dev-down:
	docker compose -f docker-compose.yml down

# Common
logs:
	docker compose logs -f

clean:
	docker compose down -v --rmi all --remove-orphans
