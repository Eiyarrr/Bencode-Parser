package bencode

func Decode(data []byte) (any, error) {
	pos := 0
	return decode(data, &pos)
}

func decode(data []byte, pos *int) (any, error) {

}
