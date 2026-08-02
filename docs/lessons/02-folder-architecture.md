# Lesson 02 — Folder architecture

## Target layout (monorepo)

```
agro-dashboard-finance/
  backend/                # Go API — see docs/backend.md
  frontend/
    public/               # static files served as-is (/logo.svg → public/logo.svg)
    src/
      app/                # routes only (App Router conventions)
        layout.tsx        # MUST be at app/layout.tsx (root layout)
        page.tsx
        globals.css
        favicon.ico
      components/         # reusable UI
      hooks/              # reusable client logic (custom hooks)
      lib/                # API clients, formatters, pure helpers
      types/              # shared TypeScript types
  docker-compose.yml      # Redis (local)
```

## What must stay at exact paths

Next.js looks for **convention files** at specific locations:

| Wrong | Correct |
| :--- | :--- |
| `app/layout/layout.tsx` | `app/layout.tsx` |
| `app/style/globals.css` | `app/globals.css` (or import from elsewhere) |
| `app/assets/favicon.ico` | `app/favicon.ico` |

Nesting `layout.tsx` under `app/layout/` breaks the root layout and can look like a `/layout` route segment.

## Outside `app/`

These folders are **not** routes. Use them freely:

- `components/` — cards, tables, chart wrappers
- `hooks/` — `useDashboard`, polling, filters
- `lib/` — `fetchQuotes`, date/money formatting
- `types/` — DTOs matching the Go API

Optional later: feature folders (`features/dashboard/...`) if the app grows.

## Import alias

`@/*` maps to `./src/*` (see `tsconfig.json`):

```ts
import { Button } from "@/components/Button"
```
