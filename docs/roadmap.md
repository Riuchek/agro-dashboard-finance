# Roadmap — delivery checklist

Last updated: **Aug 2026**.

## Phase 0 — Foundation

- [x] **#1** Backend scaffold: `net/http` (Go 1.22+ ServeMux), env config, middleware (CORS allowlist, recover, request log), bind localhost, `GET /health`
- [x] **#2** `domain` package: `Quote`, units (`BRL/@`, `BRL/bag`, `BRL/share`), shared errors
- [x] **#3** Redis cache client + typed keys + per-source TTLs ([ADR-003](ADRs/003-redis-cache.md), [ADR-006](ADRs/006-source-fetch-cadence.md))
- [x] **#4** Documented `.env.sample` (Brapi / HG Brasil / Redis) and `.gitignore` for `.env`
- [x] **#5** Docker Compose: Redis with localhost bind ([ADR-011](ADRs/011-docker-compose.md)) — API service still optional/commented

## Phase 1 — Providers + quotes API (MVP dashboard)

- [x] **#6** **Brapi** provider: B3 agri stocks (`SLCE3`, `AGRO3`, `SMTO3`…)
- [ ] **#7** **CEPEA** provider (goquery): cattle, soy, corn — stub registered; scrape logic pending in `cepea/scraper.go` (fixture: `backend/testdata/cepea-boi-gordo.html`)
- [ ] **#8** **HG Brasil** provider: USD/BRL — stub registered; API client pending in `hgbrasil/client.go`
- [x] **#9** Orchestration service: cache-aside + partial degradation when one source fails
- [x] **#10** Handlers: `GET /api/v1/quotes/stocks`, `/commodities`, `/fx`
- [x] **#11** `GET /api/v1/dashboard/overview` aggregating stocks + commodities + fx (commodities/FX populate once #7/#8 land)

## Phase 2 — Frontend MVP

- [x] **#12** Frontend scaffold (Next.js App Router + React + TS, pnpm)
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
