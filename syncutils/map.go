package syncutils

import (
	"sync"
	"sync/atomic"

	"github.com/spf13/cast"
	"golang.org/x/sync/singleflight"
)

var _ interface {
	Load(interface{}) (interface{}, bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Delete(interface{})
	Range(func(key, value interface{}) (shouldContinue bool))
} = new(OnceCreateCache)

type OnceCreateCache struct {
	m       sync.Map
	sfg     singleflight.Group
	newFunc atomic.Value
}

func NewOnceCreateCache(newFunc func() interface{}) *OnceCreateCache {
	cache := &OnceCreateCache{
		m:       sync.Map{},
		sfg:     singleflight.Group{},
		newFunc: atomic.Value{},
	}
	cache.newFunc.Store(newFunc)
	return cache
}

func (cache *OnceCreateCache) Load(key interface{}) (interface{}, bool) {
	value, ok := cache.m.Load(key)
	if ok {
		return value, ok
	}
	_, _, _ = cache.sfg.Do(cast.ToString(key), func() (interface{}, error) {
		newFunc, ok := cache.newFunc.Load().(func() interface{})
		if !ok {
			return nil, nil
		}
		cache.m.Store(key, newFunc())
		return nil, nil
	})
	return cache.m.Load(key)
}

func (cache *OnceCreateCache) Store(key, value interface{}) {
	cache.m.Store(key, value)
}

func (cache *OnceCreateCache) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return cache.m.LoadOrStore(key, value)
}

func (cache *OnceCreateCache) Delete(i interface{}) {
	cache.m.Delete(i)
}

func (cache *OnceCreateCache) Range(f func(key interface{}, value interface{}) (shouldContinue bool)) {
	cache.m.Range(f)
}

func (cache *OnceCreateCache) Clear() {
	var keys []interface{}
	cache.m.Range(func(key, _ interface{}) bool {
		keys = append(keys, key)
		return true
	})
	for _, key := range keys {
		cache.m.Delete(key)
	}
}

func (cache *OnceCreateCache) SetNewFunc(newFunc func() interface{}) {
	cache.newFunc.Store(newFunc)
	cache.Clear()
}
