// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package container

import (
	"errors"
	"fmt"
)

type (
	Handler func(*Base) any
	Closer  func() error
)

type Base struct {
	h map[string]any
	c map[string]Closer
}

func New() *Base {
	return &Base{h: make(map[string]any), c: make(map[string]Closer)}
}

func (b Base) RegisterHandler(name string, handler Handler) { b.h[name] = handler }

func (b Base) RegisterCloser(name string, closer Closer) { b.c[name] = closer }

func (b *Base) Get(name string) any {
	value, ok := b.h[name]
	if !ok {
		panic("hryzun/container: invalid name")
	}

	if handler, ok := value.(Handler); ok {
		value = handler(b)
		b.h[name] = value
	}

	return value
}

func (b Base) Close() error {
	var group error

	for name, closer := range b.c {
		if err := closer(); err != nil {
			group = errors.Join(group, fmt.Errorf("%s: %s", name, err))
		}
	}

	return group
}

func Get[T any](base *Base, name string) T { return base.Get(name).(T) }
