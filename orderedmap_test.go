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

func BenchmarkSet(b *testing.B) {
	m := New[string, int]()
	for i := 0; i < b.N; i++ {
		m.Set("foo", 1)
	}
}

func BenchmarkGet(b *testing.B) {
	m := New[string, int]()
	m.Set("foo", 1)
	for i := 0; i < b.N; i++ {
		_, _ = m.Get("foo")
	}
}
