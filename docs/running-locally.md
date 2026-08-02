# Running locally

How to start Redis, the Go API, and the Next.js frontend on your machine.

## Prerequisites

- Go 1.22+
- Docker (for Redis)
- pnpm (frontend)
- API tokens: [Brapi](https://brapi.dev), [HG Brasil](https://hgbrasil.com/status/finance) (HG Brasil provider is still a stub — token is for when you implement it)

## 1. Redis (Docker Compose)

From the repo root:

```bash
docker compose up -d redis
```

Redis is published on **`127.0.0.1:6379`** only (see [ADR-007](ADRs/007-local-hardened-access.md)).

Verify:

```bash
docker exec $(docker ps -qf "ancestor=redis:7-alpine") redis-cli PING
```

Expected: `PONG`.

## 2. Backend environment

```bash
cd backend
cp .env.sample .env
```

Edit `.env`:

| Variable | Required | Example |
| :--- | :--- | :--- |
| `HTTP_ADDR` | No (default `127.0.0.1:8080`) | `127.0.0.1:8080` |
| `REDIS_URL` | Yes for Redis cache | `redis://127.0.0.1:6379/0` |
| `BRAPI_TOKEN` | Yes for stocks | your Brapi token |
| `HG_BRASIL_TOKEN` | For FX (when implemented) | your HG Brasil key |
| `CORS_ALLOW_ORIGINS` | No | `http://localhost:3000` |

**Important:** the app loads `.env` from the **current working directory**. Always run the API from `backend/`:

```bash
cd backend
go run ./cmd/api
```

Startup logs:

- `cache: using Redis` — connected to Redis
- `cache: using in-memory store (set REDIS_URL for Redis)` — `REDIS_URL` missing or empty
- `cache: redis unavailable (...), falling back to in-memory store` — Redis not reachable

If port 8080 is busy:

```bash
lsof -i :8080 -t | xargs kill
```

## 3. Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000). The dashboard UI is not wired to the API yet — Phase 2 work.

## 4. Smoke-test the API (curl)

With the API running:

```bash
curl -s http://127.0.0.1:8080/health
```

```bash
curl -s "http://127.0.0.1:8080/api/v1/quotes/stocks?symbols=SLCE3,AGRO3"
```

```bash
curl -s "http://127.0.0.1:8080/api/v1/quotes/commodities?keys=boi-gordo,soja,milho"
```

```bash
curl -s http://127.0.0.1:8080/api/v1/quotes/fx
```

```bash
curl -s http://127.0.0.1:8080/api/v1/dashboard/overview | jq
```

### Expected behavior today

| Endpoint | Result |
| :--- | :--- |
| `/health` | `{"status":"ok"}` |
| `/quotes/stocks` | Quotes from Brapi |
| `/quotes/commodities` | Empty quotes + CEPEA "not implemented yet" errors |
| `/quotes/fx` | 502 + HG Brasil "not implemented yet" |
| `/dashboard/overview` | Stocks filled; commodities/FX empty; errors listed |

Partial failures are intentional — one broken provider must not take down the overview.

## 5. Verify Redis cache

After hitting stocks:

```bash
docker exec $(docker ps -qf "ancestor=redis:7-alpine") redis-cli KEYS '*'
```

Example keys: `brapi:SLCE3`, `brapi:AGRO3`.

Check TTL (Brapi ≈ 3 minutes):

```bash
docker exec $(docker ps -qf "ancestor=redis:7-alpine") redis-cli TTL 'brapi:SLCE3'
```

First request fetches upstream and writes Redis; subsequent requests read from cache until TTL expires.

## 6. Running the API inside Docker (optional)

The Compose `api` service is commented out. If you enable it, set:

```bash
REDIS_URL=redis://redis:6379/0
```

Use the Compose service name `redis`, not `127.0.0.1` (that points to the API container itself).

## Links

- [backend.md](backend.md) — code layout and provider stubs
- [apis.md](apis.md) — upstream sources and keys
- [lessons/07-backend-go-redis-and-providers.md](lessons/07-backend-go-redis-and-providers.md) — how the pieces fit together
