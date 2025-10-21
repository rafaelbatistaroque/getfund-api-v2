package cache

import (
	"time"
)

// Service defines the interface for a cache service.
// It provides methods for setting, getting, and deleting cache entries,
// as well as closing the connection to the cache backend.
type Service interface {
	// Set stores a value in the cache with a specific key and expiration time.
	// If the value is not a string, it will be serialized to JSON.
	Set(key string, value any, time time.Duration) error
	// Get retrieves a value from the cache by its key.
	// The value is returned as a string.
	Get(key string) (string, error)
	// Delete removes a key and its associated value from the cache.
	Delete(key string) error
	// Close terminates the connection to the cache service.
	Close() error
}
