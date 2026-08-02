import type { CommoditiesResponse } from '../types/commodities'

export function useCommodities() {
    const config = useRuntimeConfig()

    const { data: commoditiesResponse } = useFetch<CommoditiesResponse>(
        `${config.public.apiUrl}/api/v1/quotes/commodities?keys=boi-gordo,soja,milho`
    )

    return { commoditiesResponse }
}
