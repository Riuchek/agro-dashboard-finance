package cepea

import (
	"context"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/cache"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

const source = cache.SourceCEPEA

type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Source() string {
	return source
}

func (s *Scraper) FetchQuote(_ context.Context, _ string) (domain.Quote, error) {
	return domain.Quote{}, domain.ErrNotImplemented
}
