package bencode

import "fmt"

func Decode(data []byte) (any, error) {
	pos := 0
	return decode(data, &pos)
}

func decode(data []byte, pos *int) (any, error) {
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
		return nil, fmt.Errorf("Unexpected Character: " + string(data[*pos]))
	}
}

func decodeInteger(data []byte, pos *int) (int64, error) {
	*pos++ // skip prefix
	var num int64 = 0

	for data[*pos] != 'e' {
		if !is_digit(data[*pos]) {
			return num, fmt.Errorf("Unexpected Character while trying to decode an integer: " + string(data[*pos]))
		}
		// normalize byte to 0
		num = num*10 + int64(data[*pos]) - 48
		*pos++
	}
	return num, nil
}

func decodeList(data []byte, pos *int) ([]any, error)                {}
func decodeDictionary(data []byte, pos *int) (map[string]any, error) {}
func decodeString(data []byte, pos *int) (string, error)             {}

func is_digit(b byte) bool {
	return b >= '0' && b <= '9'
}
