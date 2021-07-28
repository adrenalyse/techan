package techan

import (
	"github.com/dgraph-io/ristretto"
	"log"
)

/* from https://github.com/dgraph-io/ristretto */

func NewRistrettoCache() *ristretto.Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e6,     // number of keys to track frequency of (1M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Println(err)
	}

	return cache
}
