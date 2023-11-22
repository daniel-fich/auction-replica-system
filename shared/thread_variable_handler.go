package shared

import (
	"sync"
)

func MakeThreadVariable[T any](val T) *ThreadVariable[T] {
	return &ThreadVariable[T]{
		value: val,
	}
}

type ThreadVariable[T any] struct {
	mu    sync.Mutex
	value T
}

func (t *ThreadVariable[T]) Update(val T) {
	t.mu.Lock()
	t.value = val
	t.mu.Unlock()
}

func (t *ThreadVariable[T]) Get() T {
	t.mu.Lock()
	val := t.value
	t.mu.Unlock()
	return val
}
