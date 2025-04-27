// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package config

import "encoding/json"

type Adapter interface {
	Key(s string) bool
	Unmarshal(data []byte, v any) error
}

type JSON struct{}

func (JSON) Key(s string) bool { return s == ".json" }

func (JSON) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
