package bencode

import "errors"

func Decode(data []byte) (any, error) {
	var ret any = nil
	var err error = nil

	switch data[0] {
	// ASCII 'd' -> dictionary
	case 100:
		ret, err = decode_dictionary(data)

	// ASCII 'i' -> integer
	case 105:
		ret, err = decode_integer(data)

	// ASCII 'l' -> list
	case 108:
		ret, err = decode_list(data)

	// other -> string
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
			// colon ':'
			if b != 58 {
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
		// wait for 'e' to exit loop
		if b == 101 {
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

func decode_dictionary(data []byte) (map[any]any, error) {
	return nil, nil
}

func is_digit(b byte) bool {
	return b >= '0' && b <= '9'
}
