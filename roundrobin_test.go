package roundrobin

import (
	"testing"
)

func contains[T any](slice []T, element T, equalsFunc func(T, T) bool) bool {
	for _, t := range slice {
		if equalsFunc(t, element) {
			return true
		}
	}
	return false
}
func equals(a, b string) bool {
	return a == b
}

func TestRoundRobin_Elements(t *testing.T) {
	r := New[string](equals)
	circle := []string{"google.com", "google.nl", "google.de"}
	for _, s := range circle {
		r.Add(s)
	}
	elements := r.Elements()
	for i := 0; i < len(circle); i++ {
		if !contains(elements, circle[i], equals) {
			t.Fail()
		}
	}
}
func TestRoundRobin_CircularEffect(t *testing.T) {
	r := New[string](equals)
	circle := []string{"google.com", "google.nl", "google.de"}
	for _, s := range circle {
		r.Add(s)
	}
	elements := make([]string, len(circle))
	// rotate the circle completely 4 times
	for j := 0; j < 4; j++ {
		for i := 0; i < len(elements); i++ {
			elements[i] = r.Next()
		}
		for i := 0; i < len(circle); i++ {
			if !contains(elements, circle[i], equals) {
				t.Fail()
			}
		}
	}
}

func TestRoundRobin_EmptyListElements(t *testing.T) {
	r := New[string](equals)
	elements := r.Elements()
	if len(elements) != 0 {
		t.Fail()
	}
}

func TestRoundRobin_EmptyPop(t *testing.T) {
	r := New[string](equals)
	next := r.Next()
	if next != "" {
		t.Fail()
	}
}
