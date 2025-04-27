// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package hryzun

import "fmt"

type Code uint8

func (c Code) Error() string {
	return fmt.Sprintf("%d", c)
}

type Status struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func NewStatus(code Code, message string) *Status {
	return &Status{Code: code, Message: message}
}

func (s Status) Error() string {
	return fmt.Sprintf("%d: %s", s.Code, s.Message)
}
