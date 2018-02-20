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

func getCachedJSON(obj interface{}) (*Cache, error) {
	cacher, ok := obj.(Cacher)
	if !ok {
		return nil, fmt.Errorf("%T cannot be converted to cachedjson.Object", obj)
	}
	cached_json := cacher.GetCachedJSON()
	if cached_json == nil {
		return nil, fmt.Errorf("%T.GetCachedJSON() should not return nil", obj)
	}
	return cached_json, nil
}

func Marshal(obj interface{}) ([]byte, error) {
	cached_json, err := getCachedJSON(obj)
	if err != nil {
		return nil, err
	}
	if cached_json.data != nil {
		return cached_json.data, nil
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	cached_json.data = data
	return data, nil
}

func Unmarshal(data []byte, obj interface{}) error {
	cached_json, err := getCachedJSON(obj)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	cached_json.data = data
	return nil
}

func Update(obj interface{}) error {
	cached_json, err := getCachedJSON(obj)
	if err != nil {
		return err
	}
	cached_json.data = nil
	return nil
}
