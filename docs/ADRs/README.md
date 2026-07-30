# Architecture Decision Records (ADRs)

Significant architecture decisions for Agro Dashboard Finance.  
File pattern: `NNN-topic-in-kebab-case.md` (MADR).

ADRs are historical records. If a decision changes, add a new ADR and mark the old one **Superseded**.

Project language: **English** (docs, code, commits).

| # | Decision | Status |
| :--- | :--- | :--- |
| [001](001-frontend-next-react.md) | Frontend Next.js + React | Accepted |
| [002](002-backend-go-stdlib.md) | Backend Go stdlib (`net/http`) | Accepted |
| [003](003-redis-cache.md) | Redis as cache layer | Accepted |
| [004](004-no-persistence.md) | No quote history database | Accepted |
| [005](005-goquery-scraping.md) | goquery for CEPEA HTML | Accepted |
| [006](006-source-fetch-cadence.md) | Fetch cadence and TTLs by source | Accepted |
| [007](007-local-hardened-access.md) | Local-only, hardened, no auth | Accepted |
| [008](008-pnpm.md) | pnpm for frontend packages | Accepted |
| [009](009-recharts.md) | Recharts for charts | Accepted |
| [010](010-next-app-router.md) | Next.js App Router | Accepted |
| [011](011-docker-compose.md) | Docker Compose for local stack | Accepted |
| [012](012-on-demand-refresh.md) | On-demand cache-aside (no background ticker) | Accepted |
| [013](013-docs-layout-english.md) | Docs under `/docs`, English everywhere | Accepted |
