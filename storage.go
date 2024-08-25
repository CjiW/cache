package main

import "time"

const (
	StorageAccessConst = time.Millisecond * 10
)

type Storage struct {
	data []interface{}
}

func (s *Storage) Get(key int64) interface{} {
	time.Sleep(StorageAccessConst)
	return s.data[key]
}

func (s *Storage) Set(key int64, val interface{}) {
	time.Sleep(StorageAccessConst)
	s.data[key] = val
}