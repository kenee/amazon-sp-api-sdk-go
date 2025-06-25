package auth

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// MemoryCache implements LWAAccessTokenCache using in-memory storage
type MemoryCache struct {
	cache *cache.Cache
	mu    sync.RWMutex
}

// NewMemoryCache creates a new memory cache with default settings
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: cache.New(5*time.Minute, 10*time.Minute), // Default TTL and cleanup interval
	}
}

// NewMemoryCacheWithTTL creates a new memory cache with custom TTL
func NewMemoryCacheWithTTL(defaultTTL, cleanupInterval time.Duration) *MemoryCache {
	return &MemoryCache{
		cache: cache.New(defaultTTL, cleanupInterval),
	}
}

// Get retrieves a value from cache
func (m *MemoryCache) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if value, found := m.cache.Get(key); found {
		if str, ok := value.(string); ok {
			return str, true
		}
	}
	return "", false
}

// Set stores a value in cache with TTL
func (m *MemoryCache) Set(key, value string, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Set(key, value, ttl)
}

// Delete removes a value from cache
func (m *MemoryCache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Delete(key)
}

// Flush clears all items from cache
func (m *MemoryCache) Flush() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Flush()
}

// ItemCount returns the number of items in cache
func (m *MemoryCache) ItemCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.ItemCount()
}
