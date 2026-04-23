# ADR 0004: Use Docker Compose `test` profile for containerized test execution

- Date: 2026-04-22
- Status: Accepted
- Deciders: Core backend maintainers
- Related: `docker-compose.yml`, `Makefile`

## Context

The project supports local Go runs and container runs. We need a repeatable way to run tests in a containerized environment that is close to CI/runtime tooling and easy for contributors.

## Decision

Use a dedicated Docker Compose service profile (`test`) to execute `go test ./...` in containers.

## Alternatives considered

- Local-only testing (`go test ./...`): fastest feedback, but weaker parity with containerized environments.
- Separate bespoke CI scripts without Compose: works in CI, but less discoverable for local contributors.

## Consequences

### Positive

- Improves local/CI parity for test execution.
- Keeps test workflow discoverable via `docker-compose.yml` and `make compose-test`.
- Provides a foundation for future service dependencies in integration tests.

### Negative / Trade-offs

- Slower than local-only tests due to container startup overhead.
- Requires Docker availability on contributor machines.

## Revisit trigger

Revisit when test suite size grows and we need parallelized CI-native orchestration, or when test dependencies exceed what Compose profiles handle cleanly.

