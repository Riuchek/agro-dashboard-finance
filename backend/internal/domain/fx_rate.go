package domain

import "time"

type FXRate struct {
	Key              string    `json:"key"`
	Buy              float64   `json:"buy"`
	Sell             float64   `json:"sell"`
	Mid              float64   `json:"mid"`
	VariationPercent float64   `json:"variation_percent"`
	Unit             Unit      `json:"unit"`
	Source           string    `json:"source"`
	UpdatedAt        time.Time `json:"updated_at"`
}
