// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package module

import (
	"errors"
	"fmt"
)

type (
	Handler func(*Container) any
	Closer  func() error
)

type Container struct {
	handlers map[string]any
	closers  map[string]Closer
}

func New() *Container {
	return &Container{
		handlers: make(map[string]any),
		closers:  make(map[string]Closer),
	}
}

func (c Container) RegisterHandler(name string, handler Handler) {
	c.handlers[name] = handler
}

func (c Container) RegisterCloser(name string, closer Closer) {
	c.closers[name] = closer
}

func (c *Container) Get(name string) any {
	value, ok := c.handlers[name]
	if !ok {
		panic("hryzun/module: invalid name")
	}

	if handler, ok := value.(Handler); ok {
		value = handler(c)
		c.handlers[name] = value
	}

	return value
}

func (c Container) Close() error {
	var group error

	for name, closer := range c.closers {
		if err := closer(); err != nil {
			group = errors.Join(group, fmt.Errorf("%s: %s", name, err))
		}
	}

	return group
}

func Get[T any](container *Container, name string) T {
	return container.Get(name).(T)
}
