# ADR 0002: Use multi-stage Docker build for runtime images

- Date: 2026-04-22
- Status: Accepted
- Deciders: Core backend maintainers
- Related: `Dockerfile`, `docker-compose.yml`

## Context

The repository ships a Go service and should be runnable in containers both locally and in deployment environments. We want predictable builds and small runtime images without build toolchains in production containers.

## Decision

Use a multi-stage Docker build:

1. Build in a Go builder image.
2. Copy only the compiled binary into a minimal runtime image.

## Alternatives considered

- Single-stage Go image: simpler Dockerfile, but larger image and unnecessary build tools in runtime.
- Distroless runtime immediately: tighter security posture, but can increase debugging friction in early-stage development.

## Consequences

### Positive

- Smaller runtime image and faster image distribution.
- Better separation between build-time and run-time concerns.
- Cleaner path toward production hardening.

### Negative / Trade-offs

- Slightly more Dockerfile complexity.
- Additional consideration for debugging in minimal runtime images.

## Revisit trigger

Revisit when deployment security requirements tighten (for example, move to distroless), or if debugging/operability needs a different runtime base.

