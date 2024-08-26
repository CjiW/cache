package main

type Cache interface {
	Get(key int64) interface{}
	Set(key int64, val interface{})
	Flush()
	HitRate() float64
}
