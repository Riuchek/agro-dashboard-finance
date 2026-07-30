# ADR-010: Use Next.js App Router

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: frontend, nextjs

## Context and Problem Statement

Next.js supports App Router (current default) and Pages Router. The maintainer wants to learn the current default.

## Decision Drivers

- Align with Next.js defaults and current docs
- Learning goal for App Router / React Server Components patterns
- Long-term relevance vs legacy Pages Router

## Considered Options

- App Router (`app/`)
- Pages Router (`pages/`)

## Decision Outcome

Chosen option: **App Router** (`app/` directory) for the frontend.

### Positive Consequences

- Matches official Next.js learning path
- Layouts and server/client component split available when useful

### Negative Consequences

- Steeper learning curve (what runs on server vs client)
- Dashboard interactivity still needs client components for charts/filters — expected

## Links

- [ADR-001: Next.js + React](001-frontend-next-react.md)
