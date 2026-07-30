# ADR-013: Centralize docs under `/docs` in English

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia
- **Tags**: docs, process

## Context and Problem Statement

Early notes lived under `backend/docs/` in Portuguese while ADRs started under `docs/ADRs/`. The project needs one documentation root and one language.

## Decision Drivers

- Single place for humans and future contributors
- English for code, docs, commits, and ADRs
- Clear separation: product/architecture docs vs application code

## Considered Options

- Keep docs beside each app (`backend/docs`, `frontend/docs`)
- Central `docs/` at repo root (ADRs, APIs, roadmap)
- Mixed PT/EN

## Decision Outcome

Chosen option: **all project documentation under `/docs`**, written in **English**. Application packages stay in `backend/` and `frontend/` without parallel doc trees (except code-adjacent comments if ever needed — project rule still prefers self-explanatory code).

Layout:

```
docs/
├── ADRs/
├── apis.md
└── roadmap.md
```

### Positive Consequences

- One index to link from the README
- Consistent language across the repo

### Negative Consequences

- Contributors must write English even for informal notes
- Backend-only clones still need the root `docs/` tree

## Links

- [docs/ADRs](./README.md)
- [README](../../README.md)
