package bencode

import "io"

func Unmarshal(reader io.Reader, value any) error {
	value, err := Decode(reader)
	if err != nil {
		return err
	}

	return nil
}
