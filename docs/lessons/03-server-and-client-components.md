# Lesson 03 — Server and Client Components

## The default

In the App Router, every component is a **Server Component** unless you opt into the client.

Server Components:

- Run on the **server** (at request time or during build)
- Can read secrets, hit databases/APIs without exposing keys to the browser
- Send HTML (and a lighter payload) to the client
- **Cannot** use browser-only APIs: `useState`, `useEffect`, `onClick`, `window`, most chart libs that need the DOM

## Opting into the client

Add this at the **top** of a file (before imports):

```tsx
"use client"
```

That file (and modules it imports for interactivity) become a **Client Component** boundary. It still may be pre-rendered on the server once, then hydrated in the browser so hooks and events work.

## How to decide

Ask: **does this need the browser?**

| Need | Choice |
| :--- | :--- |
| Fetch data for the first paint, render a static table/card shell | Server Component (default) |
| Filters, toggles, local UI state | Client Component |
| `useEffect`, polling, subscriptions | Client Component |
| Charts that bind to the DOM (e.g. Recharts) | Client Component |
| Layouts, metadata, fonts | Usually Server Component |
| Click handlers, controlled inputs | Client Component |

## Practical pattern for this dashboard

The Go API already exposes `GET /api/v1/dashboard/overview` (stocks live; commodities/FX when CEPEA/HG providers are done). The frontend will fetch that from Server Components or `lib/` on first paint — no browser token for upstream APIs.

Keep the **page/layout as Server Components**. Push interactivity into small client leaves:

```
app/page.tsx                    ← Server (compose data + layout)
components/QuoteCard.tsx        ← Server (presentational)
components/PriceChart.tsx       ← "use client" (Recharts)
hooks/useOnDemandRefresh.ts     ← used only from client components
```

Flow:

1. Server page fetches or receives data (or passes children that fetch).
2. Server renders structure and static bits.
3. Client islands handle charts, filters, “refresh now”.

## Rules that avoid pain

1. **Default to server.** Add `"use client"` only when you hit a browser API or hook.
2. **Push `"use client"` down.** Prefer a small chart/filter component over marking the whole page as client.
3. **Server can import Client; Client cannot import Server-only modules** that use server-only APIs (secrets, `fs`, etc.). Pass data as **props** from server parents into client children.
4. **Hooks live on the client.** Custom hooks that use `useState` / `useEffect` must be called from Client Components.

## Mental model

```
Server Component  →  “render this HTML/data on the server”
Client Component  →  “this part is interactive in the browser”
```

You are not choosing “React vs Next”. You are choosing **where the code runs**.

## Link to project decisions

- App Router: [ADR-010](../ADRs/010-next-app-router.md)
- Charts (likely client): [ADR-009](../ADRs/009-recharts.md)
- On-demand refresh (client trigger → API): [ADR-012](../ADRs/012-on-demand-refresh.md)
