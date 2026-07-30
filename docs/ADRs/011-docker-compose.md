# ADR-011: Use Docker Compose for the local stack

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: ops, docker

## Context and Problem Statement

Local development needs the Go API, Redis, and eventually the Next.js app. Cloud deploy is deferred.

## Decision Drivers

- One command to bring dependencies up (especially Redis)
- Reproducible local environment
- Deploy strategy undecided — Compose is enough for now

## Considered Options

- Docker Compose for api + redis (+ frontend later)
- Run everything only on the host (brew/local installs)
- Immediate cloud deploy (Fly, Railway, etc.)

## Decision Outcome

Chosen option: **Docker Compose** as the local orchestration path. Cloud/production deploy will get its own ADR later.

### Positive Consequences

- Redis and networking match production-ish layouts early
- Easier onboarding: clone → compose up

### Negative Consequences

- Docker required for the “happy path”
- Compose files need care so Redis/API are not accidentally published widely ([ADR-007](007-local-hardened-access.md))

## Links

- [ADR-003: Redis](003-redis-cache.md)
- [ADR-007: Local hardened access](007-local-hardened-access.md)
