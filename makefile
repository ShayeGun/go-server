# When you prefix a command with @, it prevents make from printing that command to the terminal.
# Phony targets are not associated with files but instead represent actions or commands to be run.

include .env

MIGRATE_DIR=./cmd/migrate
DB_DRIVER=postgres

up:
	@goose -dir $(MIGRATE_DIR) $(DB_DRIVER) $(DB_URI) up

down:
	@goose -dir $(MIGRATE_DIR) $(DB_DRIVER) $(DB_URI) down

status:
	@goose -dir $(MIGRATE_DIR) $(DB_DRIVER) $(DB_URI) status

.PHONY: up down status