import type { StocksUsdResponse } from '../types/stocks-usd'

const defaultSymbols = 'SLCE3,AGRO3,SMTO3'

export function useStocksUsd() {
    const config = useRuntimeConfig()

    const { data: stocksUsdResponse } = useFetch<StocksUsdResponse>(
        `${config.public.apiUrl}/api/v1/quotes/stocks-usd?symbols=${defaultSymbols}`
    )

    return { stocksUsdResponse }
}
