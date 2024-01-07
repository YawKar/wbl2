package guts

import (
	"testing"
)

type TestCase struct {
	content []string
	config  *Config
	output  []string
}

var tcases = []TestCase{
	{
		content: []string{
			"abc	def	gth",
			"",
			"lol aoe	third",
			"	eu",
		},
		config: &Config{
			FieldsRanges: []Range{
				{Leftx: 1, Rigthx: 1},
			},
		},
		output: []string{
			"abc",
			"",
			"lol aoe",
			"",
		},
	},
	{
		content: []string{
			"abc	def	gth",
			"",
			"lol aoe	third",
			"	eu",
		},
		config: &Config{
			FieldsRanges: []Range{
				{Leftx: 2, Rigthx: 2},
			},
		},
		output: []string{
			"def",
			"",
			"third",
			"eu",
		},
	},
	{
		content: []string{
			"abc	def	gth",
			"",
			"lol aoe	third",
			"	eu",
		},
		config: &Config{
			Delimiter: " ",
			FieldsRanges: []Range{
				{Leftx: 2, Rigthx: 2},
			},
		},
		output: []string{
			"",
			"",
			"aoe	third",
			"",
		},
	},
}

func TestSample(t *testing.T) {
	for tcx, tc := range tcases {
		for linex, line := range tc.content {
			if outLine, err := Cut(tc.config, line); err != nil {
				t.Logf("%#v", tc)
				t.Fatalf("#%d tcase:%d line:err: %s", tcx, linex, err)
			} else if outLine != tc.output[linex] {
				t.Logf("%#v", tc)
				t.Fatalf("#%d tcase: strings differ:\n%q\n!=\n%q", tcx, outLine, tc.output[linex])
			}
		}
	}
}
