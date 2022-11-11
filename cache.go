package cache

import (
	"time"
)

type CachedMap[T any] struct {
	m      map[string]*Data[T]
	ticker *time.Ticker
	stop   chan bool
}

type Data[T any] struct {
	Value     T
	ExpiresAt int64
}

func (tm *CachedMap[T]) Set(key string, value T, expirationDate time.Time) {
	tm.m[key] = &Data[T]{
		Value:     value,
		ExpiresAt: expirationDate.Unix(),
	}
}

func (tm *CachedMap[T]) Get(key string) (*T, bool) {
	v := tm.m[key]
	if v == nil {
		return nil, false
	}
	return &v.Value, true
}

func (tm *CachedMap[T]) clean() {
	now := time.Now().Unix()
	for key, el := range tm.m {
		if now >= el.ExpiresAt {
			delete(tm.m, key)
		}
	}
}

func (tm *CachedMap[T]) StartProcessing() {
	go func() {
		for {
			select {
			case <-tm.ticker.C:
				tm.clean()
			case <-tm.stop:
				break
			}
		}
	}()
}

func NewCacheStorage[T any](ttl time.Duration) *CachedMap[T] {
	t := &CachedMap[T]{
		m:      map[string]*Data[T]{},
		ticker: time.NewTicker(ttl),
		stop:   make(chan bool),
	}
	t.StartProcessing()
	return t
}
