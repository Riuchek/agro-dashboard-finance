import type { FxResponse } from '../types/fx'

export function useFx() {
    const config = useRuntimeConfig()

    const { data: fxResponse } = useFetch<FxResponse>(
        `${config.public.apiUrl}/api/v1/quotes/fx?keys=usd-brl`
    )

    return { fxResponse }
}
