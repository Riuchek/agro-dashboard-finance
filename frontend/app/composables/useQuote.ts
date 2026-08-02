import type { QuotesResponse } from '../types/quotes'

export function useQuote() {
    const config = useRuntimeConfig()

    const { data: quotesResponse } = useFetch<QuotesResponse>(
        `${config.public.apiUrl}/api/v1/quotes/stocks?symbols=AGRO3,SMTO3,SLCE3`
    )

    return { quotesResponse }
}
