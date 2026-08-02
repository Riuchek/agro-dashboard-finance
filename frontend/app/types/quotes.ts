export type Quote = {
    key: string
    value: number
    unit: string
    source: string
    updated_at: string
}

export type QuotesResponse = {
    quotes: Quote[]
    errors?: string[]
}
