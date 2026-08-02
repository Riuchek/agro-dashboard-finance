package cache

import (
	"fmt"
	"time"
)

const (
	SourceBrapi    = "brapi"
	SourceCEPEA    = "cepea"
	SourceHGBrasil = "hgbrasil"
)

func QuoteKey(source, key string) string {
	return fmt.Sprintf("%s:%s", source, key)
}

const (
	TTLBrapi    = 3 * time.Minute
	TTLCEPEA    = 12 * time.Hour
	TTLHGBrasil = 12 * time.Minute
)

func TTLForSource(source string) time.Duration {
	switch source {
	case SourceBrapi:
		return TTLBrapi
	case SourceCEPEA:
		return TTLCEPEA
	case SourceHGBrasil:
		return TTLHGBrasil
	default:
		return 5 * time.Minute
	}
}
