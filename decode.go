package bencode

import (
	"errors"
)

func Decode(data []byte) (any, error) {
	var ret any = nil
	var err error = nil

	switch {
	case data[0] == 'd':
		ret, err = decode_dictionary(data)

	case data[0] == 'i':
		ret, err = decode_integer(data)

	case data[0] == 'l':
		ret, err = decode_list(data)

	case is_digit(data[0]):
		ret, err = decode_string(data)

	// not a valid bencode token
	default:
		return nil, errors.New("Invalid Token")
	}

	return ret, err
}

func decode_string(data []byte) (string, error) {
	strLen := 0
	str := ""
	for _, b := range data {
		if is_digit(b) {
			// normalize byte to 0
			strLen += int(b) - 48
		} else {
			if b != ':' {
				str += string(b)
				strLen--

				if strLen <= 0 {
					break
				}
			}
		}
	}

	return str, nil
}

func decode_integer(data []byte) (int64, error) {
	var num int64 = 0
	for _, b := range data {
		if b == 'e' {
			break
		}

		if is_digit(b) {
			// normalize byte to 0
			num = num*10 + int64(b) - 48
		}
	}
	return num, nil
}

func decode_list(data []byte) ([]any, error) {
	return nil, nil
}

func decode_dictionary(data []byte) (map[string]any, error) {
	dictionary := make(map[string]any)
	key := ""
	value := ""

	is_key := true

	for _, b := range data {
		if b == 'e' {
			break
		}
		
		if b == ':' {
			if is_key {
				is_key = false
			} else {
				dictionary[key] = value
				key = ""
				value = ""
				is_key = true
			}
		}

		if is_key {
			key += string(b)
		} else {
			value += string(b)
		}
	}

	return dictionary, nil
}

func is_digit(b byte) bool {
	return b >= '0' && b <= '9'
}
