package hgbrasil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/cache"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

const source = cache.SourceHGBrasil

type Client struct {
	key    string
	client *http.Client
}

func NewClient(key string) *Client {
	return &Client{
		key: key,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Source() string {
	return source
}

func (c *Client) FetchQuote(ctx context.Context, key string) (domain.Quote, error) {
	if strings.ToLower(strings.TrimSpace(key)) != "usd-brl" {
		return domain.Quote{}, domain.ErrQuoteNotFound
	}

	rate, err := c.FetchFXRate(ctx)
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

func (c *Client) FetchFXRate(ctx context.Context) (domain.FXRate, error) {
	if c.key == "" {
		return domain.FXRate{}, domain.ErrNotConfigured
	}

	payload, err := c.fetchFinance(ctx)
	if err != nil {
		return domain.FXRate{}, err
	}

	usd, ok, err := payload.Results.Currencies.rateFor("USD")
	if err != nil {
		return domain.FXRate{}, err
	}
	if !ok || usd.Buy <= 0 {
		return domain.FXRate{}, domain.ErrQuoteNotFound
	}

	mid := usd.Buy
	if usd.Sell > 0 {
		mid = (usd.Buy + usd.Sell) / 2
	}

	return domain.FXRate{
		Key:              "usd-brl",
		Buy:              usd.Buy,
		Sell:             usd.Sell,
		Mid:              mid,
		VariationPercent: usd.Variation,
		Unit:             domain.UnitBRLPerUSD,
		Source:           source,
		UpdatedAt:        time.Now().UTC(),
	}, nil
}

func (c *Client) fetchFinance(ctx context.Context) (financeResponse, error) {
	endpoint := "https://api.hgbrasil.com/finance"

	reqURL, err := url.Parse(endpoint)
	if err != nil {
		return financeResponse{}, err
	}

	query := reqURL.Query()
	query.Set("key", c.key)
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return financeResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return financeResponse{}, fmt.Errorf("%w: %v", domain.ErrUpstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return financeResponse{}, fmt.Errorf("%w: hgbrasil status %d", domain.ErrUpstream, resp.StatusCode)
	}

	var payload financeResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return financeResponse{}, err
	}

	return payload, nil
}

type financeResponse struct {
	Results financeResults `json:"results"`
}

type financeResults struct {
	Currencies currencyMap `json:"currencies"`
}

type currencyMap map[string]json.RawMessage

func (m currencyMap) rateFor(code string) (currencyRate, bool, error) {
	raw, ok := m[code]
	if !ok {
		return currencyRate{}, false, nil
	}

	var rate currencyRate
	if err := json.Unmarshal(raw, &rate); err != nil {
		return currencyRate{}, false, err
	}

	return rate, true, nil
}

type currencyRate struct {
	Buy       float64 `json:"buy"`
	Sell      float64 `json:"sell"`
	Variation float64 `json:"variation"`
}
