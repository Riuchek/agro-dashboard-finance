export type Commodity = {
    key: string
    value: number
    unit: string
    source: string
    updated_at: string
}

export type CommoditiesResponse = {
    commodities: Commodity[]
    errors?: string[]
}
