# ADR-009: Use Recharts for dashboard charts

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia (with recommendation)
- **Tags**: frontend, charts

## Context and Problem Statement

The overview will eventually show simple series (e.g. Brapi history). We needed a React chart library that is widely used and well documented.

## Decision Drivers

- Strong React integration
- Documentation and community examples
- Good enough for MVP line/bar charts — not a Bloomberg terminal

## Considered Options

- Recharts
- Chart.js (+ react-chartjs-2)
- Apache ECharts

## Decision Outcome

Chosen option: **Recharts** as the default chart library for the Next.js dashboard.

### Positive Consequences

- Declarative React components; many dashboard tutorials
- Fast path for simple line/bar charts

### Negative Consequences

- Less ideal for extremely large datasets or exotic chart types (out of MVP scope)
- Bundle size larger than a minimal canvas chart — acceptable here

## Pros and Cons of the Options

### Recharts — Chosen

- Good: React-first DX and docs/examples
- Bad: not the most performant for huge series

### Chart.js

- Good: ubiquitous
- Bad: more imperative wrapper usage in React

### ECharts

- Good: very capable
- Bad: heavier API for simple MVP charts

## Links

- [ADR-001: Next.js](001-frontend-next-react.md)
