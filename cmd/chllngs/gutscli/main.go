package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/yawkar/wbl2/pkg/chllngs/guts"
)

func main() {
	var binName string
	if len(os.Args) > 0 {
		binName = filepath.Base(os.Args[0])
	} else {
		binName = "guts"
	}
	conf := parseConfigFromCLI()
	if err := guts.ValidateRanges(conf.FieldsRanges); err != nil {
		fmt.Fprintf(os.Stderr, "%s: err: %s\n", binName, err)
		os.Exit(1)
	}
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if output, err := guts.Cut(conf, line); err != nil {
			fmt.Fprintf(os.Stderr, "%s: err: %s\n", binName, err)
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stdout, output)
		}
	}
	if err := scan.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: err: %s\n", binName, err)
		os.Exit(1)
	}
}

func parseConfigFromCLI() *guts.Config {
	c := &guts.Config{}
	c.FieldsRanges = make([]guts.Range, 0)

	flag.BoolVar(&c.OnlyDelimited, "s", false, "do not print lines not containing delimiters")
	flag.StringVar(&c.Delimiter, "d", "", "use DELIM instead of TAB for field delimiter")
	flag.Func("f", "select only these fields (1-2,4)", func(s string) error {
		ranges := strings.Split(s, ",")
		if len(ranges) == 0 {
			return errors.New("should use at least 1 field")
		}
		for rngx, rngS := range ranges {
			if iField, err := strconv.ParseInt(rngS, 10, 32); err == nil {
				c.FieldsRanges = append(c.FieldsRanges, guts.Range{Leftx: int(iField), Rigthx: int(iField)})
			} else {
				var leftx, rigthx int
				if _, err := fmt.Sscanf(rngS, "%d-%d", &leftx, &rigthx); err != nil {
					return fmt.Errorf("couldn't parse %d range: %s", rngx, err)
				} else {
					c.FieldsRanges = append(c.FieldsRanges, guts.Range{Leftx: leftx, Rigthx: rigthx})
				}
			}
		}
		return nil
	})
	flag.Parse()

	return c
}
