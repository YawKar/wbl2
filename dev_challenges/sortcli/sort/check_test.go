/*
S - sorted
U - unique
R - reverse
N - numeric
J - just check
*/
package sort

import (
	"bytes"
	"testing"
)

func TestJS(t *testing.T) {
	for _, entry := range tdS {
		content, err := readS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJNS(t *testing.T) {
	for _, entry := range tdNS {
		content, err := readNS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{SortNumeric: true}, OtherOptions{}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJUS(t *testing.T) {
	for _, entry := range tdUS {
		content, err := readUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJRS(t *testing.T) {
	for _, entry := range tdRS {
		content, err := readRS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{Reverse: true}, OtherOptions{}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJRUS(t *testing.T) {
	for _, entry := range tdRUS {
		content, err := readRUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{Reverse: true}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJRNS(t *testing.T) {
	for _, entry := range tdRNS {
		content, err := readRNS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{Reverse: true, SortNumeric: true}, OtherOptions{}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJNUS(t *testing.T) {
	for _, entry := range tdNUS {
		content, err := readNUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if err := JustCheck(&Config{OrderingOptions{SortNumeric: true}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content)); err != nil {
			t.Fatal(err)
		}
	}
}
