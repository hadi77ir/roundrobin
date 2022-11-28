package roundrobin

import (
	"container/list"
	"sync"
)

type RoundRobin[T any] struct {
	circle           *list.List
	mutex            *sync.RWMutex
	equalityTestFunc func(T, T) bool
}

func New[T any](equalityTestFunc func(T, T) bool) *RoundRobin[T] {
	return &RoundRobin[T]{circle: list.New(), mutex: &sync.RWMutex{}, equalityTestFunc: equalityTestFunc}
}

func (r *RoundRobin[T]) Add(item T) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.circle.PushBack(item)
}

func (r *RoundRobin[T]) Next() T {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	element := r.circle.Front()
	if element == nil {
		return zero[T]()
	}
	r.circle.MoveToBack(element)
	return element.Value.(T)
}

func (r *RoundRobin[T]) Len() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.circle.Len()
}

func (r *RoundRobin[T]) Clear() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.circle.Init()
}
func (r *RoundRobin[T]) TryRemove(item T) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	element := r.circle.Front()
	result := false
	for i := 0; i < r.circle.Len(); i++ {
		if element == nil {
			return result
		}
		if r.equalityTestFunc(element.Value.(T), item) {
			// this is it. remove.
			r.circle.Remove(element)
			result = true
		}
		element = element.Next()
	}
	return result
}

func (r *RoundRobin[T]) Elements() []T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	elements := make([]T, r.circle.Len())
	element := r.circle.Front()
	for i := 0; i < r.circle.Len(); i++ {
		elements[i] = element.Value.(T)
		element = element.Next()
	}
	return elements
}

func zero[T any]() T {
	var none T
	return none
}
