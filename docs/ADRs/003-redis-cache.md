# ADR-003: Use Redis as the cache layer

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: cache, redis, backend

## Context and Problem Statement

Upstream sources have rate limits and (for CEPEA) should not be scraped aggressively. We need a shared cache with TTLs. In-memory cache would be enough for a single local process, but the maintainer wants to learn Redis and we already plan Docker Compose.

## Decision Drivers

- Learn Redis hands-on
- TTL-based caching across API restarts
- Shared cache when API runs in Compose with other services
- Align fetch policy with source cadence ([ADR-006](006-source-fetch-cadence.md))

## Considered Options

- Process memory (`sync.RWMutex` + map or go-cache)
- Redis
- No cache (hit sources every request)

## Decision Outcome

Chosen option: **Redis** as the cache-aside store for quote/indicator payloads, keyed by source + series (e.g. `cepea:boi-gordo`, `brapi:SLCE3`).

### Positive Consequences

- Real Redis experience (keys, TTL, Compose networking)
- Cache survives API process restarts
- Clear place to tune TTLs per source

### Negative Consequences

- Extra dependency and operational surface (Redis must be up)
- Overkill for a single-user local MVP — accepted as a learning trade-off
- Need failure policy when Redis is down (fail open to upstream carefully, or fail closed)

## Pros and Cons of the Options

### Redis — Chosen

- Good: learning + TTL + shared store
- Bad: more moving parts

### In-process memory

- Good: zero infra
- Bad: lost on restart; no shared learning goal

### No cache

- Good: simplest code
- Bad: hammers sources; violates CEPEA courtesy

## Links

- [ADR-006: Source fetch cadence](006-source-fetch-cadence.md)
- [ADR-011: Docker Compose](011-docker-compose.md)
- [ADR-012: On-demand refresh](012-on-demand-refresh.md)
