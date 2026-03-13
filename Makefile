.PHONY: build up deps run debug app down down-app down-deps down-deps-v down-all down-all-v logs clean

# Build app image
build:
	docker compose -f docker-compose.app.yml build

# Start everything in RUN mode (default)
up: deps run

# Start dependencies only
deps:
	docker compose -f docker-compose.deps.yml up -d

# ---- APP MODES ----
# Normal development (no debugger)
run:
	docker compose -f docker-compose.app.yml --profile run up

# Debug mode (Delve + Air)
debug:
	docker compose -f docker-compose.app.yml --profile debug up


# ---- STOP COMMANDS ----
down-app:
	docker compose -f docker-compose.app.yml down
	@echo "✅ app stopped"

down-deps:
	docker compose -f docker-compose.deps.yml down
	@echo "✅ dependencies stopped"

down-deps-v:
	docker compose -f docker-compose.deps.yml down -v
	@echo "✅ dependencies stopped and volumes deleted"

down-all:
	docker compose -f docker-compose.deps.yml -f docker-compose.app.yml down
	@echo "✅ all services stopped"

# Stop everything and delete ALL volumes
down-all-v:
	docker compose -f docker-compose.deps.yml -f docker-compose.app.yml down -v
	@echo "✅ all services stopped and all volumes deleted"

down: down-all

# ---- LOGS ----
logs:
	docker compose -f docker-compose.deps.yml -f docker-compose.app.yml logs -f

# ---- CLEAN ----
clean:
	docker compose -f docker-compose.deps.yml -f docker-compose.app.yml down -v --rmi all --remove-orphans
	@echo "✅ everything cleaned (containers, volumes, images, orphans)"
