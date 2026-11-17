MIGRATION_DIR=backend/db/migrations/version1
ENV_FILE=backend/.env
MIGRATE_CMD=cd backend && go run ./cmd/migrate

-include Makefile.local
-include $(ENV_FILE)

.PHONY: run migrate-up migrate-down migrate-reset migrate-status test test-v test-pkg test-cover

build:
	@cd backend && go build -o bin/backend cmd/api/main.go

run:
	cd backend && go run ./cmd/api

seed:
	cd backend && go run ./cmd/seed

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


# ---------- tests ----------

# Run every test in the backend.
test:
	cd backend && go test ./...

# Same, but print each test name and PASS/FAIL.
test-v:
	cd backend && go test ./... -v

# Run one package only, e.g. make test-pkg PKG=./internal/handler/vendor_auth/
test-pkg:
	cd backend && go test $(PKG) -v

# Run all tests and show how much of the code they exercise.
test-cover:
	cd backend && go test ./... -cover
