/*
S - sorted
U - unique
R - reverse
J - just check
*/
package sort

import (
	"bytes"
	"embed"
	"io/fs"
	"slices"
	"testing"
)

const (
	APPNAME  = "sort"
	TESTDATA = "testdata/"
)

//go:embed testdata
var tdRoot embed.FS

var (
	// basic samples
	td []fs.DirEntry
	// sorted samples
	tdS []fs.DirEntry
	// unique sorted samples
	tdUS []fs.DirEntry
	// reverse sorted samples
	tdRS []fs.DirEntry
)

// Read basic sample
func read(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "samples/" + name)
}

// Read sorted sample
func readS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "sorted_samples/" + name)
}

// Read reverse sorted sample
func readRS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "reverse_sorted_samples/" + name)
}

// Read unique sorted sample
func readUS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "unique_sorted_samples/" + name)
}

func init() {
	setupSamples()
}

func setupSamples() {
	var err error
	td, err = tdRoot.ReadDir("testdata/samples")
	if err != nil {
		panic(err)
	}
	tdS, err = tdRoot.ReadDir("testdata/sorted_samples")
	if err != nil {
		panic(err)
	}
	tdUS, err = tdRoot.ReadDir("testdata/unique_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdRS, err = tdRoot.ReadDir("testdata/reverse_sorted_samples")
	if err != nil {
		panic(err)
	}
}

func TestS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestUS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{UniqueOnly: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestRS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{Reverse: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readRS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestJS(t *testing.T) {
	for _, entry := range tdS {
		content, err := readS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{JustCheck: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
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
		result := bytes.Buffer{}
		if err := Sort(Config{UniqueOnly: true, JustCheck: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
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
		result := bytes.Buffer{}
		if err := Sort(Config{JustCheck: true, Reverse: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
	}
}
