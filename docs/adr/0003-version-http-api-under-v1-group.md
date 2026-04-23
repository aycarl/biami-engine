# ADR 0003: Version HTTP API under `/api/v1`

- Date: 2026-04-22
- Status: Accepted
- Deciders: Core backend maintainers
- Related: `cmd/api/main.go`

## Context

The service currently exposes one endpoint (`POST /api/v1/mobile-hook`) and will likely expand. Introducing URL-based versioning early prevents route churn and enables non-breaking evolution.

## Decision

Keep all externally consumed HTTP routes under `/api/v1` using Gin route groups.

## Alternatives considered

- No explicit versioning: simplest now, but can force breaking URL migrations later.
- Header-based versioning: flexible but less visible and harder for quick manual testing.

## Consequences

### Positive

- Clear contract boundary for clients from the first endpoint.
- Easier future introduction of `/api/v2` without breaking current integrations.
- Matches common API documentation and gateway patterns.

### Negative / Trade-offs

- Version path prefix adds mild URL verbosity.
- Requires deliberate version lifecycle ownership as APIs expand.

## Revisit trigger

Revisit if API gateway constraints require non-path versioning or if internal/external API surfaces diverge significantly.

