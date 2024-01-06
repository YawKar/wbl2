package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yawkar/wbl2/pkg/chllngs/grib"
)

func main() {
	conf, restXs := parseConfigFromCLI()
	var binName string
	if len(os.Args) > 0 {
		binName = filepath.Base(os.Args[0])
	} else {
		binName = "grib"
	}
	if len(restXs) == 0 {
		fmt.Fprintf(os.Stderr, "%s: err: pattern and file weren't provided\n", binName)
		os.Exit(1)
	}
	if len(restXs) == 1 {
		fmt.Fprintf(os.Stderr, "%s: err: file wasn't provided\n", binName)
		os.Exit(1)
	}
	conf.Patterns = restXs[:len(restXs)-1]
	filePath := restXs[len(restXs)-1]
	if file, err := os.Open(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "%s: err: %s\n", binName, err)
		os.Exit(1)
	} else {
		if err := grib.Grep(conf, file, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "%s: err: %s\n", binName, err)
			os.Exit(1)
		}
	}
}

func parseConfigFromCLI() (c *grib.Config, restArgs []string) {
	c = &grib.Config{}

	flag.BoolVar(&c.Скрепы, "скрепы", false, "who called Nyan Cat?")
	flag.BoolVar(&c.IgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&c.FixedStrings, "F", false, "patterns are strings")
	flag.BoolVar(&c.LineNumber, "n", false, "print line number")
	flag.BoolVar(&c.Count, "c", false, "print only a count of selected lines")
	flag.BoolVar(&c.InvertMatch, "v", false, "select non-matching lines")
	flag.IntVar(&c.After, "A", 0, "print NUM lines of trailing context")
	flag.IntVar(&c.Before, "B", 0, "print NUM lines of leading context")
	flag.IntVar(&c.Context, "C", 0, "print NUM lines of surrounding context (like `-A NUM -B NUM`)")

	flag.Parse()

	return c, flag.Args()
}
