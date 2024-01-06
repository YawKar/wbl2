package grib

import (
	"bytes"
	"testing"
)

type TestCase struct {
	content string
	config  *Config
	output  string
}

var tcases = []TestCase{
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				FixedStrings: true,
				Patterns:     []string{"hello", "Test"},
			},
		},
		output: "\t\tfunc \x1b[31m\x1b[1mTest\x1b[0mSample(t *testing.T) {\n" +
			"\t\t\t\t\tPatterns:     []string{\"\x1b[31m\x1b[1mhello\x1b[0m\", \"Test\"},\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				FixedStrings: true,
				IgnoreCase:   true,
				Patterns:     []string{"T"},
			},
		},
		output: "\t\tfunc \x1b[31m\x1b[1mT\x1b[0mestSample(t *testing.T) {\n" +
			"\t\t\t\tSelec\x1b[31m\x1b[1mt\x1b[0mionConfig: SelectionConfig{\n" +
			"\t\t\t\t\tFixedS\x1b[31m\x1b[1mt\x1b[0mrings: true,\n" +
			"\t\t\t\t\tPa\x1b[31m\x1b[1mt\x1b[0mterns:     []string{\"hello\", \"Test\"},\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				Patterns: []string{"\"[^\"]+\""},
			},
		},
		output: "\t\t\t\t\tPatterns:     []string{\x1b[31m\x1b[1m\"hello\"\x1b[0m, \"Test\"},\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				Patterns: []string{"\"[^\"]+\""},
			},
			ContextConfig: ContextConfig{
				Context: 2,
			},
			MiscellaneousConfig: MiscellaneousConfig{
				Скрепы: true,
			},
			OutputConfig: OutputConfig{
				LineNumber: true,
			},
		},
		output: "4:\t\t\t\tSelectionConfig: SelectionConfig{\n" +
			"5:\t\t\t\t\tFixedStrings: true,\n" +
			"6:\t\t\t\t\tPatterns:     []string{\x1b[31m\x1b[1m\"\x1b[0m\x1b[38:5:208m\x1b[1mh\x1b[0m\x1b[33m\x1b[1me\x1b[0m\x1b[32m\x1b[1ml\x1b[0m\x1b[34m\x1b[1ml\x1b[0m\x1b[35m\x1b[1mo\x1b[0m\x1b[31m\x1b[1m\"\x1b[0m, \"Test\"},\n" +
			"7:\t\t\t\t},\n" +
			"8:\t\t\t}\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				Patterns: []string{"\"[^\"]+\""},
			},
			ContextConfig: ContextConfig{
				Context: 2,
			},
			MiscellaneousConfig: MiscellaneousConfig{
				Скрепы: true,
			},
			OutputConfig: OutputConfig{
				LineNumber: true,
				Count:      true,
			},
		},
		output: "1\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				Patterns: []string{"\"[^\"]+\""},
			},
			ContextConfig: ContextConfig{
				Context: 2,
			},
			MiscellaneousConfig: MiscellaneousConfig{
				Скрепы:      true,
				InvertMatch: true,
			},
			OutputConfig: OutputConfig{
				LineNumber: true,
				Count:      true,
			},
		},
		output: "9\n",
	},
	{
		content: `
		func TestSample(t *testing.T) {
			c := &Config{
				SelectionConfig: SelectionConfig{
					FixedStrings: true,
					Patterns:     []string{"hello", "Test"},
				},
			}
		}
		`,
		config: &Config{
			SelectionConfig: SelectionConfig{
				Patterns: []string{"\"[^\"]+\""},
			},
			ContextConfig: ContextConfig{
				Context: 2,
			},
			MiscellaneousConfig: MiscellaneousConfig{
				Скрепы:      true,
				InvertMatch: true,
			},
			OutputConfig: OutputConfig{
				LineNumber: true,
			},
		},
		output: "1:\n" +
			"2:\t\tfunc TestSample(t *testing.T) {\n" +
			"3:\t\t\tc := &Config{\n" +
			"1:\n" +
			"2:\t\tfunc TestSample(t *testing.T) {\n" +
			"3:\t\t\tc := &Config{\n" +
			"4:\t\t\t\tSelectionConfig: SelectionConfig{\n" +
			"1:\n" +
			"2:\t\tfunc TestSample(t *testing.T) {\n" +
			"3:\t\t\tc := &Config{\n" +
			"4:\t\t\t\tSelectionConfig: SelectionConfig{\n" +
			"5:\t\t\t\t\tFixedStrings: true,\n" +
			"2:\t\tfunc TestSample(t *testing.T) {\n" +
			"3:\t\t\tc := &Config{\n" +
			"4:\t\t\t\tSelectionConfig: SelectionConfig{\n" +
			"5:\t\t\t\t\tFixedStrings: true,\n" +
			"6:\t\t\t\t\tPatterns:     []string{\"hello\", \"Test\"},\n" +
			"3:\t\t\tc := &Config{\n" +
			"4:\t\t\t\tSelectionConfig: SelectionConfig{\n" +
			"5:\t\t\t\t\tFixedStrings: true,\n" +
			"6:\t\t\t\t\tPatterns:     []string{\"hello\", \"Test\"},\n" +
			"7:\t\t\t\t},\n" +
			"5:\t\t\t\t\tFixedStrings: true,\n" +
			"6:\t\t\t\t\tPatterns:     []string{\"hello\", \"Test\"},\n" +
			"7:\t\t\t\t},\n" +
			"8:\t\t\t}\n" +
			"9:\t\t}\n" +
			"6:\t\t\t\t\tPatterns:     []string{\"hello\", \"Test\"},\n" +
			"7:\t\t\t\t},\n" +
			"8:\t\t\t}\n" +
			"9:\t\t}\n" +
			"10:\t\t\n" +
			"7:\t\t\t\t},\n" +
			"8:\t\t\t}\n" +
			"9:\t\t}\n" +
			"10:\t\t\n" +
			"8:\t\t\t}\n" +
			"9:\t\t}\n" +
			"10:\t\t\n",
	},
}

func TestSample(t *testing.T) {
	for tcx, tc := range tcases {
		resB := bytes.NewBufferString("")
		if err := Grep(tc.config, bytes.NewReader([]byte(tc.content)), resB); err != nil {
			t.Logf("%#v", tc)
			t.Fatalf("#%d tcase: err: %s", tcx, err)
		} else if res := resB.String(); res != tc.output {
			t.Logf("%#v", tc)
			t.Fatalf("#%d tcase: strings differ:\n%q\n!=\n%q", tcx, res, tc.output)
		}
	}
}
