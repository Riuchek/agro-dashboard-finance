<script lang="ts" setup>
const { quotesResponse } = useQuote()
const { overviewResponse } = useOverview()
const { commoditiesResponse } = useCommodities()
const { fxResponse } = useFx()
</script>

<template>
  <main class="min-h-screen bg-slate-950 px-4 py-8 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-5xl">
      <header class="mb-8">
        <p class="text-xs font-semibold uppercase tracking-[0.2em] text-emerald-500">
          Agro Dashboard
        </p>
        <h1 class="mt-2 text-2xl font-bold text-slate-100 sm:text-3xl">
          Painel Financeiro
        </h1>
        <p class="mt-1 text-sm text-slate-400">
          Acompanhe as cotações dos principais ativos do agronegócio.
        </p>
      </header>

      <QuotesCard :quote="quotesResponse" />

      <CommoditiesCard :commodities="commoditiesResponse" />

      <FxCard :fx="fxResponse" />

      <OverviewCard :overview="overviewResponse" />
    </div>
  </main>
</template>
