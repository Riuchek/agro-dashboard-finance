# ADR-007: Local-only hardened access (no login)

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: security, networking, auth

## Context and Problem Statement

The MVP is a personal local dashboard. Login is unnecessary if the stack is not exposed. Still, we want **maximum practical hardening** so a misconfiguration does not leak the API or tokens.

## Decision Drivers

- No multi-user product yet
- Secrets (Brapi, HG Brasil, etc.) must never ship to the browser
- Defense in depth for a Compose-based local stack

## Considered Options

- Open API on `0.0.0.0` with no auth
- Local-only bind + strict CORS + no public ports for Redis
- Full auth (sessions/JWT) from day one

## Decision Outcome

Chosen option: **no end-user login**, with **hardened local defaults**:

- API listens on **localhost** (or Compose internal network only); do not publish Redis to the host unless needed for debugging
- **CORS** allowlist: Next.js origin only (e.g. `http://localhost:3000`)
- Secrets only in backend env; never `NEXT_PUBLIC_*` for provider keys
- Security headers on API responses where relevant
- Rate limiting on API routes (protect even local abuse / loops)
- Frontend talks to backend via env base URL; no provider credentials in the browser

Auth can be introduced later if the app is exposed beyond the machine — that would supersede this ADR.

### Positive Consequences

- Simple UX (open dashboard, see data)
- Strong default posture for a learning project with real API keys

### Negative Consequences

- Not safe to expose on the public internet as-is
- CORS/bind mistakes can still happen — Compose/README must document the defaults

## Pros and Cons of the Options

### Hardened local, no login — Chosen

- Good: matches MVP users (you)
- Good: forces secrets to stay server-side
- Bad: no remote access without a future tunnel/auth story

### Open bind, no auth

- Good: easiest demo
- Bad: accidental exposure risk

### Auth from day one

- Good: internet-ready path
- Bad: scope creep for a single local user

## Links

- [ADR-011: Docker Compose](011-docker-compose.md)
- [README](../../README.md)
