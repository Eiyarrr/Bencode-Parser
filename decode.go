package bencode

func Decode(data []byte) (any, error) {
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
		}
	}
	return str, nil
}
