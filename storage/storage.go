package storage

import "time"

const (
	StorageAccessConst = time.Millisecond
)

type Storage struct {
	data []interface{}
}

func NewStorage(size int64) Storage {
	return Storage {
		data: make([]interface{}, size),
	}
}

func (s *Storage) Get(key int64) interface{} {
	time.Sleep(StorageAccessConst)
	return s.data[key]
}

func (s *Storage) Set(key int64, val interface{}) {
	time.Sleep(StorageAccessConst)
	s.data[key] = val
}