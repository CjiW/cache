package main

import (
	"cache/storage"
	"container/list"
)

type pair struct {
	key int64
	val interface{}
}

type LRUChche struct {
	s   *storage.Storage
	ls  *list.List
	mp  map[int64]*list.Element
	cap int
	hit, access int64
}

func NewLRUCache(sz int, s *storage.Storage) *LRUChche {
	return &LRUChche{
		s:   s,
		ls:  &list.List{},
		mp:  make(map[int64]*list.Element),
		cap: sz,
		hit: 0,
		access: 0,

	}
}

func (c *LRUChche) Get(key int64) (val interface{}) {
	e, ok := c.mp[key]
	if !ok {
		val = c.s.Get(key)
		for c.ls.Len() >= c.cap {
			p := c.ls.Remove(c.ls.Front()).(pair)
			delete(c.mp, p.key)
		}
		e = c.ls.PushBack(pair{key: key, val: val})
		c.mp[key] = e
	} else {
		val = e.Value.(pair).val
		c.ls.MoveToBack(e)
		c.hit++
	}
	c.access++
	return val
}

func (c *LRUChche) Set(key int64, val interface{}) {
	e, ok := c.mp[key]
	if !ok {
		for c.ls.Len() >= c.cap {
			p := c.ls.Remove(c.ls.Front()).(pair)
			delete(c.mp, p.key)
		}
		e = c.ls.PushBack(pair{key: key, val: val})
		c.mp[key] = e
		c.s.Set(key, val)
	} else {
		e.Value = pair{key: key, val: val}
		c.ls.MoveToBack(e)
		c.hit++
	}
	c.access++
}

func (c *LRUChche) Flush() {
}

func (c *LRUChche) HitRate() float64 {
	if c.access == 0 {
		return 0
	}
	return float64(c.hit) / float64(c.access)
}