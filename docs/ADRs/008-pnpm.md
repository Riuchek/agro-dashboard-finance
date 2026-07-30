# ADR-008: Use pnpm for frontend packages

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia
- **Tags**: frontend, tooling

## Context and Problem Statement

The Next.js app needs a package manager. Mixing managers causes lockfile drift.

## Decision Drivers

- Fast installs and strict dependency layout
- Maintainer preference
- Single lockfile convention for `frontend/`

## Considered Options

- pnpm
- yarn
- npm

## Decision Outcome

Chosen option: **pnpm** for all frontend package operations (`pnpm install`, `pnpm dev`, `pnpm build`).

### Positive Consequences

- Efficient disk use and deterministic installs
- Clear convention in README/Compose docs

### Negative Consequences

- Contributors must have pnpm available (Corepack or install)

## Links

- [ADR-001: Next.js](001-frontend-next-react.md)
