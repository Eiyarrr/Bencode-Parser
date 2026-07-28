package bencode

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func GetSHA1Info(reader io.Reader) ([20]byte, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return [20]byte{}, err
	}

	// get info section of data
	start, end, err := getStartEnd(data)

	return sha1.Sum(data[start:end]), nil
}

func getStartEnd(data []byte) (int, int, error) {
	pos := 0
	pos++ // skip prefix 'd'

	for data[pos] != 'e' {
		key, err := decodeString(data, &pos)
		if err != nil {
			return 0, 0, err
		}

		start := pos

		_, err = decode(data, &pos)
		if err != nil {
			return 0, 0, err
		}

		if key == "info" {
			return start, pos, nil
		}
	}

	pos++ // skip suffix 'e'
	return 0, 0, fmt.Errorf("No info dictionary found in .torrent file!")
}
