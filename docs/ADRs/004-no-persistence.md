# ADR-004: Do not persist quote history

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: storage, scope

## Context and Problem Statement

Dashboards often grow a time-series DB for charts. This project’s goal is to **display current aggregated quotes**, not build a historical warehouse.

## Decision Drivers

- Scope: show, don’t archive
- Avoid owning ETL/backfill/schema early
- Historical series can come from upstream APIs later (e.g. Brapi) when needed for charts

## Considered Options

- No persistence (display-only + Redis cache)
- SQLite history
- PostgreSQL history

## Decision Outcome

Chosen option: **no database for quotes**. Redis holds short-lived cached responses only. Charts that need history should call providers that already expose it, or stay out of MVP.

### Positive Consequences

- Smaller architecture and ops load
- Clear MVP boundary

### Negative Consequences

- No offline archive of past CEPEA days in our system
- Cannot build custom multi-year analytics without revisiting this ADR

## Pros and Cons of the Options

### No persistence — Chosen

- Good: matches product intent
- Bad: no owned history

### SQLite / Postgres

- Good: owned time series
- Bad: out of scope for “just show”

## Links

- [ADR-003: Redis cache](003-redis-cache.md)
- [README](../../README.md)
