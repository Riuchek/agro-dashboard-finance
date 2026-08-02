package handler

import (
	"net/http"
	"strings"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/response"
	"github.com/Riuchek/agro-dashboard-finance/internal/app/quotes"
	"github.com/Riuchek/agro-dashboard-finance/internal/domain"
)

type Quotes struct {
	service *quotes.Service
}

func NewQuotes(service *quotes.Service) *Quotes {
	return &Quotes{service: service}
}

type quotesResponse struct {
	Quotes []domain.Quote `json:"quotes"`
	Errors []string       `json:"errors,omitempty"`
}

func (h *Quotes) Stocks(w http.ResponseWriter, r *http.Request) {
	symbols := splitQueryParam(r, "symbols")
	if len(symbols) == 0 {
		response.Error(w, http.StatusBadRequest, "symbols query param is required")
		return
	}

	quotes, errs := h.service.GetStocks(r.Context(), symbols)
	response.JSON(w, http.StatusOK, quotesResponse{Quotes: quotes, Errors: errs})
}

func (h *Quotes) Commodities(w http.ResponseWriter, r *http.Request) {
	keys := splitQueryParam(r, "keys")
	if len(keys) == 0 {
		response.Error(w, http.StatusBadRequest, "keys query param is required")
		return
	}

	quotes, errs := h.service.GetCommodities(r.Context(), keys)
	response.JSON(w, http.StatusOK, quotesResponse{Quotes: quotes, Errors: errs})
}

func (h *Quotes) FX(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetFX(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, quote)
}

func splitQueryParam(r *http.Request, key string) []string {
	raw := strings.TrimSpace(r.URL.Query().Get(key))
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}
