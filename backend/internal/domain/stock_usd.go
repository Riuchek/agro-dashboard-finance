package domain

import "time"

type StockUSD struct {
	Key       string    `json:"key"`
	ValueBRL  float64   `json:"value_brl"`
	ValueUSD  float64   `json:"value_usd"`
	UnitBRL   Unit      `json:"unit_brl"`
	UnitUSD   Unit      `json:"unit_usd"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updated_at"`
}
