# biami-survey-engine

Week 1 scaffolding for the Biami survey API.

## Structure

- `cmd/api/main.go`: API entrypoint.
- `internal/handler/webhook.go`: Mobile webhook endpoint logic.
- `internal/model/survey.go`: Survey-related request/response models.
- `Dockerfile`: Multi-stage image build.
- `Makefile`: Common developer commands.

## Quick start

```bash
go mod tidy
go run ./cmd/api/main.go
```

## Test webhook

```bash
curl -X POST http://localhost:8080/api/v1/mobile-hook \
  -H "Content-Type: application/json" \
  -d '{"from": "+123456789", "text": "*123#", "session_id": "xyz123"}'
```

