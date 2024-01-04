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
	"embed"
	"io/fs"
	"slices"
	"testing"
)

const TESTDATA = "testdata/"

//go:embed testdata
var tdRoot embed.FS

var (
	// basic samples
	td []fs.DirEntry
	// sorted samples
	tdS []fs.DirEntry
	// numeric sorted samples
	tdNS []fs.DirEntry
	// unique sorted samples
	tdUS []fs.DirEntry
	// reverse sorted samples
	tdRS []fs.DirEntry
	// reverse unique sorted samples
	tdRUS []fs.DirEntry
	// reverse numeric sorted samples
	tdRNS []fs.DirEntry
	// numeric unique sorted samples
	tdNUS []fs.DirEntry
	// k1 sorted samples
	tdK1S []fs.DirEntry
	// k2 sorted samples
	tdK2S []fs.DirEntry
)

// Read basic sample
func read(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "samples/" + name)
}

// Read sorted sample
func readS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "sorted_samples/" + name)
}

// Read numeric sorted sample
func readNS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "numeric_sorted_samples/" + name)
}

// Read reverse sorted sample
func readRS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "reverse_sorted_samples/" + name)
}

// Read unique sorted sample
func readUS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "unique_sorted_samples/" + name)
}

// Read reverse unique sorted sample
func readRUS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "reverse_unique_sorted_samples/" + name)
}

// Read reverse numeric sorted sample
func readRNS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "reverse_numeric_sorted_samples/" + name)
}

// Read numeric unique sorted sample
func readNUS(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "numeric_unique_sorted_samples/" + name)
}

// Read k1 sorted sample
func readK1S(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "k1_sorted_samples/" + name)
}

// Read k2 sorted sample
func readK2S(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "k2_sorted_samples/" + name)
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
	tdNS, err = tdRoot.ReadDir("testdata/numeric_sorted_samples")
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
	tdRUS, err = tdRoot.ReadDir("testdata/reverse_unique_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdRNS, err = tdRoot.ReadDir("testdata/reverse_numeric_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdNUS, err = tdRoot.ReadDir("testdata/numeric_unique_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdK1S, err = tdRoot.ReadDir("testdata/k1_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdK2S, err = tdRoot.ReadDir("testdata/k2_sorted_samples")
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
		if err := Sort(&Config{}, bytes.NewReader(content), &result); err != nil {
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
		if err := Sort(&Config{OrderingOptions{}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content), &result); err != nil {
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

func TestNS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{SortNumeric: true}, OtherOptions{}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readNS(entry.Name())
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
		if err := Sort(&Config{OrderingOptions{Reverse: true}, OtherOptions{}}, bytes.NewReader(content), &result); err != nil {
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

func TestRUS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{Reverse: true}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readRUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestRNS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{Reverse: true, SortNumeric: true}, OtherOptions{}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readRNS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestNUS(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{SortNumeric: true}, OtherOptions{UniqueOnly: true}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readNUS(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestK1S(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{SortKey: &SortKey{TargetField: 1}}, OtherOptions{}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readK1S(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestK2S(t *testing.T) {
	for _, entry := range td {
		content, err := read(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(&Config{OrderingOptions{SortKey: &SortKey{TargetField: 2}}, OtherOptions{}}, bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readK2S(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}
