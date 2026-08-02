<script lang="ts" setup>
import type { StocksUsdResponse } from '../../types/stocks-usd'
import { formatValue, formatVariation, formatUpdatedAt } from '../../utils/utils'

const props = defineProps<{
  stocksUsd?: StocksUsdResponse | null
}>()

const variationClass = computed(() => {
  const variation = props.stocksUsd?.fx?.variation_percent ?? 0
  if (variation > 0) {
    return 'text-emerald-400'
  }
  if (variation < 0) {
    return 'text-rose-400'
  }
  return 'text-slate-300'
})
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
              d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.172-.879-1.172-2.303 0-3.182.553-.44 1.278-.659 2.003-.659.725 0 1.45.22 2.003.659M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
            />
          </svg>
        </div>
        <div>
          <h2 class="text-sm font-semibold tracking-wide text-slate-100">
            Cotações em Dólar
          </h2>
          <p class="text-xs text-slate-400">
            Ações do agro convertidas com câmbio HG Brasil
          </p>
        </div>
      </div>
      <span
        v-if="props.stocksUsd?.quotes?.length"
        class="rounded-full bg-slate-700/80 px-2.5 py-1 text-xs font-medium tabular-nums text-slate-300"
      >
        {{ props.stocksUsd.quotes.length }} ações
      </span>
    </header>

    <div
      v-if="props.stocksUsd?.fx"
      class="grid gap-4 border-b border-slate-700/60 bg-slate-800/30 px-5 py-4 sm:grid-cols-2 lg:grid-cols-4"
    >
      <div>
        <p class="text-xs font-medium uppercase tracking-wider text-slate-500">
          USD/BRL (compra)
        </p>
        <p class="mt-1 font-mono text-lg font-semibold tabular-nums text-slate-100">
          {{ formatValue(props.stocksUsd.fx.buy, 'BRL/USD') }}
        </p>
      </div>
      <div>
        <p class="text-xs font-medium uppercase tracking-wider text-slate-500">
          USD/BRL (venda)
        </p>
        <p class="mt-1 font-mono text-lg font-semibold tabular-nums text-slate-100">
          {{ formatValue(props.stocksUsd.fx.sell, 'BRL/USD') }}
        </p>
      </div>
      <div>
        <p class="text-xs font-medium uppercase tracking-wider text-slate-500">
          Variação do dia
        </p>
        <p
          class="mt-1 font-mono text-lg font-semibold tabular-nums"
          :class="variationClass"
        >
          {{ formatVariation(props.stocksUsd.fx.variation_percent) }}
        </p>
      </div>
      <div>
        <p class="text-xs font-medium uppercase tracking-wider text-slate-500">
          Taxa usada na conversão
        </p>
        <p class="mt-1 font-mono text-lg font-semibold tabular-nums text-emerald-400">
          {{ formatValue(props.stocksUsd.fx.mid, 'BRL/USD') }}
        </p>
        <p class="mt-0.5 text-xs text-slate-500">
          Fonte: {{ props.stocksUsd.fx.source }}
        </p>
      </div>
    </div>

    <div
      v-if="props.stocksUsd?.errors?.length"
      class="border-b border-amber-500/20 bg-amber-500/10 px-5 py-3"
    >
      <p
        v-for="error in props.stocksUsd.errors"
        :key="error"
        class="text-xs text-amber-300"
      >
        {{ error }}
      </p>
    </div>

    <div
      v-if="!props.stocksUsd?.quotes?.length"
      class="px-5 py-12 text-center"
    >
      <p class="text-sm text-slate-400">
        Nenhuma cotação em dólar disponível no momento.
      </p>
    </div>

    <div
      v-else
      class="overflow-x-auto"
    >
      <table class="w-full min-w-[720px] text-left text-sm">
        <thead>
          <tr class="border-b border-slate-700/60 bg-slate-800/40">
            <th class="px-5 py-3 text-xs font-semibold uppercase tracking-wider text-slate-400">
              Ativo
            </th>
            <th class="px-5 py-3 text-right text-xs font-semibold uppercase tracking-wider text-slate-400">
              Cotação (USD)
            </th>
            <th class="px-5 py-3 text-right text-xs font-semibold uppercase tracking-wider text-slate-400">
              Cotação (BRL)
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
            v-for="item in props.stocksUsd.quotes"
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
                {{ formatValue(item.value_usd, item.unit_usd) }}
              </span>
            </td>
            <td class="px-5 py-4 text-right">
              <span class="font-mono text-sm tabular-nums text-slate-400">
                {{ formatValue(item.value_brl, item.unit_brl) }}
              </span>
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
