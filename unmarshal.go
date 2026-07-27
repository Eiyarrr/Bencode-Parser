package bencode

import (
	"errors"
	"io"
	"reflect"
	"strings"
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
		dictionary, ok := value.(map[string]any)
		if !ok {
			return errors.New("unexpected error trying to fill struct")
		}

		// for i in range of number of total destinations
		for i := range destination.NumField() {
			// .Type() gives you the metadata, like names (e.g. torrent.Name)
			// .Field(i) gives you the actual field itself at the index i
			fieldType := destination.Type().Field(i)
			fieldValue := destination.Field(i)

			// try to get the key for that destination
			key := fieldType.Tag.Get("bencode")
			if key == "" {
				key = strings.ToLower(fieldType.Name)
			}

			// try to get the value at the key
			value, ok := dictionary[key]
			if !ok {
				continue
			}

			// fills it in place
			err := fill(value, fieldValue)
			if err != nil {
				return err
			}
		}

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
