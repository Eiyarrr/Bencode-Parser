package bencode

func Decode(data []byte) (any, error) {
	var ret any = nil
	var err error = nil

	switch data[0] {
	// 'd' -> dictionary
	case 100:
		ret, err = decode_dictionary(data)

	// 'i' -> integer
	case 105:
		ret, err = decode_integer(data)

	// 'l' -> list
	case 108:
		ret, err = decode_list(data)

	// other -> string
	default:
		ret, err = decode_string(data)
	}

	return ret, err
}

func decode_string(data []byte) (string, error) {
	strLen := 0
	str := ""
	for _, b := range data {
		// between 0 and 9
		if b >= 48 && b <= 57 {
			// normalize byte to 0
			strLen += int(b) - 48
			// colon ':'
		} else if b != 58 {
			str += string(b)
			strLen--

			if strLen <= 0 {
				break
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
		num = num*10 + int64(b)
	}
	return num, nil
}

func decode_list(data []byte) ([]any, error) {
	return nil, nil
}

func decode_dictionary(data []byte) (map[any]any, error) {
	return nil, nil
}
