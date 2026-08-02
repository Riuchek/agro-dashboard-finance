package port

import (
	"context"

	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

type FXRateProvider interface {
	Source() string
	FetchFXRate(ctx context.Context) (domain.FXRate, error)
}
