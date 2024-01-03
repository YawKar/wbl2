package sort

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
	"unicode"
)

func Sort(c Config, appName string, filePath string, r io.Reader, w io.Writer) error {
	lines := func() (lines []string) {
		lines = make([]string, 0)
		scanner := bufio.NewScanner(r)
		ok := scanner.Scan()
		for ok {
			lines = append(lines, scanner.Text())
			ok = scanner.Scan()
		}
		return
	}()

	sortFunc := mkSortFunc(c)

	if c.JustCheck {
		return justCheck(c, sortFunc, appName, filePath, lines, w)
	}

	slices.SortFunc(lines, sortFunc)

	if c.UniqueOnly {
		unique := make([]string, 0)
		if len(lines) > 0 {
			unique = append(unique, lines[0])
		}
		for i := 1; i < len(lines); i++ {
			if lines[i] != lines[i-1] {
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
	case c.SortKey.Enabled:
		panic("unimpl")
	case c.SortHumanNumeric:
		panic("unimpl")
	case c.SortMonth:
		panic("unimpl")
	case c.SortNumeric:
		panic("unimpl")
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

func justCheck(c Config, sortFunc func(string, string) int, appName string, filePath string, lines []string, w io.Writer) error {
	for i := 0; i < len(lines)-1; i++ {
		if sortFunc(lines[i], lines[i+1]) == 0 {
			if c.UniqueOnly {
				fmt.Printf("%s: %s:%d: disorder: %s\n", appName, filePath, i+2, lines[i+1])
				return errors.New("")
			}
		} else if sortFunc(lines[i], lines[i+1]) > 0 {
			fmt.Printf("%s: %s:%d: disorder: %s\n", appName, filePath, i+2, lines[i+1])
			return errors.New("")
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
