package cachedjson

import (
	"bytes"
	"reflect"
	"testing"
)

type testCacher struct {
	Cache
	A int
	B bool
}

func (t *testCacher) GetCachedJSON() *Cache {
	return &t.Cache
}

func TestMarshal(t *testing.T) {
	obj := &testCacher{A: 1, B: true}
	expected_data := []byte(`{"A":1,"B":true}`)
	data, err := Marshal(obj)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, expected_data) {
		t.Errorf("Expected %v but got %v", string(expected_data), string(data))
	}
}

func TestUnmarshal_ShouldUndoMarshal(t *testing.T) {
	obj := &testCacher{A: 1, B: true}
	data, err := Marshal(obj)
	if err != nil {
		t.Error(err)
	}
	obj_copy := new(testCacher)
	err = Unmarshal(data, obj_copy)
	if err != nil {
		t.Log("JSON:", string(data))
		t.Error(err)
	}
	if !reflect.DeepEqual(obj, obj_copy) {
		t.Error("The following should be equal:")
		t.Log(obj)
		t.Log(obj_copy)
	}
}

func TestMarshal_WhenDataIsChanged_ShouldStillReturnSameResult(t *testing.T) {
	cacher := &testCacher{A: 1, B: true}
	data, err := Marshal(cacher)
	if err != nil {
		t.Error(err)
	}
	cacher.A = 332
	data_copy, err := Marshal(cacher)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Error("The following should be equal:")
		t.Log("Data:", string(data))
		t.Log("Data Copy:", string(data_copy))
	}
}

func TestMarshal_WhenUpdateIsCalled_ShouldReturnUpdatedResult(t *testing.T) {
	cacher := &testCacher{A: 1, B: true}
	_, err := Marshal(cacher)
	if err != nil {
		t.Error(err)
	}
	cacher_copy := &testCacher{A: 9, B: false}
	cacher.A = cacher_copy.A
	cacher.B = cacher_copy.B
	Update(cacher)
	data, err := Marshal(cacher)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := Marshal(cacher_copy)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Error("The following should be equal:")
		t.Log("Data:", string(data))
		t.Log("Data Copy:", string(data_copy))
	}
}

func TestMarshal_AfterUnmarshalIsCalled_ShouldReturnUpdatedResult(t *testing.T) {
	cacher1 := &testCacher{A: 1, B: true}
	cacher2 := &testCacher{A: 4, B: false}
	_, err := Marshal(cacher1)
	if err != nil {
		t.Error(err)
	}
	data, err := Marshal(cacher2)
	if err != nil {
		t.Error(err)
	}
	err = Unmarshal(data, cacher1)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := Marshal(cacher1)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Error("The following should be equal:")
		t.Log("Data:", string(data))
		t.Log("Data Copy:", string(data_copy))
	}
}
