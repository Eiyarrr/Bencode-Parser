package bencode

import (
	"io"
	"reflect"
)

func Unmarshal(reader io.Reader, v any) error {
	value, err := Decode(reader)
	if err != nil {
		return err
	}

	// .Elem() gets the underlying value of the pointer
	return fill(value, reflect.ValueOf(v).Elem())
}

func fill(value any, destination reflect.Value) error { return nil }
