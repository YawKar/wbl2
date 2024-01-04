package sort

import (
	"fmt"
	"io"
	"strings"
)

func JustCheck(c *Config, r io.Reader) error {
	switch {
	case c.SortNumeric:
		if err := checkNumeric(c, r); err != nil {
			return err
		}
	default:
		if err := checkStandard(c, r); err != nil {
			return err
		}
	}
	return nil
}

func checkStandard(c *Config, r io.Reader) error {
	lines := getLines(r)
	sortFunc := strings.Compare
	if c.IgnoreLeadingBlanks {
		sortFunc = wrapIgnoreLeadingBlanks(sortFunc)
	}
	if c.Reverse {
		sortFunc = wrapReverse(sortFunc)
	}

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

func checkNumeric(c *Config, r io.Reader) error {
	lines := getLines(r)
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
	for i := 0; i < len(lines)-1; i++ {
		if numericSorter(lines[i], lines[i+1]) == 0 {
			if c.UniqueOnly {
				return fmt.Errorf("%d: disorder: %s\n", i+2, lines[i+1])
			}
		} else if numericSorter(lines[i], lines[i+1]) > 0 {
			return fmt.Errorf("%d: disorder: %s\n", i+2, lines[i+1])
		}
	}
	return nil
}
