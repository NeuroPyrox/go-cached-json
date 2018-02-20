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
	obj := &testCacher{A: 1, B: true}
	data, err := Marshal(obj)
	if err != nil {
		t.Error(err)
	}
	obj.A = 332
	data_copy, err := Marshal(obj)
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
	obj_copy := &testCacher{A: 1, B: true}
	_, err := Marshal(obj_copy)
	if err != nil {
		t.Error(err)
	}
	obj := &testCacher{A: 9, B: false}
	obj_copy.A = obj.A
	obj_copy.B = obj.B
	Update(obj_copy)
	data, err := Marshal(obj)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := Marshal(obj_copy)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Error("The following should be equal:")
		t.Log("Data:", string(data))
		t.Log("Data Copy:", string(data_copy))
	}
}

func TestMarshal_WhenItHasAlreadyBeenCalled_ShouldStillUndoUnmarshal(t *testing.T) {
	obj1 := &testCacher{A: 1, B: true}
	obj2 := &testCacher{A: 4, B: false}
	_, err := Marshal(obj1)
	if err != nil {
		t.Error(err)
	}
	data, err := Marshal(obj2)
	if err != nil {
		t.Error(err)
	}
	err = Unmarshal(data, obj1)
	if err != nil {
		t.Error(err)
	}
	data_copy, err := Marshal(obj1)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_copy) {
		t.Error("The following should be equal:")
		t.Log("Data:", string(data))
		t.Log("Data Copy:", string(data_copy))
	}
}
