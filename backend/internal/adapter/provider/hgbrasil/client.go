package hgbrasil

import (
	"context"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/cache"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

const source = cache.SourceHGBrasil

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Source() string {
	return source
}

func (c *Client) FetchQuote(_ context.Context, _ string) (domain.Quote, error) {
	if c.token == "" {
		return domain.Quote{}, domain.ErrNotConfigured
	}
	return domain.Quote{}, domain.ErrNotImplemented
}
