# ADR-005: Use goquery for CEPEA HTML scraping

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Correia
- **Tags**: scraping, backend, cepea

## Context and Problem Statement

CEPEA/ESALQ physical-market indicators have no public REST API for our use case. We must parse a small set of HTML indicator pages.

## Decision Drivers

- Few pages, CSS-selector friendly markup
- Prefer a focused library over a crawling framework
- Align with stdlib-first HTTP client usage ([ADR-002](002-backend-go-stdlib.md))

## Considered Options

- `github.com/PuerkitoBio/goquery`
- `github.com/gocolly/colly`
- Manual `net/html` only

## Decision Outcome

Chosen option: **goquery** for CEPEA (and similar one-off HTML providers). Fetch with `net/http`, parse with goquery, map into domain types.

### Positive Consequences

- jQuery-like selectors, easy fixtures in tests
- Small dependency footprint for the job

### Negative Consequences

- HTML structure changes break parsers — need fixtures and careful selectors
- Not a full crawler (acceptable; we do not need one)

## Pros and Cons of the Options

### goquery — Chosen

- Good: right size for a handful of pages
- Bad: no built-in crawl scheduling (we do not want that)

### colly

- Good: crawling toolkit
- Bad: heavier than needed

### net/html only

- Good: zero deps
- Bad: painful selectors for recurring scrapes

## Links

- [ADR-006: Source fetch cadence](006-source-fetch-cadence.md)
- [docs/apis.md](../apis.md)
