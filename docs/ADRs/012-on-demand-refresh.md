# ADR-012: On-demand cache-aside refresh (no background ticker)

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia (with recommendation)
- **Tags**: cache, backend

## Context and Problem Statement

“Background ticker” means a goroutine/timer that **periodically refreshes cache keys even when nobody is using the dashboard**. The alternative is **on-demand (cache-aside)**: on each API request, read Redis; on miss (or expiry), fetch upstream, store with TTL, return.

Given CEPEA’s daily cadence and a local single-user MVP, continuous background refresh adds complexity and unnecessary upstream traffic.

## Decision Drivers

- Respect source cadence ([ADR-006](006-source-fetch-cadence.md))
- Keep MVP simple
- Still feel “fast” via Redis TTLs after the first load

## Considered Options

- On-demand cache-aside only
- Background ticker for hot keys
- Hybrid (on-demand + optional warmer later)

## Decision Outcome

Chosen option: **on-demand cache-aside only** for the MVP.

Flow: `request → Redis GET → hit? return : fetch provider → Redis SET+TTL → return`.

A background warmer may be reconsidered later if cold starts feel bad; that would be a new ADR or a superseding update.

### Positive Consequences

- No idle scraping/API calls
- Easy to reason about and test
- Natural fit with Redis TTLs

### Negative Consequences

- First request after TTL expiry pays the upstream latency
- Overview may briefly wait on CEPEA HTML parse after a long idle period

## Pros and Cons of the Options

### On-demand — Chosen

- Good: simple, polite to upstream
- Bad: cold-request latency

### Background ticker

- Good: warm cache
- Bad: traffic even with zero users; easy to over-fetch CEPEA

## Links

- [ADR-003: Redis](003-redis-cache.md)
- [ADR-006: Source fetch cadence](006-source-fetch-cadence.md)
