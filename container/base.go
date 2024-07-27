// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package container

type Handler func(*Base) any

type Base struct{ m map[string]any }

func New() *Base {
	return &Base{m: make(map[string]any)}
}

func (b Base) Register(name string, handler Handler) { b.m[name] = handler }

func (b *Base) Get(name string) any {
	value, ok := b.m[name]
	if !ok {
		panic("hryzun/container: invalid name")
	}

	if handler, ok := value.(Handler); ok {
		value = handler(b)
		b.m[name] = value
	}

	return value
}

func Get[T any](base *Base, name string) T { return base.Get(name).(T) }
