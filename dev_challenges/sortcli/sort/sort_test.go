package sort

import (
	"bytes"
	"embed"
	"io/fs"
	"slices"
	"testing"
)

const APPNAME = "sort"
const TESTDATA = "testdata/"

//go:embed testdata
var tdRoot embed.FS

var (
	td              []fs.DirEntry
	tdsorted        []fs.DirEntry
	tduniquesorted  []fs.DirEntry
	tdreversesorted []fs.DirEntry
)

// Read basic sample
func readSample(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "samples/" + name)
}

// Read sorted sample
func readSSample(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "sorted_samples/" + name)
}

// Read reverse sorted sample
func readRSSample(name string) ([]byte, error) {
	return tdRoot.ReadFile(TESTDATA + "reverse_sorted_samples/" + name)
}

// Read unique sorted sample
func readUSSample(name string) ([]byte, error) {
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
	tdsorted, err = tdRoot.ReadDir("testdata/sorted_samples")
	if err != nil {
		panic(err)
	}
	tduniquesorted, err = tdRoot.ReadDir("testdata/unique_sorted_samples")
	if err != nil {
		panic(err)
	}
	tdreversesorted, err = tdRoot.ReadDir("testdata/reverse_sorted_samples")
	if err != nil {
		panic(err)
	}
}

func TestSorted(t *testing.T) {
	for _, entry := range td {
		content, err := readSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestUniqueSorted(t *testing.T) {
	for _, entry := range td {
		content, err := readSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{UniqueOnly: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readUSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestReverseSorted(t *testing.T) {
	for _, entry := range td {
		content, err := readSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{SortReverse: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
		validContent, err := readRSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(validContent, result.Bytes()) {
			t.Fatalf("differs from valid answer: %s: %q\n !=\n%q", entry.Name(), validContent, result.Bytes())
		}
	}
}

func TestJustCheckSorted(t *testing.T) {
	for _, entry := range tdsorted {
		content, err := readUSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{UniqueOnly: true, JustCheck: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJustCheckUniqueSorted(t *testing.T) {
	for _, entry := range tduniquesorted {
		content, err := readUSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{UniqueOnly: true, JustCheck: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
	}
}

func TestJustCheckReverseSorted(t *testing.T) {
	for _, entry := range tduniquesorted {
		content, err := readRSSample(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		result := bytes.Buffer{}
		if err := Sort(Config{JustCheck: true, SortReverse: true}, APPNAME, entry.Name(), bytes.NewReader(content), &result); err != nil {
			t.Fatal(err)
		}
	}
}
