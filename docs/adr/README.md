# Architecture Decision Records (ADRs)

This folder stores architecture decisions for `biami-engine`.

ADRs capture **why** key technical decisions were made so future changes remain consistent, reviewable, and easy to revisit.

## Workflow

- Create a new ADR for any architecture-impacting change (framework, deployment model, data/storage strategy, API boundary/versioning, security boundary, cross-cutting tooling).
- Use the template in `docs/adr/_template.md`.
- Open new ADRs with status `Proposed`.
- Move to `Accepted` once approved in PR review.
- Never edit historical intent; create a new ADR and mark old one as `Superseded` when decisions change.

## Naming and numbering

- Filename format: `NNNN-short-kebab-title.md`
- Example: `0005-add-redis-session-store.md`
- Increment the next available number in sequence.

## Status values

- `Proposed`: Draft under discussion
- `Accepted`: Approved and active decision
- `Superseded`: Replaced by a newer ADR
- `Deprecated`: No longer recommended, without a direct replacement

## ADR index

- `0001-use-gin-framework.md` - Use Gin as the HTTP framework.
- `0002-use-multi-stage-docker-build.md` - Use multi-stage Docker image builds.
- `0003-version-http-api-under-v1-group.md` - Keep HTTP routes under `/api/v1`.
- `0004-use-compose-test-profile-for-container-tests.md` - Run tests with the Compose `test` profile.

