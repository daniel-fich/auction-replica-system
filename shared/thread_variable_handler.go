package shared

import (
	"sync"
)

type ThreadVariable[T any] struct {
	mu    sync.Mutex
	Value T
}

func (t *ThreadVariable[T]) Update(val T) {
	t.mu.Lock()
	t.Value = val
	t.mu.Unlock()
}

func (t *ThreadVariable[T]) Get() T {
	t.mu.Lock()
	val := t.Value
	t.mu.Unlock()
	return val
}
