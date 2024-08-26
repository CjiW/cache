package main

import "cache/storage"

type DirectCache struct {
	s *storage.Storage
}

func NewDirectCache(s *storage.Storage) *DirectCache {
	return &DirectCache{
		s: s,
	}
}

func (c *DirectCache) Get(key int64) interface{} {
	return c.s.Get(key)
}

func (c *DirectCache) Set(key int64, val interface{}) {
	c.s.Set(key, val)
}

func (c *DirectCache) Flush() {
}

func (c *DirectCache) HitRate() float64 {
	return 0
}