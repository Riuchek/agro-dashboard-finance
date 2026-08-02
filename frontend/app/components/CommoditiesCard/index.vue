<script lang="ts" setup>
import type { CommoditiesResponse } from '../../types/commodities'

const props = defineProps<{
  commodities?: CommoditiesResponse | null
}>()


</script>

<template>
  <section class="overflow-hidden rounded-xl mt-4 border border-slate-700/60 bg-slate-900 shadow-xl shadow-black/20">
    <header class="flex items-center justify-between border-b border-slate-700/60 bg-slate-800/80 px-5 py-4">
      <div class="flex items-center gap-3">
        <div class="flex h-9 w-9 items-center justify-center rounded-lg bg-emerald-500/10 ring-1 ring-emerald-500/30">
          <svg
            class="h-5 w-5 text-emerald-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="1.5"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M2.25 18 9 11.25l4.306 4.306a11.95 11.95 0 0 1 5.814-5.518l2.74-1.22m0 0-5.94-2.281m5.94 2.28-2.28 5.941"
            />
          </svg>
        </div>
        <div>
          <h2 class="text-sm font-semibold tracking-wide text-slate-100">
            Commodities
          </h2>
          <p class="text-xs text-slate-400">
            Commodities market
          </p>
        </div>
      </div>
      <span
        v-if="props.commodities?.commodities?.length"
        class="rounded-full bg-slate-700/80 px-2.5 py-1 text-xs font-medium tabular-nums text-slate-300"
      >
        {{ props.commodities.commodities.length }} commodities
      </span>
    </header>

    <div>
      {{ props.commodities }}
    </div>
   
  </section>
</template>
