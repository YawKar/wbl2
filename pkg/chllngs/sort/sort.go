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
	lines := getLines(r)
	var linesIxs []int
	switch {
	case c.SortNumeric:
		linesIxs = sortNumeric(c, lines)
	case c.SortKey != nil:
		linesIxs = sortKey(c, lines)
	case c.SortHumanNumeric:
		panic("unimpl")
	case c.SortMonth:
		panic("unimpl")
	default:
		linesIxs = sortStandard(c, lines)
	}
	for _, pos := range linesIxs {
		_, err := fmt.Fprintln(w, lines[pos])
		if err != nil {
			return err
		}
	}
	return nil
}

func sortKey(c *Config, lines []string) []int {
	targetField := c.SortKey.TargetField
	targetFieldIsOk := func() bool {
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) >= targetField {
				return true
			}
		}
		return false
	}()
	type sIx struct {
		ix int
		s  string
	}
	sortFunc := func(s1, s2 sIx) int {
		if targetFieldIsOk {
			getTargetFieldAndRest := func(s string, targetField int) (string, string) {
				for len(s) > 0 && targetField > 0 {
					targetField--
					before, after, found := strings.Cut(s, " ")
					if targetField == 0 {
						return before, after
					} else if found {
						s = strings.TrimLeft(after, " \t")
					} else {
						return "", ""
					}
				}
				return "", ""
			}
			s1F, s1extra := getTargetFieldAndRest(s1.s, targetField)
			// fmt.Printf("%q %q %q\n", s1.s, s1F, s1extra)
			s2F, s2extra := getTargetFieldAndRest(s2.s, targetField)
			// fmt.Printf("%q %q %q\n", s2.s, s2F, s2extra)
			if s1F == s2F {
				return strings.Compare(s1extra, s2extra)
			}
			return strings.Compare(s1F, s2F)
		} else {
			return strings.Compare(s1.s, s2.s)
		}
	}
	if c.Reverse {
		prevSortFunc := sortFunc
		sortFunc = func(s1, s2 sIx) int {
			return -prevSortFunc(s1, s2)
		}
	}
	linesWithIxs := make([]sIx, len(lines))
	for i := range linesWithIxs {
		linesWithIxs[i] = sIx{i, lines[i]}
	}

	slices.SortStableFunc(linesWithIxs, sortFunc)

	if c.UniqueOnly {
		unique := make([]int, 0)
		if len(linesWithIxs) > 0 {
			unique = append(unique, linesWithIxs[0].ix)
		}
		for i := 1; i < len(linesWithIxs); i++ {
			if sortFunc(linesWithIxs[i], linesWithIxs[i-1]) != 0 {
				unique = append(unique, linesWithIxs[i].ix)
			}
		}
		return unique
	} else {
		lineIxs := make([]int, len(linesWithIxs))
		for i := range linesWithIxs {
			lineIxs[i] = linesWithIxs[i].ix
		}
		return lineIxs
	}
}

func sortStandard(c *Config, lines []string) []int {
	type sIx struct {
		ix int
		s  string
	}
	sortFunc := func(s1, s2 sIx) int {
		return strings.Compare(s1.s, s2.s)
	}
	if c.IgnoreLeadingBlanks {
		prevSortFunc := sortFunc
		sortFunc = func(s1, s2 sIx) int {
			return prevSortFunc(
				sIx{s1.ix, strings.TrimLeftFunc(s1.s, unicode.IsSpace)},
				sIx{s2.ix, strings.TrimLeftFunc(s2.s, unicode.IsSpace)},
			)
		}
	}
	if c.Reverse {
		prevSortFunc := sortFunc
		sortFunc = func(s1, s2 sIx) int {
			return -prevSortFunc(s1, s2)
		}
	}
	linesWithIxs := make([]sIx, len(lines))
	for i := range linesWithIxs {
		linesWithIxs[i] = sIx{i, lines[i]}
	}

	slices.SortStableFunc(linesWithIxs, sortFunc)

	if c.UniqueOnly {
		unique := make([]int, 0)
		if len(linesWithIxs) > 0 {
			unique = append(unique, linesWithIxs[0].ix)
		}
		for i := 1; i < len(linesWithIxs); i++ {
			if sortFunc(linesWithIxs[i], linesWithIxs[i-1]) != 0 {
				unique = append(unique, linesWithIxs[i].ix)
			}
		}
		return unique
	} else {
		lineIxs := make([]int, len(linesWithIxs))
		for i := range linesWithIxs {
			lineIxs[i] = linesWithIxs[i].ix
		}
		return lineIxs
	}
}

func sortNumeric(c *Config, lines []string) []int {
	if c.UniqueOnly {
		unique := make([]int, 0)
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
				unique = append(unique, seen[key.String()])
			}
		}
		return unique
	} else {
		type sIx struct {
			ix int
			s  string
		}
		linesWithIxs := make([]sIx, len(lines))
		for i := range lines {
			linesWithIxs[i] = sIx{i, lines[i]}
		}
		numericSorter := func(s1, s2 sIx) int {
			s1Num, s1Rest, s1Ok := getNumericPrefAndRest(s1.s)
			s2Num, s2Rest, s2Ok := getNumericPrefAndRest(s2.s)
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
			prevSortFunc := numericSorter
			numericSorter = func(s1, s2 sIx) int {
				return prevSortFunc(
					sIx{s1.ix, strings.TrimLeftFunc(s1.s, unicode.IsSpace)},
					sIx{s2.ix, strings.TrimLeftFunc(s2.s, unicode.IsSpace)},
				)
			}
		}
		if c.Reverse {
			prevSortFunc := numericSorter
			numericSorter = func(s1, s2 sIx) int {
				return -prevSortFunc(s1, s2)
			}
		}
		slices.SortStableFunc(linesWithIxs, numericSorter)
		lineIxs := make([]int, len(linesWithIxs))
		for i := range lineIxs {
			lineIxs[i] = linesWithIxs[i].ix
		}
		return lineIxs
	}
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
	sPrefNum := &big.Int{}
	if sNonNumericIx == -1 {
		if err := sPrefNum.UnmarshalText([]byte(s)); err != nil {
			return &big.Int{}, s, false
		} else {
			return sPrefNum, "", true
		}
	} else {
		if err := sPrefNum.UnmarshalText([]byte(s[:sNonNumericIx])); err != nil {
			return &big.Int{}, s, false
		} else {
			return sPrefNum, s[sNonNumericIx:], true
		}
	}
}
