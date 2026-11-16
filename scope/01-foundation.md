# Milestone 1 — Foundation

**Status:** Done (merged via PR #4, branches `1-inatialize-the-server`, `2-get-the-contract-openapi-running`)

## Goal

Repo, Go API, Next.js app, Docker Compose, PostgreSQL, Goose, configuration,
health endpoint. App starts locally and API verifies database connectivity.

## What was built

- Docker Compose for Postgres, `.env`-based config via `config.AppConfig`.
- Goose migration setup (`backend/db/migrations/version1`), `cmd/migrate` binary.
- chi router wired through oapi-codegen (`chi-server` + `strict-server`),
  contract defined in `openapi/openapi.yaml`, generated into `backend/api/gen.go`.
- uber-fx dependency graph (`internal/modules`): config, database, handler,
  server modules wired in `cmd/api/main.go`.
- `GET /v1/status` — strict handler pings the DB (`PingContext` with a 2s
  timeout) and returns 200 (`status: ok`) or 503 (`status: unhealthy`).
- CORS middleware (`go-chi/cors`) driven by `CorsConfig` env vars.
- Next.js frontend scaffold: TanStack Query provider, Orval-generated typed
  client from the shared `openapi.yaml`, Zustand cart store skeleton,
  role-based route folders (`vendor/`, `cashier/`, `manager/`), working
  status-check demo on the homepage.

## Known follow-ups (not blocking, revisit opportunistically)

- `postgres.go` hardcodes `sslmode=disable` instead of reading `cfg.SSLMode`.
- `server.go` manually builds its own `handler.Handler` via `handler.NewHandler(db)`
  instead of receiving the fx-provided instance from `HandlerModule` — two
  separate `Handler` instances currently exist in the app graph.

## Completion check

App starts locally and API verifies database connectivity. **Met.**
