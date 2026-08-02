package cache

import (
	"log"

	"github.com/Riuchek/agro-dashboard-finance/internal/port"
)

func NewStore(redisURL string) port.CacheStore {
	if redisURL == "" {
		log.Println("cache: using in-memory store (set REDIS_URL for Redis)")
		return NewMemoryCache()
	}

	store, err := NewRedisCache(redisURL)
	if err != nil {
		log.Printf("cache: redis unavailable (%v), falling back to in-memory store", err)
		return NewMemoryCache()
	}

	log.Println("cache: using Redis")
	return store
}

func Close(store port.CacheStore) {
	if closer, ok := store.(*RedisCache); ok {
		_ = closer.Close()
	}
}
