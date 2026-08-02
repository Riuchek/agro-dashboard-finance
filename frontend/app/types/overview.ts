export type Stock = {
    key: string
    value: number
    unit: string
    source: string
    updated_at: string
}

export type Commodity = {
    key: string
    value: number
    unit: string
    source: string
    updated_at: string
}

export type OverviewResponse = {
    stocks: Stock[]
    commodities: Commodity[]
    errors: string[]
}