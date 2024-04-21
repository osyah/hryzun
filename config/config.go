// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
)

type Adapter interface {
	Peek(s string) bool
	Unmarshal(data []byte, v any) error
}

func New[T any](filename string, adapters ...Adapter) (*T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	adapters = append(adapters, &JSON{})

	output, err := Unmarshal(&UnmarshalInput{
		Type:     filepath.Ext(filename),
		Data:     data,
		Adapters: adapters,
	})
	if err != nil {
		return nil, err
	}

	var result T

	config := mapstructure.DecoderConfig{
		TagName: "cfg",
		Result:  &result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
		),
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(output); err != nil {
		return nil, err
	}

	return &result, nil
}

type UnmarshalInput struct {
	Type     string
	Data     []byte
	Adapters []Adapter
}

func Unmarshal(input *UnmarshalInput) (map[string]any, error) {
	var result map[string]any

	for _, adapter := range input.Adapters {
		if !adapter.Peek(input.Type) {
			continue
		}

		if err := adapter.Unmarshal(input.Data, &result); err != nil {
			return nil, err
		}

		break
	}

	return result, nil
}
