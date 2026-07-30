# Agro Dashboard Finance

A dashboard to follow Brazilian agribusiness at a glance: physical-market prices, agri stocks on B3, FX, and public indicators — aggregated in one place.

## What it is

The **Go backend** pulls heterogeneous sources (REST APIs and HTML pages), normalizes them, and exposes a stable JSON API. The **Next.js frontend** only consumes that API and renders the dashboard.

MVP focus: a useful daily overview (cattle, soy, corn, USD, agri equities) — not a real-time trading terminal.

## High-level architecture

```
Sources (Brapi, CEPEA, Yahoo, IBGE, HG Brasil, Embrapa…)
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

## Repository layout

```
agro-dashboard-finance/
├── backend/           # Go API
├── frontend/          # Next.js app (to be created)
├── docs/
│   ├── ADRs/          # architecture decisions
│   ├── apis.md        # external data sources
│   └── roadmap.md     # delivery checklist
├── docker-compose.yml # local stack (to be created)
└── README.md
```

## Product principles

1. **Stable API** — the frontend never talks to CEPEA, Brapi, or raw HTML.
2. **Isolated sources** — one series failing must not take down the whole overview.
3. **Respect upstream limits** — especially CEPEA scraping (daily indicators; long TTLs).
4. **Local and locked down** — no login in the MVP; the stack stays private and hardened.
5. **Thin MVP** — no history DB; display current quotes only.

## Planned API (MVP)

| Method | Path | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Liveness |
| `GET` | `/api/v1/quotes/stocks?symbols=SLCE3,AGRO3` | B3 equities |
| `GET` | `/api/v1/quotes/commodities?keys=boi-gordo,soja,milho` | Physical market |
| `GET` | `/api/v1/quotes/fx` | USD/BRL |
| `GET` | `/api/v1/dashboard/overview` | Dashboard home aggregate |

## How to run

Once scaffolds exist:

```bash
docker compose up

# or separately:
cd backend && cp .env.sample .env && go run ./cmd/api
cd frontend && pnpm install && pnpm dev
```

## Documentation

| Doc | Purpose |
| :--- | :--- |
| [docs/apis.md](docs/apis.md) | External sources and units |
| [docs/roadmap.md](docs/roadmap.md) | Delivery checklist |
| [docs/ADRs](docs/ADRs/README.md) | Why we chose each technical option |