# biami-engine

Week 1 scaffolding for the Biami survey API.

## Structure

- `cmd/api/main.go`: API entrypoint.
- `internal/handler/webhook.go`: Mobile webhook endpoint logic.
- `internal/model/survey.go`: Survey-related request/response models.
- `Dockerfile`: Multi-stage image build.
- `docker-compose.yml`: Local container orchestration for API and tests.
- `Makefile`: Common developer commands.

## Quick start (local Go)

```bash
go mod tidy
go run ./cmd/api/main.go
```

## Quick start (Docker Compose)

```bash
docker compose up --build api
```

The API is exposed at `http://localhost:8080`.

## Run tests in Compose

```bash
docker compose --profile test run --rm test
```

## Makefile shortcuts

```bash
make compose-up
make compose-test
make compose-down
```

## Test webhook

```bash
curl -X POST http://localhost:8080/api/v1/mobile-hook \
  -H "Content-Type: application/json" \
  -d '{"from": "+123456789", "text": "*123#", "session_id": "xyz123"}'
```
