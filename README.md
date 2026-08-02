# Agro Dashboard Finance

A dashboard to follow Brazilian agribusiness at a glance: physical-market prices, agri stocks on B3, FX, and public indicators — aggregated in one place.

## What it is

The **Go backend** pulls heterogeneous sources (REST APIs and HTML pages), normalizes them, and exposes a stable JSON API. The **Next.js frontend** only consumes that API and renders the dashboard.

MVP focus: a useful daily overview (cattle, soy, corn, USD, agri equities) — not a real-time trading terminal.

## High-level architecture

```
Sources (Brapi, CEPEA, HG Brasil, Yahoo, IBGE, Embrapa…)
        │
        ▼
┌───────────────────────────────┐
│  Go API (aggregator)          │
│  Providers → Service → Redis  │
│  → HTTP JSON                  │
└───────────────────────────────┘
        │
        ▼
┌───────────────────────────────┐
│  Next.js dashboard            │
└───────────────────────────────┘
```

## Current status (Aug 2026)

| Area | Status |
| :--- | :--- |
| Backend API + middleware | Done |
| Redis cache-aside + in-memory fallback | Done |
| Brapi provider (B3 stocks) | Done |
| CEPEA provider (commodities) | Stub — implement in `backend/internal/adapter/provider/cepea/` |
| HG Brasil provider (USD/BRL) | Stub — implement in `backend/internal/adapter/provider/hgbrasil/` |
| Quotes + overview HTTP handlers | Done (partial degradation when a provider fails) |
| Frontend scaffold | Done (dashboard UI not built yet) |
| Docker Compose | Redis only; API service commented out |

See [docs/roadmap.md](docs/roadmap.md) for the full checklist.

## Repository layout

```
agro-dashboard-finance/
├── backend/           # Go API
├── frontend/          # Next.js App Router (pnpm)
├── docs/
│   ├── ADRs/          # architecture decisions
│   ├── lessons/       # learning notes (frontend + backend)
│   ├── apis.md        # external data sources
│   ├── backend.md     # backend layout and providers
│   ├── running-locally.md
│   └── roadmap.md
├── docker-compose.yml # Redis (local)
└── README.md
```

## API (MVP)

| Method | Path | Description | Status |
| :--- | :--- | :--- | :--- |
| `GET` | `/health` | Liveness | Live |
| `GET` | `/api/v1/quotes/stocks?symbols=SLCE3,AGRO3` | B3 equities (Brapi) | Live |
| `GET` | `/api/v1/quotes/commodities?keys=boi-gordo,soja,milho` | Physical market (CEPEA) | Stub provider |
| `GET` | `/api/v1/quotes/fx` | USD/BRL (HG Brasil) | Stub provider |
| `GET` | `/api/v1/dashboard/overview` | Stocks + commodities + FX aggregate | Partial (stocks only until providers land) |

## How to run

Full steps, env vars, curl examples, and Redis troubleshooting: **[docs/running-locally.md](docs/running-locally.md)**.

Quick start:

```bash
docker compose up -d redis

cd backend
cp .env.sample .env   # set BRAPI_TOKEN, HG_BRASIL_TOKEN
go run ./cmd/api

cd frontend
pnpm install
pnpm dev
```

The API listens on `http://127.0.0.1:8080`. The frontend dev server uses `http://localhost:3000`.

## Documentation

| Doc | Purpose |
| :--- | :--- |
| [docs/running-locally.md](docs/running-locally.md) | Redis, `.env`, curl, cache verification |
| [docs/backend.md](docs/backend.md) | Go layout, cache keys, provider contract |
| [docs/apis.md](docs/apis.md) | External sources and units |
| [docs/roadmap.md](docs/roadmap.md) | Delivery checklist |
| [docs/lessons](docs/lessons/README.md) | Learning notes |
| [docs/ADRs](docs/ADRs/README.md) | Why we chose each technical option |
