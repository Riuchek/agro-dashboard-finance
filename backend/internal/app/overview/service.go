package overview

import (
	"context"

	"github.com/Riuchek/agro-dashboard-finance/internal/app/quotes"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

var (
	defaultStockSymbols = []string{"SLCE3", "AGRO3", "SMTO3"}
	defaultCommodityKeys = []string{"boi-gordo", "soja", "milho"}
)

type Response struct {
	Stocks      []domain.Quote `json:"stocks"`
	Commodities []domain.Quote `json:"commodities"`
	FX          *domain.Quote  `json:"fx,omitempty"`
	Errors      []string       `json:"errors,omitempty"`
}

type Service struct {
	quotes *quotes.Service
}

func NewService(quotesService *quotes.Service) *Service {
	return &Service{quotes: quotesService}
}

func (s *Service) Get(ctx context.Context) Response {
	stocks, stockErrs := s.quotes.GetStocks(ctx, defaultStockSymbols)
	commodities, commodityErrs := s.quotes.GetCommodities(ctx, defaultCommodityKeys)

	var fxQuote *domain.Quote
	var fxErrs []string
	if quote, err := s.quotes.GetFX(ctx); err != nil {
		fxErrs = []string{err.Error()}
	} else {
		fxQuote = &quote
	}

	errs := append([]string{}, stockErrs...)
	errs = append(errs, commodityErrs...)
	errs = append(errs, fxErrs...)

	return Response{
		Stocks:      stocks,
		Commodities: commodities,
		FX:          fxQuote,
		Errors:      errs,
	}
}
