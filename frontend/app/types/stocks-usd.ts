export type FxRate = {
    key: string
    buy: number
    sell: number
    mid: number
    variation_percent: number
    unit: string
    source: string
    updated_at: string
}

export type StockUSD = {
    key: string
    value_brl: number
    value_usd: number
    unit_brl: string
    unit_usd: string
    source: string
    updated_at: string
}

export type StocksUsdResponse = {
    fx: FxRate
    quotes: StockUSD[]
    errors?: string[]
}
