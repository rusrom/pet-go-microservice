BIN_DIR="/home/rusrom/Code/go/trevor/mecroservices/bin"
BROKER_BINARY="brokerApp"
BROKER_BINARY_WIN="brokerApp.exe"
FRONT_BINARY=frontApp
FRONT_BINARY_WIN="frontApp.exe"

## starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up
	@echo "Docker containers were started"
	@docker compose ps

## stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Docker containers stoped and removed"
	docker compose ps

## build broker-service binaries
binary-broker:
	@echo "Building broker binaries..."
	chdir broker-service && \
	go build -o ${BIN_DIR}/broker-service/${BROKER_BINARY} ./cmd/api && \
	GOOS=windows go build -o ${BIN_DIR}/broker-service/${BROKER_BINARY_WIN} ./cmd/api
	@echo "Done!"

## build front binaries
build-front:
	@echo "Building front binaries..."
	chdir front-end && \
 	go build -o ${BIN_DIR}/front-end/${FRONT_BINARY} ./cmd/web && \
 	GOOS=windows go build -o ${BIN_DIR}/front-end/${FRONT_BINARY_WIN} ./cmd/web
	@echo "Done!"

## start: starts the front end
start-front:
	@echo "Starting front end..."
	chdir front-end && go run ./cmd/web
