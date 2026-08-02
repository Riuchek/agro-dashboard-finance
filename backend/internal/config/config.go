package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultHTTPAddr       = "127.0.0.1:8080"
	defaultCORSAllowOrigins = "http://localhost:3000"
)

type Config struct {
	HTTPAddr         string
	RedisURL         string
	BrapiToken       string
	HGBrasilToken    string
	CORSAllowOrigins []string
}

func Load() (*Config, error) {
	_ = loadEnvFile(".env")

	cfg := &Config{
		HTTPAddr:         envOrDefault("HTTP_ADDR", defaultHTTPAddr),
		RedisURL:         strings.TrimSpace(os.Getenv("REDIS_URL")),
		BrapiToken:       strings.TrimSpace(os.Getenv("BRAPI_TOKEN")),
		HGBrasilToken:    strings.TrimSpace(os.Getenv("HG_BRASIL_TOKEN")),
		CORSAllowOrigins: splitCSV(envOrDefault("CORS_ALLOW_ORIGINS", defaultCORSAllowOrigins)),
	}

	if cfg.HTTPAddr == "" {
		return nil, fmt.Errorf("HTTP_ADDR must not be empty")
	}

	return cfg, nil
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}
