# Backend (Go API)

The backend aggregates external quote sources behind a stable JSON API. It uses **cache-aside** with Redis and degrades gracefully when individual providers fail.

## Layout

```
backend/
├── cmd/api/main.go              # wiring: config, cache, providers, HTTP server
├── internal/
│   ├── config/                  # env + .env file loader
│   ├── domain/                  # Quote, units, shared errors
│   ├── port/                    # CacheStore, QuoteProvider interfaces
│   ├── app/
│   │   ├── quotes/              # cache-aside orchestration
│   │   └── overview/            # default symbols/keys aggregate
│   └── adapter/
│       ├── cache/               # Redis + in-memory + keys/TTLs
│       ├── http/                # handlers, middleware, server
│       └── provider/
│           ├── brapi/           # ✅ implemented
│           ├── cepea/           # stub — implement FetchQuote
│           └── hgbrasil/        # stub — implement FetchQuote
├── testdata/                    # HTML fixtures for CEPEA tests (later)
├── .env.sample
└── go.mod
```

## Request flow

```
HTTP handler
    → quotes.Service (or overview.Service)
        → cache.Get(key)
            hit  → return cached Quote JSON
            miss → provider.FetchQuote → cache.Set(TTL) → return
```

Cache keys: `{source}:{key}` — e.g. `brapi:SLCE3`, `cepea:boi-gordo`, `hgbrasil:usd-brl`.

## TTLs by source

| Source | TTL | Rationale |
| :--- | :--- | :--- |
| `brapi` | 3 min | Near-real-time equities |
| `cepea` | 12 h | Daily indicators ([ADR-006](ADRs/006-source-fetch-cadence.md)) |
| `hgbrasil` | 12 min | FX refresh cadence |

Defined in `internal/adapter/cache/keys.go`.

## Provider contract

Every provider implements `port.QuoteProvider`:

```go
Source() string
FetchQuote(ctx context.Context, key string) (domain.Quote, error)
```

Return a `domain.Quote` with the correct `Unit` (`BRL/share`, `BRL/@`, `BRL/bag`, etc.).

### Integration status

| Provider | Package | Keys | Status |
| :--- | :--- | :--- | :--- |
| Brapi | `adapter/provider/brapi` | Stock tickers (`SLCE3`, `AGRO3`, …) | **Done** — calls `brapi.dev/api/quote/{symbol}` |
| CEPEA | `adapter/provider/cepea` | `boi-gordo`, `soja`, `milho` | **Stub** — returns `ErrNotImplemented`; use goquery + [ADR-005](ADRs/005-goquery-scraping.md) |
| HG Brasil | `adapter/provider/hgbrasil` | `usd-brl` | **Stub** — returns `ErrNotImplemented` when token set |

Stub providers are already registered in `cmd/api/main.go`. Implementing `FetchQuote` is enough — handlers and cache wiring already exist.

## Shared errors

| Error | Meaning |
| :--- | :--- |
| `ErrNotConfigured` | Missing API token in env |
| `ErrNotImplemented` | Provider stub |
| `ErrUpstream` | HTTP/network failure from external API |
| `ErrQuoteNotFound` | Upstream responded but series missing |

Handlers surface per-key errors in JSON (`errors` array) without failing the whole response when possible.

## HTTP routes

Registered in `internal/adapter/http/server.go`:

| Method | Path |
| :--- | :--- |
| `GET` | `/health` |
| `GET` | `/api/v1/quotes/stocks?symbols=…` |
| `GET` | `/api/v1/quotes/commodities?keys=…` |
| `GET` | `/api/v1/quotes/fx` |
| `GET` | `/api/v1/dashboard/overview` |

## Configuration

Loaded from process env; unset keys can come from `backend/.env` when you run from that directory.

| Env var | Purpose |
| :--- | :--- |
| `HTTP_ADDR` | Listen address (default `127.0.0.1:8080`) |
| `REDIS_URL` | Redis connection URL; empty → in-memory cache |
| `BRAPI_TOKEN` | Brapi API token |
| `HG_BRASIL_TOKEN` | HG Brasil key |
| `CORS_ALLOW_ORIGINS` | Comma-separated frontend origins |

## Redis fallback

If `REDIS_URL` is empty, the API uses an in-memory cache (`factory.go`). If Redis is configured but unreachable at startup, it logs a warning and falls back to memory. For local dev with Docker Redis, expect `cache: using Redis` on startup.

## Related docs

- [running-locally.md](running-locally.md) — start commands and curl
- [apis.md](apis.md) — upstream URLs and units
- [lessons/07-backend-go-redis-and-providers.md](lessons/07-backend-go-redis-and-providers.md)
