# Lesson 04 — Hooks vs Vue composables

## Mapping

| Vue / Nuxt | React / Next |
| :--- | :--- |
| `components/` | `components/` |
| `composables/useFoo.ts` | `hooks/useFoo.ts` (custom hook) |

Same idea: extract reusable **logic** (state, effects, derived values) out of UI files.

## Example shape

```ts
"use client"

import { useEffect, useState } from "react"

export function useOnDemandRefresh(refresh: () => Promise<void>) {
  const [pending, setPending] = useState(false)

  async function run() {
    setPending(true)
    try {
      await refresh()
    } finally {
      setPending(false)
    }
  }

  return { pending, run }
}
```

Call it only from a Client Component (or another hook used by one).

## Where hooks do not replace Server Components

Fetching on the server for the first paint often stays in Server Components or `lib/` helpers called from them — not everything needs a hook. Use hooks when the logic is tied to **client lifecycle** (mount, intervals, user events, local state).
