package co

import "sync"

// CreateOnce ...
type CreateOnce[T any] struct {
	once sync.Once
	data T
}

// New ...
// @receiver *CreateOnce[T]
// @param T
func (t *CreateOnce[T]) New(v T) {
	t.once.Do(func() {
		t.data = v
	})
}

// FuncNew ...
// @receiver *CreateOnce[T]
// @param func() T
func (t *CreateOnce[T]) FuncNew(fn func() T) {
	t.once.Do(func() {
		t.data = fn()
	})
}

// Data ...
// @receiver *CreateOnce[T]
// @return T
func (t *CreateOnce[T]) Data() T {
	return t.data
}
