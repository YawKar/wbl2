package guts

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strings"
)

func Cut(c *Config, line string) (string, error) {
	if err := ValidateRanges(c.FieldsRanges); err != nil {
		return "", err
	}
	delimiter := "\t"
	if c.Delimiter != "" {
		delimiter = c.Delimiter
	}
	outputLineSb := strings.Builder{}
	if c.OnlyDelimited {
		if strings.Contains(line, delimiter) {
			slices := strings.Split(line, delimiter)
			for _, rng := range c.FieldsRanges {
				for i := rng.Leftx; i <= rng.Rigthx; i++ {
					i := i - 1 // 0-indexed
					if len(slices) > i {
						if outputLineSb.Len() > 0 {
							fmt.Fprint(&outputLineSb, delimiter)
						}
						fmt.Fprintf(&outputLineSb, "%s", slices[i])
					} else {
						break
					}
				}
			}
		}
	} else {
		slices := strings.Split(line, delimiter)
		for _, rng := range c.FieldsRanges {
			for i := rng.Leftx; i <= rng.Rigthx; i++ {
				i := i - 1 // 0-indexed
				if len(slices) > i {
					if outputLineSb.Len() > 0 {
						fmt.Fprint(&outputLineSb, delimiter)
					}
					fmt.Fprintf(&outputLineSb, "%s", slices[i])
				} else {
					break
				}
			}
		}
	}
	return outputLineSb.String(), nil
}

func ValidateRanges(ranges []Range) error {
	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return cmp.Compare(r1.Leftx, r2.Leftx)
	})
	if len(ranges) == 0 {
		return errors.New("no field ranges were provided")
	}
	for i := 0; i < len(ranges)-1; i++ {
		if err := validateRange(ranges[i]); err != nil {
			return err
		}
		if ranges[i].Leftx >= ranges[i+1].Leftx {
			return fmt.Errorf("left range's leftx is >= that of the right one: %v and %v", ranges[i], ranges[i+1])
		}
		if ranges[i].Rigthx >= ranges[i+1].Leftx {
			return fmt.Errorf("left range's rigthx is >= the leftx of the rigth range: %v and %v", ranges[i], ranges[i+1])
		}
	}
	return nil
}

func validateRange(rng Range) error {
	if rng.Leftx > rng.Rigthx {
		return fmt.Errorf("range has reversed borders: %v", rng)
	}
	return nil
}
