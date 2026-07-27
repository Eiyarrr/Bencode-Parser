package bencode

import (
	"fmt"
	"io"
)

func Decode(reader io.Reader) (any, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	pos := 0
	return DecodeBytes(data, &pos)
}

func DecodeBytes(data []byte, pos *int) (any, error) {
	switch data[*pos] {
	case 'i':
		return decodeInteger(data, pos)

	case 'l':
		return decodeList(data, pos)

	case 'd':
		return decodeDictionary(data, pos)

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return decodeString(data, pos)

	default:
		return nil, fmt.Errorf("unexpected character while finding correct decoder to use: %q", data[*pos])
	}
}

func decodeInteger(data []byte, pos *int) (int64, error) {
	*pos++ // skip prefix 'i'
	var num int64 = 0
	negative := false

	// check for possible negative numbers
	if data[*pos] == '-' {
		negative = true
		*pos++
	}

	for data[*pos] != 'e' {
		if !is_digit(data[*pos]) {
			return num, fmt.Errorf("unexpected character while trying to decode an integer: %q", data[*pos])
		}
		// normalize byte to 0
		num = num*10 + int64(data[*pos]) - 48
		*pos++
	}

	if negative {
		num *= -1
	}

	*pos++ // skip suffix 'e'
	return num, nil
}

func decodeList(data []byte, pos *int) ([]any, error) {
	*pos++ // skip prefix 'l'

	var list []any

	for data[*pos] != 'e' {
		value, err := DecodeBytes(data, pos)
		if err != nil {
			return nil, err
		}

		list = append(list, value)
	}

	*pos++ // skip suffix 'e'
	return list, nil
}
func decodeDictionary(data []byte, pos *int) (map[string]any, error) {
	*pos++ // skip prefix 'd'

	dictionary := make(map[string]any)

	for data[*pos] != 'e' {
		key, err := decodeString(data, pos)
		if err != nil {
			return nil, err
		}

		value, err := DecodeBytes(data, pos)
		if err != nil {
			return nil, err
		}

		dictionary[key] = value
	}

	*pos++ // skip suffix 'e'
	return dictionary, nil
}

func decodeString(data []byte, pos *int) (string, error) {
	strLen := 0
	str := ""

	// get length of string
	for data[*pos] != ':' {
		if !is_digit(data[*pos]) {
			return str, fmt.Errorf("unexpected character while trying to decode the length of a string: %q", data[*pos])
		}
		// normalize byte to 0
		strLen = strLen*10 + int(data[*pos]) - 48
		*pos++
	}

	// skip divider ':'
	*pos++

	// set str to slice from ':' to end of string
	str = string(data[*pos : *pos+strLen])
	*pos += strLen

	return str, nil
}

func is_digit(b byte) bool {
	return b >= '0' && b <= '9'
}
