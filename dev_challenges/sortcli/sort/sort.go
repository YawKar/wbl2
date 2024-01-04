package sort

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"slices"
	"strings"
	"unicode"
)

func getLines(r io.Reader) (lines []string) {
	lines = make([]string, 0)
	scanner := bufio.NewScanner(r)
	ok := scanner.Scan()
	for ok {
		lines = append(lines, scanner.Text())
		ok = scanner.Scan()
	}
	return
}

func JustCheck(c Config, r io.Reader) error {
	lines := getLines(r)
	sortFunc := mkSortFunc(c)

	for i := 0; i < len(lines)-1; i++ {
		if sortFunc(lines[i], lines[i+1]) == 0 {
			if c.UniqueOnly {
				return fmt.Errorf("%d: disorder: %s\n", i+2, lines[i+1])
			}
		} else if sortFunc(lines[i], lines[i+1]) > 0 {
			return fmt.Errorf("%d: disorder: %s\n", i+2, lines[i+1])
		}
	}
	return nil
}

func Sort(c Config, r io.Reader, w io.Writer) error {
	lines := getLines(r)
	sortFunc := mkSortFunc(c)
	slices.SortFunc(lines, sortFunc)

	if c.UniqueOnly {
		unique := make([]string, 0)
		if len(lines) > 0 {
			unique = append(unique, lines[0])
		}
		for i := 1; i < len(lines); i++ {
			if sortFunc(lines[i], lines[i-1]) != 0 {
				unique = append(unique, lines[i])
			}
		}
		lines = unique
	}

	for _, line := range lines {
		_, err := fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func mkSortFunc(c Config) func(string, string) int {
	var sortFunc func(string, string) int
	switch {
	case c.SortKey != nil:
		panic("unimpl")
	case c.SortHumanNumeric:
		panic("unimpl")
	case c.SortMonth:
		panic("unimpl")
	case c.SortNumeric:
		sortFunc = numericSort
	default:
		sortFunc = strings.Compare
	}
	if c.IgnoreLeadingBlanks {
		sortFunc = wrapIgnoreLeadingBlanks(sortFunc)
	}
	if c.Reverse {
		sortFunc = wrapReverse(sortFunc)
	}
	return sortFunc
}

func wrapIgnoreLeadingBlanks(f func(string, string) int) func(string, string) int {
	return func(s1, s2 string) int {
		return f(strings.TrimLeftFunc(s1, unicode.IsSpace), strings.TrimLeftFunc(s2, unicode.IsSpace))
	}
}

func wrapReverse(f func(string, string) int) func(string, string) int {
	return func(s1, s2 string) int {
		return -f(s1, s2)
	}
}

func numericSort(s1, s2 string) int {
	getNumericPrefAndRest := func(s string) (*big.Int, string) {
		sNonNumericIx := strings.IndexFunc(s, func(r rune) bool { return !(unicode.IsDigit(r) || r == '+' || r == '-') })
		if sNonNumericIx == -1 {
			return &big.Int{}, s
		}
		sPrefNum := &big.Int{}
		if err := sPrefNum.UnmarshalText([]byte(s[:sNonNumericIx])); err != nil {
			return &big.Int{}, s
		}
		return sPrefNum, s[sNonNumericIx:]
	}
	s1Num, s1Rest := getNumericPrefAndRest(s1)
	s2Num, s2Rest := getNumericPrefAndRest(s2)
	if numCmp := s1Num.Cmp(s2Num); numCmp != 0 {
		return numCmp
	} else {
		return strings.Compare(s1Rest, s2Rest)
	}
}
