package cachedjson

import (
	"encoding/json"
	"sync"
)

type Cache interface {
	json.Marshaler
	json.Unmarshaler
	sync.Locker
	RLock()
	RUnlock()
	Update()
}

type cache struct {
	sync.RWMutex
	obj  interface{}
	data []byte
}

func New(obj interface{}) Cache {
	return &cache{obj: obj}
}

func (c *cache) MarshalJSON() ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	if c.data != nil {
		return c.data, nil
	}
	data, err := json.Marshal(c.obj)
	if err != nil {
		return nil, err
	}
	c.data = data
	return data, nil
}

func (c *cache) UnmarshalJSON(data []byte) error {
	c.Lock()
	defer c.Unlock()
	err := json.Unmarshal(data, c.obj)
	if err != nil {
		return err
	}
	c.data = data
	return nil
}

func (c *cache) Update() {
	c.Lock()
	defer c.Unlock()
	c.data = nil
}
