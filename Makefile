include .env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=./migrations

run:
	go run ./cmd/api/main.go

up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down% 