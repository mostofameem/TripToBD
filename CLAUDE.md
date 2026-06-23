# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

TripToBD is a Go **microservices monorepo** for a Bangladesh travel-planning platform. Five independently deployable services sit side by side as top-level directories, each with its **own `go.mod` (separate Go module)** — there is no shared root module, so commands always run from inside a service directory.

- `identity-service/` — auth, JWT, RBAC, OAuth, sessions (PostgreSQL)
- `location-service/` — tourist locations, routes, comments (MongoDB)
- `hotel-service/` — hotels, reviews, pictures (PostgreSQL)
- `vehicle-service/` — vehicles, routes, fares, reviews, pictures (PostgreSQL)
- `restaurant-service/` — restaurants, branches, search, events (PostgreSQL + MongoDB + RabbitMQ)

REST is the public protocol; gRPC is reserved for inter-service calls.

## Module-name gotcha

Two services' Go **module** names differ from their **directory** names. Imports/paths use the module name, not the directory:

- `identity-service/` → module **`identity-rbac`**
- `vehicle-service/` → module **`vehicles`**
- (`hotel-service`, `location-service`, `restaurant-service` match their directory names.)

## How each service boots

All services except restaurant use **Cobra** with a root `main.go` that delegates to `cmd.Execute()`:

```bash
cd <service-dir>
go run main.go serve-rest          # identity, location, hotel, vehicle
```

Identity additionally has these subcommands (run in this order on first setup):

```bash
go run main.go serve-migrate-up    # apply SQL migrations
go run main.go serve-seeding       # seed roles/permissions
go run main.go serve-add-user      # create super-admin from user_config.json
go run main.go serve-rest          # start on :5001
```

**Restaurant is the exception** — it has no Cobra. Its root `main.go` calls `cmd.Cmd()` directly, and it reads a JSON config via the `-c` flag (defaults to `../config.json`, so always pass it explicitly):

```bash
cd restaurant-service
go run main.go -c ./config.json
```

(The `-c` flag is parsed inside `config.LoadConfig()` via the stdlib `flag` package, then validated with `go-playground/validator`.)

## Common commands

Run everything from inside the relevant service directory:

```bash
go build ./...                     # compile
go test ./...                      # run all tests
go test -v -run TestFoo ./pkg/...  # run a single test
go vet ./...                       # lint (no separate linter is configured)
go mod tidy                        # sync dependencies
```

Hot reload (optional, needs `go install github.com/air-verse/air@latest`):

```bash
make dev        # = air serve-rest (identity/location/hotel/vehicle)
```

Regenerate gRPC stubs after editing `.proto` files (needs `protoc` + plugins):

```bash
make install-proto-deps   # installs protoc-gen-go, protoc-gen-go-grpc
make build-proto          # runs protoc for that service's .proto files
```

> ⚠️ **Makefiles are stale.** Every service `Makefile` declares `MAIN:=./cmd/server`, but no `cmd/server` directory exists in any service — the real entry is the root `main.go`. Consequently `make build` / `make start` **will fail** as written. Use `go run main.go <subcommand>` directly, or fix `MAIN` to `.` before relying on a make target.

## Architecture per service (layered / hexagonal-leaning)

```
cmd/        Cobra entrypoints (serve-rest, migrations, seeding) — restaurant uses cmd.Cmd() instead
config/     Config loading — Viper+.env for most; JSON -c flag for restaurant
web/        HTTP server, routes, handlers, middlewares, swagger (/swagger)
<domain>/   Domain service: business logic + ports (interfaces)
repo/ db/ mongodb/   Persistence adapters (sqlx/squirrel for PG; mongo-driver for Mongo)
grpc/       gRPC servers + generated clients (where present)
migrations/ sql-migrate *.up.sql / *.down.sql   (identity uses internal/migrations)
logger/     slog structured-logging helpers (JSON)
```

`web` (transport) → domain service → `repo`/`db`/`mongodb` (persistence). Business logic stays independent of HTTP/DB details — add new ports as interfaces in the domain package and implement them in the persistence package.

