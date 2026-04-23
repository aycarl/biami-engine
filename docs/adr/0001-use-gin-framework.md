# ADR 0001: Use Gin as the HTTP framework

- Date: 2026-04-22
- Status: Accepted
- Deciders: Core backend maintainers
- Related: `cmd/api/main.go`, `internal/handler/webhook.go`, `go.mod`

## Context

`biami-engine` is currently a small API with one webhook endpoint and basic request validation. We need a framework that is easy to adopt now, supports middleware as complexity grows, and remains performant for mobile webhook traffic.

## Decision

Use `github.com/gin-gonic/gin` as the HTTP framework for API routing, request binding, and response handling.

## Alternatives considered

- Standard library (`net/http`) + custom helpers: minimal dependencies, but more boilerplate for routing, binding, and middleware composition.
- `chi`: lightweight and composable, but would still require additional decisions/utilities for request binding patterns currently covered by Gin.

## Consequences

### Positive

- Fast onboarding with concise route and handler code.
- Built-in JSON binding (`ShouldBindJSON`) aligns with current webhook payload handling.
- Strong middleware ecosystem for future auth, tracing, and request controls.

### Negative / Trade-offs

- Adds a framework dependency and upgrade surface area.
- Framework conventions influence handler architecture over time.

## Revisit trigger

Revisit if we need finer-grained control over allocations/latency at high throughput, or if framework constraints significantly slow feature delivery.

