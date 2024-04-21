// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import "encoding/json"

type JSON struct{}

func (JSON) Peek(s string) bool { return s == ".json" }

func (JSON) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
