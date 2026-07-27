package bencode

import (
	"errors"
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

func fill(value any, destination reflect.Value) error {
	switch destination.Kind() {
		case reflect.Struct:

		case reflect.Slice:
			
		case reflect.Int64:

		case reflect.String:
			str, ok := value.(string)
			if !ok {
				return errors.New("unexpected error trying to fill string")
			}

			destination.SetString(str)
			
	}
	return nil
}
