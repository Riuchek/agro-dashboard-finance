package domain

import "time"

type Quote struct {
	Key       string    `json:"key"`
	Value     float64   `json:"value"`
	Unit      Unit      `json:"unit"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updated_at"`
}
