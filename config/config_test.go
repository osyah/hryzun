// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config_test

import (
	"reflect"
	"testing"

	"github.com/osyah/hryzun/config"
)

type Mock struct {
	String  string `cfg:"string"`
	Integer int    `cfg:"integer"`
	Array   []any  `cfg:"array"`
	Object  Object `cfg:"object"`
}

type Object struct {
	Boolean bool `cfg:"boolean"`
}

var NewTests = map[string]struct {
	filename string
	want     *Mock
}{
	"OK": {
		filename: "fixtures/config.json",
		want: &Mock{
			String:  "Hello World!",
			Integer: 1984,
			Array:   []any{"string"},
			Object:  Object{Boolean: false},
		},
	},
}

func TestNew(t *testing.T) {
	for name, test := range NewTests {
		t.Run(name, func(t *testing.T) {
			got, err := config.New[Mock](test.filename)
			if err != nil {
				t.Fatal("hryzun/config:", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("hryzun/config: result does not match")
			}
		})
	}
}

var UnmarshalTests = map[string]struct {
	input *config.UnmarshalInput
	want  map[string]any
}{
	"OK": {
		input: &config.UnmarshalInput{
			Type:     ".json",
			Data:     []byte(`{"hello": "world"}`),
			Adapters: []config.Adapter{&config.JSON{}},
		},
		want: map[string]any{"hello": "world"},
	},
}

func TestUnmarshal(t *testing.T) {
	for name, test := range UnmarshalTests {
		t.Run(name, func(t *testing.T) {
			got, err := config.Unmarshal(test.input)
			if err != nil {
				t.Fatal("hryzun/config:", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("hryzun/config: result does not match")
			}
		})
	}
}
