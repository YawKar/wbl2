package anagramoogle

import (
	"maps"
	"slices"
	"testing"
)

func TestXxx(t *testing.T) {
	type Test struct {
		input  []string
		output map[string][]string
		err    error
	}
	tests := []Test{
		{
			input: []string{
				"пятак", "пЯтка", "тяпка",
				"ЛИСТОК", "сЛиТоК", "столик",
				"лол",
				"КРЯ", "ярк",
			},
			output: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
				"кря":    {"кря", "ярк"},
			},
		},
	}
	for i, tcase := range tests {
		if res, err := GroupIntoAnagrams(tcase.input); err != nil {
			if tcase.err == nil || tcase.err.Error() != err.Error() {
				t.Errorf("%d tcase: errors don't match: %s != %s", i, err, tcase.err)
			}
		} else if res == nil {
			t.Errorf("%d tcase: ptr to result is nil", i)
		} else if *res == nil {
			t.Errorf("%d tcase: result map is nil", i)
		} else if !maps.EqualFunc(*res, tcase.output, slices.Equal) {
			t.Errorf("%d tcase: maps differ:\nlen=%d\n%v\n!=\nlen=%d\n%v", i, len(*res), *res, len(tcase.output), tcase.output)
		}
	}
}
