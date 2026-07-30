# Lesson 05 — Tooling choices (create-next-app and local setup)

## Choices used for this frontend

| Option | Choice | Why |
| :--- | :--- | :--- |
| TypeScript | Yes | Matches ADRs and safer API contracts |
| Linter | ESLint | Default Next lint setup |
| React Compiler | Yes | Auto-memoization; less manual `useMemo` / `useCallback` |
| Tailwind CSS | Yes | Fast UI iteration for dashboards |
| `src/` directory | Yes | Keeps app code separate from config files |
| App Router | Yes | See [ADR-010](../ADRs/010-next-app-router.md) |
| Import alias | `@/*` (default) | Clean imports from `src/` |
| `AGENTS.md` | Optional | Guides coding agents on current Next patterns |

## React Compiler (short)

A React build plugin that optimizes components so you rely less on hand-written memoization. Still relatively new; fine for this MVP. If a library fights the compiler, you can revisit `reactCompiler` in `next.config.ts`.

## asdf and Node

This machine uses **asdf** for Node. Shims require a version in `.tool-versions`.

asdf **0.18+** replaced `asdf local` / `asdf global` with:

```bash
asdf set nodejs 24.18.0          # current directory → .tool-versions
asdf set -u nodejs 24.18.0       # user/global
```

Prefer `.tool-versions` at the **monorepo root** so `frontend/` can stay empty for `create-next-app` (it refuses non-empty dirs) and so backend/frontend share one Node pin when needed.

## Package manager

Frontend uses **pnpm** — see [ADR-008](../ADRs/008-pnpm.md).
