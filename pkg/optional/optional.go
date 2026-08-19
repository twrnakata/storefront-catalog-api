package optional

import (
	"bytes"
	"encoding/json"
)

type Optional[T any] struct {
	set   bool
	null  bool
	value T
}

func From[T any](value T) Optional[T] {
	return Optional[T]{set: true, value: value}
}

func (optional *Optional[T]) UnmarshalJSON(data []byte) error {
	optional.set = true
	if bytes.Equal(data, []byte("null")) {
		optional.null = true
		var zero T
		optional.value = zero
		return nil
	}
	optional.null = false
	return json.Unmarshal(data, &optional.value)
}

func (optional Optional[T]) IsSet() bool {
	return optional.set
}

func (optional Optional[T]) IsNull() bool {
	return optional.set && optional.null
}

func (optional Optional[T]) Value() T {
	return optional.value
}
