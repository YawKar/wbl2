package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/yawkar/wbl2/dev_challenges/sortcli/sort"
)

type CLIConfig struct {
	sort.Config
	justCheck bool
}

func main() {
	c := parseConfigUsingFlags()
	restArgs := flag.Args()
	var binName string
	if len(os.Args) > 0 {
		binName = filepath.Base(os.Args[0])
	} else {
		binName = "sortcli"
	}
	if len(restArgs) == 0 {
		fmt.Fprintf(os.Stderr, "%s: file wasn't provided (make sure it is the last arg)\n", binName)
		os.Exit(1)
	}
	for _, filePath := range restArgs {
		filePath := filepath.Clean(filePath)
		if file, err := os.Open(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "%s: file wasn't found: %s\n", binName, filePath)
		} else {
			if c.justCheck {
				if err := sort.JustCheck(&c.Config, file); err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s: %s\n", binName, filePath, err)
				}
			} else {
				if err := sort.Sort(&c.Config, file, os.Stdout); err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s: %s\n", binName, filePath, err)
				}
			}
		}
	}
}

func parseConfigUsingFlags() *CLIConfig {
	c := &CLIConfig{}

	flag.BoolVar(&c.SortNumeric, "n", false, "numeric sort")
	flag.BoolVar(&c.Reverse, "r", false, "reverse order")
	flag.BoolVar(&c.IgnoreLeadingBlanks, "b", false, "ignore leading blanks")
	flag.BoolVar(&c.OtherOptions.UniqueOnly, "u", false, "keep only unique")
	flag.BoolVar(&c.justCheck, "c", false, "just check the file")
	flag.Func("k", "sort by kth column (sep is space)", func(s string) error {
		if res, err := strconv.ParseInt(s, 10, 32); err != nil {
			return err
		} else if res < 1 {
			return fmt.Errorf("%v is less than 1", res)
		} else {
			c.SortKey = &sort.SortKey{TargetField: int(res)}
			return nil
		}
	})

	flag.Parse()
	return c
}
