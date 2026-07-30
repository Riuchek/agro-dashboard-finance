# ADR-001: Use Next.js and React for the frontend

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia
- **Tags**: frontend, stack

## Context and Problem Statement

The frontend consumes the Go API and renders agribusiness dashboards (cards, tables, charts). We needed to choose between Nuxt/Vue and Next/React. The backend contract is JSON only; heavy SEO and marketing pages are out of MVP scope.

## Decision Drivers

- Broad ecosystem (charts, grids, dashboard examples)
- Easier collaboration / hiring later
- First-class TypeScript
- Willingness to accept more ceremony for better libraries and learning material

## Considered Options

- Next.js + React + TypeScript
- Nuxt 4 + Vue 3 + TypeScript

## Decision Outcome

Chosen option: **Next.js + React + TypeScript**, primarily for ecosystem density and learning material, even though Nuxt can feel lighter for a simple card/table MVP.

### Positive Consequences

- More chart/grid options and reusable snippets
- Stack aligned with mainstream frontend market

### Negative Consequences

- More decisions (App Router, client vs server components) than a plain SPA needs
- Poorly structured hooks/polling can get messy quickly — discipline required

## Pros and Cons of the Options

### Next.js + React — Chosen

- Good: ecosystem and examples
- Good: hiring / collaboration
- Bad: mental overhead vs a thinner SPA

### Nuxt 4 + Vue 3

- Good: lighter templates/DX for dashboard UI
- Bad: fewer libs/examples in finance/dashboard niches

## Links

- [ADR-010: Next.js App Router](010-next-app-router.md)
- [README](../../README.md)
