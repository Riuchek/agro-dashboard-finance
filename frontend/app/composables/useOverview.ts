import type { OverviewResponse } from '../types/overview'

export function useOverview() {
    const config = useRuntimeConfig()

    const { data: overviewResponse } = useFetch<OverviewResponse>(
        `${config.public.apiUrl}/api/v1/dashboard/overview`
    )

    return { overviewResponse }
}
