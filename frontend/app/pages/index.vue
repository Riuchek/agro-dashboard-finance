<script lang="ts" setup>

const { quotesResponse } = useQuote()

</script>


<template>
  <div>
    <div
      v-for="quote in quotesResponse?.quotes ?? []"
      :key="quote.key"
      class="flex flex-row"
    >
      <h1 class="text-2xl font-bold w-1/4">{{ quote.key }}</h1>
      <h1 class="text-2xl font-bold w-1/4">{{ quote.source }}</h1>
      <h1 class="text-2xl font-bold w-1/4">{{ quote.unit }}</h1>
      <h1 class="text-2xl font-bold w-1/4">{{ quote.updated_at }}</h1>
      <h1 class="text-2xl font-bold w-1/4">{{ quote.value }}</h1>
    </div>
  </div>
</template>
