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

func TestDecodeInteger(t *testing.T) {
	result, err := Decode([]byte("i152e"))

	if err != nil {
		t.Fatal(err)
	}

	if result != int64(152) {
		t.Fatal("Wrong result")
	}
}

func TestDecodeDictionary(t *testing.T) {
	result, err := Decode([]byte("d3:cow3:moo4:spam4:eggse"))

	if err != nil {
		t.Fatal(err)
	}

	dict, ok := result.(map[string]any)
	if !ok {
		t.Fatal("Result is not a dictionary")
	}

	if dict["cow"] != "moo" {
		t.Fatal("Wrong value for cow")
	}

	if dict["spam"] != "eggs" {
		t.Fatal("Wrong value for spam")
	}
}

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

func TestDecodeInteger(t *testing.T) {
	result, err := Decode([]byte("i152e"))

	if err != nil {
		t.Fatal(err)
	}

	if result != int64(152) {
		t.Fatal("Wrong result")
	}
}

func TestDecodeDictionary(t *testing.T) {
	result, err := Decode([]byte("d3:cow3:moo4:spam4:eggse"))

	if err != nil {
		t.Fatal(err)
	}

	dict, ok := result.(map[string]any)
	if !ok {
		t.Fatal("Result is not a dictionary")
	}

	if dict["cow"] != "moo" {
		t.Fatal("Wrong value for cow")
	}

	if dict["spam"] != "eggs" {
		t.Fatal("Wrong value for spam")
	}
}

func TestDecodeList(t *testing.T) {
	result, err := Decode([]byte("l4:spami42e3:fooe"))

	if err != nil {
		t.Fatal(err)
	}

	list, ok := result.([]any)
	if !ok {
		t.Fatal("Result is not a list")
	}

	if list[0] != "spam" {
		t.Fatal("Wrong first value")
	}

	if list[1] != int64(42) {
		t.Fatal("Wrong second value")
	}

	if list[2] != "foo" {
		t.Fatal("Wrong third value")
	}
}
