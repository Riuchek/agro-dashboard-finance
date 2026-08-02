export type Fx = {
    key: string
    value: number
    unit: string
    source: string
    updated_at: string
}

export type FxResponse = {
    fx: Fx[]
    errors?: string[]
}
