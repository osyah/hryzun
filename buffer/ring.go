// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package buffer

import "sync"

type Ring[T any] struct {
	buf   []T
	size  int
	write int
	count int
	mu    sync.Mutex
}

func NewRing[T any](size int) *Ring[T] {
	return &Ring[T]{
		buf:  make([]T, size),
		size: size,
	}
}

func (r *Ring[T]) Add(value T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.buf[r.write] = value
	r.write = (r.write + 1) % r.size

	if r.count < r.size {
		r.count++
	}
}

func (r *Ring[T]) Get() []T {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]T, r.count)

	for i := 0; i < r.count; i++ {
		result[i] = r.buf[((r.write + i) % r.size)]
	}

	return result
}

func (r *Ring[T]) First() (T, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.count == 0 {
		var zero T

		return zero, false
	}

	return r.buf[r.write], true
}

func (r *Ring[T]) Last() (T, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.count == 0 {
		var zero T

		return zero, false
	}

	return r.buf[((r.write + r.count - 1) % r.size)], true
}

func (r *Ring[T]) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.count
}