**Identity uses a different layout** from the other four: it is organized under `internal/` (`internal/api/{handlers,routes,middlewares,utils,swagger}`, `internal/repo`, `internal/rbac`, `internal/token`, `internal/auth`, `internal/Mail`, `internal/entity`, `internal/migrations`). The other services use flat top-level packages. Mirror whichever style the surrounding service already uses.

## Per-service persistence & config quirks

| Service | Config source | DB | Migrations |
| --- | --- | --- | --- |
| identity | `.env` (Viper) | PostgreSQL | `internal/migrations`, explicit `serve-migrate-up/down` |
| location | `.env` (Viper) | MongoDB | n/a |
| hotel | `.env` (Viper) | PostgreSQL | `migrations/`, **auto-runs on startup** (`repo.MigrateDB`) |
| vehicle | `.env` (Viper) | PostgreSQL | `migrations/`, not auto-run in `serve-rest` |
| restaurant | **`config.json` via `-c`** | PostgreSQL + MongoDB | `migrations/` via `dbconfig.yml` |

- **gRPC at runtime:** only **restaurant** actually starts its gRPC server at boot (`grpc.Start()` in `cmd/main.go`). Location's, hotel's and vehicle's gRPC startup is present in code but **commented out** in their `cmd/rest.go` — those three are REST-only at runtime despite the README/architecture diagram implying otherwise.
- **`JWT_SECREAT` typo:** hotel and vehicle read the env var as `JWT_SECREAT` (misspelled). The root `docker-compose.yml` already maps `JWT_SECREAT: ${JWT_SECRET}`, so the stack works — but preserve the typo if editing those services' config code.
- **Restaurant MongoDB uses `mongodb+srv://`** (Atlas SRV scheme) in `restaurant-service/mongodb/connection.go`. Its `MONGO_HOST` must be an Atlas-style hostname, or that file must be changed to `mongodb://` for a standalone/replica-set Mongo. (Location uses a standard `mongodb://` URI.)
- `dbconfig.yml` holds `sql-migrate` datasource defaults for local development.

## Running the whole stack

The root `docker-compose.yml` runs all 5 services against **external** Postgres/Mongo/RabbitMQ (the compose file starts only the Go services, not the datastores). The root `.env` (copy from `.env.example`) is the single source of truth — compose maps the global vars into each service's expected names.

```bash
cp .env.example .env              # fill in external DB/Mongo/RabbitMQ hosts + JWT_SECRET
# one-time build prep (some Dockerfiles bake in config files):
touch location-service/.env hotel-service/.env vehicle-service/.env
cp restaurant-service/config-example.json restaurant-service/config.json
docker compose up --build -d
```

Default host ports: identity `5001`, location `6002` (gRPC `3345`), hotel `5005`, vehicle `5006`, restaurant `6001` (gRPC host port `3346` → container `3345`).

## CI/CD (`.github/workflows/ci-cd.yml`)

Matrix tests + builds the services, then Trivy security scan, then builds/pushes Docker images to GHCR, then deploys (`develop`→dev, `main`→prod). Two things to know before trusting CI green/red:

- **`identity-service` is NOT in the CI matrix** — it is neither tested nor built by the workflow. Only `hotel-service, location-service, restaurant-service, vehicle-service` are.
- **CI references a Python `user-service`** (in `test-python-service` and the `build-and-push` matrix) that **does not exist in this repo**. Those jobs will fail until the directory is added or the workflow is corrected.

The workflow's build step is `go build ... ./cmd/rest.go` (a single file in package `cmd`, not a `main` package) — verify a service actually produces a binary that way before relying on CI's build verdict.

## Conventions

- Branch strategy: `main` (production) ← `develop` (staging) ← `feature/*` / `hotfix/*`. PRs target `develop`.
- `.env`, `config.json`, and built `main` binaries are gitignored — never commit real config.
- Each service exposes `/health-check` or `/hello` for liveness and `/swagger` for interactive API docs.
- Structured logging goes through the service's `logger` package (wraps `log/slog`) — use `logger.Extra(...)` for fields rather than raw `fmt`/`log`.
