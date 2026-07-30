# ADR-006: Source fetch cadence and cache TTLs

- **Date**: 2026-07-30
- **Status**: Accepted
- **Deciders**: João Marinho
- **Tags**: cache, scraping, rate-limit, sources

## Context and Problem Statement

We must not hammer upstream sources. Cadence differs a lot: CEPEA publishes a **single daily** indicator on business days; equity quotes move during market hours. TTLs and scrape limits should follow that reality.

## Decision Drivers

- CEPEA methodology: collection through the day, indicator closed ~17:00–18:00 BRT, published after ~18:00 (about **18:41** for boi gordo per CEPEA FAQ); one official value per business day — not intraday for the main index
- Equity/FX APIs update much more often and have their own free-tier limits
- Courtesy User-Agent and low request volume for HTML scraping

## Considered Options

- Same short TTL for every source
- Per-source TTLs + hard min interval between upstream calls
- Background polling every minute for all keys

## Decision Outcome

Chosen option: **per-source Redis TTLs** and a **minimum interval** before re-hitting the same upstream URL/key (cache-aside; see [ADR-012](012-on-demand-refresh.md)).

### Baseline policy (MVP)

| Source | Upstream cadence (practical) | Redis TTL | Min interval between upstream hits |
| :--- | :--- | :--- | :--- |
| CEPEA (physical) | ~1× per business day after ~18:00 BRT | **12h** | **6h** per indicator URL |
| Brapi (B3 stocks) | Intraday while market open | **2–5 min** | Same as TTL |
| HG Brasil (FX) | Intraday / daily free-tier limits | **10–15 min** | Same as TTL |
| Yahoo (futures/stocks) | Intraday | **2–5 min** | Same as TTL |
| IBGE SIDRA | Infrequent statistical releases | **24h** | **24h** |

CEPEA scrapes: identifiable User-Agent, no parallel fan-out across many pages on cold start without spacing, prefer warming only overview keys.

### Positive Consequences

- Matches how sources actually change
- Protects CEPEA and free API quotas
- Easy to tune via config/env

### Negative Consequences

- Overview may show yesterday’s CEPEA value until after publication + first cache miss
- TTL tables need maintenance when adding sources

## Pros and Cons of the Options

### Per-source TTL — Chosen

- Good: honest to upstream behavior
- Bad: more config knobs

### Uniform short TTL

- Good: simple
- Bad: abusive for CEPEA

### Aggressive background polling

- Good: always “warm”
- Bad: wasteful and rude to scraped sites

## Links

- [ADR-003: Redis](003-redis-cache.md)
- [ADR-005: goquery](005-goquery-scraping.md)
- [CEPEA boi gordo FAQ](https://cepea.org.br/br/faq-do-indicador-do-boi-gordo-cepea-esalq.aspx)
