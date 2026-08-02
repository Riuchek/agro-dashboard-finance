# Frontend — Agro Dashboard Finance

Next.js App Router dashboard that consumes the local Go API. See [docs/lessons](../docs/lessons/README.md) for stack learning notes.

## Prerequisites

- Node (see repo `.tool-versions` or `frontend/.tool-versions`)
- Backend running at `http://127.0.0.1:8080` — [docs/running-locally.md](../docs/running-locally.md)

## Run

```bash
pnpm install
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000).

## Scripts

| Command | Purpose |
| :--- | :--- |
| `pnpm dev` | Dev server |
| `pnpm build` | Production build |
| `pnpm start` | Serve production build |
| `pnpm lint` | ESLint |

## Stack

- Next.js 16 App Router, React 19, TypeScript
- Tailwind CSS 4, React Compiler
- pnpm — [ADR-008](../docs/ADRs/008-pnpm.md)

## Project layout

Target structure (see [lesson 02](../docs/lessons/02-folder-architecture.md)):

```
src/
  app/           # routes
  components/    # UI (to add)
  hooks/         # client logic (to add)
  lib/           # API client, formatters (to add)
  types/         # DTOs matching Go API (to add)
```

## API integration (planned)

The overview page will call:

- `GET http://127.0.0.1:8080/api/v1/dashboard/overview`

Until Phase 2 UI work lands, use curl against the backend ([running-locally.md](../docs/running-locally.md)).

CORS is configured via backend `CORS_ALLOW_ORIGINS=http://localhost:3000`.

## Docs

- [docs/roadmap.md](../docs/roadmap.md) — Phase 2 checklist (#13–#14)
- [docs/ADRs/010-next-app-router.md](../docs/ADRs/010-next-app-router.md)
- [AGENTS.md](./AGENTS.md) — agent notes for this Next.js version
