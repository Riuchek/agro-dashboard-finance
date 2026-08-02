package brapi

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

const source = cache.SourceBrapi

type Client struct {
	token  string
	client *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Source() string {
	return source
}

func (c *Client) FetchQuote(ctx context.Context, key string) (domain.Quote, error) {
	if c.token == "" {
		return domain.Quote{}, domain.ErrNotConfigured
	}

	symbol := strings.ToUpper(strings.TrimSpace(key))
	endpoint := fmt.Sprintf("https://brapi.dev/api/quote/%s", url.PathEscape(symbol))
	reqURL, err := url.Parse(endpoint)
	if err != nil {
		return domain.Quote{}, err
	}

	query := reqURL.Query()
	query.Set("token", c.token)
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return domain.Quote{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("%w: %v", domain.ErrUpstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Quote{}, fmt.Errorf("%w: brapi status %d", domain.ErrUpstream, resp.StatusCode)
	}

	var payload quoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return domain.Quote{}, err
	}
	if len(payload.Results) == 0 {
		return domain.Quote{}, domain.ErrQuoteNotFound
	}

	result := payload.Results[0]
	updatedAt := time.Now().UTC()
	if result.RegularMarketTime != "" {
		if parsed, err := time.Parse(time.RFC3339, result.RegularMarketTime); err == nil {
			updatedAt = parsed.UTC()
		}
	}

	return domain.Quote{
		Key:       symbol,
		Value:     result.RegularMarketPrice,
		Unit:      domain.UnitBRLPerShare,
		Source:    source,
		UpdatedAt: updatedAt,
	}, nil
}

type quoteResponse struct {
	Results []quoteResult `json:"results"`
}

type quoteResult struct {
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	RegularMarketTime  string  `json:"regularMarketTime"`
}
