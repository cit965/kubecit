package data

import (
	"reflect"
	"testing"
)

type Source struct {
	String string
	Int    int
	Map    map[string]string
	List   []string
	Struct A
	Ptr    *A
}

type A struct {
	Foo string
	Bar int64
}

type Dest struct {
	String string
	Int    int64
	Map    map[string]string
	List   []string
	Struct A
	Ptr    *A
}

func TestConvertType(t *testing.T) {
	s := &Source{
		String: "test string",
		Int:    100,
		Map:    map[string]string{"abc": "123"},
		List:   []string{"zoo", "foo", "xoo"},
		Struct: A{Foo: "foo", Bar: 32},
		Ptr:    &A{Foo: "foo ptr", Bar: 64},
	}
	d := &Dest{}
	err := ConvertType(s, d)
	if err != nil && !reflect.DeepEqual(s.String, d.String) && !reflect.DeepEqual(s.Int, d.Int) && !reflect.DeepEqual(s.Map, d.Map) &&
		!reflect.DeepEqual(s.List, d.List) && !reflect.DeepEqual(s.Struct, d.Struct) && !reflect.DeepEqual(s.Ptr, &d.Ptr) {
		t.Fatal(err)
	}
}
