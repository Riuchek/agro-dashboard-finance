# External APIs and data sources — Brazilian agribusiness

Reference for APIs and sources used to monitor agri equities, futures, physical commodity prices, and public indicators.

---

## 1. Financial market and equities (B3 and futures)

| Source / API | Access | Auth / cost | Provides | Examples |
| :--- | :--- | :--- | :--- | :--- |
| **Brapi**<br>`brapi.dev` | REST (JSON) | Free (token / soft limits) | Near-real-time and historical quotes for B3-listed stocks | `GET https://brapi.dev/api/quote/SLCE3`<br>`GET https://brapi.dev/api/quote/AGRO3,SMTO3` |
| **Yahoo Finance** | REST / scraping | Free (no key) | B3 equities and global commodity futures | **B3 tickers:** `SLCE3.SA`, `AGRO3.SA`, `JBSS3.SA`, `BEEF3.SA`, `MRFG3.SA`, `SMTO3.SA`, `TTEN3.SA`<br>**CME futures:** `ZC=F` (corn), `ZS=F` (soy) |
| **HG Brasil Finance** | REST (JSON) | Free (key / daily limits) | B3 stocks, FX (USD/BRL), macro indicators | `GET https://api.hgbrasil.com/finance?key=YOUR_KEY` |

---

## 2. Physical market and commodities (spot)

| Source / API | Access | Auth / cost | Provides | How to consume |
| :--- | :--- | :--- | :--- | :--- |
| **CEPEA / Esalq (USP)** | **Web scraping** (HTML) | Free (no official public REST for this use) | National physical-market reference: cattle arroba, calf, soy/corn bags, coffee, sugar, cotton | HTML via CSS selectors.<br>Example: `https://www.cepea.esalq.usp.br/br/indicador/boi-gordo.aspx` |

**Cadence:** main CEPEA/ESALQ indicators are **daily on business days**, typically published after ~18:00 BRT (one official value per day — not intraday for the headline index). See [ADR-006](ADRs/006-source-fetch-cadence.md).

---

## 3. Government data, climate, and macro indicators

| Source / API | Access | Auth / cost | Provides | Examples |
| :--- | :--- | :--- | :--- | :--- |
| **IBGE SIDRA** | REST (JSON) | Free, public | Consolidated statistics: livestock (PPM), agricultural production (PAM), harvest surveys (LSPA) | `GET https://api.sidra.ibge.gov.br/values/t/1612/p/last%201/v/all/N1/all` |
| **Embrapa AgroAPI** | REST (JSON) | Free (registration / token) | Agricultural climate risk zoning (ZARC), soils, agrochemicals, climate | `GET https://api.cnptia.embrapa.br/agroapi/v1/...` |

---

## 4. Units glossary

| Asset | Unit | Notes |
| :--- | :--- | :--- |
| **Cattle arroba (`@`)** | **BRL / @** | 1 arroba = **15 kg** of carcass. Physical-market quote (e.g. SP average). |
| **Corn and soy (physical)** | **BRL / bag** | Standard bag **60 kg** (e.g. Paranaguá, Rio Verde, Campinas). |
| **B3 equities** | **BRL / share** | Unit price of listed agri-sector companies. |
| **CME / B3 futures** | Tickers / US cents | Contracts with specific expiries (e.g. BGI cattle, CCM corn on B3). |
