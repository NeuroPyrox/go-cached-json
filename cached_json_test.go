package cachedjson

import (
	"bytes"
	"testing"
)

type testStruct struct {
	cache Cache
	A     int
	B     string
}

func TestCache_ShouldBeEmpty_ByDefault(t *testing.T) {
	var obj testStruct
	if !obj.cache.IsEmpty() {
		t.Error("Cache is not empty by default")
	}
}

func TestCache_ShoudNotBeEmpty_AfterUnmarshal(t *testing.T) {
	var obj testStruct
	data := []byte(`{"A":1,"B":"hi"}`)
	err := obj.cache.Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.cache.IsEmpty() {
		t.Error("Cache is empty after unmarshalling")
	}
}

func TestCache_ShoudNotBeEmpty_AfterUpdate(t *testing.T) {
	obj := testStruct{A: 1, B: "hi"}
	err := obj.cache.Update(obj)
	if err != nil {
		t.Error(err)
	}
	if obj.cache.IsEmpty() {
		t.Error("Cache is empty after updating")
	}
}

func TestCache_ShouldBeEmpty_AfterClear(t *testing.T) {
	obj := testStruct{A: 1, B: "hi"}
	err := obj.cache.Update(obj)
	if err != nil {
		t.Error(err)
	}
	obj.cache.Clear()
	if !obj.cache.IsEmpty() {
		t.Error("Cache is not empty after clearing")
	}
}

func TestCache_MarshalShouldReturnX_AfterXIsUnmarshalled(t *testing.T) {
	var obj testStruct
	data := []byte(`{"A":1,"B":"hi"}`)
	err := obj.cache.Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := obj.cache.Marshal()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Errorf("%v != %v", string(data), string(data_copy))
	}
}

func TestCache_MarshalShouldReturnError_IfCacheIsEmpty(t *testing.T) {
	var obj testStruct
	if _, err := obj.cache.Marshal(); err == nil {
		t.Error("Marshal does not return error when cache is empty")
	}
}

func TestCache_UnmarshalShouldReturnError_IfInvalidJSON(t *testing.T) {
	var obj testStruct
	invalid_json := []byte("w8345uiwj8ur")
	err := obj.cache.Unmarshal(invalid_json, &obj)
	if err == nil {
		t.Error("Expected error when passing invalid json")
	}
}

func TestCache_MarshalShouldBeSame_AfterUnmarshallingInvalidJSON(t *testing.T) {
	obj := testStruct{A: 1, B: "hi"}
	err := obj.cache.Update(obj)
	if err != nil {
		t.Error(err)
	}
	data, err := obj.cache.Marshal()
	if err != nil {
		t.Error(err)
	}
	invalid_json := []byte("w8345uiwj8ur")
	_ = obj.cache.Unmarshal(invalid_json, &obj)
	data_copy, err := obj.cache.Marshal()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Errorf("%v != %v", string(data), string(data_copy))
	}
}

func TestCache_MarshalShouldRepresentX_AfterUpdateWithX(t *testing.T) {
	obj := testStruct{A: 1, B: "hi"}
	data := []byte(`{"A":1,"B":"hi"}`)
	err := obj.cache.Update(obj)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := obj.cache.Marshal()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Errorf("%v != %v", string(data), string(data_copy))
	}
}

func TestCache_UpdateReturnsError_WhenObjectHasNoJSONRepresentation(t *testing.T) {
	no_json_representation := make(chan int)
	cache := Cache{}
	err := cache.Update(no_json_representation)
	if err == nil {
		t.Error("Expected error when calling update with an object that has no json representation")
	}
}
