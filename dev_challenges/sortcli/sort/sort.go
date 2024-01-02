package sort

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
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

	if c.JustCheck {
		return justCheck(c, appName, filePath, lines, w)
	}

	slices.SortFunc(lines, basicSort(c.SortReverse))

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

func justCheck(c Config, appName string, filePath string, lines []string, w io.Writer) error {
	for i := 0; i < len(lines)-1; i++ {
		if lines[i] == lines[i+1] {
			if c.UniqueOnly {
				fmt.Printf("%s: %s:%d: disorder: %s\n", appName, filePath, i+2, lines[i+1])
				return errors.New("")
			}
		} else if lines[i] < lines[i+1] {
			if c.SortReverse {
				fmt.Printf("%s: %s:%d: disorder: %s\n", appName, filePath, i+2, lines[i+1])
				return errors.New("")
			}
		} else { // if lines[i] > lines[i+1]
			if !c.SortReverse {
				fmt.Printf("%s: %s:%d: disorder: %s\n", appName, filePath, i+2, lines[i+1])
				return errors.New("")
			}
		}
	}
	return nil
}

func basicSort(reverse bool) func(string, string) int {
	if reverse {
		return func(s1, s2 string) int {
			return -strings.Compare(s1, s2)
		}
	}
	return strings.Compare
}
