# ADR-002: Backend Go with stdlib (no HTTP framework)

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: backend, go, http

## Context and Problem Statement

The backend aggregates external sources (REST and scraping), normalizes data, and serves JSON. The MVP HTTP surface is small (`/health`, quotes, overview). The maintainer prefers a “purist” Go style: no web frameworks.

## Decision Drivers

- Small, stable API surface
- Full control of the request path (custom middleware)
- Fewer dependencies and upgrade surface
- Maintainer preference for stdlib

## Considered Options

- `net/http` + ServeMux (Go 1.22+)
- `go-chi/chi`
- Gin / Echo

## Decision Outcome

Chosen option: **native `net/http` with ServeMux (Go 1.22+)**, no HTTP framework. Point libraries (e.g. goquery for HTML) are allowed when they solve a specific problem; they do not replace the HTTP server.

### Positive Consequences

- No coupling to a third-party router
- Explicit middleware and routing
- Minimal module dependencies

### Negative Consequences

- More boilerplate for middleware, params, and route grouping
- Conventions must live in our code/docs (not “gifted” by a framework)

## Pros and Cons of the Options

### net/http ServeMux — Chosen

- Good: stdlib, stable, method-aware routing
- Bad: fewer ready-made helpers

### chi

- Good: ergonomic middleware and subrouters
- Bad: extra dependency without critical MVP gain

### Gin / Echo

- Good: fast DX
- Bad: abstraction/magic conflicts with the project’s stdlib-first line

## Links

- [ADR-005: goquery](005-goquery-scraping.md)
- [README](../../README.md)
