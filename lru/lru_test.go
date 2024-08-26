package main

import (
	"cache/storage"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	storage := storage.NewStorage(1024)
	c := NewLRUCache(100, &storage)
	for i := range 10000 {
		key := int64(rand.Uint64()%1000)
		c.Set(key, i)
		assert.Equal(t, c.Get(key), i)
		t.Logf("hit rate: %.6f\r", c.HitRate())
	}
}