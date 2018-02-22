package cachedjson

import (
	"encoding/json"
	"fmt"
)

type Cache struct {
	data []byte
}

func (cache *Cache) Clear() {
	cache.data = nil
}

func (cache *Cache) IsEmpty() bool {
	return cache.data == nil
}

func (cache *Cache) Marshal() ([]byte, error) {
	if cache.IsEmpty() {
		return nil, fmt.Errorf("JSON cache is empty")
	}
	return cache.data, nil
}

func (cache *Cache) Unmarshal(data []byte, obj interface{}) error {
	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	cache.data = data
	return nil
}

func (cache *Cache) Update(obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	cache.data = data
	return nil
}

func (cache *Cache) UpdateIfEmpty(obj interface{}) error {
	if cache.IsEmpty() {
		return cache.Update(obj)
	}
	return nil
}
