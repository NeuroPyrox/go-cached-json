package cachedjson

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Cache struct {
	mutex sync.RWMutex
	data  []byte
}

func (cache *Cache) Clear() {
	cache.mutex.Lock()
	cache.data = nil
	cache.mutex.Unlock()
}

func (cache *Cache) IsEmpty() bool {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	return cache.isEmpty()
}

func (cache *Cache) isEmpty() bool {
	return cache.data == nil
}

func (cache *Cache) Marshal() ([]byte, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if cache.isEmpty() {
		return nil, fmt.Errorf("JSON cache is empty")
	}
	return cache.data, nil
}

func (cache *Cache) Unmarshal(data []byte, obj interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	cache.data = data
	return nil
}

func (cache *Cache) Update(obj interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	return cache.update(obj)
}

func (cache *Cache) update(obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	cache.data = data
	return nil
}

func (cache *Cache) UpdateIfEmpty(obj interface{}) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.isEmpty() {
		return cache.update(obj)
	}
	return nil
}
