.PHONY: build up down logs clean downv

# Development
build:
	docker compose -f docker-compose.yml build

up:
	docker compose -f docker-compose.yml up

down:
	docker compose -f docker-compose.yml down

downv:
	docker compose -f docker-compose.yml down -v
# Common
logs:
	docker compose logs -f

clean:
	docker compose down -v --rmi all --remove-orphans
