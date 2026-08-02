package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/handler"
	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/middleware"
)

type Dependencies struct {
	Quotes           *handler.Quotes
	Overview         *handler.Overview
	CORSAllowOrigins []string
}

type Server struct {
	addr   string
	server *http.Server
}

func New(addr string, deps Dependencies) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /api/v1/quotes/stocks", deps.Quotes.Stocks)
	mux.HandleFunc("GET /api/v1/quotes/commodities", deps.Quotes.Commodities)
	mux.HandleFunc("GET /api/v1/quotes/fx", deps.Quotes.FX)
	mux.HandleFunc("GET /api/v1/quotes/stocks-usd", deps.Quotes.StocksUSD)
	mux.HandleFunc("GET /api/v1/dashboard/overview", deps.Overview.Get)

	var h http.Handler = mux
	h = middleware.CORS(deps.CORSAllowOrigins)(h)
	h = middleware.Recover(h)
	h = middleware.Log(h)

	return &Server{
		addr: addr,
		server: &http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}

func (s *Server) Run() error {
	fmt.Printf("api listening on http://%s\n", s.addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
