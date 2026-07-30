# Roadmap — delivery checklist

Check items as you go. Each can become a GitHub Issue.

---

## Phase 0 — Foundation

- [ ] **#1** Backend scaffold: `net/http` (Go 1.22+ ServeMux), env config, middleware (CORS allowlist, recover, request log), bind localhost, `GET /health`
- [ ] **#2** `domain` package: `Quote`, `Indicator`, units (`BRL/@`, `BRL/bag`, etc.), shared errors
- [ ] **#3** Redis cache client + typed keys + per-source TTLs ([ADR-003](ADRs/003-redis-cache.md), [ADR-006](ADRs/006-source-fetch-cadence.md))
- [ ] **#4** Documented `.env.sample` (Brapi / HG Brasil / Embrapa / Redis) and `.gitignore` for `.env`
- [ ] **#5** Docker Compose: API + Redis (hardened ports) ([ADR-011](ADRs/011-docker-compose.md))

## Phase 1 — Providers + quotes API (MVP dashboard)

- [ ] **#6** **Brapi** provider: B3 agri stocks (`SLCE3`, `AGRO3`, `SMTO3`…)
- [ ] **#7** **CEPEA** provider (goquery): cattle, soy, corn (later coffee/sugar/cotton)
- [ ] **#8** **HG Brasil** provider: USD/BRL
- [ ] **#9** Orchestration service: cache-aside + partial degradation when one source fails
- [ ] **#10** Handlers: `GET /api/v1/quotes/stocks`, `/commodities`, `/fx`
- [ ] **#11** `GET /api/v1/dashboard/overview` aggregating stocks + commodities + fx

## Phase 2 — Frontend MVP

- [ ] **#12** Frontend scaffold (Next.js App Router + React + TS, pnpm) calling local API
- [ ] **#13** Overview page: price cards (cattle, soy, corn, USD) + agri stocks table
- [ ] **#14** Simple charts with Recharts (history only when upstream provides it — e.g. Brapi)

## Phase 3 — More sources and robustness

- [ ] **#15** **Yahoo Finance** provider (CME futures `ZC=F`, `ZS=F` and equity fallback)
- [ ] **#16** **IBGE SIDRA** provider (selected PAM/PPM series)
- [ ] **#17** Unit tests for CEPEA HTML fixtures and handlers with `httptest`
- [ ] **#18** Basic API rate limiting ([ADR-007](ADRs/007-local-hardened-access.md))

## Phase 4 — Nice to have

- [ ] **#19** Embrapa AgroAPI provider (climate / ZARC) if token is available
- [ ] **#20** Compose profile including frontend
- [ ] **#21** Optional background cache warmer (only if cold starts hurt — would need a new ADR)

---

## Suggested order

`#1 → #2 → #3 → #4 → #5 → #6 → #7 → #10 → #11 → #12 → #13`

That path yields a useful local dashboard. Everything else is on demand.
