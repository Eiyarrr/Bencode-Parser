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

func decodeInteger(data, pos) (i64, error) {}
func decodeList(data, pos) ([]any, error) {}
func decodeDictionary(data, pos) (map[string]any, error) {}
func decodeString(data, pos) (string, error) {}
