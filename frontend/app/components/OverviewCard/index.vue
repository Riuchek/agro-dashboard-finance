<script lang="ts" setup>
import type { OverviewResponse } from '../../types/overview'
import { formatValue, formatUpdatedAt } from '../../utils/utils'

const props = defineProps<{
  overview?: OverviewResponse | null
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
            Visão Geral
          </h2>
          <p class="text-xs text-slate-400">
            Acompanhe as cotações dos principais ativos do agronegócio.
          </p>
        </div>
      </div>
      <span
        v-if="props.overview?.stocks?.length"
        class="rounded-full bg-slate-700/80 px-2.5 py-1 text-xs font-medium tabular-nums text-slate-300"
      >
        {{ props.overview.stocks.length }} ações
      </span>
    </header>

    <div
      v-if="props.overview?.errors?.length"
      class="border-b border-amber-500/20 bg-amber-500/10 px-5 py-3"
    >
      <p
        v-for="error in props.overview.errors"
        :key="error"
        class="text-xs text-amber-300"
      >
        {{ error }}
      </p>
    </div>

    <div
      v-if="!props.overview?.stocks?.length"
      class="px-5 py-12 text-center"
    >
      <p class="text-sm text-slate-400">
        Nenhuma cotação disponível no momento.
      </p>
    </div>

    <div
      v-else
      class="overflow-x-auto"
    >
      <table class="w-full min-w-[640px] text-left text-sm">
        <thead>
          <tr class="border-b border-slate-700/60 bg-slate-800/40">
            <th class="px-5 py-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
              Ativo
            </th>
            <th class="px-5 py-3 text-right text-xs font-semibold uppercase tracking-wider text-slate-400">
              Cotação
            </th>
            <th class="px-5 py-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
              Unidade
            </th>
            <th class="px-5 py-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
              Fonte
            </th>
            <th class="px-5 py-3 text-right text-xs font-semibold uppercase tracking-wider text-slate-400">
              Atualizado
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-700/40">
          <tr
            v-for="item in props.overview.stocks"
            :key="item.key"
            class="transition-colors hover:bg-slate-800/50"
          >
            <td class="px-5 py-4">
              <span class="inline-flex items-center rounded-md bg-slate-800 px-2.5 py-1 font-mono text-xs font-semibold tracking-wide text-emerald-400 ring-1 ring-slate-600">
                {{ item.key }}
              </span>
            </td>
            <td class="px-5 py-4 text-right">
              <span class="font-mono text-base font-semibold tabular-nums text-slate-100">
                {{ formatValue(item.value, item.unit) }}
              </span>
            </td>
            <td class="px-5 py-4 text-slate-400">
              {{ item.unit }}
            </td>
            <td class="px-5 py-4">
              <span class="rounded-full bg-slate-800/80 px-2.5 py-0.5 text-xs font-medium text-slate-300">
                {{ item.source }}
              </span>
            </td>
            <td class="px-5 py-4 text-right font-mono text-xs tabular-nums text-slate-500">
              {{ formatUpdatedAt(item.updated_at) }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
