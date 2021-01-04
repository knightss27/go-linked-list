package list

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	expected := reflect.TypeOf(&List{})
	actual := reflect.TypeOf(NewList())

	if actual != expected {
		t.Errorf("Test failed for %v", "NewList")
	}
}

func TestAddFirst(t *testing.T) {
	var tests = []struct {
		expected, actual *List
	}{
		{&List{1, &node{1, nil, nil}, &node{1, nil, nil}}, NewList().AddFirst(1)},
		{&List{2, &node{2, nil, nil}, &node{1, nil, nil}}, NewList().AddFirst(1).AddFirst(2)},
		{&List{3, &node{3, nil, nil}, &node{1, nil, nil}}, NewList().AddFirst(1).AddFirst(2).AddFirst(3)},
	}

	current := "AddFirst"

	for _, test := range tests {
		t.Run(current, func(t *testing.T) {
			if test.actual.size != test.expected.size {
				t.Errorf("%v list.%v = %d, wanted %d", current, "size", test.actual.size, test.expected.size)
			} else if test.actual.head.data != test.expected.head.data {
				t.Errorf("%v list.%v = %d, wanted %d", current, "head", test.actual.head.data, test.expected.head.data)
			} else if test.actual.tail.data != test.expected.tail.data {
				t.Errorf("%v list.%v = %d, wanted %d", current, "tail", test.actual.tail.data, test.expected.tail.data)
			}
		})
	}
}

func TestAddLast(t *testing.T) {
	var tests = []struct {
		expected, actual *List
	}{
		{&List{1, &node{1, nil, nil}, &node{1, nil, nil}}, NewList().AddLast(1)},
		{&List{2, &node{1, nil, nil}, &node{2, nil, nil}}, NewList().AddLast(1).AddLast(2)},
		{&List{3, &node{1, nil, nil}, &node{3, nil, nil}}, NewList().AddLast(1).AddLast(2).AddLast(3)},
	}

	current := "TestAddLast"

	for _, test := range tests {
		t.Run(current, func(t *testing.T) {
			if test.actual.size != test.expected.size {
				t.Errorf("%v list.%v = %d, wanted %d", current, "size", test.actual.size, test.expected.size)
			} else if test.actual.head.data != test.expected.head.data {
				t.Errorf("%v list.%v = %d, wanted %d", current, "head", test.actual.head.data, test.expected.head.data)
			} else if test.actual.tail.data != test.expected.tail.data {
				t.Errorf("%v list.%v = %d, wanted %d", current, "tail", test.actual.tail.data, test.expected.tail.data)
			}
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	var tests = []struct {
		expected, actual *List
	}{
		{&List{0, nil, nil}, NewList().AddFirst(1).RemoveFirst()},
	}

	current := "TestRemoveFirst"

	for _, test := range tests {
		t.Run(current, func(t *testing.T) {
			if test.actual.size != test.expected.size {
				t.Errorf("%v list.%v = %d, wanted %d", current, "size", test.actual.size, test.expected.size)
			} else if test.actual.head != test.expected.head {
				t.Errorf("%v list.%v = %v, wanted %v", current, "head", test.actual.head, test.expected.head)
			} else if test.actual.tail != test.expected.tail {
				t.Errorf("%v list.%v = %v, wanted %v", current, "tail", test.actual.tail, test.expected.tail)
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	var tests = []struct {
		expected, actual *List
	}{
		{&List{0, nil, nil}, NewList().AddLast(1).RemoveLast()},
	}

	current := "TestRemoveLast"

	for _, test := range tests {
		t.Run(current, func(t *testing.T) {
			if test.actual.size != test.expected.size {
				t.Errorf("%v list.%v = %d, wanted %d", current, "size", test.actual.size, test.expected.size)
			} else if test.actual.head != test.expected.head {
				t.Errorf("%v list.%v = %v, wanted %v", current, "head", test.actual.head, test.expected.head)
			} else if test.actual.tail != test.expected.tail {
				t.Errorf("%v list.%v = %v, wanted %v", current, "tail", test.actual.tail, test.expected.tail)
			}
		})
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		expected, actual *List
	}{
		{&List{0, nil, nil}, NewList().AddFirst(1).Clear()},
	}

	current := "TestClear"

	for _, test := range tests {
		t.Run(current, func(t *testing.T) {
			if test.actual.size != test.expected.size {
				t.Errorf("%v list.%v = %d, wanted %d", current, "size", test.actual.size, test.expected.size)
			} else if test.actual.head != test.expected.head {
				t.Errorf("%v list.%v = %v, wanted %v", current, "head", test.actual.head, test.expected.head)
			} else if test.actual.tail != test.expected.tail {
				t.Errorf("%v list.%v = %v, wanted %v", current, "tail", test.actual.tail, test.expected.tail)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		name             string
		expected, actual bool
	}{
		{"Size0", true, NewList().IsEmpty()},
		{"Size1", false, NewList().AddFirst(1).IsEmpty()},
	}

	current := "TestIsEmpty"

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.actual != test.expected {
				t.Errorf("%v list.%v = %v, wanted %v", current, "IsEmpty", test.actual, test.expected)
			}
		})
	}
}
