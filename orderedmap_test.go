package orderedmap

import (
	"testing"
)

func TestBasicSetGetLen(t *testing.T) {
	testcases := []struct {
		k         string
		v         int
		expectLen int
	}{
		{k: "foo", v: 1, expectLen: 1},
		{k: "bar", v: 2, expectLen: 2},
		{k: "hey", v: 3, expectLen: 3},
		{k: "foo", v: 9, expectLen: 3}, // reset
	}

	m := New[string, int]()
	for _, tt := range testcases {
		m.Set(tt.k, tt.v)
		want := tt.v
		got, ok := m.Get(tt.k)
		if !ok || got != want {
			t.Errorf("got %v, want %v, ok = %v", got, want, ok)
		}
		if l := m.Len(); l != tt.expectLen {
			t.Errorf("got Len() = %v, expected %v", l, tt.expectLen)
		}
	}
}

func TestClear(t *testing.T) {
	m := New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("hey", 3)

	m.Clear()
	if l := m.Len(); l != 0 {
		t.Errorf("after Clear: got Len() = %v, want 0", l)
	}
	if _, ok := m.Get("foo"); ok {
		t.Error("after Clear: Get(\"foo\") returned ok = true, want false")
	}

	// The map must remain usable after clearing.
	m.Set("baz", 4)
	if l := m.Len(); l != 1 {
		t.Errorf("after Clear+Set: got Len() = %v, want 1", l)
	}
	if got, ok := m.Get("baz"); !ok || got != 4 {
		t.Errorf("after Clear+Set: got %v, ok = %v, want 4, true", got, ok)
	}
}

func TestDelete(t *testing.T) {
	m := New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("hey", 3)

	// Deleting an existing key removes it and decrements Len.
	m.Delete("bar")
	if _, ok := m.Get("bar"); ok {
		t.Error("Get(\"bar\") after Delete returned ok = true, want false")
	}
	if l := m.Len(); l != 2 {
		t.Errorf("after Delete: got Len() = %v, want 2", l)
	}

	// Deleting an absent key is a no-op.
	m.Delete("missing")
	if l := m.Len(); l != 2 {
		t.Errorf("after no-op Delete: got Len() = %v, want 2", l)
	}

	// Surviving keys are untouched.
	if got, ok := m.Get("foo"); !ok || got != 1 {
		t.Errorf("got %v, %v; want 1, true", got, ok)
	}
}

func TestLenNil(t *testing.T) {
	var m *OrderedMap[string, int]
	if l := m.Len(); l != 0 {
		t.Errorf("nil map Len() = %v, want 0", l)
	}
}

func BenchmarkSet(b *testing.B) {
	m := New[string, int]()
	for b.Loop() {
		m.Set("foo", 1)
	}
}

func BenchmarkGet(b *testing.B) {
	m := New[string, int]()
	m.Set("foo", 1)
	for b.Loop() {
		_, _ = m.Get("foo")
	}
}
