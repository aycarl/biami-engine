# biami-engine

`biami-engine` is the backend API foundation for Biami's mobile survey flows.

Today, this repository is a **Week 1 scaffold**: it boots a Gin HTTP server, exposes one mobile webhook endpoint, validates JSON payload shape, and returns a placeholder USSD-style response. It is structured for incremental growth (handlers, models, containerized runs, and basic tests) but does not yet include persistence, session orchestration, or business workflows.

## Current repository state

### Implemented

- Gin API server in `cmd/api/main.go` listening on `:8080`.
- `POST /api/v1/mobile-hook` endpoint in `internal/handler/webhook.go`.
- Request binding for `from`, `text`, and `session_id` via `internal/model/survey.go`.
- Basic handler tests in `internal/handler/webhook_test.go` (success + bad payload).
- Local and containerized developer workflows via `Makefile`, `Dockerfile`, and `docker-compose.yml`.

### Not implemented yet

- Session-aware mobile flow/state management.
- Database or external storage integration.
- Authentication/authorization and production hardening.
- Domain-level survey execution logic (current response is static placeholder text).

## Project structure

- `cmd/api/main.go`: API entrypoint and route wiring.
- `internal/handler/webhook.go`: Mobile webhook HTTP handler.
- `internal/handler/webhook_test.go`: Handler unit tests.
- `internal/model/survey.go`: Request/response model definitions.
- `tests/mobile-hook.http`: HTTP request snippet for manual webhook testing.
- `Dockerfile`: Multi-stage image build.
- `docker-compose.yml`: Compose services for API and test profile.
- `Makefile`: Common developer commands.
- `docs/roadmap/week-1.md`: Week 1 implementation/roadmap context.

## Run locally

```bash
go mod tidy
go run ./cmd/api/main.go
```

API base URL: `http://localhost:8080`

## Run with Docker Compose

```bash
docker compose up --build api
```

## Run tests

### Local

```bash
go test ./...
```

### Compose test profile

```bash
docker compose --profile test run --rm test
```

## Makefile shortcuts

```bash
make run
make build
make test
make compose-up
make compose-test
make compose-down
```

## Example webhook request

```bash
curl -X POST http://localhost:8080/api/v1/mobile-hook \
  -H "Content-Type: application/json" \
  -d '{"from": "+123456789", "text": "*123#", "session_id": "xyz123"}'
```

Expected behavior right now: HTTP `200` with a static menu-like response string on valid JSON, or HTTP `400` for invalid JSON payloads.

