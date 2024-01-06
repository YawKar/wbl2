package grib

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
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

type grepMatch struct {
	line   int
	startx int
	endx   int // exclusive
}

func Grep(c *Config, r io.Reader, w io.Writer) error {
	var (
		lines   = getLines(r)
		err     error
		matches []grepMatch
	)
	if c.SelectionConfig.FixedStrings {
		if matches, err = grepF(c, lines); err != nil {
			return err
		}
	} else {
		if matches, err = grepG(c, lines); err != nil {
			return err
		}
	}
	if err := printMatches(c, w, matches, lines); err != nil {
		return err
	}
	return nil
}

func formatMatchSnippet(c *Config, lines []string, g grepMatch) string {
	const (
		cRed    = "\033[31m"
		cGreen  = "\033[32m"
		cYellow = "\033[33m"
		cBlue   = "\033[34m"
		cPurple = "\033[35m"
		cCyan   = "\033[36m"
		cWhite  = "\033[37m"

		cOrange = "\033[38:5:208m"

		cBold  = "\033[1m"
		cReset = "\033[0m"
	)
	var (
		matching  string
		matchLine string = lines[g.line]
	)
	if c.Скрепы {
		curx := 0
		rainbow := []string{cRed, cOrange, cYellow, cGreen, cBlue, cPurple}
		matchingSb := strings.Builder{}
		for _, r := range matchLine[g.startx:g.endx] {
			matchingSb.WriteString(rainbow[curx])
			matchingSb.WriteString(cBold)
			curx = (curx + 1) % len(rainbow)
			matchingSb.WriteRune(r)
			matchingSb.WriteString(cReset)
		}
		matching = matchingSb.String()
	} else {
		matching = fmt.Sprintf("%s%s%s%s", cRed, cBold, matchLine[g.startx:g.endx], cReset)
	}

	snippetSb := strings.Builder{}
	var (
		beforeCtx int
		afterCtx  int
	)
	if c.Context > 0 {
		beforeCtx, afterCtx = c.Context, c.Context
	} else {
		beforeCtx, afterCtx = c.Before, c.After
	}
	// add before context
	for leadingx := g.line - beforeCtx; leadingx < g.line; leadingx++ {
		if leadingx < 0 {
			continue
		}
		if c.LineNumber {
			snippetSb.WriteString(fmt.Sprintf("%d:%s\n", leadingx+1, lines[leadingx]))
		} else {
			snippetSb.WriteString(fmt.Sprintf("%s\n", lines[leadingx]))
		}
	}
	// add the matching line
	if c.LineNumber {
		snippetSb.WriteString(fmt.Sprintf("%d:%s%s%s\n", g.line+1, matchLine[:g.startx], matching, matchLine[g.endx:]))
	} else {
		snippetSb.WriteString(fmt.Sprintf("%s%s%s\n", matchLine[:g.startx], matching, matchLine[g.endx:]))
	}
	// add after context
	for trailingx := g.line + 1; trailingx < g.line+afterCtx+1; trailingx++ {
		if trailingx >= len(lines) {
			break
		}
		if c.LineNumber {
			snippetSb.WriteString(fmt.Sprintf("%d:%s\n", trailingx+1, lines[trailingx]))
		} else {
			snippetSb.WriteString(fmt.Sprintf("%s\n", lines[trailingx]))
		}
	}
	return snippetSb.String()
}

func printMatches(c *Config, w io.Writer, matches []grepMatch, lines []string) error {
	if c.Count {
		if c.InvertMatch {
			fmt.Fprintf(w, "%d\n", len(lines)-len(matches))
		} else {
			fmt.Fprintf(w, "%d\n", len(matches))
		}
		return nil
	}
	if c.InvertMatch {
		matches := matches
		for linex := range lines {
			if len(matches) > 0 {
				if matches[0].line == linex {
					matches = matches[1:]
				} else {
					if _, err := fmt.Fprintf(w, "%s", formatMatchSnippet(c, lines, grepMatch{line: linex})); err != nil {
						return err
					}
				}
			} else {
				if _, err := fmt.Fprintf(w, "%s", formatMatchSnippet(c, lines, grepMatch{line: linex})); err != nil {
					return err
				}
			}
		}
	} else {
		for _, gMatch := range matches {
			if _, err := fmt.Fprintf(w, "%s", formatMatchSnippet(c, lines, gMatch)); err != nil {
				return err
			}
		}
	}
	return nil
}

func grepF(c *Config, lines []string) ([]grepMatch, error) {
	gMatches := make([]grepMatch, 0)
	for linex, line := range lines {
		match := false
		startx := -1
		endx := -1
		if c.IgnoreCase {
			line = strings.ToLower(line)
		}
		for _, pattern := range c.Patterns {
			if c.IgnoreCase {
				if startx = strings.Index(line, strings.ToLower(pattern)); startx != -1 {
					match = true
					endx = startx + len(pattern)
					break
				}
			} else {
				if startx = strings.Index(line, pattern); startx != -1 {
					match = true
					endx = startx + len(pattern)
					break
				}
			}
		}
		if match {
			gMatches = append(gMatches, grepMatch{
				line:   linex,
				startx: startx,
				endx:   endx,
			})
		}
	}
	return gMatches, nil
}

func grepG(c *Config, lines []string) ([]grepMatch, error) {
	gMatches := make([]grepMatch, 0)
	for linex, line := range lines {
		match := false
		startx := -1
		endx := -1
		for _, pattern := range c.Patterns {
			if compiledPattern, err := regexp.Compile(pattern); err != nil {
				return nil, err
			} else if loc := compiledPattern.FindStringIndex(line); loc != nil {
				match = true
				startx, endx = loc[0], loc[1]
				break
			}
		}
		if match {
			gMatches = append(gMatches, grepMatch{
				line:   linex,
				startx: startx,
				endx:   endx,
			})
		}
	}
	return gMatches, nil
}
