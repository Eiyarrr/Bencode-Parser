package bencode

import (
	"bytes"
	"errors"
	"fmt"
)

func Encode(value any) ([]byte, error) {
	switch v := value.(type) {
	case int64:
		return encodeInteger(v)

	case []any:
		return encodeList(v)

	case map[string]any:
		return encodeDictionary(v)

	case string:
		return encodeString(v)

	default:
		return nil, errors.New("unexpected value type during encoding")
	}
}

func encodeInteger(value int64) ([]byte, error) {
	// "i%de" -> "i" + numbers + "e"
	return []byte(fmt.Sprintf("i%de", value)), nil
}

func encodeList(list []any) ([]byte, error)                {
	var buf bytes.Buffer

	// add prefix
	buf.WriteByte('l')

	for _, item := range list {
		encoded, err := Encode(item)
		if err != nil {
			return nil, err
		}

		buf.Write(encoded)
	}

	// add suffix
	buf.WriteByte('e')

	return buf.Bytes(), nil
}

func encodeDictionary(value map[string]any) ([]byte, error) { return nil, nil }

func encodeString(value string) ([]byte, error) {
	// "%d:%s" -> numbers + ":" + string
	return []byte(fmt.Sprintf("%d:%s", len(value), value)), nil
}
