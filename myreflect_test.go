package leetcode

import (
	"reflect"
	"testing"
)

type ReflectTestA struct {
	A int
	B string
}

type ReflectTestB struct {
	A int
	B string
}

func TestReflect(t *testing.T) {
	a := &ReflectTestA{
		A: 10,
		B: "test",
	}

	t.Log(reflect.TypeOf(*a).Name())
	t.Log(reflect.ValueOf(*a).Type().Name())
}
