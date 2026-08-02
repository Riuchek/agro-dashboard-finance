# Lesson 06 — Components, Tailwind, and hooks (vs Nuxt SFC)

## Vue / Nuxt single-file component

```vue
<script setup>
const { pending, run } = useRefresh()
</script>

<template>
  <button @click="run">Refresh</button>
</template>

<style scoped>
button { /* ... */ }
</style>
```

Three blocks in one file: script, template, CSS.

## React / Next component

One function. Markup is JSX **inside** the return. There is no separate `<template>` or `<style scoped>` block.

```tsx
"use client"

import { useOnDemandRefresh } from "@/hooks/useOnDemandRefresh"

type Props = {
  onRefresh: () => Promise<void>
}

export function RefreshButton({ onRefresh }: Props) {
  const { pending, run } = useOnDemandRefresh(onRefresh)

  return (
    <button
      type="button"
      disabled={pending}
      onClick={() => void run()}
      className="rounded bg-zinc-900 px-4 py-2 text-sm text-white disabled:opacity-50"
    >
      {pending ? "Refreshing…" : "Refresh"}
    </button>
  )
}
```

Mental map:

| Nuxt | Next / React |
| :--- | :--- |
| `<script setup>` | Top of the function + imports |
| `<template>` | `return ( ... JSX ... )` |
| `<style scoped>` | Usually **Tailwind `className`**, or a CSS module / global CSS |
| `composables/useX` | `hooks/useX` called inside the component |

## Where does Tailwind go?

On the element, via `className` (not `class` — React uses `className`).

You do **not** need a `<style>` block per component for the common case. Shared tokens / resets live in `app/globals.css`.

Patterns:

1. **Utilities on the page or component** — fine for small UI
2. **Extract a component** when the same markup + classes repeat
3. **Avoid giant pages** — page composes components; components own their `className`s

Page (server) composing children:

```tsx
import { QuoteCard } from "@/components/QuoteCard"

export default function HomePage() {
  return (
    <main className="mx-auto max-w-5xl p-6">
      <h1 className="text-2xl font-semibold">Agro dashboard</h1>
      <QuoteCard title="Soy" price={120} />
    </main>
  )
}
```

Presentational child (can stay a Server Component if it has no hooks/events):

```tsx
type Props = { title: string; price: number }

export function QuoteCard({ title, price }: Props) {
  return (
    <article className="rounded border border-zinc-200 p-4">
      <h2 className="text-sm text-zinc-500">{title}</h2>
      <p className="text-xl font-medium">{price}</p>
    </article>
  )
}
```

## How to wire hooks

1. Put logic in `src/hooks/useSomething.ts`
2. Import it in a **Client Component**
3. Call the hook at the **top level** of that function (same rules as Vue: no hooks inside `if` / loops)

```tsx
"use client"

import { useState } from "react"
import { useQuotes } from "@/hooks/useQuotes"

export function QuotesPanel() {
  const { quotes, loading, error, reload } = useQuotes()
  const [filter, setFilter] = useState("")

  if (loading) return <p className="text-sm text-zinc-500">Loading…</p>
  if (error) return <p className="text-sm text-red-600">{error}</p>

  return (
    <div className="space-y-4">
      <input
        className="w-full rounded border px-3 py-2"
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      />
      <button type="button" className="underline" onClick={() => void reload()}>
        Reload
      </button>
      <ul>
        {quotes
          .filter((q) => q.name.includes(filter))
          .map((q) => (
            <li key={q.id}>{q.name}</li>
          ))}
      </ul>
    </div>
  )
}
```

Hook file:

```ts
"use client"

import { useCallback, useEffect, useState } from "react"

export function useQuotes() {
  const [quotes, setQuotes] = useState<{ id: string; name: string }[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const reload = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch("http://127.0.0.1:8080/api/v1/dashboard/overview")
      if (!res.ok) throw new Error("Failed to load")
      setQuotes(await res.json())
    } catch (e) {
      setError(e instanceof Error ? e.message : "Unknown error")
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    void reload()
  }, [reload])

  return { quotes, loading, error, reload }
}
```

## Server page + client island

Keep data shell on the server; mount interactive pieces as children:

```tsx
import { QuotesPanel } from "@/components/QuotesPanel"

export default function DashboardPage() {
  return (
    <main className="p-6">
      <h1 className="mb-4 text-2xl font-semibold">Dashboard</h1>
      <QuotesPanel />
    </main>
  )
}
```

`QuotesPanel` has `"use client"` + hooks. The page can stay a Server Component.

## Short answers

- **Styling with Tailwind?** Yes — mostly `className` on JSX in pages and components (same idea as Nuxt + Tailwind utility classes, without SFC `<style>`).
- **script / template / css blocks?** No — one TSX function: logic above, JSX in `return`, styles as utilities (or CSS modules if you really need them).
- **Hooks?** `hooks/useX.ts` → import → call at top of a Client Component → use returned values in JSX / handlers.
