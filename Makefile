MIGRATION_DIR=backend/db/migrations/version1
ENV_FILE=backend/.env
MIGRATE_CMD=cd backend && go run ./cmd/migrate

-include Makefile.local
-include $(ENV_FILE)

.PHONY: run migrate-up migrate-down migrate-reset migrate-status

build:
	@cd backend && go build -o bin/backend cmd/api/main.go

run:
	cd backend && go run ./cmd/api

migrate-up:
	$(MIGRATE_CMD) -action=up

migrate-down:
	$(MIGRATE_CMD) -action=down

migrate-reset:
	$(MIGRATE_CMD) -action=reset

migrate-status:
	$(MIGRATE_CMD) -action=status

gen-api:
	@cd backend && oapi-codegen --config cfg.yaml ../openapi/openapi.yaml
	@echo "Generated api/gen.go from openapi.yaml"

