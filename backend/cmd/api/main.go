package main

import (
	"log"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/cache"
	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/handler"
	httpserver "github.com/Riuchek/agro-dashboard-finance/internal/adapter/http"
	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/provider/brapi"
	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/provider/cepea"
	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/provider/hgbrasil"
	"github.com/Riuchek/agro-dashboard-finance/internal/app/overview"
	"github.com/Riuchek/agro-dashboard-finance/internal/app/quotes"
	"github.com/Riuchek/agro-dashboard-finance/internal/config"
	"github.com/Riuchek/agro-dashboard-finance/internal/port"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	cacheStore := cache.NewStore(cfg.RedisURL)
	defer cache.Close(cacheStore)

	providers := map[string]port.QuoteProvider{
		cache.SourceBrapi:    brapi.NewClient(cfg.BrapiToken),
		cache.SourceCEPEA:    cepea.NewScraper(),
		cache.SourceHGBrasil: hgbrasil.NewClient(cfg.HGBrasilToken),
	}

	hgClient := hgbrasil.NewClient(cfg.HGBrasilToken)
	quotesService := quotes.NewService(cacheStore, providers, hgClient)
	overviewService := overview.NewService(quotesService)

	server := httpserver.New(cfg.HTTPAddr, httpserver.Dependencies{
		Quotes:           handler.NewQuotes(quotesService),
		Overview:         handler.NewOverview(overviewService),
		CORSAllowOrigins: cfg.CORSAllowOrigins,
	})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
