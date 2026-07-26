package bencode

import "testing"

func TestDecodeString(t *testing.T) {
	result, err := Decode([]byte("4:spam"))

	if err != nil {
		t.Fatal(err)
	}

	if result != "spam" {
		t.Fatal("Wrong result")
	}
}
