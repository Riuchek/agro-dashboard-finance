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
}

func NewService(cacheStore port.CacheStore, providers map[string]port.QuoteProvider) *Service {
	return &Service{
		cache:     cacheStore,
		providers: providers,
	}
}

func (s *Service) GetStocks(ctx context.Context, symbols []string) ([]domain.Quote, []string) {
	return s.fetchMany(ctx, cache.SourceBrapi, symbols)
}

func (s *Service) GetCommodities(ctx context.Context, keys []string) ([]domain.Quote, []string) {
	return s.fetchMany(ctx, cache.SourceCEPEA, keys)
}

func (s *Service) GetFX(ctx context.Context) (domain.Quote, error) {
	quotes, errs := s.fetchMany(ctx, cache.SourceHGBrasil, []string{"usd-brl"})
	if len(errs) > 0 {
		return domain.Quote{}, fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	if len(quotes) == 0 {
		return domain.Quote{}, domain.ErrQuoteNotFound
	}
	return quotes[0], nil
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
