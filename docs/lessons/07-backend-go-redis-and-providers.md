# Lesson 07 — Backend: Go API, Redis, and providers

Notes from standing up the Go backend, wiring Redis, and testing the first live provider (Brapi).

## Monorepo split

| Part | Role |
| :--- | :--- |
| `backend/` | Go API — sources, cache, JSON |
| `frontend/` | Next.js — consumes backend only |
| `docker-compose.yml` | Redis for local dev |

The frontend never calls Brapi, CEPEA, or HG Brasil directly (see root [README](../../README.md)).

## Layering (rough hexagonal shape)

```
adapter/http     → handlers (HTTP in)
adapter/provider → brapi, cepea, hgbrasil (external out)
adapter/cache    → Redis / memory
app/quotes       → orchestration (cache-aside)
domain           → Quote, units, errors
port             → interfaces (CacheStore, QuoteProvider)
```

`cmd/api/main.go` is the composition root: load config, create cache, register providers, start server.

## Environment and `.env`

Config reads `os.Getenv` after optionally loading `backend/.env`:

- Run **`go run ./cmd/api` from `backend/`** so `.env` is found.
- Existing shell env vars win over `.env` (useful in CI).

Critical for cache:

```bash
REDIS_URL=redis://127.0.0.1:6379/0
```

Docker Compose maps Redis to `127.0.0.1:6379`. Starting Redis alone does **not** set this variable — you still need it in `.env`.

## Cache-aside flow

On each quote request:

1. Build key: `{source}:{key}` (e.g. `brapi:SLCE3`)
2. `GET` from Redis
3. On hit → unmarshal JSON → return
4. On miss → call provider → `SET` with source TTL → return

TTLs differ by upstream cadence — see [ADR-006](../ADRs/006-source-fetch-cadence.md) and `cache/keys.go`.

If Redis is down at startup, the factory logs and falls back to in-memory cache so you can still develop (cache lost on restart).

## Provider interface

Each source implements:

```go
Source() string
FetchQuote(ctx context.Context, key string) (domain.Quote, error)
```

Register new providers in `main.go`'s `providers` map. Handlers and cache keys already route by source name.

### What is done vs stub

| Provider | Status |
| :--- | :--- |
| Brapi | Fetches `https://brapi.dev/api/quote/{symbol}`, maps to `BRL/share` |
| CEPEA | Returns `ErrNotImplemented` — scrape with goquery ([ADR-005](../ADRs/005-goquery-scraping.md)); fixture in `testdata/cepea-boi-gordo.html` |
| HG Brasil | Returns `ErrNotImplemented` — call finance API with `HG_BRASIL_TOKEN` |

Partial degradation: `/dashboard/overview` returns stocks even when CEPEA/HG fail; errors appear in the `errors` JSON field.

## Testing manually

1. `docker compose up -d redis`
2. `cd backend && go run ./cmd/api` → expect `cache: using Redis`
3. `curl` stocks endpoint → quotes in JSON
4. `redis-cli KEYS '*'` inside the container → `brapi:…` keys

See [running-locally.md](../running-locally.md) for full curl list.

## Common pitfalls

| Symptom | Cause |
| :--- | :--- |
| `set REDIS_URL for Redis` | No `.env` or running outside `backend/` |
| `redis unavailable … falling back` | Redis container not up or wrong port |
| `listen tcp … address already in use` | Previous API process still on 8080 |
| Commodities/FX always error | Expected until CEPEA/HG `FetchQuote` is implemented |

## Links

- [backend.md](../backend.md) — folder reference
- [ADR-002](../ADRs/002-backend-go-stdlib.md) — stdlib HTTP
- [ADR-003](../ADRs/003-redis-cache.md) — Redis (implemented: `go-redis`, cache-aside, in-memory fallback)
- [ADR-012](../ADRs/012-on-demand-refresh.md) — no background ticker
