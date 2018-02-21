package cachedjson

import (
	"encoding/json"
	"fmt"
)

type Cacher interface {
	GetCachedJSON() *Cache
}

type Cache struct {
	data []byte
}

func getCachedJSON(cacher Cacher) (*Cache, error) {
	cached_json := cacher.GetCachedJSON()
	if cached_json == nil {
		return nil, fmt.Errorf("%T.GetCachedJSON() should not return nil", cacher)
	}
	return cached_json, nil
}

func Marshal(cacher Cacher) ([]byte, error) {
	cached_json, err := getCachedJSON(cacher)
	if err != nil {
		return nil, err
	}
	if cached_json.data != nil {
		return cached_json.data, nil
	}
	data, err := json.Marshal(cacher)
	if err != nil {
		return nil, err
	}
	cached_json.data = data
	return data, nil
}

func Unmarshal(data []byte, cacher Cacher) error {
	cached_json, err := getCachedJSON(cacher)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, cacher)
	if err != nil {
		return err
	}
	cached_json.data = data
	return nil
}

func Update(cacher Cacher) error {
	cached_json, err := getCachedJSON(cacher)
	if err != nil {
		return err
	}
	cached_json.data = nil
	return nil
}
