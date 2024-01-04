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

func Sort(c *Config, r io.Reader, w io.Writer) error {
	switch {
	case c.SortNumeric:
		if err := sortNumeric(c, r, w); err != nil {
			return err
		}
	default:
		if err := sortStandard(c, r, w); err != nil {
			return err
		}
	}
	return nil
}

func sortStandard(c *Config, r io.Reader, w io.Writer) error {
	lines := getLines(r)
	sortFunc := strings.Compare
	if c.IgnoreLeadingBlanks {
		sortFunc = wrapIgnoreLeadingBlanks(sortFunc)
	}
	if c.Reverse {
		sortFunc = wrapReverse(sortFunc)
	}
	slices.SortStableFunc(lines, sortFunc)

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

func sortNumeric(c *Config, r io.Reader, w io.Writer) error {
	lines := getLines(r)
	if c.UniqueOnly {
		unique := make([]string, 0)
		{
			seen := make(map[string]int)
			for i := range lines {
				s1N, _, _ := getNumericPrefAndRest(lines[i])
				if _, found := seen[s1N.String()]; !found {
					seen[s1N.String()] = i
				}
			}
			seenKeys := make([]*big.Int, 0)
			for k := range seen {
				b := big.NewInt(0)
				err := b.UnmarshalText([]byte(k))
				if err != nil {
					panic(err)
				}
				seenKeys = append(seenKeys, b)
			}
			slices.SortFunc(seenKeys, (*big.Int).Cmp)
			for _, key := range seenKeys {
				unique = append(unique, lines[seen[key.String()]])
			}
		}
		lines = unique
	} else {
		numericSorter := func(s1, s2 string) int {
			s1Num, s1Rest, s1Ok := getNumericPrefAndRest(s1)
			s2Num, s2Rest, s2Ok := getNumericPrefAndRest(s2)
			if numCmp := s1Num.Cmp(s2Num); numCmp != 0 {
				return numCmp
			} else if s1Ok != s2Ok {
				if s1Ok && !s2Ok {
					return -1
				} else {
					return 1
				}
			} else {
				return strings.Compare(s1Rest, s2Rest)
			}
		}
		if c.IgnoreLeadingBlanks {
			numericSorter = wrapIgnoreLeadingBlanks(numericSorter)
		}
		if c.Reverse {
			numericSorter = wrapReverse(numericSorter)
		}
		slices.SortStableFunc(lines, numericSorter)
	}
	for _, line := range lines {
		_, err := fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}
	return nil
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

func getNumericPrefAndRest(s string) (*big.Int, string, bool) {
	sNonNumericIx := strings.IndexFunc(s, func(r rune) bool { return !(unicode.IsDigit(r) || r == '+' || r == '-') })
	if sNonNumericIx == -1 {
		return &big.Int{}, s, false
	}
	sPrefNum := &big.Int{}
	if err := sPrefNum.UnmarshalText([]byte(s[:sNonNumericIx])); err != nil {
		return &big.Int{}, s, false
	}
	return sPrefNum, s[sNonNumericIx:], true
}
