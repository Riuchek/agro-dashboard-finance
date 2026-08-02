package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/cache"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
	"github.com/Riuchek/agro-dashboard-finance/internal/port"
)

type Service struct {
	cache      port.CacheStore
	providers  map[string]port.QuoteProvider
	fxProvider port.FXRateProvider
}

type StocksUSDResponse struct {
	FX     domain.FXRate     `json:"fx"`
	Quotes []domain.StockUSD `json:"quotes"`
	Errors []string          `json:"errors,omitempty"`
}

func NewService(cacheStore port.CacheStore, providers map[string]port.QuoteProvider, fxProvider port.FXRateProvider) *Service {
	return &Service{
		cache:      cacheStore,
		providers:  providers,
		fxProvider: fxProvider,
	}
}

func (s *Service) GetStocks(ctx context.Context, symbols []string) ([]domain.Quote, []string) {
	return s.fetchMany(ctx, cache.SourceBrapi, symbols)
}

func (s *Service) GetCommodities(ctx context.Context, keys []string) ([]domain.Quote, []string) {
	return s.fetchMany(ctx, cache.SourceCEPEA, keys)
}

func (s *Service) GetFX(ctx context.Context) (domain.Quote, error) {
	rate, err := s.GetFXRate(ctx)
	if err != nil {
		return domain.Quote{}, err
	}

	return domain.Quote{
		Key:       rate.Key,
		Value:     rate.Mid,
		Unit:      rate.Unit,
		Source:    rate.Source,
		UpdatedAt: rate.UpdatedAt,
	}, nil
}

func (s *Service) GetFXRate(ctx context.Context) (domain.FXRate, error) {
	return s.fetchFXRate(ctx)
}

func (s *Service) GetStocksUSD(ctx context.Context, symbols []string) (StocksUSDResponse, error) {
	fxRate, err := s.fetchFXRate(ctx)
	if err != nil {
		return StocksUSDResponse{}, err
	}

	stocks, errs := s.GetStocks(ctx, symbols)
	quotes := make([]domain.StockUSD, 0, len(stocks))
	for _, stock := range stocks {
		quotes = append(quotes, domain.StockUSD{
			Key:       stock.Key,
			ValueBRL:  stock.Value,
			ValueUSD:  stock.Value / fxRate.Mid,
			UnitBRL:   stock.Unit,
			UnitUSD:   domain.UnitUSDPerShare,
			Source:    stock.Source,
			UpdatedAt: stock.UpdatedAt,
		})
	}

	return StocksUSDResponse{
		FX:     fxRate,
		Quotes: quotes,
		Errors: errs,
	}, nil
}

func (s *Service) fetchMany(ctx context.Context, source string, keys []string) ([]domain.Quote, []string) {
	provider, ok := s.providers[source]
	if !ok {
		return nil, []string{fmt.Sprintf("%s: provider not registered", source)}
	}

	quotes := make([]domain.Quote, 0, len(keys))
	var errs []string

	for _, key := range keys {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		quote, err := s.fetchOne(ctx, provider, source, key)
		if err != nil {
			log.Printf("quotes: %s/%s: %v", source, key, err)
			errs = append(errs, fmt.Sprintf("%s/%s: %v", source, key, err))
			continue
		}
		quotes = append(quotes, quote)
	}

	return quotes, errs
}

func (s *Service) fetchOne(ctx context.Context, provider port.QuoteProvider, source, key string) (domain.Quote, error) {
	cacheKey := cache.QuoteKey(source, key)

	if cached, err := s.cache.Get(ctx, cacheKey); err != nil {
		return domain.Quote{}, err
	} else if cached != "" {
		var quote domain.Quote
		if err := json.Unmarshal([]byte(cached), &quote); err != nil {
			return domain.Quote{}, err
		}
		return quote, nil
	}

	quote, err := provider.FetchQuote(ctx, key)
	if err != nil {
		return domain.Quote{}, err
	}

	payload, err := json.Marshal(quote)
	if err != nil {
		return domain.Quote{}, err
	}

	if err := s.cache.Set(ctx, cacheKey, string(payload), cache.TTLForSource(source)); err != nil {
		log.Printf("quotes: cache set %s: %v", cacheKey, err)
	}

	return quote, nil
}

func (s *Service) fetchFXRate(ctx context.Context) (domain.FXRate, error) {
	cacheKey := cache.QuoteKey(cache.SourceHGBrasil, "usd-brl-detail")

	if cached, err := s.cache.Get(ctx, cacheKey); err != nil {
		return domain.FXRate{}, err
	} else if cached != "" {
		var rate domain.FXRate
		if err := json.Unmarshal([]byte(cached), &rate); err != nil {
			return domain.FXRate{}, err
		}
		return rate, nil
	}

	rate, err := s.fxProvider.FetchFXRate(ctx)
	if err != nil {
		return domain.FXRate{}, err
	}

	payload, err := json.Marshal(rate)
	if err != nil {
		return domain.FXRate{}, err
	}

	if err := s.cache.Set(ctx, cacheKey, string(payload), cache.TTLForSource(cache.SourceHGBrasil)); err != nil {
		log.Printf("quotes: cache set %s: %v", cacheKey, err)
	}

	return rate, nil
}
