package cachedjson

import (
	"testing"
)

type testFields struct {
	A int
	B string
}

type testStruct struct {
	Cache
	testFields
}

func newTestStruct() *testStruct {
	ts := new(testStruct)
	ts.Cache = New(&ts.testFields)
	return ts
}

func TestCache_MarshalJSON_WhenFirstCreated(t *testing.T) {
	ts := newTestStruct()
	ts.A = 4
	ts.B = "h"
	actual, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	expected := `{"A":4,"B":"h"}`
	if string(actual) != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestCache_MarshalJSON_AfterObjModified(t *testing.T) {
	ts := newTestStruct()
	ts.A = 4
	ts.B = "h"
	ts.MarshalJSON()
	ts.A = 6
	ts.B = "d"
	actual, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	expected := `{"A":4,"B":"h"}`
	if string(actual) != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestCache_MarshalJSON_AfterObjModifiedAndUpdate(t *testing.T) {
	ts := newTestStruct()
	ts.A = 4
	ts.B = "h"
	ts.MarshalJSON()
	ts.A = 6
	ts.B = "d"
	ts.Update()
	actual, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	expected := `{"A":6,"B":"d"}`
	if string(actual) != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestCache_MarshalJSON_OnUnsupportedType(t *testing.T) {
	ch := make(chan int)
	c := New(ch)
	_, err := c.MarshalJSON()
	if err == nil {
		t.Error("Expected error!")
	}
	t.Logf("Returned:\n%v", err)
}

func TestCache_UnmarshalJSON_OnValidJSON(t *testing.T) {
	ts := newTestStruct()
	err := ts.UnmarshalJSON([]byte(`{"A":8,"B":"hi"}`))
	if err != nil {
		t.Error(err)
	}
	expected := testFields{8, "hi"}
	if ts.testFields != expected {
		t.Errorf("%v != %v", ts.testFields, expected)
	}
}

func TestCache_UnmarshalJSON_OnInvalidJSON(t *testing.T) {
	ts := newTestStruct()
	ts.A = 4
	ts.B = "h"
	err := ts.UnmarshalJSON([]byte(`{"A":8,"B":re6yt43trfe&^%$R}`))
	if err == nil {
		t.Error("Expected error!")
	}
	t.Logf("Returned:\n%v", err)
	expected := testFields{4, "h"}
	if ts.testFields != expected {
		t.Errorf("%v != %v", ts.testFields, expected)
	}
}
