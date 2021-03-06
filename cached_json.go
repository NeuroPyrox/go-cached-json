package cachedjson

import (
	"encoding/json"
)

type Cache interface {
	json.Marshaler
	json.Unmarshaler
	Update()
}

type cache struct {
	obj  interface{}
	data []byte
}

func New(obj interface{}) Cache {
	return &cache{obj: obj}
}

func (c *cache) MarshalJSON() ([]byte, error) {
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
	err := json.Unmarshal(data, c.obj)
	if err != nil {
		return err
	}
	c.data = data
	return nil
}

func (c *cache) Update() {
	c.data = nil
}
