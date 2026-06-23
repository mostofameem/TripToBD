# TripToBD

A travel-planning microservices platform for exploring Bangladesh — covering hotels, restaurants, vehicles/transport routes, and tourist locations, with a dedicated identity & RBAC service for authentication and authorization.

The platform is built as a **Go monorepo** of independently deployable services that communicate over REST (primary) and gRPC (inter-service), backed by **PostgreSQL** and **MongoDB**, with **RabbitMQ** for asynchronous messaging.

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Monorepo Structure](#monorepo-structure)
- [Services](#services)
  - [Identity Service (Auth & RBAC)](#identity-service-auth--rbac)
  - [Location Service](#location-service)
  - [Hotel Service](#hotel-service)
  - [Vehicle Service](#vehicle-service)
  - [Restaurant Service](#restaurant-service)
- [Tech Stack](#tech-stack)
- [Shared Patterns](#shared-patterns)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Inter-service Communication (gRPC)](#inter-service-communication-grpc)
- [CI/CD](#cicd)
- [Contributing](#contributing)
- [Credits](#credits)

---

## Overview

TripToBD ("Trip to Bangladesh") is a domain-driven microservices backend that lets users discover and review:

- **Locations** — tourist destinations, routes, and comments.
- **Hotels** — listings with ratings, reviews, and pictures.
- **Vehicles** — transport options (buses, cars, etc.) with routes, fares, reviews, and pictures.
- **Restaurants** — listings with branches, reviews, and location-based search.

All access is gated by the **identity service**, which provides JWT authentication, role-based access control (RBAC), email-invite onboarding, OAuth social login, and session auditing.

---

## Architecture

```
                            ┌──────────────────────────┐
                            │      Identity Service     │
                            │   (Auth · JWT · RBAC)     │
                            └─────────────┬────────────┘
                                          │ issues / verifies JWT
   ┌──────────────────┬───────────────────┼───────────────────┬──────────────────┐
   ▼                  ▼                   ▼                   ▼                  ▼
┌─────────┐      ┌─────────┐         ┌─────────┐         ┌─────────┐       ┌─────────────┐
│ Location│◄─────│ Vehicle │         │  Hotel  │         │Restaur- │       │   Clients   │
│ Service │ gRPC │ Service │         │ Service │         │  ant    │       │ (Web/Mobile)│
└────┬────┘      └────┬────┘         └────┬────┘         └────┬────┘       └─────────────┘
     │                │                   │                   │
     ▼                ▼                   ▼                   ▼
 MongoDB        PostgreSQL          PostgreSQL         PostgreSQL + MongoDB
                                                     + RabbitMQ (events)
```

- **REST** is the public-facing protocol for every service.
- **gRPC** is used for typed inter-service calls (location exposes a `PostService`; restaurant exposes a `RestaurantsService`).
- Each service owns its own database (no shared schemas) and can be built, tested, and deployed in isolation.

---

## Monorepo Structure

```
TripToBD/
├── identity-service/      # Auth, RBAC, users, sessions, OAuth
├── location-service/      # Tourist locations & routes (MongoDB + gRPC)
├── hotel-service/         # Hotels, reviews, pictures, stats
├── vehicle-service/       # Vehicles, routes, fares, reviews, pictures
├── restaurant-service/    # Restaurants, branches, search (PG + Mongo + RabbitMQ + gRPC)
├── docs/                  # CI/CD pipeline & setup guides
├── .github/workflows/     # GitHub Actions CI/CD
└── README.md
```

Each service follows a consistent, layered (hexagonal-leaning) layout:

```
service/
├── cmd/            # Cobra CLI entrypoints (serve-rest, migrations, seeding)
├── config/         # Config loading (Viper / JSON + validator)
├── web/            # HTTP server, routes, handlers, middlewares, swagger
├── <domain>/       # Domain service (business logic + ports)
├── repo/ | db/ | mongodb/   # Data-access layer
├── grpc/           # gRPC servers & protobuf-generated clients (where used)
├── migrations/     # SQL migrations (sql-migrate)
├── logger/         # Structured (slog) logging helpers
├── metrics/        # gopsutil-based host metrics (where used)
├── rabbitmq/       # Resilient RabbitMQ client (restaurant only)
├── main.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── dbconfig.yml
└── .env.example
```

---

## Services

### Identity Service (Auth & RBAC)

Authentication, authorization, and identity management for the whole platform.

| Property | Value |
| --- | --- |
| Module | `identity-rbac` |
| HTTP port | `5001` |
| Database | PostgreSQL |
| Key libs | `golang-jwt/v5`, `markbates/goth` (OAuth), `spf13/cobra`, `spf13/viper`, `sqlx`, `sql-migrate` |

**Responsibilities:** user registration/login, JWT access & refresh tokens, roles & permissions (RBAC), email-invite user onboarding, OAuth social login, session tracking, and audit logging.

**REST endpoints**

| Method | Path | Notes |
| --- | --- | --- |
| `GET` | `/health-check` | Liveness |
| `POST` | `/api/v1/register` | Invite-token registration |
| `POST` | `/api/v1/login` | Credential login |
| `GET` | `/api/v1/token/refresh` | Refresh access token |
| `PATCH` | `/api/v1/reset-password` | JWT-authenticated |
| `POST` | `/api/v1/roles` | Create role (authz) |
| `GET` | `/api/v1/roles` | List roles |
| `GET` | `/api/v1/permissions` | List permissions |
| `POST` | `/api/v1/roles/assign-permission` | Attach permission → role |
| `POST` | `/api/v1/users/assign-role` | Attach role → user |
| `POST` | `/api/v1/users/invite` | Invite a user by email |
| `GET` | `/api/v1/users` | List users |
| `GET` | `/api/v1/users/me/permissions` | Current user's permissions |
| `GET`/`GET` | `/auth/{provider}/login`, `/auth/{provider}/callback` | OAuth flow |

**CLI commands** (via Cobra): `serve-rest`, `serve-migrate-up`, `serve-migrate-down`, `serve-seeding`, `serve-add-user`.

**Schema highlights:** `users`, `roles`, `permissions`, `user_roles`, `role_permissions`, `user_sessions`, `audit_logs`, `user_onboarding_process`.

> See [identity-service/README.md](identity-service/README.md) for its detailed setup (seeding, super-admin bootstrap, etc.).

---

### Location Service

Tourist destinations, travel routes, and community comments.

| Property | Value |
| --- | --- |
| Module | `location-service` |
| HTTP port | `6002` |
| gRPC port | `3345` |
| Database | MongoDB |
| Key libs | `mongo-driver`, `grpc`/`protobuf`, `cobra`, `viper` |

**REST endpoints**

| Method | Path |
| --- | --- |
| `GET` | `/locations/hello` |
| `POST` | `/locations/add-location` |
| `GET` | `/locations/get-location` |
| `GET` | `/locations/get-locations` |
| `POST` | `/add-route` |
| `GET` | `/get-route` |

Exposes a gRPC `PostService` for other services to query location data. Documents are stored in MongoDB (`Location` with nested `Comment[]` and `Route`).

---

### Hotel Service

Hotel listings with reviews, pictures, and live host stats.

| Property | Value |
| --- | --- |
| Module | `hotel-service` |
| HTTP port | `5005` (gRPC `3355` provisioned) |
| Database | PostgreSQL |
| Key libs | `sqlx`, `sql-migrate`, `squirrel`, `gopsutil` (metrics), `cobra` |

**REST endpoints**

| Method | Path |
| --- | --- |
| `GET` | `/hello` |
| `GET` | `/api/v1/stats` | Host metrics |
| `POST` | `/hotel` | Add hotel |
| `GET` | `/hotel` | List hotels |
| `POST` | `/review` | Add review |
| `GET` | `/review` | Get review |
| `POST` | `/pics` | Add pictures |
| `GET` | `/pics` | Get pictures |

Auto-migrates its schema on startup (`vehicles`, `routes`, `reviews`, `pictures` tables).

---

### Vehicle Service

Transport options (buses/cars/etc.) with routes, fares, reviews, and pictures.

| Property | Value |
| --- | --- |
| Module | `vehicles` |
| HTTP port | `5005` (gRPC provisioned) |
| Database | PostgreSQL |
| Key libs | `sqlx`, `sql-migrate`, `squirrel`, `gopsutil`, `cobra` |

**REST endpoints**

| Method | Path |
| --- | --- |
| `GET` | `/hello` |
| `GET` | `/api/v1/stats` |
| `POST` | `/add-vehicle` |
| `GET` | `/get-vehicle` |
| `GET` | `/get-vehicles` |
| `POST` | `/add-routes` |
| `GET` | `/get-routes/vehicle` | Routes for a vehicle |
| `GET` | `/get-routes/location` | Routes for a destination |
| `POST` | `/add-review` |
| `GET` | `/get-review` |
| `POST` | `/add-pics` |
| `GET` | `/get-pics` |

**Schema:** `vehicles`, `routes` (src/dest/category/cost), `reviews`, `pictures` (text-array of URLs).

---

### Restaurant Service

Restaurant discovery, branches, reviews, and search — the richest service in the stack.

| Property | Value |
| --- | --- |
| Module | `restaurant-service` |
| HTTP port | `6001` |
| gRPC port | `3345` |
| Databases | PostgreSQL (restaurants) + MongoDB (comments/posts) |
| Messaging | RabbitMQ (resilient reconnect client) |
| Key libs | `mongo-driver`, `grpc`/`protobuf`, `rabbitmq/amqp091-go`, `sql-migrate` |

**REST endpoints**

| Method | Path |
| --- | --- |
| `GET` | `/hello` |
| `POST` | `/add-restaurant` |
| `GET` | `/get-restaurant/{id}` |
| `GET` | `/get-restaurants` |
| `GET` | `/search` |
| `GET` | `/get-restaurants/location/{id}` |
| `POST` | `/review-restaurant` |
| `GET` | `/Update-restaurant/{id}` |
| `POST` | `/restaurant/add-branch` |

Exposes a gRPC `RestaurantsService`. Unlike the others, this service is configured from a **JSON config file** (`-c ./config.json`) rather than environment variables, and ships a self-healing RabbitMQ client (auto-reconnect, in-memory message retry queue).

---

## Tech Stack

| Concern | Choice |
| --- | --- |
| Language | Go 1.22–1.24 |
| HTTP / routing | `net/http` (Go 1.22+ method routing) + `rs/cors` |
| CLI | `spf13/cobra` |
| Config | `spf13/viper`, `joho/godotenv`, JSON |
| Validation | `go-playground/validator/v10` |
| Relational DB | PostgreSQL via `lib/pq`, `jmoiron/sqlx`, `rubenv/sql-migrate` |
| Query builder | `Masterminds/squirrel` |
| Document DB | MongoDB (`go.mongodb.org/mongo-driver`) |
| Messaging | RabbitMQ (`rabbitmq/amqp091-go`) |
| RPC | gRPC + Protocol Buffers |
| Auth | `golang-jwt/v5`, `markbates/goth` (OAuth) |
| Logging | `log/slog` (structured JSON) |
| Metrics | `shirou/gopsutil` |
| Docs | Swagger UI per service |
| Containers | Multi-stage Docker builds + `docker-compose` |
| CI/CD | GitHub Actions (test → security scan → build → deploy) |

---

## Shared Patterns

- **Cobra-based entrypoints** — every service boots through `cmd` (`go run main.go serve-rest`).
- **Layered design** — `web` (transport) → domain service → `repo`/`db`/`mongodb` (persistence), keeping business logic independent of HTTP/DB details.
- **Structured logging** — a small `logger` package wraps `slog` with JSON serialization helpers, consistent across services.
- **CORS + middleware manager** — a uniform `middlewares.Manager` chains handlers identically in each service.
- **Swagger** — each service exposes interactive API docs at `/swagger`.
- **SQL migrations** — `sql-migrate` with versioned `*.up.sql` / `*.down.sql` files.

---

## Prerequisites

- **Go** 1.23+ (services target 1.22–1.24)
- **PostgreSQL** (identity, hotel, vehicle, restaurant)
- **MongoDB** (location, restaurant comments)
- **RabbitMQ** (restaurant)
- **Docker** & **Docker Compose** (for containerized runs)
- **`air`** (optional, hot reload): `go install github.com/air-verse/air@latest`
- **`protoc`** + plugins (only if regenerating gRPC stubs): see each service's `Makefile` → `make install-proto-deps && make build-proto`

---

## Getting Started

Each service is self-contained. From a service directory:

```bash
# 1. Configure environment
cp .env.example .env        # then fill in DB/secret values
# (restaurant-service uses config.json instead — see below)

# 2. Run with hot reload
make dev                    # installs deps, runs `air serve-rest`

# …or build & run a static binary
make build
go run main.go serve-rest

# …or via Docker
docker-compose up --build
```

### Identity service bootstrap

The identity service needs schema + seed data + an admin user before first run:

```bash
cd identity-service
cp .env.example .env
cp user_config.example.json user_config.json   # fill admin credentials
go run main.go serve-migrate-up                 # apply migrations
go run main.go serve-seeding                    # seed roles/permissions
go run main.go serve-add-user                   # create the super admin
go run main.go serve-rest                       # start on :5001
```

### Restaurant service

Configured from JSON, not `.env`:

```bash
cd restaurant-service
cp config-example.json config.json    # fill in DB / Mongo / RabbitMQ / gRPC values
go run main.go -c ./config.json
```

---

## Configuration

- **`.env`** (Viper/godotenv) — identity, location, hotel, vehicle. Copy from each `.env.example`.
- **`config.json`** — restaurant service (loaded with the `-c` flag; validated with `validator`).
- **`dbconfig.yml`** — `sql-migrate` datasource defaults for local development.

Common environment knobs: `MODE`, `SERVICE_NAME`, `HTTP_PORT`, `GRPC_PORT`, `JWT_SECRET`, DB host/port/name/user/pass, `MIGRATION_SOURCE`, `GRPC_REQ_TIMEOUT_IN_SECOND`, and (restaurant) `RABBITMQ_URL`.

> ⚠️ Never commit real `.env` / `config.json` files — they are gitignored.

---

## Inter-service Communication (gRPC)

- **location-service** listens on gRPC `:3345` and registers a `PostService`.
- **restaurant-service** listens on gRPC `:3345` and registers a `RestaurantsService`.
- **hotel-service** / **vehicle-service** have gRPC scaffolding present but currently commented out (REST-only at runtime).

Regenerate stubs (when `.proto` files change) from the relevant service:

```bash
make install-proto-deps
make build-proto
```

---

## CI/CD

GitHub Actions (`.github/workflows/ci-cd.yml`) automates the path from push to deploy:

1. **Test** — matrix build across all Go services (`go test ./...`, `go build`).
2. **Security scan** — Trivy filesystem scan, results uploaded to the GitHub Security tab.
3. **Build & push** — multi-stage Docker images pushed to **GitHub Container Registry** (`ghcr.io`), tagged by branch/commit.
4. **Deploy** — auto-deploy to **development** (from `develop`) and **production** (from `main`), with smoke/health checks.

See [docs/CI-CD-PIPELINE.md](docs/CI-CD-PIPELINE.md) and [docs/SETUP-GUIDE.md](docs/SETUP-GUIDE.md) for full details, branching strategy, and required GitHub secrets (`GITHUB_TOKEN`, `DOCKER_REGISTRY_TOKEN`, `KUBECONFIG`, optional `SLACK_WEBHOOK`).

---

## Contributing

1. Fork the repo and create a branch: `git checkout -b feature/xyz`
2. Make your changes and commit: `git commit -m "feat: ..."`
3. Push and open a Pull Request against `develop` (or `main` for hotfixes).
4. Ensure CI passes (tests + security scan) before requesting review.

> **Branch strategy:** `main` (production) ← `develop` (staging) ← `feature/*` / `hotfix/*`.

---

## Credits

- **Backend:** Go (microservices)
- **Databases:** PostgreSQL, MongoDB
- **Messaging:** RabbitMQ
- **Developed by:** Mostofa Meem
