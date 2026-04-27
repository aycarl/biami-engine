# AI Agent Instructions for Biami Engine

Welcome! This document provides context, guidelines, and architectural details for AI coding agents and assistants working in the `biami-engine` repository.

## Project Context
`biami-engine` is the backend API foundation for Biami's mobile survey flows. It is designed to handle high-concurrency mobile webhooks for USSD and SMS sessions, processing dynamic survey and respondent metadata. The roadmap is heavily focused on moving towards a production-grade, autonomous cloud architecture.

## Tech Stack
- **Language**: Go (v1.25.0+)
- **Web Framework**: Gin (`github.com/gin-gonic/gin`)
- **Containerization**: Docker (multi-stage builds) & Docker Compose
- **Infrastructure (Roadmap)**: Terraform, AWS (ECS Fargate, Lambda, SQS, RDS PostgreSQL, ElastiCache Redis)

## Architecture & Directory Structure
The repository strictly adheres to a standard Go project layout to separate concerns and ensure maintainability:
- `cmd/api/main.go`: Application entry point and route wiring. Keep this file minimal.
- `internal/handler/`: HTTP handlers containing webhook logic (e.g., `webhook.go`). Handlers should orchestrate, not contain heavy business logic.
- `internal/model/`: Data structures and JSON schemas (e.g., `survey.go`). Uses `json.RawMessage` for dynamic, unstructured survey payload data.
- `docs/roadmap/`: Contains the 13-week roadmap, architectural context, and milestones. Reference this to understand the project's trajectory.
- `tests/`: Resources for integration and manual testing (e.g., `.http` files).

## Coding Standards & Guidelines
When contributing to `biami-engine`, AI agents must adhere to the following rules:

1. **Go Best Practices**: Write idiomatic Go code. Avoid global state. Keep functions small, testable, and strictly typed. Handle errors gracefully and explicitly.
2. **Gin Framework**: Utilize `gin` contexts for request binding (`ShouldBindJSON`), response generation, and middleware integration.
3. **Dynamic Data Handling**: Surveys are highly dynamic. Retain support for unstructured data by utilizing `json.RawMessage` in payloads, rather than locking down schemas unnecessarily.
4. **Design for Serverless/Cloud**: Write code anticipating a decoupled, event-driven architecture. Complex domain logic (like quality scoring) is slated to be offloaded to worker queues (e.g., SQS/Lambda). Ensure tight boundaries around your domains.
5. **Testing**: Write unit tests for all new code, especially `internal` packages. Ensure `make test` executes cleanly. 
6. **Developer Tooling**: Use the existing `Makefile` (`make run`, `make build`, `make test`, `make compose-up`). Ensure any new dependencies or environment variables are reflected in `docker-compose.yml` and the `Dockerfile`.

## Roadmap Awareness
The project is split into distinct phases:
- **Phase 0 (Current)**: The Core App (Go/Gin API, mobile webhooks, basic JSON validation).
- **Phase 1**: Infrastructure & Orchestration (Terraform, RDS PostgreSQL, Redis, ECS Fargate).
- **Phase 2**: Serverless & Security (Lambda, SQS, WAF, S3/CloudFront).
- **Phase 3**: Observability & CI/CD (CloudWatch, GitHub Actions).

*Before introducing a new design pattern or dependency, verify it aligns with the architectural goals outlined in `docs/roadmap/index.md`.*
