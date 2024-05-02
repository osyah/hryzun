// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package status

import "fmt"

type Code uint8

func (c Code) Error() string {
	return fmt.Sprintf("%d", c)
}

type Status struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func New(code Code, message string) *Status {
	return &Status{Code: code, Message: message}
}

func (s Status) Error() string {
	return fmt.Sprintf("%d: %s", s.Code, s.Message)
}
