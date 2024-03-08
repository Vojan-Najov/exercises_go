package main

import (
	"reflect"
	"testing"
)

func TestStackAll(t *testing.T) {
	itemsInt := []int{3, 6, 2, 6, 2}
	testStack(t, itemsInt)

	itemsString := []string{"s", "sdfsf", "444"}
	testStack(t, itemsString)

	itemsStringEmpty := []string{}
	testStack(t, itemsStringEmpty)

	itemsStruct := []UselessStruct{{
		a: 0,
		b: 0,
	}, {
		a: 4,
		b: 1,
	}}
	testStack(t, itemsStruct)

	itemsPointers := []*UselessStruct{{
		a: 0,
		b: 0,
	}, {
		a: 4,
		b: 1,
	}}
	testStack(t, itemsPointers)
}

type UselessStruct struct {
	a int
	b int
}

func testStack[T any](t *testing.T, items []T) {
	t.Helper()
	stack := Stack[T]{}
	for _, item := range items {
		stack.Push(item)
	}

	for i := len(items) - 1; i > -1; i-- {
		item, err := stack.Pop()
		if err != nil {
			t.Fatalf("stack.Pop(): %s", err)
		}
		if !reflect.DeepEqual(item, items[i]) {
			t.Fatalf("expected: %v, got: %v", items[i], item)
		}
	}
	_, err := stack.Pop()
	if err == nil {
		t.Fatalf("expecting error on Pop(), got none")
	}

}
