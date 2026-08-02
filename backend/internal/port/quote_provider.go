package port

import (
	"context"

	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

type QuoteProvider interface {
	Source() string
	FetchQuote(ctx context.Context, key string) (domain.Quote, error)
}
